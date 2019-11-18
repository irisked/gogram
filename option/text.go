package option

import (
	"github.com/irisked/gogram/telegram/method"
)

// Text returns struct for setting optional Text field.
func Text(t string) interface {
	textOptionSetter
} {
	return &text{t}
}

// LastName returns struct for setting optional LastName field.
func LastName(t string) interface {
	method.SendContactOption
} {
	return &text{t}
}

// textOptionSetter returns struct for setting optional Text field.
type textOptionSetter interface {
	method.AnswerCallbackQueryOption
}

type text struct {
	text string
}

type textOption interface {
	ApplyTextOption(*text)
}

func (t *text) ApplyAnswerCallbackQueryOption(m *method.AnswerCallbackQuery) {
	m.Text = t.text
}

func (t *text) ApplySendContactOption(m *method.SendContact) {
	m.LastName = t.text
}
