package telegram

// MessageTarget identifies the message an edit request applies to:
// either a chat message (ChatID and MessageID) or an inline message
// (InlineMessageID). Not to be confused with MessageID, the result of
// copyMessage.
type MessageTarget struct {
	ChatID          int64  `json:"chat_id,omitempty"`
	MessageID       int    `json:"message_id,omitempty"`
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

// ChatMessage targets message messageID in chat chatID.
func ChatMessage(chatID int64, messageID int) MessageTarget {
	return MessageTarget{ChatID: chatID, MessageID: messageID}
}

// InlineMessage targets a message sent via inline mode.
func InlineMessage(inlineMessageID string) MessageTarget {
	return MessageTarget{InlineMessageID: inlineMessageID}
}
