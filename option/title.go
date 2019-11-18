package option

import (
	"github.com/irisked/gogram/params"
)

// Title returns struct for setting optional Title field.
func Title(title string) interface {
	titleOptionSetter
} {
	return &titleOption{title}
}

// titleOptionSetter is an interface for setting title option.
type titleOptionSetter interface {
	params.AudioConfigOption
}

type titleOption struct {
	title string
}

func (to *titleOption) ApplyAudioConfigOption(parameter *params.AudioConfig) {
	parameter.Title = to.title
}
