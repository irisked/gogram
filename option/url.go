package option

import (
	"github.com/irisked/gogram/telegram/method"
)

// URL returns struct for setting optional URL field.
func URL(url string) interface {
	urlOptionSetter
} {
	return &urlOption{url}
}

// urlOptionSetter is an interface for setting URL option.
type urlOptionSetter interface {
	method.AnswerCallbackQueryOption
}

type urlOption struct {
	url string
}

func (uo *urlOption) ApplyAnswerCallbackQueryOption(acq *method.AnswerCallbackQuery) {
	acq.URL = uo.url
}
