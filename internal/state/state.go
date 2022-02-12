package state

import "github.com/robotscone/adventure/internal/input"

type Controller interface {
	Switch(name string, data interface{})
	Push(name string, data interface{})
	Pop()
}

type State interface {
	Init(controller Controller)
	Enter(controller Controller, data interface{})
	Resume(controller Controller)
	Input(controller Controller, device *input.Device)
	Update(controller Controller, delta float64)
	Render()
	Exit()
}

type Base struct{}

func (*Base) Init(controller Controller)                        {}
func (*Base) Enter(controller Controller, data interface{})     {}
func (*Base) Resume(controller Controller)                      {}
func (*Base) Input(controller Controller, device *input.Device) {}
func (*Base) Update(controller Controller, delta float64)       {}
func (*Base) Render()                                           {}
func (*Base) Exit()                                             {}
