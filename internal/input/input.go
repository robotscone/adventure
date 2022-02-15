package input

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var unmappedButton = &Button{}

type Button struct {
	Value        float64
	IsDown       bool
	IsPressed    bool
	IsReleased   bool
	PressedAt    time.Time
	ReleasedAt   time.Time
	DownDuration time.Duration
}

// BindingMap is a map of action strings to input names that we can look up from SDL.
type BindingMap map[string][]string

// ButtonMap is a map of action strings to actual button state.
type ButtonMap map[string]*Button

// Device represents an input device and must be created using NewDevice.
type Device struct {
	bindings BindingMap
	previous ButtonMap
	current  ButtonMap
}

var keyboard struct {
	current  []Button
	previous []Button
}

// Registered devices.
var devices []*Device

func NewDevice(bindings BindingMap) *Device {
	if bindings == nil {
		bindings = make(BindingMap)
	}

	device := &Device{
		bindings: bindings,
		previous: make(ButtonMap),
		current:  make(ButtonMap),
	}

	for action := range bindings {
		device.previous[action] = &Button{}
		device.current[action] = &Button{}
	}

	devices = append(devices, device)

	return device
}

func (d *Device) Get(action string) *Button {
	if button := d.current[action]; button != nil {
		return button
	}

	return unmappedButton
}

func setButtonState(current, previous *Button, value float64, now time.Time) {
	current.Value = value
	current.IsDown = current.Value != 0
	current.IsPressed = !previous.IsDown && current.IsDown
	current.IsReleased = previous.IsDown && !current.IsDown

	if current.IsPressed {
		current.PressedAt = now
	}

	if current.IsReleased {
		current.ReleasedAt = now
	}

	// If the button was pressed later than it was released then it
	// means the button is still being held down, so we should update the
	// down duration value
	if current.ReleasedAt.Before(current.PressedAt) {
		current.DownDuration = time.Since(current.PressedAt)
	}
}

func Update() {
	now := time.Now()
	keyboardState := sdl.GetKeyboardState()

	// If the current keyboard state's length is less than SDL is reporting
	// then we just re-make the slice and populate with blank button states
	if len(keyboard.current) < len(keyboardState) {
		keyboard.previous = make([]Button, len(keyboardState))
		keyboard.current = make([]Button, len(keyboardState))
	}

	// Save the last keyboard state so we can do comparisons
	for i, button := range keyboard.current {
		keyboard.previous[i] = button
	}

	// Loop over the current keyboard state to update the current button values
	for i, value := range keyboardState {
		setButtonState(&keyboard.current[i], &keyboard.previous[i], float64(value), now)
	}

	// Loop over all registered devices and update button pointers based
	// on their internal binding maps
	for _, device := range devices {
		for action, buttons := range device.bindings {
			// Save the last device state so we can do comparisons
			*device.previous[action] = *device.current[action]

			previous := device.previous[action]
			current := device.current[action]

			current.Value = 0

			for _, name := range buttons {
				code, ok := scancodes[name]
				if !ok {
					continue
				}

				if button := &keyboard.current[code]; button.Value != 0 {
					current.Value = button.Value
				}
			}

			setButtonState(current, previous, current.Value, now)
		}
	}
}
