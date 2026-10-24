package telegram

import "strings"

// UpdateType is the kind of an Update: the name of its one set field,
// such as "message" or "callback_query". The Update* constants list
// them.
type UpdateType string

func (ut UpdateType) String() string { return string(ut) }

// IsFromInlineMessage reports whether the button was attached to a
// message sent in inline mode.
func (cq *CallbackQuery) IsFromInlineMessage() bool {
	return cq.InlineMessageID != ""
}

// Route returns the part of Data before ":", as set by CallbackButton.
func (cq *CallbackQuery) Route() string {
	route, _, _ := strings.Cut(cq.Data, ":")
	return route
}

// Args returns the comma-separated arguments after ":" in Data, as set
// by CallbackButton.
func (cq *CallbackQuery) Args() []string {
	_, args, ok := strings.Cut(cq.Data, ":")
	if !ok || args == "" {
		return nil
	}
	return strings.Split(args, ",")
}

// Target returns the message the pressed button belongs to, for use
// with the edit methods. ok is false if that message is unknown.
func (cq *CallbackQuery) Target() (target MessageTarget, ok bool) {
	if cq.IsFromInlineMessage() {
		return InlineMessage(cq.InlineMessageID), true
	}
	if cq.Message == nil || cq.Message.Chat == nil {
		return MessageTarget{}, false
	}
	return ChatMessage(cq.Message.Chat.ID, cq.Message.MessageID), true
}

// IsPrivate reports whether the chat is a private chat with a user.
func (c *Chat) IsPrivate() bool { return c.Type == "private" }

// IsGroup reports whether the chat is a basic group.
func (c *Chat) IsGroup() bool { return c.Type == "group" }

// IsSupergroup reports whether the chat is a supergroup.
func (c *Chat) IsSupergroup() bool { return c.Type == "supergroup" }

// IsChannel reports whether the chat is a channel.
func (c *Chat) IsChannel() bool { return c.Type == "channel" }

// String returns the username, or the full name if there is none.
func (u *User) String() string {
	if u.Username != "" {
		return u.Username
	}
	if u.LastName != "" {
		return u.FirstName + " " + u.LastName
	}
	return u.FirstName
}
