package event

import (
	"fmt"
	"reflect"
)

type Event interface{}

type Listener interface{}

type Broker struct {
	listeners map[string][]Listener
	queue     []Event
}

func NewBroker() *Broker {
	return &Broker{listeners: make(map[string][]Listener)}
}

func (b *Broker) Listen(listener Listener) {
	listenerType := reflect.TypeOf(listener)
	if listenerType.Kind() != reflect.Func {
		panic("listener must be a function")
	}

	paramIfaces := []interface{}{(*Event)(nil)}
	paramCount := len(paramIfaces)

	// Ensure the listener has the correct number of parameters
	if listenerType.NumIn() != paramCount {
		panic(fmt.Sprintf("listener must have %v parameters, got %v", paramCount, listenerType.NumIn()))
	}

	// Manually type check the listener's parameters to ensure they implement required interfaces
	for i, iface := range paramIfaces {
		if !listenerType.In(i).Implements(reflect.TypeOf(iface).Elem()) {
			panic(fmt.Sprintf("listener parameter %v does not implement %T", i, iface))
		}
	}

	// The first parameter of a listener should be a struct
	if i := 0; listenerType.In(i).Kind() != reflect.Struct {
		panic(fmt.Sprintf("listener parameter %v must be a struct", i))
	}

	// We use the name of the first parameter's type as the event key
	// This also allows us to statically type event data at the same time
	eventType := listenerType.In(0)
	if eventType.Kind() != reflect.Struct {
		panic("listener parameter  must be a struct")
	}

	key := eventType.Name()
	b.listeners[key] = append(b.listeners[key], listener)
}

func (b *Broker) Dispatch(event Event) {
	b.fire(event)
}

func (b *Broker) Queue(event Event) {
	b.queue = append(b.queue, event)
}

func (b *Broker) Process() {
	for _, event := range b.queue {
		b.fire(event)
	}

	b.queue = b.queue[:0]
}

func (b *Broker) fire(event Event) {
	eventType := reflect.TypeOf(event)
	if eventType.Kind() != reflect.Struct {
		panic("event must be a struct")
	}

	key := eventType.Name()
	listenerFuncs, ok := b.listeners[key]
	if !ok {
		panic(fmt.Sprintf("unknown event %v", key))
	}

	callArgs := []reflect.Value{reflect.ValueOf(event)}

	for _, listenerFunc := range listenerFuncs {
		listener := reflect.ValueOf(listenerFunc)
		listener.Call(callArgs)
	}
}
