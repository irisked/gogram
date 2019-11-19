package args

// TextData text data.
type TextData struct {
	Text      string
	ParseMode string
}

// Text simple text.
func Text(text string) TextData {
	return TextData{Text: text}
}

// Markdown markdown text.
func Markdown(text string) TextData {
	return TextData{Text: text, ParseMode: "markdown"}
}

// HTML html text.
func HTML(text string) TextData {
	return TextData{Text: text, ParseMode: "HTML"}
}
