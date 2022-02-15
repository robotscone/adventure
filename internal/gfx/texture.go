package gfx

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type Texture struct {
	renderer *sdl.Renderer
	texture  *sdl.Texture
	width    int
	height   int
}

func (t *Texture) SetAlphaMod(a float64) {
	t.texture.SetAlphaMod(uint8(math.MaxUint8 * a))
}

func (t *Texture) SetColorMod(r, g, b float64) {
	t.texture.SetColorMod(uint8(math.MaxUint8*r), uint8(math.MaxUint8*g), uint8(math.MaxUint8*b))
}

func (t *Texture) Draw(src *sdl.Rect, dst *sdl.FRect, flip Flip) {
	var rendererFlip sdl.RendererFlip
	switch flip {
	case FlipHorizontal:
		rendererFlip = sdl.FLIP_HORIZONTAL
	case FlipVertical:
		rendererFlip = sdl.FLIP_VERTICAL
	default:
		rendererFlip = sdl.FLIP_NONE
	}

	t.renderer.CopyExF(t.texture, src, dst, 0.0, nil, rendererFlip)
}

func (t *Texture) Destroy() {
	t.texture.Destroy()
}
