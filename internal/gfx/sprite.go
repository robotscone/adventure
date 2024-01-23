package gfx

import "fmt"

type Sprite struct {
	*Texture
	src        Rect
	dst        FRect
	animation  *Animation
	animations map[string]*Animation
}

func (s *Sprite) SetFPS(fps float64) {
	s.animation.SetFPS(fps)
}

func NewSprite(texture *Texture, x, y, width, height int) *Sprite {
	s := &Sprite{
		Texture: texture,
		dst: FRect{
			Width:  float64(texture.width),
			Height: float64(texture.height),
		},
		animations: make(map[string]*Animation),
	}

	s.Crop(x, y, width, height)

	return s
}

func (s *Sprite) Crop(x, y, width, height int) {
	s.src.X = x
	s.src.Y = y
	s.src.Width = width
	s.src.Height = height

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
	s.dst.X = x
	s.dst.Y = y

	if s.animation != nil {
		s.dst.Width = float64(s.animation.frame.src.Width)
		s.dst.Height = float64(s.animation.frame.src.Height)

		s.Texture.DrawRect(&s.animation.frame.src, &s.dst, s.animation.frame.flip)
	} else {
		s.dst.Width = float64(s.src.Width)
		s.dst.Height = float64(s.src.Height)

		s.Texture.DrawRect(&s.src, &s.dst, FlipNone)
	}
}
