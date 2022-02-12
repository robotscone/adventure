package state

import (
	"fmt"

	"github.com/robotscone/adventure/internal/input"
)

type group struct {
	state State
	data  interface{}
}

type FSM struct {
	next    State
	data    interface{}
	state   State
	stack   []*group
	stackOp stackOp
	states  map[string]State
}

func NewFSM() *FSM {
	return &FSM{
		state:  &Base{},
		states: make(map[string]State),
	}
}

func (f *FSM) RegisterState(name string, state State) {
	if _, ok := f.states[name]; ok {
		panic(fmt.Sprintf("duplicate state registration for %q", name))
	}

	state.Init(f)

	f.states[name] = state
}

func (f *FSM) Switch(name string, data interface{}) {
	state, ok := f.states[name]
	if !ok {
		fmt.Printf("attempted to set unknown state %q\n", name)

		return
	}

	f.next = state
	f.data = data
	f.stackOp = stackNone
}

func (f *FSM) Push(name string, data interface{}) {
	state, ok := f.states[name]
	if !ok {
		fmt.Printf("attempted to push unknown state %q\n", name)

		return
	}

	f.next = state
	f.data = data
	f.stackOp = stackPush
}

func (f *FSM) Pop() {
	if len(f.stack) == 0 {
		fmt.Println("attempted to pop an empty state stack")

		return
	}

	f.next = f.stack[len(f.stack)-1].state
	f.data = f.stack[len(f.stack)-1].data
	f.stackOp = stackPop
}

func (f *FSM) Init() {
	f.state.Init(f)
}

func (f *FSM) Input(device *input.Device) {
	f.state.Input(f, device)

	f.transition()
}

func (f *FSM) Update(delta float64) {
	f.state.Update(f, delta)

	f.transition()
}

func (f *FSM) Render() {
	f.state.Render()
}

func (f *FSM) transition() bool {
	if f.next == nil || f.next == f.state {
		f.next = nil
		f.data = nil
		f.stackOp = stackNone

		return false
	}

	f.state.Exit()

	switch f.stackOp {
	case stackNone:
		// Do nothing
	case stackPush:
		f.stack = append(f.stack, &group{
			state: f.state,
			data:  f.data,
		})
	case stackPop:
		if len(f.stack) > 0 {
			n := len(f.stack) - 1
			f.stack[n] = nil
			f.stack = f.stack[:n]
		}
	}

	f.state = f.next

	f.state.Enter(f.data)

	f.next = nil
	f.data = nil
	f.stackOp = stackNone

	return true
}
