package method

import (
	"os"

	"github.com/irisked/gogram/params"
	"github.com/irisked/gogram/types/markup/keyboard"
)

// GetFile contains information about getFile telegram api method.
type GetFile struct {
	FileID string `json:"file_id"`
	SimpleMethod
}

// NewGetFile creates GetFile.
func NewGetFile(fileID string) *GetFile {
	method := new(GetFile)
	method.FileID = fileID
	return method
}

// Method returns telegram api method endpoint.
func (m *GetFile) Method() string {
	return "getFile"
}

// SendPhoto contains information about sendPhoto telegram api method.
type SendPhoto struct {
	ChatID              int64             `json:"chat_id" form:"chat_id"`
	File                *os.File          `json:"-"`
	Photo               string            `json:"photo" form:"photo"`
	Caption             string            `json:"caption,omitempty" form:"caption,omitempty"`
	ParseMode           string            `json:"parse_mode,omitempty" form:"parse_mode,omitempty"`
	DisableNotification bool              `json:"disable_notification,omitempty" form:"disable_notification,omitempty"`
	ReplyToMessageID    int               `json:"reply_to_message_id,omitempty" form:"reply_to_message_id,omitempty"`
	ReplyMarkup         keyboard.Keyboard `json:"reply_markup,omitempty" form:"reply_markup,omitempty"`
}

// Method returns telegram api method endpoint.
func (m *SendPhoto) Method() string {
	return "sendPhoto"
}

// IsMultipart returns true if method contains file to upload
// and false otherwise.
func (m *SendPhoto) IsMultipart() bool {
	return m.File != nil
}

// NewSendPhoto creates SendPhoto.
func NewSendPhoto(chatID int64, photo *params.PhotoConfig, options ...SendPhotoOption) *SendPhoto {
	method := new(SendPhoto)
	method.ChatID = chatID
	method.Photo = photo.Reference
	method.File = photo.Local
	method.Caption = photo.Caption
	method.ParseMode = photo.ParseMode
	for _, option := range options {
		option.ApplySendPhotoOption(method)
	}
	return method
}

// SendPhotoOption its a functional options interface.
type SendPhotoOption interface {
	ApplySendPhotoOption(*SendPhoto)
}

// SendAudio contains information about sendMessage telegram api method.
type SendAudio struct {
	ChatID              int64             `json:"chat_id" form:"chat_id"`
	File                *os.File          `json:"-" form:"file"`
	ThumbFile           *os.File          `json:"-" form:"thumb"`
	Audio               string            `json:"audio"`
	Thumb               string            `json:"thumb,omitempty" form:"thumb,omitempty"`
	Caption             string            `json:"caption,omitempty" form:"caption,omitempty"`
	Duration            int               `json:"duration,omitempty" form:"caption,omitempty"`
	Performer           string            `json:"performer,omitempty" form:"performer,omitempty"`
	Title               string            `json:"title,omitempty" form:"title,omitempty"`
	ParseMode           string            `json:"parse_mode,omitempty" form:"parse_mode,omitempty"`
	DisableNotification bool              `json:"disable_notification,omitempty" form:"disable_notification,omitempty"`
	ReplyToMessageID    int               `json:"reply_to_message_id,omitempty" form:"reply_to_message_id,omitempty"`
	ReplyMarkup         keyboard.Keyboard `json:"reply_markup,omitempty" form:"reply_markup,omitempty"`
}

// Method returns telegram api method endpoint.
func (m *SendAudio) Method() string {
	return "sendAudio"
}

// IsMultipart returns true if method contains a files to upload
// and false otherwise.
func (m *SendAudio) IsMultipart() bool {
	return m.File != nil && m.ThumbFile != nil
}

// NewSendAudio creates SendAudio.
func NewSendAudio(chatID int64, audio *params.AudioConfig, options ...SendAudioOption) *SendAudio {
	method := new(SendAudio)
	method.ChatID = chatID
	method.File = audio.Local
	method.Audio = audio.Reference
	method.Thumb = audio.Thumb
	method.ThumbFile = audio.ThumbFile
	method.Caption = audio.Caption
	method.ParseMode = audio.ParseMode
	method.Duration = audio.Duration
	method.Performer = audio.Performer
	method.Title = audio.Title
	for _, opt := range options {
		opt.ApplySendAudioOption(method)
	}
	return method
}

// SendAudioOption its a functional options interface.
type SendAudioOption interface {
	ApplySendAudioOption(*SendAudio)
}

// SendDocument contains information about sendDocument telegram api method.
type SendDocument struct {
	ChatID              int64             `json:"chat_id" form:"chat_id"`
	File                *os.File          `json:"-" form:"file"`
	ThumbFile           *os.File          `json:"-" form:"thumb"`
	Document            string            `json:"document" form:"document"`
	Thumb               string            `json:"thumb,omitempty" form:"thumb,omitempty"`
	Caption             string            `json:"caption,omitempty" form:"caption,omitempty"`
	ParseMode           string            `json:"parse_mode,omitempty" form:"parse_mode,omitempty"`
	DisableNotification bool              `json:"disable_notification,omitempty" form:"disable_notification,omitempty"`
	ReplyToMessageID    int               `json:"reply_to_message_id,omitempty" form:"reply_to_message_id,omitempty"`
	ReplyMarkup         keyboard.Keyboard `json:"reply_markup,omitempty" form:"reply_markup,omitempty"`
}

// Method returns telegram api method endpoint.
func (m *SendDocument) Method() string {
	return "sendDocument"
}

// IsMultipart returns true if method contains file to upload
// and false otherwise.
func (m *SendDocument) IsMultipart() bool {
	return m.File != nil && m.ThumbFile != nil
}

// NewSendDocument creates SendDocument.
func NewSendDocument(chatID int64, document *params.DocumentConfig, options ...SendDocumentOption) *SendDocument {
	method := new(SendDocument)
	method.ChatID = chatID
	method.File = document.Local
	method.Document = document.Reference
	method.ThumbFile = document.ThumbFile
	method.Thumb = document.Thumb
	method.Caption = document.Caption
	method.ParseMode = document.ParseMode
	for _, opt := range options {
		opt.ApplySendDocumentOption(method)
	}
	return method
}

// SendDocumentOption its a functional options interface.
type SendDocumentOption interface {
	ApplySendDocumentOption(*SendDocument)
}

// SendVideo contains information about sendVideo telegram api method.
type SendVideo struct {
	ChatID              int64             `json:"chat_id" form:"chat_id"`
	File                *os.File          `json:"-" form:"file"`
	ThumbFile           *os.File          `json:"-" form:"thumb"`
	Video               string            `json:"video" form:"video"`
	Thumb               string            `json:"thumb,omitempty" form:"thumb,omitempty"`
	Caption             string            `json:"caption,omitempty" form:"caption,omitempty"`
	Width               int               `json:"width,omitempty" form:"width,omitempty"`
	Height              int               `json:"height,omitempty" form:"height,omitempty"`
	Duration            int               `json:"duration,omitempty" form:"duration,omitempty"`
	SupportsStreaming   bool              `json:"supports_streaming,omitempty" form:"supports_streaming,omitempty"`
	ParseMode           string            `json:"parse_mode,omitempty" form:"parse_mode,omitempty"`
	DisableNotification bool              `json:"disable_notification,omitempty" form:"disable_notification,omitempty"`
	ReplyToMessageID    int               `json:"reply_to_message_id,omitempty" form:"reply_to_message_id,omitempty"`
	ReplyMarkup         keyboard.Keyboard `json:"reply_markup,omitempty" form:"reply_markup,omitempty"`
}

// Method returns telegram api method endpoint.
func (m *SendVideo) Method() string {
	return "sendVideo"
}

// IsMultipart returns true if method contains file to upload
// and false otherwise.
func (m *SendVideo) IsMultipart() bool {
	return m.File != nil && m.ThumbFile != nil
}

// NewSendVideo creates SendVideo.
func NewSendVideo(chatID int64, video *params.VideoConfig, options ...SendVideoOption) *SendVideo {
	method := new(SendVideo)
	method.ChatID = chatID
	method.File = video.Local
	method.Video = video.Reference
	method.Thumb = video.Thumb
	method.ThumbFile = video.ThumbFile
	method.Caption = video.Caption
	method.ParseMode = video.ParseMode
	method.Duration = video.Duration
	method.Height = video.Height
	method.Width = video.Width
	method.SupportsStreaming = video.SupportsStreaming
	for _, opt := range options {
		opt.ApplySendVideoOption(method)
	}
	return method
}

// SendVideoOption its a functional options interface.
type SendVideoOption interface {
	ApplySendVideoOption(*SendVideo)
}

// SendAnimation contains information about sendAnimation telegram api method.
type SendAnimation struct {
	ChatID              int64             `json:"chat_id" form:"chat_id"`
	File                *os.File          `json:"-" form:"file"`
	ThumbFile           *os.File          `json:"-" form:"thumb"`
	Animation           string            `json:"animation" form:"animation"`
	Caption             string            `json:"caption,omitempty" form:"caption,omitempty"`
	Thumb               string            `json:"thumb,omitempty" form:"thumb,omitempty"`
	Width               int               `json:"width,omitempty" form:"width,omitempty"`
	Height              int               `json:"height,omitempty" form:"height,omitempty"`
	Duration            int               `form:"duration,omitempty" form:"duration,omitempty"`
	ParseMode           string            `json:"parse_mode,omitempty" form:"parse_mode,omitempty"`
	DisableNotification bool              `json:"disable_notification,omitempty" form:"disable_notification,omitempty"`
	ReplyToMessageID    int               `json:"reply_to_message_id,omitempty" form:"reply_to_message_id,omitempty"`
	ReplyMarkup         keyboard.Keyboard `json:"reply_markup,omitempty" form:"reply_markup,omitempty"`
}

// Method returns telegram api method endpoint.
func (m *SendAnimation) Method() string {
	return "sendAnimation"
}

// IsMultipart returns true if method contains file to upload
// and false otherwise.
func (m *SendAnimation) IsMultipart() bool {
	return m.File != nil && m.ThumbFile != nil
}

// NewSendAnimation creates SendAnimation.
func NewSendAnimation(chatID int64, animation *params.AnimationConfig, options ...SendAnimationOption) *SendAnimation {
	method := new(SendAnimation)
	method.ChatID = chatID
	method.File = animation.Local
	method.Animation = animation.Reference
	method.Thumb = animation.Thumb
	method.ThumbFile = animation.ThumbFile
	method.Caption = animation.Caption
	method.ParseMode = animation.ParseMode
	method.Duration = animation.Duration
	method.Width = animation.Width
	method.Height = animation.Height
	for _, opt := range options {
		opt.ApplySendAnimationOption(method)
	}
	return method
}

// SendAnimationOption its a functional options interface.
type SendAnimationOption interface {
	ApplySendAnimationOption(*SendAnimation)
}

// SendVoice contains information about sendVoice telegram api method.
type SendVoice struct {
	ChatID              int64             `json:"chat_id" form:"chat_id"`
	File                *os.File          `json:"-" form:"file"`
	Voice               string            `json:"audio" form:"audio"`
	Caption             string            `json:"caption,omitempty" form:"caption,omitempty"`
	Duration            int               `json:"duration,omitempty" form:"duration,omitempty"`
	ParseMode           string            `json:"parse_mode,omitempty" form:"parse_mode,omitempty"`
	DisableNotification bool              `json:"disable_notification,omitempty" form:"disable_notification,omitempty"`
	ReplyToMessageID    int               `json:"reply_to_message_id,omitempty" form:"reply_to_message_id,omitempty"`
	ReplyMarkup         keyboard.Keyboard `json:"reply_markup,omitempty" form:"reply_markup,omitempty"`
}

// Method returns telegram api method endpoint.
func (m *SendVoice) Method() string {
	return "sendVoice"
}

// IsMultipart returns true if method contains file to upload
// and false otherwise.
func (m *SendVoice) IsMultipart() bool {
	return m.File != nil
}

// NewSendVoice creates SendVoice.
func NewSendVoice(chatID int64, voice *params.VoiceConfig, options ...SendVoiceOption) *SendVoice {
	method := new(SendVoice)
	method.ChatID = chatID
	method.File = voice.Local
	method.Voice = voice.Reference
	method.Caption = voice.Caption
	method.ParseMode = voice.ParseMode
	method.Duration = voice.Duration
	for _, opt := range options {
		opt.ApplySendVoiceOption(method)
	}
	return method
}

// SendVoiceOption its a functional options interface.
type SendVoiceOption interface {
	ApplySendVoiceOption(*SendVoice)
}

// SendVideoNote contains information about sendVideoNote telegram api method.
type SendVideoNote struct {
	ChatID              int64             `json:"chat_id" form:"chat_id"`
	File                *os.File          `json:"-" form:"file"`
	ThumbFile           *os.File          `json:"-" form:"thumb"`
	VideoNote           string            `json:"video_note" form:"video_note"`
	Thumb               string            `json:"thumb,omitempty" form:"thumb,omitempty"`
	Length              int               `json:"length,omitempty" form:"length,omitempty"`
	Duration            int               `json:"duration,omitempty" form:"duration,omitempty"`
	DisableNotification bool              `json:"disable_notification,omitempty" form:"disable_notification,omitempty"`
	ReplyToMessageID    int               `json:"reply_to_message_id,omitempty" form:"reply_to_message_id,omitempty"`
	ReplyMarkup         keyboard.Keyboard `json:"reply_markup,omitempty" form:"reply_markup,omitempty"`
}

// Method returns telegram api method endpoint.
func (m *SendVideoNote) Method() string {
	return "sendVideoNote"
}

// IsMultipart returns true if method contains file to upload
// and false otherwise.
func (m *SendVideoNote) IsMultipart() bool {
	return m.File != nil && m.ThumbFile != nil
}

// NewSendVideoNote creates SendVideoNote.
func NewSendVideoNote(chatID int64, videoNote *params.VideoNoteConfig, options ...SendVideoNoteOption) *SendVideoNote {
	method := new(SendVideoNote)
	method.ChatID = chatID
	method.File = videoNote.Local
	method.VideoNote = videoNote.Reference
	method.Thumb = videoNote.Thumb
	method.ThumbFile = videoNote.ThumbFile
	method.Length = videoNote.Length
	method.Duration = videoNote.Duration
	for _, opt := range options {
		opt.ApplySendVideoNoteOption(method)
	}
	return method
}

// SendVideoNoteOption its a functional options interface.
type SendVideoNoteOption interface {
	ApplySendVideoNoteOption(*SendVideoNote)
}

// SendMediaGroup contains information about sendMediaGroup telegram api method.
type SendMediaGroup struct {
	Files               []*os.File `json:"-" form:"file"`
	ThumbFiles          []*os.File `json:"-" form:"thumb"`
	ChatID              int64      `json:"chat_id" form:"chat_id"`
	Media               []string   `json:"media" form:"media"`
	DisableNotification bool       `json:"disable_notification,omitempty" form:"disable_notification,omitempty"`
	ReplyToMessageID    int        `json:"reply_to_message_id,omitempty" form:"reply_to_message_id,omitempty"`
}

// Method returns telegram api method endpoint.
func (m *SendMediaGroup) Method() string {
	return "sendMediaGroup"
}

// IsMultipart returns true if method contains file to upload
// and false otherwise.
func (m *SendMediaGroup) IsMultipart() bool {
	return false
}

// NewSendMediaGroup creates SendMediaGroup.
func NewSendMediaGroup(chatID int64, medias []params.MediaGroupConfig, options ...SendMediaGroupOption) *SendMediaGroup {
	method := new(SendMediaGroup)
	method.ChatID = chatID
	for _, media := range medias {
		stringified, files := media.GetMediaGroupConfig()
		method.Media = append(method.Media, stringified)
		if file, ok := files["file"]; ok {
			method.Files = append(method.Files, file)
		}
		if file, ok := files["thumb"]; ok {
			method.Files = append(method.ThumbFiles, file)
		}
	}
	for _, opt := range options {
		opt.ApplySendMediaGroupOption(method)
	}
	return method
}

// SendMediaGroupOption its a functional options interface.
type SendMediaGroupOption interface {
	ApplySendMediaGroupOption(*SendMediaGroup)
}

// EditMessageMedia contains information about editMessageMedia telegram api method.
type EditMessageMedia struct {
	File            *os.File          `json:"-" form:"file"`
	ThumbFile       *os.File          `json:"-" form:"thumb"`
	ChatID          int64             `json:"chat_id,omitempty" form:"chat_id,omitempty"`
	MessageID       int               `json:"message_id,omitempty" form:"message_id,omitempty"`
	InlineMessageID string            `json:"inline_message_id,omitempty" form:"inline_message_id,omitempty"`
	Media           string            `json:"media" form:"media"`
	ReplyMarkup     keyboard.Keyboard `json:"reply_markup,omitempty"`
}

// Method returns telegram api method endpoint.
func (m *EditMessageMedia) Method() string {
	return "editMessageMedia"
}

// IsMultipart returns true if method contains file to upload
// and false otherwise.
func (m *EditMessageMedia) IsMultipart() bool {
	return m.File != nil && m.ThumbFile != nil
}

// NewEditMessageMedia creates EditMessageMedia.
func NewEditMessageMedia(message *params.MessageID, media params.InputMediaConfig, options ...EditMessageMediaOption) *EditMessageMedia {
	method := new(EditMessageMedia)
	method.ChatID = message.ChatID
	method.MessageID = message.MessageID
	method.InlineMessageID = message.InlineMessageID
	stringified, files := media.GetInputMediaConfig()
	if file, ok := files["file"]; ok {
		method.File = file
	}
	if file, ok := files["thumb"]; ok {
		method.ThumbFile = file
	}
	method.Media = stringified
	for _, opt := range options {
		opt.ApplyEditMessageMediaOption(method)
	}
	return method
}

// EditMessageMediaOption its a functional options interface.
type EditMessageMediaOption interface {
	ApplyEditMessageMediaOption(*EditMessageMedia)
}
