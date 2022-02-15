package gfx

import (
	"bytes"
	"image"
	"log"
	"os"

	_ "image/png" // Register the png format

	"github.com/veandco/go-sdl2/sdl"
)

var textureCache = map[ScaleQuality]map[string]*Texture{
	ScaleNearest:     make(map[string]*Texture),
	ScaleLinear:      make(map[string]*Texture),
	ScaleAnisotropic: make(map[string]*Texture),
}

type ScaleQuality string

const (
	ScaleNearest     ScaleQuality = "nearest"
	ScaleLinear      ScaleQuality = "linear"
	ScaleAnisotropic ScaleQuality = "best"
)

type Renderer struct {
	*sdl.Renderer
	pixelFormat uint32
}

func NewRenderer(window *sdl.Window) (*Renderer, error) {
	pixelFormat, err := window.GetPixelFormat()
	if err != nil {
		return nil, err
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, err
	}

	r := &Renderer{
		Renderer:    renderer,
		pixelFormat: pixelFormat,
	}

	return r, nil
}

func (rn *Renderer) NewTexture(imagePath string, scaleQuality ScaleQuality) *Texture {
	if t := textureCache[scaleQuality][imagePath]; t != nil {
		return t
	}

	b, err := os.ReadFile(imagePath)
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}

	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, string(scaleQuality))

	bounds := img.Bounds()

	texture, err := rn.CreateTexture(rn.pixelFormat, sdl.TEXTUREACCESS_TARGET, int32(bounds.Max.X), int32(bounds.Max.Y))
	if err != nil {
		panic(err)
	}

	texture.SetBlendMode(sdl.BLENDMODE_BLEND)

	rn.SetRenderTarget(texture)
	rn.SetDrawColor(255, 255, 255, 0)
	rn.Clear()

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()

			rn.SetDrawColor(uint8(r), uint8(g), uint8(b), uint8(a))
			rn.DrawPoint(int32(x), int32(y))
		}
	}

	rn.SetRenderTarget(nil)

	t := &Texture{
		renderer: rn.Renderer,
		texture:  texture,
		width:    bounds.Max.X,
		height:   bounds.Max.Y,
	}

	textureCache[scaleQuality][imagePath] = t

	return t
}
