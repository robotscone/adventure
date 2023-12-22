package gfx

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"os"
	"unsafe"

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

type Rect struct {
	X      int
	Y      int
	Width  int
	Height int
}

type FRect struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

type Renderer struct{ *sdl.Renderer }

func NewRenderer(window *sdl.Window) (*Renderer, error) {
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, err
	}

	renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)

	return &Renderer{Renderer: renderer}, nil
}

func (rn *Renderer) NewTexture(img image.Image, scaleQuality ScaleQuality) *Texture {
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, string(scaleQuality))

	bounds := img.Bounds()

	texture, err := rn.CreateTexture(uint32(sdl.PIXELFORMAT_RGBA32), sdl.TEXTUREACCESS_STATIC, int32(bounds.Max.X), int32(bounds.Max.Y))
	if err != nil {
		panic(err)
	}

	texture.SetBlendMode(sdl.BLENDMODE_BLEND)

	bpp := 4
	pitch := bounds.Max.X * bpp
	pixels := make([]byte, pitch*bounds.Max.Y)

	if gray, ok := img.(*image.Gray); ok {
		for y := 0; y < bounds.Max.Y; y++ {
			for x := 0; x < bounds.Max.X; x++ {
				offset := y*pitch + x*bpp

				pixels[offset+0] = 0xFF
				pixels[offset+1] = 0xFF
				pixels[offset+2] = 0xFF
				pixels[offset+3] = gray.GrayAt(x, y).Y
			}
		}
	} else {
		for y := 0; y < bounds.Max.Y; y++ {
			for x := 0; x < bounds.Max.X; x++ {
				r, g, b, a := img.At(x, y).RGBA()
				offset := y*pitch + x*bpp

				pixels[offset+0] = byte(r)
				pixels[offset+1] = byte(g)
				pixels[offset+2] = byte(b)
				pixels[offset+3] = byte(a)
			}
		}
	}

	texture.Update(nil, unsafe.Pointer(&pixels[0]), pitch)

	t := &Texture{
		renderer: rn,
		texture:  texture,
		width:    bounds.Max.X,
		height:   bounds.Max.Y,
	}

	return t
}

func (rn *Renderer) NewTextureFromFile(imagePath string, scaleQuality ScaleQuality) *Texture {
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

	t := rn.NewTexture(img, scaleQuality)

	textureCache[scaleQuality][imagePath] = t

	return t
}

func (rn *Renderer) NewSizedTexture(texture *sdl.Texture, width, height int) *Texture {
	return &Texture{
		renderer: rn,
		texture:  texture,
		width:    width,
		height:   height,
	}
}
