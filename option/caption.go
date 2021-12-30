package option

import (
	"gogram/args"
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
	args.PhotoConfigOption
	args.AudioConfigOption
	args.DocumentConfigOption
	args.VideoConfigOption
	args.AnimationConfigOption
	args.VoiceConfigOption
}

type caption struct {
	text      string
	parseMode string
}

func (c *caption) ApplyPhotoConfigOption(parameter *args.PhotoConfig) {
	parameter.Caption = c.text
	parameter.ParseMode = c.parseMode
}

func (c *caption) ApplyAudioConfigOption(parameter *args.AudioConfig) {
	parameter.Caption = c.text
	parameter.ParseMode = c.parseMode
}

func (c *caption) ApplyDocumentConfigOption(parameter *args.DocumentConfig) {
	parameter.Caption = c.text
	parameter.ParseMode = c.parseMode
}

func (c *caption) ApplyVideoConfigOption(parameter *args.VideoConfig) {
	parameter.Caption = c.text
	parameter.ParseMode = c.parseMode
}

func (c *caption) ApplyAnimationConfigOption(parameter *args.AnimationConfig) {
	parameter.Caption = c.text
	parameter.ParseMode = c.parseMode
}

func (c *caption) ApplyVoiceConfigOption(parameter *args.VoiceConfig) {
	parameter.Caption = c.text
	parameter.ParseMode = c.parseMode
}
