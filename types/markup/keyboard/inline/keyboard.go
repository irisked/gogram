package inline

import (
	"encoding/json"
)

// Inline represents an inline keyboard that
// appears right next to the message it belongs to
type Inline struct {
	Buttons [][]*Button `json:"inline_keyboard"`
}

// NewKeyboard creates Inline Keyboard
func NewKeyboard() *Inline {
	keyboard := new(Inline)
	keyboard.Buttons = [][]*Button{}
	return keyboard
}

// AddRow adds buttons[] to the keyboard
func (k *Inline) AddRow(buttons []*Button) {
	k.Buttons = append(k.Buttons, buttons)
}

// AddButton adds button to the row number
// if row > len(k.Buttons) button will be added to last row
func (k *Inline) AddButton(row int, button *Button) {
	if row < len(k.Buttons) {
		k.Buttons[row] = append(k.Buttons[row], button)
	} else {
		k.Buttons[len(k.Buttons)-1] = append(k.Buttons[len(k.Buttons)-1], button)
	}
}

// Serialize makes json
func (k *Inline) Serialize() []byte {
	res, _ := json.Marshal(k)
	return res
}
