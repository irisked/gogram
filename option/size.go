package option

import (
	"github.com/irisked/gogram/params"
)

// Sizes returns struct for setting optional height and width field.
func Sizes(height, width int) interface {
	sizeOptionSetter
} {
	return &sizeOption{height, width}
}

// Length returns struct for setting optional length field.
func Length(l int) interface {
	params.VideoNoteConfigOption
} {
	return &sizeOption{height: l}
}

// sizeOptionSetter is an interface for setting duration option.
type sizeOptionSetter interface {
	params.VideoConfigOption
	params.AnimationConfigOption
}

type sizeOption struct {
	height int
	width  int
}

func (so *sizeOption) ApplyVideoConfigOption(parameter *params.VideoConfig) {
	parameter.Height = so.height
	parameter.Width = so.width
}

func (so *sizeOption) ApplyAnimationConfigOption(parameter *params.AnimationConfig) {
	parameter.Height = so.height
	parameter.Width = so.width
}

func (so *sizeOption) ApplyVideoNoteConfigOption(parameter *params.VideoNoteConfig) {
	parameter.Length = so.height
}
