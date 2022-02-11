package gfx

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type ScaleQuality string

const (
	ScaleNearest     ScaleQuality = "nearest"
	ScaleLinear      ScaleQuality = "linear"
	ScaleAnisotropic ScaleQuality = "best"
)

var textures = map[ScaleQuality]map[string]*Texture{
	ScaleNearest:     make(map[string]*Texture),
	ScaleLinear:      make(map[string]*Texture),
	ScaleAnisotropic: make(map[string]*Texture),
}

type Texture struct {
	renderer *sdl.Renderer
	texture  *sdl.Texture
	width    int
	height   int
}

func NewTexture(renderer *sdl.Renderer, scaleQuality ScaleQuality, path string) *Texture {
	if t, ok := textures[scaleQuality][path]; ok {
		return t
	}

	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, string(scaleQuality))

	image, err := img.Load(path)
	if err != nil {
		panic(err)
	}
	defer image.Free()

	texture, err := renderer.CreateTextureFromSurface(image)
	if err != nil {
		panic(err)
	}

	texture.SetBlendMode(sdl.BLENDMODE_BLEND)

	t := &Texture{
		renderer: renderer,
		texture:  texture,
		width:    int(image.W),
		height:   int(image.H),
	}

	textures[scaleQuality][path] = t

	return t
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
