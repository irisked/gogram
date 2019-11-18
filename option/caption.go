package option

import (
	"github.com/irisked/gogram/params"
)

// SimpleCaption simple text.
func SimpleCaption(text string) interface{ captionSetter } {
	return &caption{text: text}
}

// MarkdownCaption markdown text.
func MarkdownCaption(text string) interface{ captionSetter } {
	return &caption{text: text, parseMode: "markdown"}
}

// HTMLCaption html text.
func HTMLCaption(text string) interface{ captionSetter } {
	return &caption{text: text, parseMode: "HTML"}
}

// captionSetter is an interface for setting caption option.
type captionSetter interface {
	params.PhotoConfigOption
	params.AudioConfigOption
	params.DocumentConfigOption
	params.VideoConfigOption
	params.AnimationConfigOption
	params.VoiceConfigOption
}

type caption struct {
	text      string
	parseMode string
}

func (c *caption) ApplyPhotoConfigOption(parameter *params.PhotoConfig) {
	parameter.Caption = c.text
	parameter.ParseMode = c.parseMode
}

func (c *caption) ApplyAudioConfigOption(parameter *params.AudioConfig) {
	parameter.Caption = c.text
	parameter.ParseMode = c.parseMode
}

func (c *caption) ApplyDocumentConfigOption(parameter *params.DocumentConfig) {
	parameter.Caption = c.text
	parameter.ParseMode = c.parseMode
}

func (c *caption) ApplyVideoConfigOption(parameter *params.VideoConfig) {
	parameter.Caption = c.text
	parameter.ParseMode = c.parseMode
}

func (c *caption) ApplyAnimationConfigOption(parameter *params.AnimationConfig) {
	parameter.Caption = c.text
	parameter.ParseMode = c.parseMode
}

func (c *caption) ApplyVoiceConfigOption(parameter *params.VoiceConfig) {
	parameter.Caption = c.text
	parameter.ParseMode = c.parseMode
}
