package telegram

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

// ChatID identifies the chat a request is for: a numeric chat ID, or the
// username of a public channel or supergroup.
//
// Client methods and request constructors take numeric IDs, which cover
// nearly every use. To address a chat by username instead, set the
// field of the request and send it with Client.Do:
//
//	req := telegram.NewSendMessage(0, telegram.Text("Hello, channel"))
//	req.ChatID = telegram.ChatUsername("mychannel")
//	err := client.Do(ctx, req, nil)
type ChatID struct {
	ID       int64
	Username string // without the leading "@"
}

// ChatUsername returns the ChatID of a public channel or supergroup. The
// leading "@" is optional.
func ChatUsername(username string) ChatID {
	return ChatID{Username: strings.TrimPrefix(username, "@")}
}

// IsZero reports whether c is unset.
func (c ChatID) IsZero() bool { return c.ID == 0 && c.Username == "" }

// String returns "@username" or the numeric ID.
func (c ChatID) String() string {
	if c.Username != "" {
		return "@" + c.Username
	}
	return strconv.FormatInt(c.ID, 10)
}

// MarshalJSON encodes c as Telegram expects: a number, or "@username".
func (c ChatID) MarshalJSON() ([]byte, error) {
	if c.Username != "" {
		return json.Marshal("@" + c.Username)
	}
	return strconv.AppendInt(nil, c.ID, 10), nil
}

// UnmarshalJSON decodes a number or "@username".
func (c *ChatID) UnmarshalJSON(b []byte) error {
	b = bytes.TrimSpace(b)
	if len(b) > 0 && b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		*c = ChatUsername(s)
		return nil
	}
	*c = ChatID{}
	if string(b) == "null" {
		return nil
	}
	return json.Unmarshal(b, &c.ID)
}
