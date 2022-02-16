package input

import "github.com/veandco/go-sdl2/sdl"

func newMouseButtons() map[string]*Button {
	return map[string]*Button{
		"left":   {},
		"middle": {},
		"right":  {},
		"extra1": {},
		"extra2": {},
	}
}

var mouseMasks = map[string]uint32{
	"left":   sdl.ButtonLMask(),
	"middle": sdl.ButtonMMask(),
	"right":  sdl.ButtonRMask(),
	"extra1": sdl.ButtonX1Mask(),
	"extra2": sdl.ButtonX2Mask(),
}
