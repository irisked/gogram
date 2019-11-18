package reply

import (
	"encoding/json"
)

// Reply represents represents a custom keyboard with reply options
type Reply struct {
	Buttons         [][]*Button `json:"keyboard"`
	ResizeKeyboard  bool        `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool        `json:"one_time_keyboard,omitempty"`
	Selective       bool        `json:"selective,omitempty"`
}

// NewKeyboard creates Reply
func NewKeyboard(options ...func(*Reply)) *Reply {
	keyboard := new(Reply)
	for _, option := range options {
		option(keyboard)
	}
	keyboard.Buttons = [][]*Button{}
	return keyboard
}

// Resize sets ResizeKeyboard to true.
func Resize() func(*Reply) {
	return func(repl *Reply) {
		repl.ResizeKeyboard = true
	}
}

// OneTime sets OneTimeKeyboard to true.
func OneTime() func(*Reply) {
	return func(repl *Reply) {
		repl.OneTimeKeyboard = true
	}
}

// Selective sets Selective to true.
func Selective() func(*Reply) {
	return func(repl *Reply) {
		repl.Selective = true
	}
}

// AddRow adds buttons[] to the keyboard
func (k *Reply) AddRow(buttons []*Button) {
	k.Buttons = append(k.Buttons, buttons)
}

// AddButton adds button to the row number
// if row > len(k.Reply) button will be added to last row
func (k *Reply) AddButton(row int, button *Button) {
	if row < len(k.Buttons) {
		k.Buttons[row] = append(k.Buttons[row], button)
	} else {
		k.Buttons[len(k.Buttons)-1] = append(k.Buttons[len(k.Buttons)-1], button)
	}
}

// Serialize makes json
func (k *Reply) Serialize() []byte {
	res, _ := json.Marshal(k)
	return res
}
