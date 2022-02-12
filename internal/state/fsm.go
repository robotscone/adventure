package state

import (
	"fmt"

	"github.com/robotscone/adventure/internal/input"
)

var base = &Base{}

type FSM struct {
	stack  []State
	states map[string]State
}

func NewFSM() *FSM {
	return &FSM{
		stack:  []State{base},
		states: make(map[string]State),
	}
}

func (f *FSM) RegisterState(name string, state State) {
	if _, ok := f.states[name]; ok {
		panic(fmt.Sprintf("duplicate state registration for %q", name))
	}

	f.states[name] = state

	state.Init(f)
}

func (f *FSM) Switch(name string, data interface{}) {
	state, ok := f.states[name]
	if !ok {
		fmt.Printf("attempted to switch to unknown state %q\n", name)

		return
	}

	n := len(f.stack) - 1
	top := f.stack[n]

	top.Exit()

	f.stack[n] = state

	state.Enter(f, data)
}

func (f *FSM) Push(name string, data interface{}) {
	state, ok := f.states[name]
	if !ok {
		fmt.Printf("attempted to push unknown state %q\n", name)

		return
	}

	f.stack = append(f.stack, state)

	state.Enter(f, data)
}

func (f *FSM) Pop() {
	if len(f.stack) == 0 {
		fmt.Println("attempted to pop an empty state stack")

		return
	}

	n := len(f.stack) - 1
	top := f.stack[n]

	top.Exit()

	f.stack[n] = nil
	f.stack = f.stack[:n]

	f.stack[len(f.stack)-1].Resume(f)
}

func (f *FSM) Input(device *input.Device) {
	f.stack[len(f.stack)-1].Input(f, device)
}

func (f *FSM) Update(delta float64) {
	f.stack[len(f.stack)-1].Update(f, delta)
}

func (f *FSM) Render() {
	for i := 0; i < len(f.stack); i++ {
		f.stack[i].Render()
	}
}
