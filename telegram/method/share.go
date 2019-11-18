package method

import (
	"github.com/irisked/gogram/types"
	"github.com/irisked/gogram/types/markup/keyboard"
)

// SendContact contains information about sendContact telegram api method.
type SendContact struct {
	ChatID              int64             `json:"chat_id"`
	PhoneNumber         string            `json:"phone_number"`
	FirstName           string            `json:"first_name"`
	LastName            string            `json:"last_name,omitempty"`
	VCard               string            `json:"vcard,omitempty"`
	DisableNotification bool              `json:"disable_notification,omitempty"`
	ReplyToMessageID    int               `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         keyboard.Keyboard `json:"reply_markup,omitempty"`
	SimpleMethod
}

// SendContactOption interface for optional parameters.
type SendContactOption interface {
	ApplySendContactOption(*SendContact)
}

// NewSendContact creates SendContact.
func NewSendContact(chatID int64, contact types.Contact, options ...SendContactOption) *SendContact {
	method := new(SendContact)
	method.ChatID = chatID
	method.PhoneNumber = contact.PhoneNumber
	method.FirstName = contact.FirstName
	method.LastName = contact.LastName
	for _, option := range options {
		option.ApplySendContactOption(method)
	}
	return method
}

// Method returns telegram api method endpoint.
func (r *SendContact) Method() string {
	return "sendContact"
}

// SendLocation contains information about sendLocation telegram api method.
type SendLocation struct {
	ChatID              int64             `json:"chat_id"`
	Latitude            float64           `json:"latitude"`
	Longitude           float64           `json:"longitude"`
	LivePeriod          int               `json:"live_period,omitempty"`
	DisableNotification bool              `json:"disable_notification,omitempty"`
	ReplyToMessageID    int               `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         keyboard.Keyboard `json:"reply_markup,omitempty"`
	SimpleMethod
}

// SendLocationOption interface for optional parameters.
type SendLocationOption interface {
	ApplySendLocationOption(*SendLocation)
}

// NewSendLocation creates SendLocation.
func NewSendLocation(chatID int64, location types.Location, options ...SendLocationOption) *SendLocation {
	method := new(SendLocation)
	method.ChatID = chatID
	method.Latitude = location.Latitude
	method.Longitude = location.Longitude
	for _, option := range options {
		option.ApplySendLocationOption(method)
	}
	return method
}

// Method returns telegram api method endpoint.
func (r *SendLocation) Method() string {
	return "sendLocation"
}

// SendVenue contains information about sendVenue telegram api method.
type SendVenue struct {
	ChatID              int64             `json:"chat_id"`
	Latitude            float64           `json:"latitude"`
	Longitude           float64           `json:"longitude"`
	Title               string            `json:"title"`
	Address             string            `json:"address"`
	FoursquareID        string            `json:"foursquare_id,omitempty"`
	FoursquareType      string            `json:"foursquare_type,omitempty"`
	DisableNotification bool              `json:"disable_notification,omitempty"`
	ReplyToMessageID    int               `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         keyboard.Keyboard `json:"reply_markup,omitempty"`
	SimpleMethod
}

// SendVenueOption interface for optional parameters.
type SendVenueOption interface {
	ApplySendVenueOption(*SendVenue)
}

// NewSendVenue creates SendVenue.
func NewSendVenue(chatID int64, venue types.Venue, options ...SendVenueOption) *SendVenue {
	method := new(SendVenue)
	method.ChatID = chatID
	method.Latitude = venue.Location.Latitude
	method.Longitude = venue.Location.Longitude
	method.Title = venue.Title
	method.Address = venue.Address
	for _, option := range options {
		option.ApplySendVenueOption(method)
	}
	return method
}

// Method returns telegram api method endpoint.
func (r *SendVenue) Method() string {
	return "sendVenue"
}
