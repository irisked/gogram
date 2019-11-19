package option

import (
	"github.com/irisked/gogram/args"
)

// Duration returns struct for setting optional Duration field.
func Duration(time int) interface {
	durationOptionSetter
} {
	return &durationOption{time}
}

// durationOptionSetter is an interface for setting duration option.
type durationOptionSetter interface {
	args.AudioConfigOption
	args.VideoConfigOption
	args.AnimationConfigOption
	args.VoiceConfigOption
	args.VideoNoteConfigOption
}

type durationOption struct {
	time int
}

func (d *durationOption) ApplyAudioConfigOption(parameter *args.AudioConfig) {
	parameter.Duration = d.time
}

func (d *durationOption) ApplyVideoConfigOption(parameter *args.VideoConfig) {
	parameter.Duration = d.time
}

func (d *durationOption) ApplyAnimationConfigOption(parameter *args.AnimationConfig) {
	parameter.Duration = d.time
}

func (d *durationOption) ApplyVoiceConfigOption(parameter *args.VoiceConfig) {
	parameter.Duration = d.time
}

func (d *durationOption) ApplyVideoNoteConfigOption(parameter *args.VideoNoteConfig) {
	parameter.Duration = d.time
}
