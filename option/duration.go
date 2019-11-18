package option

import (
	"github.com/irisked/gogram/params"
)

// Duration returns struct for setting optional Duration field.
func Duration(time int) interface {
	durationOptionSetter
} {
	return &durationOption{time}
}

// durationOptionSetter is an interface for setting duration option.
type durationOptionSetter interface {
	params.AudioConfigOption
	params.VideoConfigOption
	params.AnimationConfigOption
	params.VoiceConfigOption
	params.VideoNoteConfigOption
}

type durationOption struct {
	time int
}

func (d *durationOption) ApplyAudioConfigOption(parameter *params.AudioConfig) {
	parameter.Duration = d.time
}

func (d *durationOption) ApplyVideoConfigOption(parameter *params.VideoConfig) {
	parameter.Duration = d.time
}

func (d *durationOption) ApplyAnimationConfigOption(parameter *params.AnimationConfig) {
	parameter.Duration = d.time
}

func (d *durationOption) ApplyVoiceConfigOption(parameter *params.VoiceConfig) {
	parameter.Duration = d.time
}

func (d *durationOption) ApplyVideoNoteConfigOption(parameter *params.VideoNoteConfig) {
	parameter.Duration = d.time
}
