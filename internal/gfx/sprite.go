package gfx

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Sprite struct {
	*Texture
	src        sdl.Rect
	dst        sdl.FRect
	animation  *Animation
	animations map[string]*Animation
}

func NewSprite(texture *Texture, x, y, width, height int) *Sprite {
	s := &Sprite{
		Texture: texture,
		dst: sdl.FRect{
			W: float32(texture.width),
			H: float32(texture.height),
		},
		animations: make(map[string]*Animation),
	}

	s.Crop(x, y, width, height)

	return s
}

func (s *Sprite) Crop(x, y, width, height int) {
	s.src.X = int32(x)
	s.src.Y = int32(y)
	s.src.W = int32(width)
	s.src.H = int32(height)

	s.animation = nil
}

func (s *Sprite) RegisterAnimation(name string, animation Animation) {
	if _, ok := s.animations[name]; ok {
		panic(fmt.Sprintf("duplicate animation registration for %q", name))
	}

	s.animations[name] = &animation
}

func (s *Sprite) SetAnimation(name string) {
	animation := s.animations[name]
	if animation == nil {
		fmt.Printf("attempted to set unknown animation %q\n", name)

		return
	}

	if s.animation == animation {
		return
	}

	s.animation = animation

	s.animation.Reset()
}

func (s *Sprite) Update(delta float64) {
	if s.animation != nil {
		s.animation.Update(delta)
	}
}

func (s *Sprite) Draw(x, y float64) {
	s.dst.X = float32(x)
	s.dst.Y = float32(y)

	if s.animation != nil {
		s.dst.W = float32(s.animation.frame.src.W)
		s.dst.H = float32(s.animation.frame.src.H)

		s.Texture.Draw(&s.animation.frame.src, &s.dst, s.animation.frame.flip)
	} else {
		s.dst.W = float32(s.src.W)
		s.dst.H = float32(s.src.H)

		s.Texture.Draw(&s.src, &s.dst, FlipNone)
	}
}
