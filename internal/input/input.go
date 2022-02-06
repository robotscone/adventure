package input

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var unmappedButton = &Button{}

type Button struct {
	IsDown       bool
	IsPressed    bool
	IsReleased   bool
	PressedAt    time.Time
	ReleasedAt   time.Time
	DownDuration time.Duration
}

var keyboard struct {
	previous []*Button
	current  []*Button
}

// Registered devices.
var devices []*Device

// BindingMap is a map of action strings to input names that we can look up from SDL.
type BindingMap map[string]string

// ButtonMap is a map of action strings to actual button state.
type ButtonMap map[string]*Button

// Device represents an input device and must be created using NewDevice.
type Device struct {
	bindings BindingMap
	buttons  ButtonMap
}

func NewDevice(bindings BindingMap) *Device {
	if bindings == nil {
		bindings = make(BindingMap)
	}

	device := &Device{
		bindings: bindings,
		buttons:  make(ButtonMap),
	}

	devices = append(devices, device)

	return device
}

func (d *Device) Get(action string) *Button {
	if button, ok := d.buttons[action]; ok {
		return button
	}

	return unmappedButton
}

func Update() {
	now := time.Now()
	state := sdl.GetKeyboardState()

	// If the current keyboard state's length is less than SDL is reporting
	// then we just re-make the slice and populate with blank button states
	if len(keyboard.current) < len(state) {
		keyboard.current = make([]*Button, len(state))
		for i := range keyboard.current {
			keyboard.current[i] = &Button{}
		}

		keyboard.previous = make([]*Button, len(state))
		for i := range keyboard.previous {
			keyboard.previous[i] = &Button{}
		}
	}

	// Save the last keyboard state so we can do comparisons
	for i, button := range keyboard.current {
		*keyboard.previous[i] = *button
	}

	// Loop over the current keyboard state to update the current button values
	for i, value := range state {
		previous := keyboard.previous[i]
		button := keyboard.current[i]
		isDown := value != 0

		button.IsDown = isDown
		button.IsPressed = !previous.IsDown && isDown
		button.IsReleased = previous.IsDown && !isDown

		if button.IsPressed {
			button.PressedAt = now
		}

		if button.IsReleased {
			button.ReleasedAt = now
		}

		// If the button was pressed later than it was released then it
		// means the button is still being held down, so we should update the
		// down duration value
		if button.ReleasedAt.Before(button.PressedAt) {
			button.DownDuration = time.Since(button.PressedAt)
		}
	}

	// Loop over all registered devices and update button pointers based
	// on their internal binding maps
	for _, device := range devices {
		for action, button := range device.bindings {
			device.buttons[action] = keyboard.current[scancodes[button]]
		}
	}
}
