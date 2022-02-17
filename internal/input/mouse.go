package input

import "github.com/veandco/go-sdl2/sdl"

type mouseButton struct {
	Button
	mask uint32
}

func newMouseButtons() map[string]*mouseButton {
	return map[string]*mouseButton{
		"left":   {mask: sdl.ButtonLMask()},
		"middle": {mask: sdl.ButtonMMask()},
		"right":  {mask: sdl.ButtonRMask()},
		"extra1": {mask: sdl.ButtonX1Mask()},
		"extra2": {mask: sdl.ButtonX2Mask()},
	}
}
