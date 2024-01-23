package input

import "github.com/veandco/go-sdl2/sdl"

type controllerButton struct {
	Button
	code       int
	isAxis     bool
	isNegative bool
}

func (button *Button) setButtonDeadZone(deadZoneValue float64) {
	button.deadZone = deadZoneValue
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
		"trigger:left":   {code: sdl.CONTROLLER_AXIS_TRIGGERLEFT, isAxis: true, Button: Button{deadZone: 200}},
		"trigger:right":  {code: sdl.CONTROLLER_AXIS_TRIGGERRIGHT, isAxis: true, Button: Button{deadZone: 200}},
		"dpad:up":        {code: sdl.CONTROLLER_BUTTON_DPAD_UP},
		"dpad:down":      {code: sdl.CONTROLLER_BUTTON_DPAD_DOWN},
		"dpad:left":      {code: sdl.CONTROLLER_BUTTON_DPAD_LEFT},
		"dpad:right":     {code: sdl.CONTROLLER_BUTTON_DPAD_RIGHT},
		"lstick:left":    {code: sdl.CONTROLLER_AXIS_LEFTX, isAxis: true, isNegative: true, Button: Button{deadZone: 2000}},
		"lstick:right":   {code: sdl.CONTROLLER_AXIS_LEFTX, isAxis: true, Button: Button{deadZone: 2000}},
		"lstick:up":      {code: sdl.CONTROLLER_AXIS_LEFTY, isAxis: true, isNegative: true, Button: Button{deadZone: 2000}},
		"lstick:down":    {code: sdl.CONTROLLER_AXIS_LEFTY, isAxis: true, Button: Button{deadZone: 2000}},
		"rstick:left":    {code: sdl.CONTROLLER_AXIS_RIGHTX, isAxis: true, isNegative: true, Button: Button{deadZone: 2000}},
		"rstick:right":   {code: sdl.CONTROLLER_AXIS_RIGHTX, isAxis: true, Button: Button{deadZone: 2000}},
		"rstick:up":      {code: sdl.CONTROLLER_AXIS_RIGHTY, isAxis: true, isNegative: true, Button: Button{deadZone: 2000}},
		"rstick:down":    {code: sdl.CONTROLLER_AXIS_RIGHTY, isAxis: true, Button: Button{deadZone: 2000}},
	}
}
