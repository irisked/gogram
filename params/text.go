package params

// Text text data.
type Text struct {
	Text      string
	ParseMode string
}

// SimpleText simple text.
func SimpleText(text string) *Text {
	return &Text{Text: text}
}

// MarkdownText markdown text.
func MarkdownText(text string) *Text {
	return &Text{Text: text, ParseMode: "markdown"}
}

// HTMLText html text.
func HTMLText(text string) *Text {
	return &Text{Text: text, ParseMode: "HTML"}
}
