package reply

import (
	"encoding/json"
)

// Keyboard represents represents a custom keyboard with reply options
type Keyboard struct {
	Buttons         [][]Button `json:"keyboard"`
	ResizeKeyboard  bool        `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool        `json:"one_time_keyboard,omitempty"`
	Selective       bool        `json:"selective,omitempty"`
}

// NewKeyboard creates Keyboard
func NewKeyboard(options ...func(*Keyboard)) *Keyboard {
	keyboard := new(Keyboard)
	for _, option := range options {
		option(keyboard)
	}
	keyboard.Buttons = make([][]Button, 0)
	for i := range keyboard.Buttons {
    keyboard.Buttons[i] = make([]Button, 0)
}
	return keyboard
}

// Resize sets ResizeKeyboard to true.
func Resize() func(*Keyboard) {
	return func(repl *Keyboard) {
		repl.ResizeKeyboard = true
	}
}

// OneTime sets OneTimeKeyboard to true.
func OneTime() func(*Keyboard) {
	return func(repl *Keyboard) {
		repl.OneTimeKeyboard = true
	}
}

// Selective sets Selective to true.
func Selective() func(*Keyboard) {
	return func(repl *Keyboard) {
		repl.Selective = true
	}
}

// AddRow adds buttons[] to the keyboard
func (k *Keyboard) AddRow(buttons ...Button) {
	k.Buttons = append(k.Buttons, buttons)
}

// Serialize makes json
func (k *Keyboard) Serialize() []byte {
	res, _ := json.Marshal(k)
	return res
}
