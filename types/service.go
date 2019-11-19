package types

import (
	"strings"
	"github.com/irisked/gogram/args"
)

// ChatAction it's a type of action to broadcast.
type ChatAction string

const (
	// Typing chat action for text messages.
	Typing ChatAction = "typing"
	// UploadingPhoto chat action for photo uploading.
	UploadingPhoto ChatAction = "upload_photo"
	// RecordingVideo chat action for uploading video.
	RecordingVideo ChatAction = "record_video"
	// UploadingAudio chat action for uploading audio.
	UploadingAudio ChatAction = "upload_audio"
	// UploadingDocument chat action for uploading document.
	UploadingDocument ChatAction = "upload_document"
	// RecordingVideNote chat action for uploading video note.
	RecordingVideNote ChatAction = "record_video_note"
	// UploadingVideoNote chat action for uploading video note.
	UploadingVideoNote ChatAction = "upload_video_note"
)

// UpdateType represents all available update types.
type UpdateType string

func (ut UpdateType) String() string {
	return string(ut)
}

const (
	// IncomingMessage message update type.
	IncomingMessage UpdateType = "message"
	// IncomingEditedMessage edited_message update type.
	IncomingEditedMessage UpdateType = "edited_message"
	// IncomingChannelPost channel_post update type.
	IncomingChannelPost UpdateType = "channel_post"
	// IncomingEditedChannelPost edited_channel_post update type.
	IncomingEditedChannelPost UpdateType = "edited_channel_post"
	// IncomingInlineQuery inline_query update type.
	IncomingInlineQuery UpdateType = "inline_query"
	// IncomingChosenInlineResult chosen_inline_result update type.
	IncomingChosenInlineResult UpdateType = "chosen_inline_result"
	// IncomingCallbackQuery callback_query update type.
	IncomingCallbackQuery UpdateType = "callback_query"
	// IncomingShippingQuery shipping_query update type.
	IncomingShippingQuery UpdateType = "shipping_query"
	// IncomingPreCheckoutQuery pre_checkout_query update type.
	IncomingPreCheckoutQuery UpdateType = "pre_checkout_query"
)

// Update is an update response, from GetUpdates.
type Update struct {
	UpdateID           int                 `json:"update_id"`
	Message            *Message            `json:"message"`
	EditedMessage      *Message            `json:"edited_message"`
	ChannelPost        *Message            `json:"channel_post"`
	EditedChannelPost  *Message            `json:"edited_channel_post"`
	InlineQuery        *InlineQuery        `json:"inline_query"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`
	CallbackQuery      *CallbackQuery      `json:"callback_query"`
	ShippingQuery      *ShippingQuery      `json:"shipping_query"`
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query"`
}

// Path returns update path.
func (u *Update) Path() string {
	switch ut := u.Type(); ut {
	case IncomingMessage:
		return "/" + ut.String() + u.Message.Path()
	case IncomingEditedMessage:
		return "/" + ut.String() + u.EditedMessage.Path()
	case IncomingChannelPost:
		return "/" + ut.String() + u.ChannelPost.Path()
	case IncomingEditedChannelPost:
		return "/" + ut.String() + u.EditedChannelPost.Path()
	case IncomingCallbackQuery:
		return "/" + ut.String() + u.CallbackQuery.Path()
	
	default:
		return ""
	}
}

// IsMessage returns update type.
func (u *Update) IsMessage() bool {
	return u.Message != nil || u.EditedMessage != nil ||
		u.ChannelPost != nil || u.EditedChannelPost != nil
}

// Type returns update type.
func (u *Update) Type() UpdateType {
	if u.EditedMessage != nil {
		return IncomingEditedMessage
	}
	if u.ChannelPost != nil {
		return IncomingChannelPost
	}
	if u.EditedChannelPost != nil {
		return IncomingEditedChannelPost
	}
	if u.InlineQuery != nil {
		return IncomingInlineQuery
	}
	if u.CallbackQuery != nil {
		return IncomingCallbackQuery
	}
	if u.ShippingQuery != nil {
		return IncomingShippingQuery
	}
	if u.PreCheckoutQuery != nil {
		return IncomingPreCheckoutQuery
	}
	return IncomingMessage
}

// CallbackQuery is data sent when a keyboard button with callback data
// is clicked.
type CallbackQuery struct {
	ID              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message"`           // optional
	InlineMessageID string   `json:"inline_message_id"` // optional
	ChatInstance    string   `json:"chat_instance"`

	Data          string `json:"data"`            // optional
	GameShortName string `json:"game_short_name"` // optional
}

// IsFromInlineMessage returns true if callback query comes from inline message
// and false otherwise.
func (cq *CallbackQuery) IsFromInlineMessage() bool {
	return cq.InlineMessageID != ""
}

func (cq *CallbackQuery) Path() string {
	return "/" + strings.Split(cq.Data, ":")[0]
}

// MessageID returns id of message.
func (cq *CallbackQuery) MessageID() *args.MessageID {
	if cq.IsFromInlineMessage() {
		return args.InlineMessageID(cq.InlineMessageID)
	}
	return args.ChatMessage(cq.Message.MessageID, cq.Message.Chat.ID)
}

// InlineQuery is a Query from domain for an inline method.
type InlineQuery struct {
	ID       string    `json:"id"`
	From     *User     `json:"from"`
	Location *Location `json:"location"` // optional
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

// ChosenInlineResult is an inline query result chosen by a User
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            *User     `json:"from"`
	Location        *Location `json:"location"`
	InlineMessageID string    `json:"inline_message_id"`
	Query           string    `json:"query"`
}

// WebhookInfo is information about a currently set webhook.
type WebhookInfo struct {
	URL                  string `json:"url"`
	HasCustomCertificate bool   `json:"has_custom_certificate"`
	PendingUpdateCount   int    `json:"pending_update_count"`
	LastErrorDate        int    `json:"last_error_date"`    // optional
	LastErrorMessage     string `json:"last_error_message"` // optional
}
