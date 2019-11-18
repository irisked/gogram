package params

// MessageID struct.
type MessageID struct {
	ChatID          int64  `json:"chat_id,omitempty"`
	MessageID       int    `json:"message_id,omitempty"`
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

// InlineMessageID inline message id.
func InlineMessageID(id string) *MessageID {
	return &MessageID{InlineMessageID: id}
}

// ChatMessage chat message.
func ChatMessage(msgID int, chatID int64) *MessageID {
	return &MessageID{MessageID: msgID, ChatID: chatID}
}
