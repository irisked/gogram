package option

import (
	"gogram/telegram/method"
)

// Limit returns struct for setting optional Limit field.
func Limit(o int) interface {
	limitOptionSetter
} {
	return &limitOption{o}
}

// MaxConnections returns struct for setting optional MaxConnections field.
func MaxConnections(o int) interface {
	maxConnectionsoptionSetter
} {
	return &limitOption{o}
}

type maxConnectionsoptionSetter interface {
	method.SetWebhookOption
}

type limitOptionSetter interface {
	method.GetUserProfilePhotosOption
	method.GetUpdatesOption
}

type limitOption struct {
	limit int
}

func (sa *limitOption) ApplyGetUserProfilePhotosOption(acq *method.GetUserProfilePhotos) {
	acq.Limit = sa.limit
}

func (sa *limitOption) ApplyGetUpdatesOption(acq *method.GetUpdates) {
	acq.Limit = sa.limit
}

func (sa *limitOption) ApplySetWebhookOption(sw *method.SetWebhook) {
	sw.MaxConnections = sa.limit
}
