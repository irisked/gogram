package types

import (
	"errors"
	"net/url"
	"strings"
	"time"
)

// MessageType represents all available message types.
type MessageType string

func (mt MessageType) String() string {
	return string(mt)
}

const (
	// TextMessage represents text message sub type.
	TextMessage MessageType = "text"
	// AudioMessage represents audio message sub type.
	AudioMessage MessageType = "audio"
	// DocumentMessage represents document message sub type.
	DocumentMessage MessageType = "document"
	// AnimationMessage represents animation message sub type.
	AnimationMessage MessageType = "animation"

	// TODO: Game?

	// CommandMessage represents command message sub type.
	CommandMessage MessageType = "command_message"
	// PhotoMessage represents photo message sub type.
	PhotoMessage MessageType = "photo"
	// StickerMessage represents sticker message sub type.
	StickerMessage MessageType = "sticker"
	// VideoMessage represents video message sub type.
	VideoMessage MessageType = "video"
	// VideoNoteMessage represents video_note message sub type.
	VideoNoteMessage MessageType = "video_note"
	// VoiceMessage represents voice message sub type.
	VoiceMessage MessageType = "voice"
	// ContactMessage represents voice message sub type.
	ContactMessage MessageType = "contact"
	// LocationMessage represents voice message sub type.
	LocationMessage MessageType = "location"
	// VenueMessage represents venue message sub type.
	VenueMessage MessageType = "venue"
	// NewChatMembersMessage represents new_chat_members message sub type.
	NewChatMembersMessage MessageType = "new_chat_members"
	// LeftChatMemberMessage represents left_chat_member message sub type.
	LeftChatMemberMessage MessageType = "left_chat_member"
	// NewChatPhotoMessage represents new_chat_photo message sub type.
	NewChatPhotoMessage MessageType = "new_chat_photo"
	// DeleteChatPhotoMessage represents delete_chat_photo message sub type.
	DeleteChatPhotoMessage MessageType = "delete_chat_photo"
	// InvoiceMessage represents invoice message sub type.
	InvoiceMessage MessageType = "invoice"
	// SuccessfulPaymentMessage represents successful_payment message sub type.
	SuccessfulPaymentMessage MessageType = "successful_payment"
	// OtherMessage represents successful_payment message sub type.
	OtherMessage MessageType = "other"

	// TODO: connected website?
	// TODO: passport data?
)

// Message is returned by almost every method, and contains data about
// almost anything.
type Message struct {
	messageType           MessageType
	MessageID             int                `json:"message_id"`
	From                  *User              `json:"from"`
	Date                  int                `json:"date"`
	Chat                  *Chat              `json:"chat"`
	ForwardFrom           *User              `json:"forward_from"`
	ForwardFromChat       *Chat              `json:"forward_from_chat"`
	ForwardFromMessageID  int                `json:"forward_from_message_id"`
	ForwardDate           int                `json:"forward_date"`
	ReplyToMessage        *Message           `json:"reply_to_message"`
	EditDate              int                `json:"edit_date"`
	Text                  string             `json:"text"`
	Entities              *[]MessageEntity   `json:"entities"`
	Audio                 *Audio             `json:"audio"`
	Document              *Document          `json:"document"`
	Animation             *ChatAnimation     `json:"animation"`
	Game                  *Game              `json:"game"`
	Photo                 *[]PhotoSize       `json:"photo"`
	Sticker               *Sticker           `json:"sticker"`
	Video                 *Video             `json:"video"`
	VideoNote             *VideoNote         `json:"video_note"`
	Voice                 *Voice             `json:"voice"`
	Caption               string             `json:"caption"`
	Contact               *Contact           `json:"contact"`
	Location              *Location          `json:"location"`
	Venue                 *Venue             `json:"venue"`
	NewChatMembers        *[]User            `json:"new_chat_members"`
	LeftChatMember        *User              `json:"left_chat_member"`
	NewChatTitle          string             `json:"new_chat_title"`
	NewChatPhoto          *[]PhotoSize       `json:"new_chat_photo"`
	DeleteChatPhoto       bool               `json:"delete_chat_photo"`
	GroupChatCreated      bool               `json:"group_chat_created"`
	SuperGroupChatCreated bool               `json:"supergroup_chat_created"`
	ChannelChatCreated    bool               `json:"channel_chat_created"`
	MigrateToChatID       int64              `json:"migrate_to_chat_id"`
	MigrateFromChatID     int64              `json:"migrate_from_chat_id"`
	PinnedMessage         *Message           `json:"pinned_message"`
	Invoice               *Invoice           `json:"invoice"`
	SuccessfulPayment     *SuccessfulPayment `json:"successful_payment"`
}

// Path returns message path.
func (m *Message) Path() string {
	path := "/"
	path = path + m.Type().String()
	if m.IsCommand() {
		path = path + "/" + m.Command()
	}
	return path;
}

// Type returns message sub type.
func (m *Message) Type() MessageType {
	if m.IsCommand() {
		return CommandMessage
	}
	if m.Text != "" {
		return TextMessage
	}
	if m.Audio != nil {
		return AudioMessage
	}
	if m.Document != nil {
		return DocumentMessage
	}
	if m.Animation != nil {
		return AnimationMessage
	}
	if m.Photo != nil {
		return PhotoMessage
	}
	if m.Sticker != nil {
		return StickerMessage
	}
	if m.VideoNote != nil {
		return VideoNoteMessage
	}
	if m.Voice != nil {
		return VoiceMessage
	}
	if m.Contact != nil {
		return ContactMessage
	}
	if m.Location != nil {
		return LocationMessage
	}
	if m.LeftChatMember != nil {
		return LeftChatMemberMessage
	}
	if m.NewChatPhoto != nil {
		return NewChatPhotoMessage
	}
	if m.Invoice != nil {
		return InvoiceMessage
	}
	if m.SuccessfulPayment != nil {
		return SuccessfulPaymentMessage
	}
	return OtherMessage
}

// Time converts the message timestamp into a Time.
func (m *Message) Time() time.Time {
	return time.Unix(int64(m.Date), 0)
}

// IsCommand returns true if message starts with a "bot_command" entity.
func (m *Message) IsCommand() bool {
	if m.Entities == nil || len(*m.Entities) == 0 {
		return false
	}

	entity := (*m.Entities)[0]
	return entity.Offset == 0 && entity.Type == "bot_command"
}

// Command checks if the message was a command and if it was, returns the
// command. If the Message was not a command, it returns an empty string.
//
// If the command contains the at name syntax, it is removed. Use
// CommandWithAt() if you do not want that.
func (m *Message) Command() string {
	command := m.CommandWithAt()

	if i := strings.Index(command, "@"); i != -1 {
		command = command[:i]
	}

	return command
}

// CommandWithAt checks if the message was a command and if it was, returns the
// command. If the Message was not a command, it returns an empty string.
//
// If the command contains the at name syntax, it is not removed. Use Command()
// if you want that.
func (m *Message) CommandWithAt() string {
	if !m.IsCommand() {
		return ""
	}

	// IsCommand() checks that the message begins with a bot_command entity
	entity := (*m.Entities)[0]
	return m.Text[1:entity.Length]
}

// CommandArguments checks if the message was a command and if it was,
// returns all text after the command name. If the Message was not a
// command, it returns an empty string.
//
// Note: The first character after the command name is omitted:
// - "/foo bar baz" yields "bar baz", not " bar baz"
// - "/foo-bar baz" yields "bar baz", too
// Even though the latter is not a command conforming to the spec, the API
// marks "/foo" as command entity.
func (m *Message) CommandArguments() string {
	if !m.IsCommand() {
		return ""
	}

	// IsCommand() checks that the message begins with a bot_command entity
	entity := (*m.Entities)[0]
	if len(m.Text) == entity.Length {
		return "" // The command makes up the whole message
	}

	return m.Text[entity.Length+1:]
}

// MessageEntity contains information about data in a Message.
type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	URL    string `json:"url"`  // optional
	User   *User  `json:"user"` // optional
}

// ParseURL attempts to parse a URL contained within a MessageEntity.
func (entity MessageEntity) ParseURL() (*url.URL, error) {
	if entity.URL == "" {
		return nil, errors.New("Bad url")
	}

	return url.Parse(entity.URL)
}
