package option

import (
	"gogram/telegram/method"
)

// DisableWebPagePreview returns struct for setting optional DisableWebPagePreview field.
func DisableWebPagePreview() interface {
	disableWebPagePreviewSetter
} {
	return &disableWebPagePreviewOption{true}
}

// disableWebPagePreviewSetter is an interface for setting disableWebPagePreview option.
type disableWebPagePreviewSetter interface {
	method.SendMessageOption
	method.EditMessageTextOption
}

type disableWebPagePreviewOption struct {
	disable bool
}

func (d *disableWebPagePreviewOption) ApplySendMessageOption(msg *method.SendMessage) {
	msg.DisableWebPagePreview = d.disable
}

func (d *disableWebPagePreviewOption) ApplyEditMessageTextOption(msg *method.EditMessageText) {
	msg.DisableWebPagePreview = d.disable
}
