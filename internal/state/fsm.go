package state

import (
	"fmt"
)

var base = &Base{}

type FSM struct {
	data   *Data
	stack  []State
	states map[string]State
}

func NewFSM(data *Data) *FSM {
	fsm := &FSM{
		data:   data,
		stack:  []State{base},
		states: make(map[string]State),
	}

	return fsm
}

func (f *FSM) RegisterState(name string, state State) {
	if _, ok := f.states[name]; ok {
		panic(fmt.Sprintf("duplicate state registration for %q", name))
	}

	f.states[name] = state

	state.Init(f, f.data)
}

func (f *FSM) Switch(name string, message interface{}) {
	state, ok := f.states[name]
	if !ok {
		fmt.Printf("attempted to switch to unknown state %q\n", name)

		return
	}

	n := len(f.stack) - 1

	f.stack[n].Exit()

	f.stack[n] = state

	state.Enter(f, f.data, message)
}

func (f *FSM) Push(name string, message interface{}) {
	state, ok := f.states[name]
	if !ok {
		fmt.Printf("attempted to push unknown state %q\n", name)

		return
	}

	f.stack[len(f.stack)-1].Pause()

	f.stack = append(f.stack, state)

	state.Enter(f, f.data, message)
}

func (f *FSM) Pop() {
	if len(f.stack) == 0 {
		fmt.Println("attempted to pop an empty state stack")

		return
	}

	n := len(f.stack) - 1

	f.stack[n].Exit()

	f.stack[n] = nil
	f.stack = f.stack[:n]

	f.stack[len(f.stack)-1].Resume(f, f.data)
}

func (f *FSM) Input() {
	f.stack[len(f.stack)-1].Input(f, f.data)
}

func (f *FSM) Update() {
	f.stack[len(f.stack)-1].Update(f, f.data)
}

func (f *FSM) Render() {
	for i := 0; i < len(f.stack); i++ {
		f.stack[i].Render()
	}
}
