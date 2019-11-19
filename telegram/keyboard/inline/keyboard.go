package inline

import (
	"encoding/json"
)

// Keyboard represents an inline keyboard that
// appears right next to the message it belongs to
type Keyboard struct {
	Buttons [][]Button `json:"inline_keyboard"`
}

// NewKeyboard creates Keyboard Keyboard
func NewKeyboard(buttons ...Button) *Keyboard {
	keyboard := new(Keyboard)
	keyboard.Buttons = [][]Button{}
	keyboard.Buttons = append(keyboard.Buttons, buttons)
	return keyboard
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
