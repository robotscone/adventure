package gfx

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Sprite struct {
	texture    *Texture
	dst        sdl.FRect
	animation  *Animation
	animations map[string]*Animation
}

func NewSprite(texture *Texture) *Sprite {
	var animation Animation

	animation.AddFrame(0, 0, texture.width, texture.height, FlipNone)

	return &Sprite{
		texture: texture,
		dst: sdl.FRect{
			W: float32(texture.width),
			H: float32(texture.height),
		},
		animation:  &animation,
		animations: make(map[string]*Animation),
	}
}

func (s *Sprite) AddAnimation(name string, animation *Animation) {
	s.animations[name] = animation
}

func (s *Sprite) SetAnimation(name string) {
	s.animation = s.animations[name]

	s.animation.Reset()
}

func (s *Sprite) Update(delta float64) {
	s.animation.Step(delta)
}

func (s *Sprite) Draw(x, y float64) {
	s.dst.X = float32(x)
	s.dst.Y = float32(y)
	s.dst.W = float32(s.animation.frame.src.W)
	s.dst.H = float32(s.animation.frame.src.H)

	s.texture.Draw(&s.animation.frame.src, &s.dst, s.animation.frame.flip)
}
