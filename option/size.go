package option

import (
	"github.com/irisked/gogram/args"
)

// Sizes returns struct for setting optional height and width field.
func Sizes(height, width int) interface {
	sizeOptionSetter
} {
	return &sizeOption{height, width}
}

// Length returns struct for setting optional length field.
func Length(l int) interface {
	args.VideoNoteConfigOption
} {
	return &sizeOption{height: l}
}

// sizeOptionSetter is an interface for setting duration option.
type sizeOptionSetter interface {
	args.VideoConfigOption
	args.AnimationConfigOption
}

type sizeOption struct {
	height int
	width  int
}

func (so *sizeOption) ApplyVideoConfigOption(parameter *args.VideoConfig) {
	parameter.Height = so.height
	parameter.Width = so.width
}

func (so *sizeOption) ApplyAnimationConfigOption(parameter *args.AnimationConfig) {
	parameter.Height = so.height
	parameter.Width = so.width
}

func (so *sizeOption) ApplyVideoNoteConfigOption(parameter *args.VideoNoteConfig) {
	parameter.Length = so.height
}
