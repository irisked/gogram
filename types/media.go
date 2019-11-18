package types

// File contains information about a file to download from domain.
type File struct {
	FileID   string `json:"file_id"`
	FileSize int    `json:"file_size"` // optional
	FilePath string `json:"file_path"` // optional
}

// Audio contains information about audio.
type Audio struct {
	FileID    string     `json:"file_id"`
	Duration  int        `json:"duration"`
	Performer string     `json:"performer"` // optional
	Title     string     `json:"title"`     // optional
	MimeType  string     `json:"mime_type"` // optional
	FileSize  int        `json:"file_size"` // optional
	Thumb     *PhotoSize `json:"thumb"`
}

// Animation is a GIF animation demonstrating the game.
type Animation struct {
	FileID   string     `json:"file_id"`
	Thumb    *PhotoSize `json:"thumb"`
	FileName string     `json:"file_name"`
	MimeType string     `json:"mime_type"`
	FileSize int        `json:"file_size"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Duration int        `json:"duration"`
}

// Document contains information about a document.
type Document struct {
	FileID    string     `json:"file_id"`
	Thumbnail *PhotoSize `json:"thumb"`     // optional
	FileName  string     `json:"file_name"` // optional
	MimeType  string     `json:"mime_type"` // optional
	FileSize  int        `json:"file_size"` // optional
}

// Sticker contains information about a sticker.
type Sticker struct {
	FileID    string     `json:"file_id"`
	Width     int        `json:"width"`
	Height    int        `json:"height"`
	Thumbnail *PhotoSize `json:"thumb"`     // optional
	Emoji     string     `json:"emoji"`     // optional
	FileSize  int        `json:"file_size"` // optional
	SetName   string     `json:"set_name"`  // optional
}

// ChatAnimation contains information about an animation.
type ChatAnimation struct {
	FileID    string     `json:"file_id"`
	Width     int        `json:"width"`
	Height    int        `json:"height"`
	Duration  int        `json:"duration"`
	Thumbnail *PhotoSize `json:"thumb"`     // optional
	FileName  string     `json:"file_name"` // optional
	MimeType  string     `json:"mime_type"` // optional
	FileSize  int        `json:"file_size"` // optional
}

// Video contains information about a video.
type Video struct {
	FileID    string     `json:"file_id"`
	Width     int        `json:"width"`
	Height    int        `json:"height"`
	Duration  int        `json:"duration"`
	Thumbnail *PhotoSize `json:"thumb"`     // optional
	MimeType  string     `json:"mime_type"` // optional
	FileSize  int        `json:"file_size"` // optional
}

// VideoNote contains information about a video.
type VideoNote struct {
	FileID    string     `json:"file_id"`
	Length    int        `json:"length"`
	Duration  int        `json:"duration"`
	Thumbnail *PhotoSize `json:"thumb"`     // optional
	FileSize  int        `json:"file_size"` // optional
}

// PhotoSize contains information about photos.
type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size"` // optional
}

// Voice contains information about a voice.
type Voice struct {
	FileID   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"` // optional
	FileSize int    `json:"file_size"` // optional
}
