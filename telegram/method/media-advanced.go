package method

import (
	"github.com/irisked/gogram/args"
	"github.com/irisked/gogram/telegram/keyboard"
	"github.com/irisked/gogram/types"
)

// CopyMessage contains information about copyMessage telegram api method.
type CopyMessage struct {
	ChatID              int64           `json:"chat_id" form:"chat_id"`
	FromChatID          int64           `json:"from_chat_id" form:"from_chat_id"`
	MessageID           int             `json:"message_id" form:"message_id"`
	Caption             string          `json:"caption,omitempty" form:"caption,omitempty"`
	ParseMode           string          `json:"parse_mode,omitempty" form:"parse_mode,omitempty"`
	DisableNotification bool            `json:"disable_notification,omitempty" form:"disable_notification,omitempty"`
	ReplyToMessageID    int             `json:"reply_to_message_id,omitempty" form:"reply_to_message_id,omitempty"`
	AllowCopy           bool            `json:"allow_sending_without_reply,omitempty" form:"allow_sending_without_reply,omitempty"`
	ReplyMarkup         keyboard.Markup `json:"reply_markup,omitempty" form:"reply_markup,omitempty"`
	SimpleMethod
}

// NewCopyMessage creates CopyMessage.
func NewCopyMessage(chatID, fromChatID int64, messageID int, opts ...Option[*CopyMessage]) *CopyMessage {
	m := &CopyMessage{
		ChatID:     chatID,
		FromChatID: fromChatID,
		MessageID:  messageID,
	}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *CopyMessage) Method() string {
	return "copyMessage"
}

// SetCaption sets caption
func (r *CopyMessage) SetCaption(c string) {
	r.Caption = c
}

// SetParseMode sets parse_mode
func (r *CopyMessage) SetParseMode(m string) {
	r.ParseMode = m
}

// SetDisableNotification sets disable_notification
func (r *CopyMessage) SetDisableNotification(v bool) {
	r.DisableNotification = v
}

// SetReplyToMessageID sets reply_to_message_id
func (r *CopyMessage) SetReplyToMessageID(id int) {
	r.ReplyToMessageID = id
}

// SetAllowCopy sets allow_sending_without_reply
func (r *CopyMessage) SetAllowCopy(v bool) {
	r.AllowCopy = v
}

// SetReplyMarkup sets reply_markup
func (r *CopyMessage) SetReplyMarkup(m keyboard.Markup) {
	r.ReplyMarkup = m
}

// CopyMessages contains information about copyMessages telegram api method.
type CopyMessages struct {
	ChatID              int64 `json:"chat_id" form:"chat_id"`
	FromChatID          int64 `json:"from_chat_id" form:"from_chat_id"`
	MessageIDs          []int `json:"message_ids" form:"message_ids"`
	DisableNotification bool  `json:"disable_notification,omitempty" form:"disable_notification,omitempty"`
	RemoveSender        bool  `json:"remove_sender,omitempty" form:"remove_sender,omitempty"`
	SimpleMethod
}

// NewCopyMessages creates CopyMessages.
func NewCopyMessages(chatID, fromChatID int64, messageIDs []int, opts ...Option[*CopyMessages]) *CopyMessages {
	m := &CopyMessages{
		ChatID:     chatID,
		FromChatID: fromChatID,
		MessageIDs: messageIDs,
	}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *CopyMessages) Method() string {
	return "copyMessages"
}

// SetDisableNotification sets disable_notification
func (r *CopyMessages) SetDisableNotification(v bool) {
	r.DisableNotification = v
}

// SetRemoveSender sets remove_sender
func (r *CopyMessages) SetRemoveSender(v bool) {
	r.RemoveSender = v
}

// DeleteMessages contains information about deleteMessages telegram api method.
type DeleteMessages struct {
	ChatID     int64 `json:"chat_id" form:"chat_id"`
	MessageIDs []int `json:"message_ids" form:"message_ids"`
	SimpleMethod
}

// NewDeleteMessages creates DeleteMessages.
func NewDeleteMessages(chatID int64, messageIDs []int) *DeleteMessages {
	m := &DeleteMessages{
		ChatID:     chatID,
		MessageIDs: messageIDs,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *DeleteMessages) Method() string {
	return "deleteMessages"
}

// StopPoll contains information about stopPoll telegram api method.
type StopPoll struct {
	ChatID      int64           `json:"chat_id" form:"chat_id"`
	MessageID   int             `json:"message_id" form:"message_id"`
	ReplyMarkup keyboard.Markup `json:"reply_markup,omitempty" form:"reply_markup,omitempty"`
	SimpleMethod
}

// NewStopPoll creates StopPoll.
func NewStopPoll(chatID int64, messageID int, opts ...Option[*StopPoll]) *StopPoll {
	m := &StopPoll{
		ChatID:    chatID,
		MessageID: messageID,
	}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *StopPoll) Method() string {
	return "stopPoll"
}

// SetReplyMarkup sets reply_markup
func (r *StopPoll) SetReplyMarkup(kb keyboard.Markup) {
	r.ReplyMarkup = kb
}

// EditMessageLiveLocation contains information about editMessageLiveLocation telegram api method.
type EditMessageLiveLocation struct {
	ChatID          int64           `json:"chat_id,omitempty" form:"chat_id,omitempty"`
	MessageID       int             `json:"message_id,omitempty" form:"message_id,omitempty"`
	InlineMessageID string          `json:"inline_message_id,omitempty" form:"inline_message_id,omitempty"`
	Latitude        float64         `json:"latitude" form:"latitude"`
	Longitude       float64         `json:"longitude" form:"longitude"`
	ReplyMarkup     keyboard.Markup `json:"reply_markup,omitempty" form:"reply_markup,omitempty"`
	SimpleMethod
}

// NewEditMessageLiveLocation creates EditMessageLiveLocation.
func NewEditMessageLiveLocation(message *args.MessageID, location types.Location, opts ...Option[*EditMessageLiveLocation]) *EditMessageLiveLocation {
	m := &EditMessageLiveLocation{
		ChatID:          message.ChatID,
		MessageID:       message.MessageID,
		InlineMessageID: message.InlineMessageID,
		Latitude:        location.Latitude,
		Longitude:       location.Longitude,
	}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *EditMessageLiveLocation) Method() string {
	return "editMessageLiveLocation"
}

// SetReplyMarkup sets reply_markup
func (r *EditMessageLiveLocation) SetReplyMarkup(kb keyboard.Markup) {
	r.ReplyMarkup = kb
}

// StopMessageLiveLocation contains information about stopMessageLiveLocation telegram api method.
type StopMessageLiveLocation struct {
	ChatID          int64           `json:"chat_id,omitempty" form:"chat_id,omitempty"`
	MessageID       int             `json:"message_id,omitempty" form:"message_id,omitempty"`
	InlineMessageID string          `json:"inline_message_id,omitempty" form:"inline_message_id,omitempty"`
	ReplyMarkup     keyboard.Markup `json:"reply_markup,omitempty" form:"reply_markup,omitempty"`
	SimpleMethod
}

// NewStopMessageLiveLocation creates StopMessageLiveLocation.
func NewStopMessageLiveLocation(message *args.MessageID, opts ...Option[*StopMessageLiveLocation]) *StopMessageLiveLocation {
	m := &StopMessageLiveLocation{
		ChatID:          message.ChatID,
		MessageID:       message.MessageID,
		InlineMessageID: message.InlineMessageID,
	}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *StopMessageLiveLocation) Method() string {
	return "stopMessageLiveLocation"
}

// SetReplyMarkup sets reply_markup
func (r *StopMessageLiveLocation) SetReplyMarkup(kb keyboard.Markup) {
	r.ReplyMarkup = kb
}

// DeleteWebhook contains information about deleteWebhook telegram api method.
type DeleteWebhook struct {
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty" form:"drop_pending_updates,omitempty"`
	SimpleMethod
}

// NewDeleteWebhook creates DeleteWebhook.
func NewDeleteWebhook(opts ...Option[*DeleteWebhook]) *DeleteWebhook {
	m := &DeleteWebhook{}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *DeleteWebhook) Method() string {
	return "deleteWebhook"
}

// SetDropPendingUpdates sets drop_pending_updates
func (r *DeleteWebhook) SetDropPendingUpdates(v bool) {
	r.DropPendingUpdates = v
}
