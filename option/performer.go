package option

import (
	"github.com/irisked/gogram/args"
)

// Performer returns struct for setting optional Performer field.
func Performer(performer string) interface {
	performerOptionSetter
} {
	return &performerOption{performer}
}

// performerOptionSetter is an interface for setting performer option.
type performerOptionSetter interface {
	args.AudioConfigOption
}

type performerOption struct {
	performer string
}

func (p *performerOption) ApplyAudioConfigOption(parameter *args.AudioConfig) {
	parameter.Performer = p.performer
}
