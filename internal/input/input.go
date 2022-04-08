package input

import (
	"math"
	"strings"
	"time"

	"github.com/robotscone/adventure/internal/gfx"
	"github.com/robotscone/adventure/internal/linalg"
	"github.com/veandco/go-sdl2/sdl"
)

var unmappedButton = &Button{}

type Button struct {
	Value        float64
	deadZone     float64
	IsDown       bool
	IsPressed    bool
	IsReleased   bool
	PressedAt    time.Time
	ReleasedAt   time.Time
	DownDuration time.Duration
}

var Mouse = struct {
	*Device
	Position linalg.Vec2
	Delta    linalg.Vec2
	current  map[string]*mouseButton
	previous map[string]*mouseButton
}{
	Device: NewDevice(BindingMap{
		"left":   {"mouse:left"},
		"middle": {"mouse:middle"},
		"right":  {"mouse:right"},
		"extra1": {"mouse:extra1"},
		"extra2": {"mouse:extra2"},
	}),
	current:  newMouseButtons(),
	previous: newMouseButtons(),
}

type controller struct {
	id sdl.JoystickID
	*sdl.GameController
	current  map[string]*controllerButton
	previous map[string]*controllerButton
}

var controllers []*controller

var keyboard struct {
	current  []Button
	previous []Button
}

// BindingMap is a map of action strings to input names that we can look up from SDL.
type BindingMap map[string][]string

// ButtonMap is a map of action strings to actual button state.
type ButtonMap map[string]*Button

// Device represents an input device and must be created using NewDevice.
type Device struct {
	bindings   BindingMap
	previous   ButtonMap
	current    ButtonMap
	controller *controller
}

// Registered devices.
var devices []*Device

func Init() {

	for i := 0; i < sdl.NumJoysticks(); i++ {
		if !sdl.IsGameController(i) {
			continue
		}

		controller := sdl.GameControllerOpen(i)
		if controller == nil {
			continue
		}

		AddController(controller.Joystick().InstanceID())
	}
}

func NewDevice(bindings BindingMap) *Device {
	if bindings == nil {
		bindings = make(BindingMap)
	}

	device := &Device{
		bindings: bindings,
		previous: make(ButtonMap),
		current:  make(ButtonMap),
	}

	for action := range bindings {
		device.previous[action] = &Button{}
		device.current[action] = &Button{}
	}

	devices = append(devices, device)

	return device
}

func AddController(id sdl.JoystickID) {
	device := sdl.GameControllerFromInstanceID(id)
	if device == nil {
		return
	}

	c := &controller{
		id:             id,
		GameController: device,
		current:        newControllerButtons(),
		previous:       newControllerButtons(),
	}

	controllers = append(controllers, c)

	for devicenumber, device := range devices {
		// Device #0 is the mouse, skip over this one.
		if devicenumber == 0 {
			continue
		}

		if device.controller == nil {
			device.controller = c
			break
		}
	}
}

func (d *Device) Get(action string) *Button {
	if button := d.current[action]; button != nil {
		return button
	}

	return unmappedButton
}

func setButtonState(current, previous *Button, value float64, now time.Time) {
	current.Value = value
	current.IsDown = current.Value != 0
	current.IsPressed = !previous.IsDown && current.IsDown
	current.IsReleased = previous.IsDown && !current.IsDown

	if current.IsPressed {
		current.PressedAt = now
	}

	if current.IsReleased {
		current.ReleasedAt = now
	}

	// If the button was pressed later than it was released then it
	// means the button is still being held down, so we should update the
	// down duration value
	if current.ReleasedAt.Before(current.PressedAt) {
		current.DownDuration = time.Since(current.PressedAt)
	}
}

func Update(renderer *gfx.Renderer) {
	now := time.Now()
	mouseX, mouseY, mouseState := sdl.GetMouseState()
	keyboardState := sdl.GetKeyboardState()

	if renderer != nil {
		logicalMouseX, logicalMouseY := renderer.RenderWindowToLogical(int(mouseX), int(mouseY))

		mouseX, mouseY = int32(logicalMouseX), int32(logicalMouseY)
	}

	mousePreviousX := Mouse.Position.X
	mousePreviousY := Mouse.Position.Y
	Mouse.Position.X = float64(mouseX)
	Mouse.Position.Y = float64(mouseY)
	Mouse.Delta.X = Mouse.Position.X - mousePreviousX
	Mouse.Delta.Y = Mouse.Position.Y - mousePreviousY

	// Save the last mouse state so we can do comparisons
	for name, button := range Mouse.current {
		*Mouse.previous[name] = *button
	}

	for name, button := range Mouse.current {
		var value float64
		if mouseState&button.mask != 0 {
			// If the button is down we set value to 1
			value = 1
		}

		setButtonState(&Mouse.current[name].Button, &Mouse.previous[name].Button, value, now)
	}

	for _, controller := range controllers {
		for name, button := range controller.current {
			var value float64
			if button.isAxis {
				value = float64(controller.Axis(sdl.GameControllerAxis(button.code)))

				switch {
				case !button.isNegative && value > button.deadZone:
					value = (value - button.deadZone) / (math.MaxInt16 - button.deadZone)
				case button.isNegative && value < -button.deadZone:
					value = (value + button.deadZone) / (math.MinInt16 + button.deadZone)
				default:
					value = 0
				}
			} else if controller.Button(sdl.GameControllerButton(button.code)) != 0 {
				value = 1
			}

			setButtonState(&controller.current[name].Button, &controller.previous[name].Button, value, now)
		}
	}

	// If the current keyboard state's length is less than SDL is reporting
	// then we just re-make the slice and populate with blank button states
	if len(keyboard.current) < len(keyboardState) {
		keyboard.previous = make([]Button, len(keyboardState))
		keyboard.current = make([]Button, len(keyboardState))
	}

	// Save the last keyboard state so we can do comparisons
	for i, button := range keyboard.current {
		keyboard.previous[i] = button
	}

	// Loop over the current keyboard state to update the current button values
	for i, value := range keyboardState {
		setButtonState(&keyboard.current[i], &keyboard.previous[i], float64(value), now)
	}

	// Loop over all registered devices and update button pointers based
	// on their internal binding maps
	for _, device := range devices {
		for action, buttons := range device.bindings {
			// Save the last device state so we can do comparisons
			*device.previous[action] = *device.current[action]

			previous := device.previous[action]
			current := device.current[action]

			current.Value = 0

			for _, name := range buttons {
				switch {
				case strings.HasPrefix(name, "mouse:"):
					parts := strings.Split(name, ":")
					key := parts[len(parts)-1]

					if button := Mouse.current[key]; button.Value != 0 {
						current.Value = button.Value
					}
				case strings.HasPrefix(name, "keyboard:"):
					code, ok := scancodes[name]
					if !ok {
						continue
					}

					if button := &keyboard.current[code]; button.Value != 0 {
						current.Value = button.Value
					}
				case strings.HasPrefix(name, "gamepad:"):
					parts := strings.Split(name, ":")
					key := strings.Join(parts[1:], ":")

					if device.controller == nil {
						continue
					}

					button, ok := device.controller.current[key]
					if !ok {
						continue
					}
					if button.Value != 0 {
						current.Value = button.Value
					}
				}
			}

			setButtonState(current, previous, current.Value, now)
		}
	}
}
