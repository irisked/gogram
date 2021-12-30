package option

import (
	"gogram/telegram/method"
)

// Offset returns struct for setting optional Offset field.
func Offset(o int) interface {
	hasOffset
} {
	return &offsetOption{offset: o}
}

type hasOffset interface {
	method.GetUserProfilePhotosOption
	method.GetUpdatesOption
}

type offsetOption struct {
	offset int
}

func (sa *offsetOption) ApplyGetUserProfilePhotosOption(acq *method.GetUserProfilePhotos) {
	acq.Offset = sa.offset
}

func (sa *offsetOption) ApplyGetUpdatesOption(gu *method.GetUpdates) {
	gu.Offset = sa.offset
}
