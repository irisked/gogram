package telegram

// Parse modes for formatted text.
const (
	ParseModeMarkdown   = "Markdown"
	ParseModeMarkdownV2 = "MarkdownV2"
	ParseModeHTML       = "HTML"
)

// TextData is message text together with its parse mode.
type TextData struct {
	Text      string
	ParseMode string
}

// Text is plain text without formatting.
func Text(text string) TextData {
	return TextData{Text: text}
}

// Markdown is text formatted with legacy Markdown.
func Markdown(text string) TextData {
	return TextData{Text: text, ParseMode: ParseModeMarkdown}
}

// MarkdownV2 is text formatted with MarkdownV2.
func MarkdownV2(text string) TextData {
	return TextData{Text: text, ParseMode: ParseModeMarkdownV2}
}

// HTML is text formatted with HTML tags.
func HTML(text string) TextData {
	return TextData{Text: text, ParseMode: ParseModeHTML}
}

// Chat actions for SendChatAction: the status shown to the users of a
// chat while the bot prepares a reply.
const (
	ChatActionTyping          = "typing"
	ChatActionUploadPhoto     = "upload_photo"
	ChatActionRecordVideo     = "record_video"
	ChatActionUploadVideo     = "upload_video"
	ChatActionRecordVoice     = "record_voice"
	ChatActionUploadVoice     = "upload_voice"
	ChatActionUploadDocument  = "upload_document"
	ChatActionChooseSticker   = "choose_sticker"
	ChatActionFindLocation    = "find_location"
	ChatActionRecordVideoNote = "record_video_note"
	ChatActionUploadVideoNote = "upload_video_note"
)
