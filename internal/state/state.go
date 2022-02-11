package state

import "github.com/robotscone/adventure/internal/input"

type stackOp byte

const (
	stackNone stackOp = iota
	stackPush
	stackPop
)

type Controller interface {
	Switch(name string, data interface{})
	Push(name string, data interface{})
	Pop()
}

type State interface {
	Init()
	Enter(data interface{})
	Input(controller Controller, device *input.Device)
	Update(controller Controller, delta float64)
	Render()
	Exit()
}

type Base struct{}

func (*Base) Init()                                             {}
func (*Base) Enter(data interface{})                            {}
func (*Base) Input(controller Controller, device *input.Device) {}
func (*Base) Update(controller Controller, delta float64)       {}
func (*Base) Render()                                           {}
func (*Base) Exit()                                             {}
