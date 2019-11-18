package option

import (
	"github.com/irisked/gogram/telegram/method"
)

// Timeout returns struct for setting optional Timeout field.
func Timeout(o int) interface {
	timeoutOptionSetter
} {
	return &timeoutOption{o}
}

// UntilDate returns struct for setting optional UntilDate field.
func UntilDate(o int) interface {
	method.KickChatMemberOption
} {
	return &timeoutOption{o}
}

type timeoutOptionSetter interface {
	method.GetUpdatesOption
}

func (sa *timeoutOption) ApplyGetUpdatesOption(acq *method.GetUpdates) {
	acq.Timeout = sa.timeout
}

func (sa *timeoutOption) ApplyKickChatMemberOption(acq *method.KickChatMember) {
	acq.UntilDate = sa.timeout
}

type timeoutOption struct {
	timeout int
}
