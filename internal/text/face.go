package text

import (
	"image"
	"math"
	"os"
	"strings"

	"github.com/robotscone/adventure/internal/gfx"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
)

type Glyph struct {
	gfx.Rect
	Advance      fixed.Int26_6
	BearingLeft  fixed.Int26_6
	BearingRight fixed.Int26_6
	Ascent       fixed.Int26_6
	Descent      fixed.Int26_6
}

type GlyphRect struct {
	Src *Glyph
	Dst Glyph
}

type Text struct {
	gfx.FRect
	Glyphs      []GlyphRect
	CursorY     int
	GlyphHeight int
	LineHeight  int
}

type Atlas struct {
	*gfx.Texture
	Width      int
	Height     int
	CellWidth  int
	CellHeight int
}

type Face struct {
	*Atlas
	ptSize  float64
	dpi     float64
	font    *sfnt.Font
	face    font.Face
	metrics font.Metrics
	glyphs  map[rune]*Glyph
	empty   *Glyph
	missing *Glyph
}

func NewFace(fontPath string, ptSize, dpi float64, renderer *gfx.Renderer, scaleQuality gfx.ScaleQuality, charset string) (*Face, error) {
	b, err := os.ReadFile(fontPath)
	if err != nil {
		return nil, err
	}

	collection, err := opentype.ParseCollection(b)
	if err != nil {
		return nil, err
	}

	tt, err := collection.Font(0)
	if err != nil {
		return nil, err
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    ptSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, err
	}

	fc := &Face{
		ptSize:  ptSize,
		dpi:     dpi,
		font:    tt,
		face:    face,
		metrics: face.Metrics(),
	}

	fc.generateAtlas(renderer, scaleQuality, charset)

	return fc, nil
}

func (fc *Face) generateAtlas(renderer *gfx.Renderer, scaleQuality gfx.ScaleQuality, charset string) {
	missing := '□'
	charset = strings.NewReplacer("\n", "", "\r", "", "\t", "").Replace(charset) + string(missing)

	runes := []rune(charset)
	cells := int(math.Round(math.Sqrt(float64(len(runes))) + 0.5))
	glyphs := make(map[rune]*Glyph, cells)
	maxCellWidth, maxCellHeight := 0, 0

	for _, r := range runes {
		bounds, advance, ok := fc.face.GlyphBounds(r)
		if !ok {
			bounds, advance, _ = fc.face.GlyphBounds(missing)
		}

		glyph := &Glyph{
			Rect: gfx.Rect{
				Width:  (bounds.Max.X - bounds.Min.X).Ceil(),
				Height: (fc.metrics.Ascent + fc.metrics.Descent).Ceil(),
			},
			Advance:      advance,
			BearingLeft:  bounds.Min.X,
			BearingRight: advance - bounds.Max.X,
			Ascent:       -bounds.Min.Y,
			Descent:      bounds.Min.Y,
		}

		if glyph.Width > maxCellWidth {
			maxCellWidth = glyph.Width + 1
		}

		if glyph.Height > maxCellHeight {
			maxCellHeight = glyph.Height + 1
		}

		glyphs[r] = glyph
	}

	atlasWidth := maxCellWidth * cells
	atlasHeight := maxCellHeight * cells
	img := image.NewGray(image.Rect(0, 0, atlasWidth, atlasHeight))
	drawer := font.Drawer{
		Dst:  img,
		Src:  image.White,
		Face: fc.face,
	}

atlasLoop:
	for y := 0; y < cells; y++ {
		for x := 0; x < cells; x++ {
			i := cells*y + x

			if i >= len(runes) {
				break atlasLoop
			}

			r := runes[i]
			glyph := glyphs[r]

			glyph.X = maxCellWidth * x
			glyph.Y = maxCellHeight * y

			drawer.Dot = fixed.Point26_6{
				X: fixed.I(glyph.X) - glyph.BearingLeft,
				Y: fixed.I(glyph.Y) + fc.metrics.Ascent,
			}

			drawer.DrawString(string(r))
		}
	}

	fc.glyphs = glyphs
	fc.empty = &Glyph{}
	fc.missing = fc.glyphs[missing]
	fc.Atlas = &Atlas{
		Texture:    renderer.NewTexture(img, scaleQuality),
		Width:      atlasWidth,
		Height:     atlasHeight,
		CellWidth:  maxCellWidth - 1,
		CellHeight: maxCellHeight - 1,
	}
}

func (fc *Face) DrawGlyphs(runes []rune) Text {
	text := Text{
		GlyphHeight: (fc.metrics.Ascent + fc.metrics.Descent).Ceil(),
		LineHeight:  fc.metrics.Height.Ceil(),
	}

	if len(runes) == 0 {
		return text
	}

	text.Glyphs = make([]GlyphRect, 0, len(runes))

	// The scale is based on units per em to get as close as possible to a
	// point size that would represent 100% for the given font
	// The multiplication by 64 is required before converting float64 to Int26_6
	scale := fc.dpi / float64(fc.font.UnitsPerEm())
	kernScale := fixed.Int26_6(fc.ptSize * scale * 64)

	var textWidth fixed.Int26_6

	prev := rune(-1)
	startX, startY := fixed.I(0), fixed.I(0)
	x, y := startX, startY

	for _, r := range runes {
		glyph := fc.glyphs[r]
		if glyph == nil {
			glyph = fc.missing
		}

		var bearingLeft fixed.Int26_6
		if prev >= 0 {
			x += fc.face.Kern(prev, r).Mul(kernScale)
			bearingLeft = glyph.BearingLeft
		}
		dstX := x + bearingLeft

		dst := GlyphRect{
			Src: glyph,
			Dst: Glyph{
				Rect: gfx.Rect{
					X:      dstX.Ceil(),
					Y:      y.Ceil(),
					Width:  glyph.Width,
					Height: glyph.Height,
				},
			},
		}

		switch r {
		case ' ':
			dst.Dst.Width = glyph.Advance.Ceil()
		case '\n':
			dst.Src = fc.empty
			dst.Dst.Width = int(float64(text.GlyphHeight) / 2.5) // 2.5 is arbitrary
			dst.Dst.Height = text.GlyphHeight
		}

		text.Glyphs = append(text.Glyphs, dst)

		if r != '\n' {
			x += glyph.Advance
			if prev < 0 {
				x -= glyph.BearingLeft
			}
		}

		lineWidth := (x - startX) - glyph.BearingRight
		if lineWidth > textWidth {
			textWidth = lineWidth
		}

		prev = r

		if r == '\n' {
			x = startX
			y += fc.metrics.Height
			prev = -1
		}
	}

	dst := GlyphRect{
		Src: fc.empty,
		Dst: Glyph{
			Rect: gfx.Rect{
				X:      x.Ceil(),
				Y:      y.Ceil(),
				Width:  int(float64(text.GlyphHeight) / 2.5), // 2.5 is arbitrary
				Height: text.GlyphHeight,
			},
		},
	}

	text.Glyphs = append(text.Glyphs, dst)

	text.Width = float64(textWidth.Ceil())
	text.Height = float64(((y - startY) + fc.metrics.Ascent + fc.metrics.Descent).Ceil())

	return text
}

func (fc *Face) DrawText(x, y, scale float64, runes []rune) Text {
	text := fc.DrawGlyphs(runes)

	for _, g := range text.Glyphs {
		dstX := x + float64(g.Dst.X)*scale
		dstY := y + float64(g.Dst.Y)*scale
		dstWidth := float64(g.Dst.Width) * scale
		dstHeight := float64(g.Dst.Height) * scale

		fc.Atlas.Draw(g.Src.X, g.Src.Y, g.Src.Width, g.Src.Height, dstX, dstY, dstWidth, dstHeight, gfx.FlipNone)
	}

	text.Width *= scale
	text.Height *= scale

	return text
}
