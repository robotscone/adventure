package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"time"

	"github.com/robotscone/adventure/internal/ease"
	"github.com/robotscone/adventure/internal/event"
	"github.com/robotscone/adventure/internal/gfx"
	"github.com/robotscone/adventure/internal/input"
	"github.com/robotscone/adventure/internal/linalg"
	"github.com/robotscone/adventure/internal/state"
	"github.com/robotscone/adventure/internal/timer"
	"github.com/veandco/go-sdl2/sdl"
)

type Entity struct {
	Position   linalg.Vec2
	Sprite     *gfx.Sprite
	Controller *state.FSM
}

type EntityBaseState struct {
	state.Base
	Entity *Entity
}

func (s *EntityBaseState) Render() {
	s.Entity.Sprite.Draw(s.Entity.Position.X, s.Entity.Position.Y)
}

type IdleState struct {
	*EntityBaseState
}

func (s *IdleState) Enter(controller state.Controller, data *state.Data, message interface{}) {
	s.Entity.Sprite.SetAnimation("idle")
}

type MoveState struct {
	*EntityBaseState
	direction linalg.Vec2
}

func (s *MoveState) Enter(controller state.Controller, data *state.Data, message interface{}) {
	s.direction, _ = message.(linalg.Vec2)
	s.direction = s.direction.Norm()
}

func (s *MoveState) Update(controller state.Controller, data *state.Data) {
	switch {
	case s.direction.Y < 0:
		s.Entity.Sprite.SetAnimation("walk:up")
	case s.direction.Y > 0:
		s.Entity.Sprite.SetAnimation("walk:down")
	case s.direction.X < 0:
		s.Entity.Sprite.SetAnimation("walk:left")
	case s.direction.X > 0:
		s.Entity.Sprite.SetAnimation("walk:right")
	default:
		s.Entity.Sprite.SetAnimation("idle")
	}

	speed := 70.0
	s.Entity.Position.X += speed * data.Delta * s.direction.X
	s.Entity.Position.Y += speed * data.Delta * s.direction.Y
}

type HeroIdleState struct {
	*IdleState
}

func (s *HeroIdleState) Input(controller state.Controller, data *state.Data) {
	up := data.Device.Get("up")
	down := data.Device.Get("down")
	left := data.Device.Get("left")
	right := data.Device.Get("right")

	canMoveY := !(up.IsDown && down.IsDown) && (up.IsDown || down.IsDown)
	canMoveX := !(left.IsDown && right.IsDown) && (left.IsDown || right.IsDown)

	if !canMoveX && !canMoveY {
		return
	}

	controller.Switch("move", nil)
}

type HeroMoveState struct {
	*MoveState
}

func (s *HeroMoveState) Input(controller state.Controller, data *state.Data) {
	up := data.Device.Get("up")
	down := data.Device.Get("down")
	left := data.Device.Get("left")
	right := data.Device.Get("right")

	// If all inputs are down then we can't move so we switch to idle
	allDown := up.IsDown && down.IsDown && left.IsDown && right.IsDown
	noneDown := !up.IsDown && !down.IsDown && !left.IsDown && !right.IsDown
	if allDown || noneDown {
		controller.Switch("idle", nil)

		return
	}

	switch {
	case up.IsDown && down.IsDown:
		s.direction.Y = 0
	case up.IsDown:
		s.direction.Y = -1
	case down.IsDown:
		s.direction.Y = 1
	default:
		s.direction.Y = 0
	}

	switch {
	case left.IsDown && right.IsDown:
		s.direction.X = 0
	case left.IsDown:
		s.direction.X = -1
	case right.IsDown:
		s.direction.X = right.Value // [0, 1] ---> 0.5
	default:
		s.direction.X = 0
	}

	s.direction = s.direction.Norm()
}

type NPCIdleState struct {
	*IdleState
	timer time.Duration
}

func (s *NPCIdleState) Enter(controller state.Controller, data *state.Data, message interface{}) {
	s.IdleState.Enter(controller, data, message)

	s.timer = 0
}

func (s *NPCIdleState) Update(controller state.Controller, data *state.Data) {
	s.timer += time.Duration(data.Delta * float64(time.Second))

	if s.timer >= 2*time.Second {
		var direction linalg.Vec2

		for direction.X == 0 && direction.Y == 0 {
			direction.X = float64(rand.Intn(3) - 1)
			direction.Y = float64(rand.Intn(3) - 1)
		}

		controller.Switch("stroll", direction)
	}
}

type NPCStrollState struct {
	*MoveState
	timer time.Duration
}

func (s *NPCStrollState) Enter(controller state.Controller, data *state.Data, message interface{}) {
	s.MoveState.Enter(controller, data, message)

	s.timer = 0
}

func (s *NPCStrollState) Update(controller state.Controller, data *state.Data) {
	s.timer += time.Duration(data.Delta * float64(time.Second))

	s.MoveState.Update(controller, data)

	if s.timer >= 500*time.Millisecond {
		controller.Switch("idle", nil)
	}
}

type FadeInState struct {
	state.Base
	renderer *gfx.Renderer
	tween    *ease.Tween
	rect     sdl.FRect
	color    int
}

func (s *FadeInState) Init(controller state.Controller, data *state.Data) {
	s.tween = ease.NewTween(1, 0, 2*time.Second, ease.Linear)
	s.tween.OnFinished(func(t *ease.Tween) {
		controller.Pop()
	})

	s.rect.X = 0
	s.rect.Y = 0
	s.rect.W = 1000
	s.rect.H = 1000
}

func (s *FadeInState) Enter(controller state.Controller, data *state.Data, message interface{}) {
	s.color, _ = message.(int)
	s.color = s.color<<8 | 0xFF
	s.tween.Reset()
}

func (s *FadeInState) Update(controller state.Controller, data *state.Data) {
	s.tween.Update(data.Delta)
}

func (s *FadeInState) Render() {
	r := uint8(s.color >> 24 & 0xFF)
	g := uint8(s.color >> 16 & 0xFF)
	b := uint8(s.color >> 8 & 0xFF)
	a := uint8(float64(s.color&0xFF) * s.tween.Value)

	s.renderer.SetDrawColor(r, g, b, a)
	s.renderer.FillRectF(&s.rect)
}

type ExploreState struct {
	state.Base
	renderer *gfx.Renderer
	timer    *timer.Timer
	entities []*Entity
}

func (s *ExploreState) Init(controller state.Controller, data *state.Data) {
	s.timer = timer.New()

	s.timer.Every(1*time.Second, 4, func() { fmt.Println("Tick...") })
	s.timer.After(5*time.Second, func() { fmt.Println("Boom!") })

	var walkDown gfx.Animation
	walkDown.AddFrame(30*0, 30*1, 22, 24, gfx.FlipNone)
	walkDown.AddFrame(30*1, 30*1, 22, 24, gfx.FlipNone)
	walkDown.AddFrame(30*2, 30*1, 22, 24, gfx.FlipNone)
	walkDown.AddFrame(30*3, 30*1, 22, 24, gfx.FlipNone)
	walkDown.AddFrame(30*4, 30*1, 22, 24, gfx.FlipNone)
	walkDown.AddFrame(30*5, 30*1, 22, 24, gfx.FlipNone)
	walkDown.AddFrame(30*6, 30*1, 22, 24, gfx.FlipNone)
	walkDown.AddFrame(30*7, 30*1, 22, 24, gfx.FlipNone)
	walkDown.SetFPS(24)

	var walkUp gfx.Animation
	walkUp.AddFrame(30*0, 30*4, 22, 24, gfx.FlipNone)
	walkUp.AddFrame(30*1, 30*4, 22, 24, gfx.FlipNone)
	walkUp.AddFrame(30*2, 30*4, 22, 24, gfx.FlipNone)
	walkUp.AddFrame(30*3, 30*4, 22, 24, gfx.FlipNone)
	walkUp.AddFrame(30*4, 30*4, 22, 24, gfx.FlipNone)
	walkUp.AddFrame(30*5, 30*4, 22, 24, gfx.FlipNone)
	walkUp.AddFrame(30*6, 30*4, 22, 24, gfx.FlipNone)
	walkUp.AddFrame(30*7, 30*4, 22, 24, gfx.FlipNone)
	walkUp.SetFPS(24)

	var walkLeft gfx.Animation
	walkLeft.AddFrame(30*8, 30*1, 22, 24, gfx.FlipNone)
	walkLeft.AddFrame(30*9, 30*1, 22, 24, gfx.FlipNone)
	walkLeft.AddFrame(30*10, 30*1, 22, 24, gfx.FlipNone)
	walkLeft.AddFrame(30*11, 30*1, 22, 24, gfx.FlipNone)
	walkLeft.AddFrame(30*12, 30*1, 22, 24, gfx.FlipNone)
	walkLeft.AddFrame(30*13, 30*1, 22, 24, gfx.FlipNone)
	walkLeft.SetFPS(24)

	var walkRight gfx.Animation
	walkRight.AddFrame(30*8, 30*4, 22, 24, gfx.FlipNone)
	walkRight.AddFrame(30*9, 30*4, 22, 24, gfx.FlipNone)
	walkRight.AddFrame(30*10, 30*4, 22, 24, gfx.FlipNone)
	walkRight.AddFrame(30*11, 30*4, 22, 24, gfx.FlipNone)
	walkRight.AddFrame(30*12, 30*4, 22, 24, gfx.FlipNone)
	walkRight.AddFrame(30*13, 30*4, 22, 24, gfx.FlipNone)
	walkRight.SetFPS(24)

	var idle gfx.Animation
	idle.AddFrame(30*1, 30*0, 22, 24, gfx.FlipNone)
	idle.SetFPS(24)

	{
		hero := &Entity{
			Sprite:     gfx.NewSprite(s.renderer.NewTexture("assets/link.png", gfx.ScaleNearest)),
			Controller: state.NewFSM(data),
		}

		hero.Sprite.RegisterAnimation("walk:up", walkUp)
		hero.Sprite.RegisterAnimation("walk:down", walkDown)
		hero.Sprite.RegisterAnimation("walk:left", walkLeft)
		hero.Sprite.RegisterAnimation("walk:right", walkRight)
		hero.Sprite.RegisterAnimation("idle", idle)
		hero.Sprite.SetAnimation("idle")

		base := &EntityBaseState{Entity: hero}
		hero.Controller.RegisterState("idle", &HeroIdleState{IdleState: &IdleState{EntityBaseState: base}})
		hero.Controller.RegisterState("move", &HeroMoveState{MoveState: &MoveState{EntityBaseState: base}})
		hero.Controller.Switch("idle", nil)

		s.entities = append(s.entities, hero)
	}

	{
		npc := &Entity{
			Position:   linalg.Vec2{X: 75, Y: 75},
			Sprite:     gfx.NewSprite(s.renderer.NewTexture("assets/link.png", gfx.ScaleNearest)),
			Controller: state.NewFSM(data),
		}

		npc.Sprite.RegisterAnimation("walk:up", walkUp)
		npc.Sprite.RegisterAnimation("walk:down", walkDown)
		npc.Sprite.RegisterAnimation("walk:left", walkLeft)
		npc.Sprite.RegisterAnimation("walk:right", walkRight)
		npc.Sprite.RegisterAnimation("idle", idle)
		npc.Sprite.SetAnimation("idle")

		base := &EntityBaseState{Entity: npc}
		npc.Controller.RegisterState("idle", &NPCIdleState{IdleState: &IdleState{EntityBaseState: base}})
		npc.Controller.RegisterState("stroll", &NPCStrollState{MoveState: &MoveState{EntityBaseState: base}})
		npc.Controller.Switch("idle", nil)

		s.entities = append(s.entities, npc)
	}
}

func (s *ExploreState) Enter(controller state.Controller, data *state.Data, message interface{}) {
	controller.Push("fade", nil)
}

func (s *ExploreState) Resume(controller state.Controller, data *state.Data) {
	for _, entity := range s.entities {
		entity.Sprite.SetAlphaMod(1)
		entity.Sprite.SetColorMod(1, 1, 1)
	}
}

func (s *ExploreState) Input(controller state.Controller, data *state.Data) {
	if down := data.Device.Get("right"); down.IsDown {
		fmt.Println("Right down for", down.DownDuration)
	}

	if data.Device.Get("pause").IsPressed {
		for _, entity := range s.entities {
			entity.Sprite.SetAlphaMod(0.5)
			entity.Sprite.SetColorMod(0.9, 0.3, 0.45)
		}

		controller.Push("pause", nil)
	}

	for _, entity := range s.entities {
		entity.Controller.Input()
	}
}

func (s *ExploreState) Update(controller state.Controller, data *state.Data) {
	s.timer.Update(data.Delta)

	for _, entity := range s.entities {
		entity.Controller.Update()
		entity.Sprite.Update(data.Delta)
	}
}

func (s *ExploreState) Render() {
	sort.SliceStable(s.entities, func(i, j int) bool { return s.entities[i].Position.Y < s.entities[j].Position.Y })

	for _, entity := range s.entities {
		entity.Controller.Render()
	}
}

type PauseState struct {
	state.Base
	renderer *gfx.Renderer
	tween    *ease.Tween
	rect     sdl.FRect
}

func (s *PauseState) Init(controller state.Controller, data *state.Data) {
	s.tween = ease.NewTween(-100, 0, 1*time.Second, ease.BounceOut)
	s.tween.OnFinished(func(t *ease.Tween) {
		if t.IsBackward() {
			controller.Pop()
		}
	})
}

func (s *PauseState) Enter(controller state.Controller, data *state.Data, message interface{}) {
	s.tween.Reset()

	s.rect.W = 100
	s.rect.H = 150
	s.rect.X = float32(s.tween.From)
	s.rect.Y = 0
}

func (s *PauseState) Input(controller state.Controller, data *state.Data) {
	if data.Device.Get("pause").IsPressed {
		s.tween.InvertOrReverse(true)
	}
}

func (s *PauseState) Update(controller state.Controller, data *state.Data) {
	s.tween.Update(data.Delta)

	s.rect.X = float32(s.tween.Value)
}

func (s *PauseState) Render() {
	s.renderer.SetDrawColor(155, 175, 198, 150)
	s.renderer.FillRectF(&s.rect)
}

func init() {
	// Ensure that the main function runs on the main thread
	// This will prevent any crashes where certain SDL2 functions expect to be
	// on the main thread
	runtime.LockOSThread()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_GAMECONTROLLER); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "0")

	window, err := sdl.CreateWindow("Adventure", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_OPENGL|sdl.WINDOW_RESIZABLE)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := gfx.NewRenderer(window)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	renderer.SetLogicalSize(200, 150)

	var bindings input.BindingMap
	if b, err := os.ReadFile("assets/bindings.json"); err == nil {
		err := json.Unmarshal(b, &bindings)
		if err != nil {
			log.Println(err)
		}
	}

	device := input.NewDevice(bindings)
	broker := event.NewBroker()
	data := state.Data{Device: device}

	game := state.NewFSM(&data)
	game.RegisterState("fade", &FadeInState{renderer: renderer})
	game.RegisterState("explore", &ExploreState{renderer: renderer})
	game.RegisterState("pause", &PauseState{renderer: renderer})
	game.Switch("explore", nil)

	lastFrame := time.Now()

	input.Init()
	var quit bool
	for !quit {
		data.Delta = time.Since(lastFrame).Seconds()
		lastFrame = time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				quit = true
			}
		}

		input.Update(renderer)

		game.Input()
		game.Update()

		renderer.SetDrawColor(0, 0, 0, 0xFF)
		renderer.Clear()

		game.Render()

		renderer.Present()

		broker.Process()
	}
}
