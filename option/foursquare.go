package option

import (
	"gogram/telegram/method"
)

// Foursquare returns struct for setting optional Foursquare field.
func Foursquare(id, vType string) interface {
	foursquareOptionSetter
} {
	return &foursquareOption{id, vType}
}

type foursquareOptionSetter interface {
	method.SendVenueOption
}

type foursquareOption struct {
	id    string
	vType string
}

func (fo *foursquareOption) ApplySendVenueOption(sv *method.SendVenue) {
	sv.FoursquareID = fo.id
	sv.FoursquareType = fo.vType
}
