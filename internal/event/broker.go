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

	if want := 1; listenerType.NumIn() != want {
		panic(fmt.Sprintf("listener must have %v parameters, got %v", want, listenerType.NumIn()))
	}

	key := eventTypeName(listenerType.In(0))
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
	key := eventTypeName(reflect.TypeOf(event))
	callArgs := []reflect.Value{reflect.ValueOf(event)}

	for _, listenerFunc := range b.listeners[key] {
		listener := reflect.ValueOf(listenerFunc)
		listener.Call(callArgs)
	}
}

func eventTypeName(typ reflect.Type) string {
	var name string
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		name = "*"
	}

	return name + typ.PkgPath() + "." + typ.Name()
}
