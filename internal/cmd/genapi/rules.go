package main

import (
	"go/token"
	"regexp"
	"strings"
	"unicode"
)

// initialisms are words spelled in capitals in Go names.
var initialisms = map[string]string{
	"api": "API", "html": "HTML", "http": "HTTP", "https": "HTTPS", "id": "ID", "ids": "IDs",
	"ip": "IP", "json": "JSON", "ttl": "TTL", "uri": "URI", "url": "URL", "utc": "UTC",
}

// skipTypes are spec types implemented by hand.
var skipTypes = map[string]bool{
	"InputFile":                true, // input_file.go
	"MaybeInaccessibleMessage": true, // decoded as *Message; see Message.IsAccessible
}

// fieldTypes overrides the Go type of a field, keyed by "Owner.field"
// or "*.field" for every owner.
var fieldTypes = map[string]string{
	// An empty query is meaningful, so it must be distinguishable from unset.
	"InlineKeyboardButton.switch_inline_query":              "*string",
	"InlineKeyboardButton.switch_inline_query_current_chat": "*string",
	"*.allowed_updates": "[]UpdateType",
}

// unionNames names unions that the spec writes inline as "A or B or C"
// without naming them, keyed by "Owner.field".
var unionNames = map[string]string{
	"InputRichMessageMedia.media": "InputRichMedia",
}

// replyMarkupTypes is the inline union of reply_markup, implemented by
// hand as ReplyMarkup.
var replyMarkupTypes = "InlineKeyboardMarkup | ReplyKeyboardMarkup | ReplyKeyboardRemove | ForceReply"

// optionNames renames the option of a field, keyed by "method.field",
// where the default name would clash with an option of another type.
// (reply_markup fields that accept only inline keyboards get no option
// of their own: WithInlineKeyboard covers them.)
var optionNames = map[string]string{
	"sendGift.text":                  "GiftText",
	"giftPremiumSubscription.text":   "GiftText",
	"sendPoll.description":           "PollDescription",
	"getBusinessAccountGifts.offset": "GiftsOffset",
	"getChatGifts.offset":            "GiftsOffset",
	"getUserGifts.offset":            "GiftsOffset",
}

// positional makes optional parameters positional arguments, keyed by
// "method.field": parameters that are nearly always passed. They are
// still omitted from the request when empty.
var positional = map[string]bool{
	"editMessageText.text":          true, // optional only when editing rich content
	"editEphemeralMessageText.text": true,
	"sendMessageDraft.text":         true,
}

// extraOptions are convenience options built on another field.
var extraOptions = []extraOption{
	{
		Name:  "ReplyToMessageID",
		Field: "reply_parameters",
		Arg:   "messageID int",
		Doc:   "replies to the message with the given ID in the same chat. Use WithReplyParameters for more control.",
		Set:   "r.ReplyParameters = &ReplyParameters{MessageID: o.v}",
	},
	{
		Name:  "DisableWebPagePreview",
		Field: "link_preview_options",
		Doc:   "disables link previews. Use WithLinkPreviewOptions for more control.",
		Set:   "r.LinkPreviewOptions = &LinkPreviewOptions{IsDisabled: true}",
	},
	{
		Name:  "InlineKeyboard",
		Field: "reply_markup",
		Arg:   "keyboard *InlineKeyboardMarkup",
		Doc:   "attaches an inline keyboard. Unlike WithReplyMarkup, it also works with methods that accept only inline keyboards, such as the edit methods.",
		Set:   "if o.v != nil { r.ReplyMarkup = o.v }",
	},
}

type extraOption struct {
	Name  string // option name without "With"
	Field string // json name of the field it sets
	Arg   string // "name type", or "" for a flag
	Doc   string // continues "With<Name> "
	Set   string // statement applying o.v to r
}

// chatShortcuts are methods offered on gogram.Update for the update's
// chat, besides the send methods, which become Reply shortcuts.
var chatShortcuts = []string{
	"sendChatAction", "deleteMessage", "deleteMessages", "setMessageReaction",
	"banChatMember", "unbanChatMember", "restrictChatMember", "promoteChatMember",
	"approveChatJoinRequest", "declineChatJoinRequest",
	"pinChatMessage", "unpinChatMessage", "unpinAllChatMessages",
	"getChat", "getChatMember", "getChatAdministrators", "leaveChat",
}

// replyExclude are send methods that are not replies.
var replyExclude = map[string]bool{
	"sendChatAction": true,
}

var re64 = regexp.MustCompile(`52 significant bits|64-bit`)

// is64 reports whether an Integer field needs 64 bits: identifiers of
// users and chats, Unix times and anything the spec says can exceed 32
// bits.
func is64(f apiField) bool {
	n := f.Name
	for _, suffix := range []string{"user_id", "chat_id", "user_ids", "chat_ids", "date"} {
		if n == suffix || strings.HasSuffix(n, "_"+suffix) {
			return true
		}
	}
	return re64.MatchString(f.Description)
}

// exported converts a snake_case spec name into an exported Go name.
func exported(snake string) string {
	var b strings.Builder
	for _, w := range strings.Split(snake, "_") {
		if w == "" {
			continue
		}
		if s, ok := initialisms[w]; ok {
			b.WriteString(s)
			continue
		}
		b.WriteString(strings.ToUpper(w[:1]) + w[1:])
	}
	return b.String()
}

// unexported converts a snake_case spec name into a Go parameter name.
func unexported(snake string) string {
	first, rest, _ := strings.Cut(snake, "_")
	name := first + exported(rest)
	if name == "type" {
		return "typ"
	}
	if token.IsKeyword(name) || reservedArgs[name] {
		return name + "Arg"
	}
	return name
}

// reservedArgs are identifiers the generated method bodies use.
var reservedArgs = map[string]bool{
	"ctx": true, "opts": true, "req": true, "res": true, "raw": true,
	"err": true, "c": true, "o": true, "r": true, "u": true, "target": true,
}

// typeName converts a CamelCase spec type name into a Go type name,
// fixing initialisms: "MessageId" becomes "MessageID".
func typeName(s string) string {
	var b strings.Builder
	start := 0
	flush := func(end int) {
		w := s[start:end]
		if up, ok := initialisms[strings.ToLower(w)]; ok {
			w = up
		}
		b.WriteString(w)
		start = end
	}
	for i := 1; i < len(s); i++ {
		if unicode.IsUpper(rune(s[i])) {
			flush(i)
		}
	}
	flush(len(s))
	return b.String()
}

// methodName converts a spec method name into a Go name: "sendMessage"
// becomes "SendMessage".
func methodName(s string) string {
	return typeName(strings.ToUpper(s[:1]) + s[1:])
}

var reTag = regexp.MustCompile(`(?:always|must be) [“"]?\*?([a-z0-9_]+)`)

// docText cleans a spec description for a Go comment.
func docText(s string) string {
	s = strings.TrimPrefix(s, "Optional. ")
	return strings.ReplaceAll(s, "\n", " ")
}

// typeDoc rewrites the first sentence of a type description so it starts
// with the type name, as Go doc comments do.
func typeDoc(name string, lines []string) []string {
	if len(lines) == 0 {
		return []string{name + " is a Bot API object."}
	}
	out := append([]string(nil), lines...)
	first := out[0]
	switch {
	case strings.HasPrefix(first, "This object represents "):
		first = name + " represents " + strings.TrimPrefix(first, "This object represents ")
	case strings.HasPrefix(first, "This object "):
		first = name + " " + strings.TrimPrefix(first, "This object ")
	case strings.HasPrefix(first, "Represents "):
		first = name + " represents " + strings.TrimPrefix(first, "Represents ")
	case strings.HasPrefix(first, "Describes "):
		first = name + " describes " + strings.TrimPrefix(first, "Describes ")
	case strings.HasPrefix(first, "Contains "):
		first = name + " contains " + strings.TrimPrefix(first, "Contains ")
	default:
		first = name + ": " + first
	}
	out[0] = first
	return out
}
