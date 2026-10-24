package method

import ()

// ForumTopicCreated contains information about forum_topic_created service message.
type ForumTopicCreated struct {
	Name      string `json:"name"`
	IconColor int    `json:"icon_color,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty"`
}

// ForumTopicClosed contains information about forum_topic_closed service message.
type ForumTopicClosed struct{}

// ForumTopicEdited contains information about forum_topic_edited service message.
type ForumTopicEdited struct {
	Name      string `json:"name,omitempty"`
	IconColor int    `json:"icon_color,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty"`
}

// ForumTopicReopened contains information about forum_topic_reopened service message.
type ForumTopicReopened struct{}

// GeneralForumTopicHidden contains information about general_forum_topic_hidden service message.
type GeneralForumTopicHidden struct{}

// GeneralForumTopicUnhidden contains information about general_forum_topic_unhidden service message.
type GeneralForumTopicUnhidden struct{}

// ForumTopic contains information about a forum topic.
type ForumTopic struct {
	MessageID       int    `json:"message_id"`
	TopicName       string `json:"topic_name"`
	IconColor       int    `json:"icon_color"`
	IconCustomEmoji string `json:"icon_custom_emoji,omitempty"`
}

// CreateForumTopic contains information about createForumTopic telegram api method.
type CreateForumTopic struct {
	ChatID      int64  `json:"chat_id" form:"chat_id"`
	Name        string `json:"name" form:"name"`
	IconColor   int    `json:"icon_color,omitempty" form:"icon_color,omitempty"`
	IconEmojiID string `json:"icon_custom_emoji_id,omitempty" form:"icon_custom_emoji_id,omitempty"`
	SimpleMethod
}

// NewCreateForumTopic creates CreateForumTopic.
func NewCreateForumTopic(chatID int64, name string, opts ...Option[*CreateForumTopic]) *CreateForumTopic {
	m := &CreateForumTopic{
		ChatID: chatID,
		Name:   name,
	}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *CreateForumTopic) Method() string {
	return "createForumTopic"
}

// SetIconColor sets icon_color
func (r *CreateForumTopic) SetIconColor(c int) {
	r.IconColor = c
}

// SetIconEmojiID sets icon_custom_emoji_id
func (r *CreateForumTopic) SetIconEmojiID(e string) {
	r.IconEmojiID = e
}

// EditForumTopic contains information about editForumTopic telegram api method.
type EditForumTopic struct {
	ChatID      int64  `json:"chat_id" form:"chat_id"`
	MessageID   int    `json:"message_id" form:"message_id"`
	Name        string `json:"name,omitempty" form:"name,omitempty"`
	IconEmojiID string `json:"icon_custom_emoji_id,omitempty" form:"icon_custom_emoji_id,omitempty"`
	SimpleMethod
}

// NewEditForumTopic creates EditForumTopic.
func NewEditForumTopic(chatID int64, messageID int, opts ...Option[*EditForumTopic]) *EditForumTopic {
	m := &EditForumTopic{
		ChatID:    chatID,
		MessageID: messageID,
	}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *EditForumTopic) Method() string {
	return "editForumTopic"
}

// SetName sets name
func (r *EditForumTopic) SetName(n string) {
	r.Name = n
}

// SetIconEmojiID sets icon_custom_emoji_id
func (r *EditForumTopic) SetIconEmojiID(e string) {
	r.IconEmojiID = e
}

// DeleteForumTopic contains information about deleteForumTopic telegram api method.
type DeleteForumTopic struct {
	ChatID    int64 `json:"chat_id" form:"chat_id"`
	MessageID int   `json:"message_id" form:"message_id"`
	SimpleMethod
}

// NewDeleteForumTopic creates DeleteForumTopic.
func NewDeleteForumTopic(chatID int64, messageID int) *DeleteForumTopic {
	m := &DeleteForumTopic{
		ChatID:    chatID,
		MessageID: messageID,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *DeleteForumTopic) Method() string {
	return "deleteForumTopic"
}

// GetForumTopic contains information about getForumTopic telegram api method.
type GetForumTopic struct {
	ChatID    int64 `json:"chat_id" form:"chat_id"`
	MessageID int   `json:"message_id" form:"message_id"`
	SimpleMethod
}

// NewGetForumTopic creates GetForumTopic.
func NewGetForumTopic(chatID int64, messageID int) *GetForumTopic {
	m := &GetForumTopic{
		ChatID:    chatID,
		MessageID: messageID,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *GetForumTopic) Method() string {
	return "getForumTopic"
}

// CloseForumTopic contains information about closeForumTopic telegram api method.
type CloseForumTopic struct {
	ChatID    int64 `json:"chat_id" form:"chat_id"`
	MessageID int   `json:"message_id" form:"message_id"`
	SimpleMethod
}

// NewCloseForumTopic creates CloseForumTopic.
func NewCloseForumTopic(chatID int64, messageID int) *CloseForumTopic {
	m := &CloseForumTopic{
		ChatID:    chatID,
		MessageID: messageID,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *CloseForumTopic) Method() string {
	return "closeForumTopic"
}

// ReopenForumTopic contains information about reopenForumTopic telegram api method.
type ReopenForumTopic struct {
	ChatID    int64 `json:"chat_id" form:"chat_id"`
	MessageID int   `json:"message_id" form:"message_id"`
	SimpleMethod
}

// NewReopenForumTopic creates ReopenForumTopic.
func NewReopenForumTopic(chatID int64, messageID int) *ReopenForumTopic {
	m := &ReopenForumTopic{
		ChatID:    chatID,
		MessageID: messageID,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *ReopenForumTopic) Method() string {
	return "reopenForumTopic"
}

// GetForumTopicIconStickers contains information about getForumTopicIconStickers telegram api method.
type GetForumTopicIconStickers struct {
	SimpleMethod
}

// NewGetForumTopicIconStickers creates GetForumTopicIconStickers.
func NewGetForumTopicIconStickers() *GetForumTopicIconStickers {
	return new(GetForumTopicIconStickers)
}

// Method returns telegram api method endpoint.
func (r *GetForumTopicIconStickers) Method() string {
	return "getForumTopicIconStickers"
}

// UnpinAllForumTopicMessages contains information about unpinAllForumTopicMessages telegram api method.
type UnpinAllForumTopicMessages struct {
	ChatID    int64 `json:"chat_id" form:"chat_id"`
	MessageID int   `json:"message_id" form:"message_id"`
	SimpleMethod
}

// NewUnpinAllForumTopicMessages creates UnpinAllForumTopicMessages.
func NewUnpinAllForumTopicMessages(chatID int64, messageID int) *UnpinAllForumTopicMessages {
	m := &UnpinAllForumTopicMessages{
		ChatID:    chatID,
		MessageID: messageID,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *UnpinAllForumTopicMessages) Method() string {
	return "unpinAllForumTopicMessages"
}

// EditGeneralForumTopic contains information about editGeneralForumTopic telegram api method.
type EditGeneralForumTopic struct {
	ChatID int64  `json:"chat_id" form:"chat_id"`
	Name   string `json:"name" form:"name"`
	SimpleMethod
}

// NewEditGeneralForumTopic creates EditGeneralForumTopic.
func NewEditGeneralForumTopic(chatID int64, name string) *EditGeneralForumTopic {
	m := &EditGeneralForumTopic{
		ChatID: chatID,
		Name:   name,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *EditGeneralForumTopic) Method() string {
	return "editGeneralForumTopic"
}

// CloseGeneralForumTopic contains information about closeGeneralForumTopic telegram api method.
type CloseGeneralForumTopic struct {
	ChatID int64 `json:"chat_id" form:"chat_id"`
	SimpleMethod
}

// NewCloseGeneralForumTopic creates CloseGeneralForumTopic.
func NewCloseGeneralForumTopic(chatID int64) *CloseGeneralForumTopic {
	m := &CloseGeneralForumTopic{
		ChatID: chatID,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *CloseGeneralForumTopic) Method() string {
	return "closeGeneralForumTopic"
}

// ReopenGeneralForumTopic contains information about reopenGeneralForumTopic telegram api method.
type ReopenGeneralForumTopic struct {
	ChatID int64 `json:"chat_id" form:"chat_id"`
	SimpleMethod
}

// NewReopenGeneralForumTopic creates ReopenGeneralForumTopic.
func NewReopenGeneralForumTopic(chatID int64) *ReopenGeneralForumTopic {
	m := &ReopenGeneralForumTopic{
		ChatID: chatID,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *ReopenGeneralForumTopic) Method() string {
	return "reopenGeneralForumTopic"
}

// HideGeneralForumTopic contains information about hideGeneralForumTopic telegram api method.
type HideGeneralForumTopic struct {
	ChatID int64 `json:"chat_id" form:"chat_id"`
	SimpleMethod
}

// NewHideGeneralForumTopic creates HideGeneralForumTopic.
func NewHideGeneralForumTopic(chatID int64) *HideGeneralForumTopic {
	m := &HideGeneralForumTopic{
		ChatID: chatID,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *HideGeneralForumTopic) Method() string {
	return "hideGeneralForumTopic"
}

// UnhideGeneralForumTopic contains information about unhideGeneralForumTopic telegram api method.
type UnhideGeneralForumTopic struct {
	ChatID int64 `json:"chat_id" form:"chat_id"`
	SimpleMethod
}

// NewUnhideGeneralForumTopic creates UnhideGeneralForumTopic.
func NewUnhideGeneralForumTopic(chatID int64) *UnhideGeneralForumTopic {
	m := &UnhideGeneralForumTopic{
		ChatID: chatID,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *UnhideGeneralForumTopic) Method() string {
	return "unhideGeneralForumTopic"
}

// UnpinAllGeneralForumTopicMessages contains information about unpinAllGeneralForumTopicMessages telegram api method.
type UnpinAllGeneralForumTopicMessages struct {
	ChatID int64 `json:"chat_id" form:"chat_id"`
	SimpleMethod
}

// NewUnpinAllGeneralForumTopicMessages creates UnpinAllGeneralForumTopicMessages.
func NewUnpinAllGeneralForumTopicMessages(chatID int64) *UnpinAllGeneralForumTopicMessages {
	m := &UnpinAllGeneralForumTopicMessages{
		ChatID: chatID,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *UnpinAllGeneralForumTopicMessages) Method() string {
	return "unpinAllGeneralForumTopicMessages"
}
