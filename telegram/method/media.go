package method

import (
	"os"

	"github.com/irisked/gogram/args"
	"github.com/irisked/gogram/telegram/keyboard"
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
	ReplyMarkup         keyboard.Markup `json:"reply_markup,omitempty" form:"reply_markup,omitempty"`
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
func NewSendPhoto(chatID int64, photo *args.PhotoConfig, opts ...Option[*SendPhoto]) *SendPhoto {
	method := new(SendPhoto)
	method.ChatID = chatID
	method.Photo = photo.Reference
	method.File = photo.Local
	method.Caption = photo.Caption
	method.ParseMode = photo.ParseMode
	for _, o := range opts {
		o(method)
	}
	return method
}

// SetCaption sets caption
func (m *SendPhoto) SetCaption(c string) {
	m.Caption = c
}

// SetParseMode sets parse_mode
func (m *SendPhoto) SetParseMode(m2 string) {
	m.ParseMode = m2
}

// SetDisableNotification sets disable_notification
func (m *SendPhoto) SetDisableNotification(v bool) {
	m.DisableNotification = v
}

// SetReplyToMessageID sets reply_to_message_id
func (m *SendPhoto) SetReplyToMessageID(id int) {
	m.ReplyToMessageID = id
}

// SetReplyMarkup sets reply_markup
func (m *SendPhoto) SetReplyMarkup(kb keyboard.Markup) {
	m.ReplyMarkup = kb
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
	ReplyMarkup         keyboard.Markup `json:"reply_markup,omitempty" form:"reply_markup,omitempty"`
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
func NewSendAudio(chatID int64, audio *args.AudioConfig, opts ...Option[*SendAudio]) *SendAudio {
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
	for _, o := range opts {
		o(method)
	}
	return method
}

// SetThumb sets thumb
func (m *SendAudio) SetThumb(t string) {
	m.Thumb = t
}

// SetCaption sets caption
func (m *SendAudio) SetCaption(c string) {
	m.Caption = c
}

// SetDuration sets duration
func (m *SendAudio) SetDuration(d int) {
	m.Duration = d
}

// SetPerformer sets performer
func (m *SendAudio) SetPerformer(p string) {
	m.Performer = p
}

// SetTitle sets title
func (m *SendAudio) SetTitle(t string) {
	m.Title = t
}

// SetParseMode sets parse_mode
func (m *SendAudio) SetParseMode(m2 string) {
	m.ParseMode = m2
}

// SetDisableNotification sets disable_notification
func (m *SendAudio) SetDisableNotification(v bool) {
	m.DisableNotification = v
}

// SetReplyToMessageID sets reply_to_message_id
func (m *SendAudio) SetReplyToMessageID(id int) {
	m.ReplyToMessageID = id
}

// SetReplyMarkup sets reply_markup
func (m *SendAudio) SetReplyMarkup(kb keyboard.Markup) {
	m.ReplyMarkup = kb
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
	ReplyMarkup         keyboard.Markup `json:"reply_markup,omitempty" form:"reply_markup,omitempty"`
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
func NewSendDocument(chatID int64, document *args.DocumentConfig, opts ...Option[*SendDocument]) *SendDocument {
	method := new(SendDocument)
	method.ChatID = chatID
	method.File = document.Local
	method.Document = document.Reference
	method.ThumbFile = document.ThumbFile
	method.Thumb = document.Thumb
	method.Caption = document.Caption
	method.ParseMode = document.ParseMode
	for _, o := range opts {
		o(method)
	}
	return method
}

// SetThumb sets thumb
func (m *SendDocument) SetThumb(t string) {
	m.Thumb = t
}

// SetCaption sets caption
func (m *SendDocument) SetCaption(c string) {
	m.Caption = c
}

// SetParseMode sets parse_mode
func (m *SendDocument) SetParseMode(m2 string) {
	m.ParseMode = m2
}

// SetDisableNotification sets disable_notification
func (m *SendDocument) SetDisableNotification(v bool) {
	m.DisableNotification = v
}

// SetReplyToMessageID sets reply_to_message_id
func (m *SendDocument) SetReplyToMessageID(id int) {
	m.ReplyToMessageID = id
}

// SetReplyMarkup sets reply_markup
func (m *SendDocument) SetReplyMarkup(kb keyboard.Markup) {
	m.ReplyMarkup = kb
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
	ReplyMarkup         keyboard.Markup `json:"reply_markup,omitempty" form:"reply_markup,omitempty"`
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
func NewSendVideo(chatID int64, video *args.VideoConfig, opts ...Option[*SendVideo]) *SendVideo {
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
	for _, o := range opts {
		o(method)
	}
	return method
}

// SetThumb sets thumb
func (m *SendVideo) SetThumb(t string) {
	m.Thumb = t
}

// SetCaption sets caption
func (m *SendVideo) SetCaption(c string) {
	m.Caption = c
}

// SetWidth sets width
func (m *SendVideo) SetWidth(w int) {
	m.Width = w
}

// SetHeight sets height
func (m *SendVideo) SetHeight(h int) {
	m.Height = h
}

// SetDuration sets duration
func (m *SendVideo) SetDuration(d int) {
	m.Duration = d
}

// SetSupportsStreaming sets supports_streaming
func (m *SendVideo) SetSupportsStreaming(s bool) {
	m.SupportsStreaming = s
}

// SetParseMode sets parse_mode
func (m *SendVideo) SetParseMode(m2 string) {
	m.ParseMode = m2
}

// SetDisableNotification sets disable_notification
func (m *SendVideo) SetDisableNotification(v bool) {
	m.DisableNotification = v
}

// SetReplyToMessageID sets reply_to_message_id
func (m *SendVideo) SetReplyToMessageID(id int) {
	m.ReplyToMessageID = id
}

// SetReplyMarkup sets reply_markup
func (m *SendVideo) SetReplyMarkup(kb keyboard.Markup) {
	m.ReplyMarkup = kb
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
	ReplyMarkup         keyboard.Markup `json:"reply_markup,omitempty" form:"reply_markup,omitempty"`
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
func NewSendAnimation(chatID int64, animation *args.AnimationConfig, opts ...Option[*SendAnimation]) *SendAnimation {
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
	for _, o := range opts {
		o(method)
	}
	return method
}

// SetThumb sets thumb
func (m *SendAnimation) SetThumb(t string) {
	m.Thumb = t
}

// SetCaption sets caption
func (m *SendAnimation) SetCaption(c string) {
	m.Caption = c
}

// SetWidth sets width
func (m *SendAnimation) SetWidth(w int) {
	m.Width = w
}

// SetHeight sets height
func (m *SendAnimation) SetHeight(h int) {
	m.Height = h
}

// SetDuration sets duration
func (m *SendAnimation) SetDuration(d int) {
	m.Duration = d
}

// SetParseMode sets parse_mode
func (m *SendAnimation) SetParseMode(m2 string) {
	m.ParseMode = m2
}

// SetDisableNotification sets disable_notification
func (m *SendAnimation) SetDisableNotification(v bool) {
	m.DisableNotification = v
}

// SetReplyToMessageID sets reply_to_message_id
func (m *SendAnimation) SetReplyToMessageID(id int) {
	m.ReplyToMessageID = id
}

// SetReplyMarkup sets reply_markup
func (m *SendAnimation) SetReplyMarkup(kb keyboard.Markup) {
	m.ReplyMarkup = kb
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
	ReplyMarkup         keyboard.Markup `json:"reply_markup,omitempty" form:"reply_markup,omitempty"`
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
func NewSendVoice(chatID int64, voice *args.VoiceConfig, opts ...Option[*SendVoice]) *SendVoice {
	method := new(SendVoice)
	method.ChatID = chatID
	method.File = voice.Local
	method.Voice = voice.Reference
	method.Caption = voice.Caption
	method.ParseMode = voice.ParseMode
	method.Duration = voice.Duration
	for _, o := range opts {
		o(method)
	}
	return method
}

// SetCaption sets caption
func (m *SendVoice) SetCaption(c string) {
	m.Caption = c
}

// SetDuration sets duration
func (m *SendVoice) SetDuration(d int) {
	m.Duration = d
}

// SetParseMode sets parse_mode
func (m *SendVoice) SetParseMode(m2 string) {
	m.ParseMode = m2
}

// SetDisableNotification sets disable_notification
func (m *SendVoice) SetDisableNotification(v bool) {
	m.DisableNotification = v
}

// SetReplyToMessageID sets reply_to_message_id
func (m *SendVoice) SetReplyToMessageID(id int) {
	m.ReplyToMessageID = id
}

// SetReplyMarkup sets reply_markup
func (m *SendVoice) SetReplyMarkup(kb keyboard.Markup) {
	m.ReplyMarkup = kb
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
	ReplyMarkup         keyboard.Markup `json:"reply_markup,omitempty" form:"reply_markup,omitempty"`
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
func NewSendVideoNote(chatID int64, videoNote *args.VideoNoteConfig, opts ...Option[*SendVideoNote]) *SendVideoNote {
	method := new(SendVideoNote)
	method.ChatID = chatID
	method.File = videoNote.Local
	method.VideoNote = videoNote.Reference
	method.Thumb = videoNote.Thumb
	method.ThumbFile = videoNote.ThumbFile
	method.Length = videoNote.Length
	method.Duration = videoNote.Duration
	for _, o := range opts {
		o(method)
	}
	return method
}

// SetThumb sets thumb
func (m *SendVideoNote) SetThumb(t string) {
	m.Thumb = t
}

// SetLength sets length
func (m *SendVideoNote) SetLength(l int) {
	m.Length = l
}

// SetDuration sets duration
func (m *SendVideoNote) SetDuration(d int) {
	m.Duration = d
}

// SetDisableNotification sets disable_notification
func (m *SendVideoNote) SetDisableNotification(v bool) {
	m.DisableNotification = v
}

// SetReplyToMessageID sets reply_to_message_id
func (m *SendVideoNote) SetReplyToMessageID(id int) {
	m.ReplyToMessageID = id
}

// SetReplyMarkup sets reply_markup
func (m *SendVideoNote) SetReplyMarkup(kb keyboard.Markup) {
	m.ReplyMarkup = kb
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
func NewSendMediaGroup(chatID int64, medias []args.MediaGroupConfig, opts ...Option[*SendMediaGroup]) *SendMediaGroup {
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
	for _, o := range opts {
		o(method)
	}
	return method
}

// SetDisableNotification sets disable_notification
func (m *SendMediaGroup) SetDisableNotification(v bool) {
	m.DisableNotification = v
}

// SetReplyToMessageID sets reply_to_message_id
func (m *SendMediaGroup) SetReplyToMessageID(id int) {
	m.ReplyToMessageID = id
}

// EditMessageMedia contains information about editMessageMedia telegram api method.
type EditMessageMedia struct {
	File            *os.File          `json:"-" form:"file"`
	ThumbFile       *os.File          `json:"-" form:"thumb"`
	ChatID          int64             `json:"chat_id,omitempty" form:"chat_id,omitempty"`
	MessageID       int               `json:"message_id,omitempty" form:"message_id,omitempty"`
	InlineMessageID string            `json:"inline_message_id,omitempty" form:"inline_message_id,omitempty"`
	Media           string            `json:"media" form:"media"`
	ReplyMarkup     keyboard.Markup `json:"reply_markup,omitempty"`
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
func NewEditMessageMedia(message *args.MessageID, media args.InputMediaConfig, opts ...Option[*EditMessageMedia]) *EditMessageMedia {
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
	for _, o := range opts {
		o(method)
	}
	return method
}

// SetThumb sets thumb
func (m *EditMessageMedia) SetThumb(t string) {
}

// SetCaption sets caption
func (m *EditMessageMedia) SetCaption(c string) {
}

// SetWidth sets width
func (m *EditMessageMedia) SetWidth(w int) {
}

// SetHeight sets height
func (m *EditMessageMedia) SetHeight(h int) {
}

// SetDuration sets duration
func (m *EditMessageMedia) SetDuration(d int) {
}

// SetSupportsStreaming sets supports_streaming
func (m *EditMessageMedia) SetSupportsStreaming(s bool) {
}

// SetParseMode sets parse_mode
func (m *EditMessageMedia) SetParseMode(m2 string) {
}

// SetDisableNotification sets disable_notification
func (m *EditMessageMedia) SetDisableNotification(v bool) {
}

// SetReplyToMessageID sets reply_to_message_id
func (m *EditMessageMedia) SetReplyToMessageID(id int) {
}

// SetReplyMarkup sets reply_markup
func (m *EditMessageMedia) SetReplyMarkup(kb keyboard.Markup) {
	m.ReplyMarkup = kb
}
