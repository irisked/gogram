package content

// MessageTarget identifies which message an edit-style method (e.g.
// EditMessageText, EditMessageReplyMarkup) applies to -- either a
// chat_id+message_id pair, or an inline_message_id. Not to be
// confused with types.MessageID, which mirrors Telegram's MessageId
// API object (the response to copyMessage).
type MessageTarget struct {
	ChatID          int64  `json:"chat_id,omitempty"`
	MessageID       int    `json:"message_id,omitempty"`
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

// InlineMessageID targets a message by its inline_message_id.
func InlineMessageID(id string) *MessageTarget {
	return &MessageTarget{InlineMessageID: id}
}

// ChatMessage targets a message by its chat_id and message_id.
func ChatMessage(msgID int, chatID int64) *MessageTarget {
	return &MessageTarget{MessageID: msgID, ChatID: chatID}
}
