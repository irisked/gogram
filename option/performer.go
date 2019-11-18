package option

import (
	"github.com/irisked/gogram/params"
)

// Performer returns struct for setting optional Performer field.
func Performer(performer string) interface {
	performerOptionSetter
} {
	return &performerOption{performer}
}

// performerOptionSetter is an interface for setting performer option.
type performerOptionSetter interface {
	params.AudioConfigOption
}

type performerOption struct {
	performer string
}

func (p *performerOption) ApplyAudioConfigOption(parameter *params.AudioConfig) {
	parameter.Performer = p.performer
}
