package gfx

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type Texture struct {
	renderer *Renderer
	texture  *sdl.Texture
	width    int
	height   int
}

func (t *Texture) Renderer() *Renderer {
	return t.renderer
}

func (t *Texture) Width() int {
	return t.width
}

func (t *Texture) Height() int {
	return t.height
}

func (t *Texture) SetAlphaMod(a float64) {
	t.texture.SetAlphaMod(uint8(math.MaxUint8 * a))
}

func (t *Texture) SetColorMod(r, g, b float64) {
	t.texture.SetColorMod(uint8(math.MaxUint8*r), uint8(math.MaxUint8*g), uint8(math.MaxUint8*b))
}

func (t *Texture) DrawRect(src *Rect, dst *FRect, flip Flip) {
	t.Draw(src.X, src.Y, src.Width, src.Height, dst.X, dst.Y, dst.Width, dst.Height, flip)
}

func (t *Texture) Draw(srcX, srcY, srcWidth, srcHeight int, dstX, dstY, dstWidth, dstHeight float64, flip Flip) {
	src := sdl.Rect{
		X: int32(srcX),
		Y: int32(srcY),
		W: int32(srcWidth),
		H: int32(srcHeight),
	}

	dst := sdl.FRect{
		X: float32(dstX),
		Y: float32(dstY),
		W: float32(dstWidth),
		H: float32(dstHeight),
	}

	t.draw(&src, &dst, flip)
}

func (t *Texture) DrawAt(dstX, dstY float64, flip Flip) {
	dst := sdl.FRect{
		X: float32(dstX),
		Y: float32(dstY),
		W: float32(t.width),
		H: float32(t.height),
	}

	t.draw(nil, &dst, flip)
}

func (t *Texture) DrawStretchedAt(dstX, dstY, dstWidth, dstHeight float64, flip Flip) {
	dst := sdl.FRect{
		X: float32(dstX),
		Y: float32(dstY),
		W: float32(dstWidth),
		H: float32(dstHeight),
	}

	t.draw(nil, &dst, flip)
}

func (t *Texture) draw(src *sdl.Rect, dst *sdl.FRect, flip Flip) {
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
	dst.X += 0.01
	dst.Y += 0.01

	t.renderer.CopyExF(t.texture, src, dst, 0.0, nil, rendererFlip)
}

func (t *Texture) Destroy() {
	t.texture.Destroy()
}
