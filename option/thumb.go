package option

import (
	"fmt"
	"os"

	"github.com/irisked/gogram/params"
)

// Thumb returns struct for setting optional Thumb field.
func Thumb(file *os.File) interface {
	thumbOptionSetter
} {
	return &thumbOption{file: file}
}

func attachment(name string) string {
	return fmt.Sprintf("attachment://%s", name)
}

// thumbOptionSetter is an interface for setting cacheTime option.
type thumbOptionSetter interface {
	params.AudioConfigOption
	params.DocumentConfigOption
	params.VideoConfigOption
	params.AnimationConfigOption
	params.VideoNoteConfigOption
}

type thumbOption struct {
	file *os.File
}

func (to *thumbOption) ApplyAudioConfigOption(parameter *params.AudioConfig) {
	parameter.Thumb = attachment(to.file.Name())
	parameter.ThumbFile = to.file
}

func (to *thumbOption) ApplyDocumentConfigOption(parameter *params.DocumentConfig) {
	parameter.Thumb = attachment(to.file.Name())
	parameter.ThumbFile = to.file
}

func (to *thumbOption) ApplyVideoConfigOption(parameter *params.VideoConfig) {
	parameter.Thumb = attachment(to.file.Name())
	parameter.ThumbFile = to.file
}

func (to *thumbOption) ApplyAnimationConfigOption(parameter *params.AnimationConfig) {
	parameter.Thumb = attachment(to.file.Name())
	parameter.ThumbFile = to.file
}

func (to *thumbOption) ApplyVideoNoteConfigOption(parameter *params.VideoNoteConfig) {
	parameter.Thumb = attachment(to.file.Name())
	parameter.ThumbFile = to.file
}
