package option

import (
	"gogram/telegram/method"
)

// DisableEditMessage returns struct for setting optional DisableEditMessage field.
func DisableEditMessage() interface {
	disableEditMessageSetter
} {
	return &disableEditMessageOption{true}
}

// Force returns struct for setting optional Force field.
func Force() interface {
	forceOptionSetter
} {
	return &forceOption{true}
}

// disableEditMessageSetter is an interface for setting disableEditMesage option.
type disableEditMessageSetter interface {
	method.SetGameScoresOption
}

type disableEditMessageOption struct {
	disable bool
}

func (d *disableEditMessageOption) ApplySetGameScoresOption(msg *method.SetGameScores) {
	msg.DisableEditMessage = d.disable
}

// forceOptionSetter is an interface for setting force option.
type forceOptionSetter interface {
	method.SetGameScoresOption
}

type forceOption struct {
	force bool
}

func (d *forceOption) ApplySetGameScoresOption(msg *method.SetGameScores) {
	msg.Force = d.force
}
