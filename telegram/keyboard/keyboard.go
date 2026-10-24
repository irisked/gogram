package keyboard

// Markup is implemented by any reply-markup type (inline keyboard,
// reply keyboard, etc.) that can be serialized for the Telegram API.
type Markup interface {
	Serialize() []byte
}
