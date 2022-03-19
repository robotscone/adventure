package input

import "github.com/veandco/go-sdl2/sdl"

type controllerButton struct {
	Button
	code       int
	isAxis     bool
	isNegative bool
}

func newControllerButtons() map[string]*controllerButton {
	return map[string]*controllerButton{
		"a":              {code: sdl.CONTROLLER_BUTTON_A},
		"b":              {code: sdl.CONTROLLER_BUTTON_B},
		"x":              {code: sdl.CONTROLLER_BUTTON_X},
		"y":              {code: sdl.CONTROLLER_BUTTON_Y},
		"back":           {code: sdl.CONTROLLER_BUTTON_BACK},
		"guide":          {code: sdl.CONTROLLER_BUTTON_GUIDE},
		"start":          {code: sdl.CONTROLLER_BUTTON_START},
		"stick:left":     {code: sdl.CONTROLLER_BUTTON_LEFTSTICK},
		"stick:right":    {code: sdl.CONTROLLER_BUTTON_RIGHTSTICK},
		"shoulder:left":  {code: sdl.CONTROLLER_BUTTON_LEFTSHOULDER},
		"shoulder:right": {code: sdl.CONTROLLER_BUTTON_RIGHTSHOULDER},
		"dpad:up":        {code: sdl.CONTROLLER_BUTTON_DPAD_UP},
		"dpad:down":      {code: sdl.CONTROLLER_BUTTON_DPAD_DOWN},
		"dpad:left":      {code: sdl.CONTROLLER_BUTTON_DPAD_LEFT},
		"dpad:right":     {code: sdl.CONTROLLER_BUTTON_DPAD_RIGHT},
		"lstick:left":    {code: sdl.CONTROLLER_AXIS_LEFTX, isAxis: true, isNegative: true},
		"lstick:right":   {code: sdl.CONTROLLER_AXIS_LEFTX, isAxis: true},
		"joystick:lefty": {code: sdl.CONTROLLER_AXIS_LEFTY, isAxis: true},
	}
}
