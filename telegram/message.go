package telegram

import (
	"errors"
	"net/url"
	"strings"
	"time"
)

// MessageType is the kind of content a Message carries; see Message.Type.
type MessageType string

func (mt MessageType) String() string { return string(mt) }

// Message types.
const (
	MessageTypeCommand           MessageType = "command"
	MessageTypeText              MessageType = "text"
	MessageTypeAnimation         MessageType = "animation"
	MessageTypeAudio             MessageType = "audio"
	MessageTypeDocument          MessageType = "document"
	MessageTypePhoto             MessageType = "photo"
	MessageTypeSticker           MessageType = "sticker"
	MessageTypeStory             MessageType = "story"
	MessageTypeVideo             MessageType = "video"
	MessageTypeVideoNote         MessageType = "video_note"
	MessageTypeVoice             MessageType = "voice"
	MessageTypePaidMedia         MessageType = "paid_media"
	MessageTypeContact           MessageType = "contact"
	MessageTypeDice              MessageType = "dice"
	MessageTypeGame              MessageType = "game"
	MessageTypePoll              MessageType = "poll"
	MessageTypeChecklist         MessageType = "checklist"
	MessageTypeVenue             MessageType = "venue"
	MessageTypeLocation          MessageType = "location"
	MessageTypeNewChatMembers    MessageType = "new_chat_members"
	MessageTypeLeftChatMember    MessageType = "left_chat_member"
	MessageTypeNewChatTitle      MessageType = "new_chat_title"
	MessageTypeNewChatPhoto      MessageType = "new_chat_photo"
	MessageTypeDeleteChatPhoto   MessageType = "delete_chat_photo"
	MessageTypePinnedMessage     MessageType = "pinned_message"
	MessageTypeInvoice           MessageType = "invoice"
	MessageTypeSuccessfulPayment MessageType = "successful_payment"
	MessageTypeWebAppData        MessageType = "web_app_data"
	MessageTypeOther             MessageType = "other"
)

// Type reports what the message carries. A message starting with a bot
// command is MessageTypeCommand, even though it is also text.
func (m *Message) Type() MessageType {
	switch {
	case m.IsCommand():
		return MessageTypeCommand
	case m.Text != "":
		return MessageTypeText
	// Animations also set Document, so check them first.
	case m.Animation != nil:
		return MessageTypeAnimation
	case m.Audio != nil:
		return MessageTypeAudio
	case m.Document != nil:
		return MessageTypeDocument
	case len(m.Photo) > 0:
		return MessageTypePhoto
	case m.Sticker != nil:
		return MessageTypeSticker
	case m.Story != nil:
		return MessageTypeStory
	case m.Video != nil:
		return MessageTypeVideo
	case m.VideoNote != nil:
		return MessageTypeVideoNote
	case m.Voice != nil:
		return MessageTypeVoice
	case m.PaidMedia != nil:
		return MessageTypePaidMedia
	case m.Contact != nil:
		return MessageTypeContact
	case m.Dice != nil:
		return MessageTypeDice
	case m.Game != nil:
		return MessageTypeGame
	case m.Poll != nil:
		return MessageTypePoll
	case m.Checklist != nil:
		return MessageTypeChecklist
	// Venues also set Location, so check them first.
	case m.Venue != nil:
		return MessageTypeVenue
	case m.Location != nil:
		return MessageTypeLocation
	case len(m.NewChatMembers) > 0:
		return MessageTypeNewChatMembers
	case m.LeftChatMember != nil:
		return MessageTypeLeftChatMember
	case m.NewChatTitle != "":
		return MessageTypeNewChatTitle
	case len(m.NewChatPhoto) > 0:
		return MessageTypeNewChatPhoto
	case m.DeleteChatPhoto:
		return MessageTypeDeleteChatPhoto
	case m.PinnedMessage != nil:
		return MessageTypePinnedMessage
	case m.Invoice != nil:
		return MessageTypeInvoice
	case m.SuccessfulPayment != nil:
		return MessageTypeSuccessfulPayment
	case m.WebAppData != nil:
		return MessageTypeWebAppData
	default:
		return MessageTypeOther
	}
}

// Time returns the message date.
func (m *Message) Time() time.Time {
	return time.Unix(m.Date, 0)
}

// IsAccessible reports whether the message content is available. Fields
// that may hold a message the bot can no longer access, such as
// CallbackQuery.Message, then carry only Chat and MessageID, with a zero
// Date.
func (m *Message) IsAccessible() bool {
	return m.Date != 0
}

// IsCommand reports whether the message starts with a bot command.
func (m *Message) IsCommand() bool {
	if len(m.Entities) == 0 {
		return false
	}
	e := m.Entities[0]
	return e.Offset == 0 && e.Type == "bot_command" && e.Length <= len(m.Text)
}

// Command returns the command without the leading slash and without a
// "@botname" suffix: "/start@my_bot payload" yields "start". It returns
// "" if the message is not a command.
func (m *Message) Command() string {
	command, _, _ := strings.Cut(m.CommandWithAt(), "@")
	return command
}

// CommandWithAt is like Command but keeps the "@botname" suffix.
func (m *Message) CommandWithAt() string {
	if !m.IsCommand() {
		return ""
	}
	return m.Text[1:m.Entities[0].Length]
}

// CommandBot returns the bot a command is addressed to: "my_bot" for
// "/start@my_bot". It returns "" for commands without a "@botname"
// suffix, which are addressed to every bot in the chat.
func (m *Message) CommandBot() string {
	_, bot, _ := strings.Cut(m.CommandWithAt(), "@")
	return bot
}

// CommandArguments returns the text after the command: "/start foo bar"
// yields "foo bar". It returns "" if the message is not a command.
func (m *Message) CommandArguments() string {
	if !m.IsCommand() {
		return ""
	}
	return strings.TrimSpace(m.Text[m.Entities[0].Length:])
}

// ParseURL parses the URL of a "text_link" entity.
func (e MessageEntity) ParseURL() (*url.URL, error) {
	if e.URL == "" {
		return nil, errors.New("telegram: entity has no URL")
	}
	return url.Parse(e.URL)
}
