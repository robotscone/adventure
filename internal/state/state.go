package state

import "github.com/robotscone/adventure/internal/input"

type Controller interface {
	Switch(name string, message any)
	Push(name string, message any)
	Pop()
}

type Data struct {
	Device *input.Device
	Delta  float64
}

type State interface {
	Init(controller Controller, data *Data)
	Enter(controller Controller, data *Data, message any)
	Resume(controller Controller, data *Data)
	Input(controller Controller, data *Data)
	Update(controller Controller, data *Data)
	Draw()
	Pause()
	Exit()
}

type Base struct{}

func (*Base) Init(controller Controller, data *Data)               {}
func (*Base) Enter(controller Controller, data *Data, message any) {}
func (*Base) Resume(controller Controller, data *Data)             {}
func (*Base) Input(controller Controller, data *Data)              {}
func (*Base) Update(controller Controller, data *Data)             {}
func (*Base) Draw()                                                {}
func (*Base) Pause()                                               {}
func (*Base) Exit()                                                {}
