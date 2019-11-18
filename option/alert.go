package option

import (
	"github.com/irisked/gogram/telegram/method"
)

// ShowAlert returns struct for setting optional ShowAlert field.
func ShowAlert() interface {
	alertOptionSetter
} {
	return &showAlertOption{true}
}

// alertOptionSetter is an interface for setting showAlert option.
type alertOptionSetter interface {
	method.AnswerCallbackQueryOption
}

type showAlertOption struct {
	showAlert bool
}

func (sa *showAlertOption) ApplyAnswerCallbackQueryOption(acq *method.AnswerCallbackQuery) {
	acq.ShowAlert = sa.showAlert
}
