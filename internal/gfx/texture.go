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
	dst      sdl.FRect
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

	// The value added to the destination X and Y here is a hack to try and
	// prevent texture bleeding when low logical renderer sizes are stretched
	// to fit high resolution window sizes
	//
	// If these values don't solve the problem for certain resolutions then they
	// can be changed, but the higher they are the more noticeable the offset
	// of the final render will become
	// For example, if they're set to 0.5 then it's obvious that there's half
	// a pixel offset of the final render from the top and left sides because
	// you can see a small black line
	//
	// The more technically correct way to fix this issue would be to create
	// a render target texture that's the same size as the logical size, render
	// everything into that texture, then remove the render target and copy
	// the entire render texture to the window
	// The problem with that approach is that movement of sprites in the world
	// becomes noticeably jittery/jagged, which is why we're taking this
	// approach instead
	t.dst.X = dst.X + 0.01
	t.dst.Y = dst.Y + 0.01
	t.dst.W = dst.W
	t.dst.H = dst.H

	t.renderer.CopyExF(t.texture, src, &t.dst, 0.0, nil, rendererFlip)
}

func (t *Texture) Destroy() {
	t.texture.Destroy()
}
