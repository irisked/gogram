package option

import (
	"fmt"
	"os"

	"github.com/irisked/gogram/args"
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
	args.AudioConfigOption
	args.DocumentConfigOption
	args.VideoConfigOption
	args.AnimationConfigOption
	args.VideoNoteConfigOption
}

type thumbOption struct {
	file *os.File
}

func (to *thumbOption) ApplyAudioConfigOption(parameter *args.AudioConfig) {
	parameter.Thumb = attachment(to.file.Name())
	parameter.ThumbFile = to.file
}

func (to *thumbOption) ApplyDocumentConfigOption(parameter *args.DocumentConfig) {
	parameter.Thumb = attachment(to.file.Name())
	parameter.ThumbFile = to.file
}

func (to *thumbOption) ApplyVideoConfigOption(parameter *args.VideoConfig) {
	parameter.Thumb = attachment(to.file.Name())
	parameter.ThumbFile = to.file
}

func (to *thumbOption) ApplyAnimationConfigOption(parameter *args.AnimationConfig) {
	parameter.Thumb = attachment(to.file.Name())
	parameter.ThumbFile = to.file
}

func (to *thumbOption) ApplyVideoNoteConfigOption(parameter *args.VideoNoteConfig) {
	parameter.Thumb = attachment(to.file.Name())
	parameter.ThumbFile = to.file
}
