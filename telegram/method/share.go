package method

import (
	"github.com/irisked/gogram/telegram/keyboard"
	"github.com/irisked/gogram/types"
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

// NewSendContact creates SendContact.
func NewSendContact(chatID int64, contact types.Contact, opts ...Option[*SendContact]) *SendContact {
	method := new(SendContact)
	method.ChatID = chatID
	method.PhoneNumber = contact.PhoneNumber
	method.FirstName = contact.FirstName
	method.LastName = contact.LastName
	for _, o := range opts {
		o(method)
	}
	return method
}

// Method returns telegram api method endpoint.
func (r *SendContact) Method() string {
	return "sendContact"
}

// SetLastName sets last_name
func (r *SendContact) SetLastName(s string) {
	r.LastName = s
}

// SetVCard sets vcard
func (r *SendContact) SetVCard(s string) {
	r.VCard = s
}

// SetDisableNotification sets disable_notification
func (r *SendContact) SetDisableNotification(v bool) {
	r.DisableNotification = v
}

// SetReplyToMessageID sets reply_to_message_id
func (r *SendContact) SetReplyToMessageID(id int) {
	r.ReplyToMessageID = id
}

// SetReplyMarkup sets reply_markup
func (r *SendContact) SetReplyMarkup(kb keyboard.Keyboard) {
	r.ReplyMarkup = kb
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

// NewSendLocation creates SendLocation.
func NewSendLocation(chatID int64, location types.Location, opts ...Option[*SendLocation]) *SendLocation {
	method := new(SendLocation)
	method.ChatID = chatID
	method.Latitude = location.Latitude
	method.Longitude = location.Longitude
	for _, o := range opts {
		o(method)
	}
	return method
}

// Method returns telegram api method endpoint.
func (r *SendLocation) Method() string {
	return "sendLocation"
}

// SetLivePeriod sets live_period
func (r *SendLocation) SetLivePeriod(v int) {
	r.LivePeriod = v
}

// SetDisableNotification sets disable_notification
func (r *SendLocation) SetDisableNotification(v bool) {
	r.DisableNotification = v
}

// SetReplyToMessageID sets reply_to_message_id
func (r *SendLocation) SetReplyToMessageID(id int) {
	r.ReplyToMessageID = id
}

// SetReplyMarkup sets reply_markup
func (r *SendLocation) SetReplyMarkup(kb keyboard.Keyboard) {
	r.ReplyMarkup = kb
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

// NewSendVenue creates SendVenue.
func NewSendVenue(chatID int64, venue types.Venue, opts ...Option[*SendVenue]) *SendVenue {
	method := new(SendVenue)
	method.ChatID = chatID
	method.Latitude = venue.Location.Latitude
	method.Longitude = venue.Location.Longitude
	method.Title = venue.Title
	method.Address = venue.Address
	for _, o := range opts {
		o(method)
	}
	return method
}

// Method returns telegram api method endpoint.
func (r *SendVenue) Method() string {
	return "sendVenue"
}

// SetFoursquareID sets foursquare_id
func (r *SendVenue) SetFoursquareID(s string) {
	r.FoursquareID = s
}

// SetFoursquareType sets foursquare_type
func (r *SendVenue) SetFoursquareType(s string) {
	r.FoursquareType = s
}

// SetDisableNotification sets disable_notification
func (r *SendVenue) SetDisableNotification(v bool) {
	r.DisableNotification = v
}

// SetReplyToMessageID sets reply_to_message_id
func (r *SendVenue) SetReplyToMessageID(id int) {
	r.ReplyToMessageID = id
}

// SetReplyMarkup sets reply_markup
func (r *SendVenue) SetReplyMarkup(kb keyboard.Keyboard) {
	r.ReplyMarkup = kb
}
