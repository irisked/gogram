package option

import (
	"gogram/telegram/method"
)

// CacheTime returns struct for setting optional CacheTime field.
func CacheTime(time int) interface {
	cacheTimeOptionSetter
} {
	return &cacheTimeOption{time}
}

// cacheTimeOptionSetter is an interface for setting cacheTime option.
type cacheTimeOptionSetter interface {
	method.AnswerCallbackQueryOption
}

type cacheTimeOption struct {
	time int
}

func (ct *cacheTimeOption) ApplyAnswerCallbackQueryOption(acq *method.AnswerCallbackQuery) {
	acq.CacheTime = ct.time
}
