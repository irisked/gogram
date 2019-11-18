package params

import (
	"encoding/json"
	"fmt"
	"os"
)

func attachment(name string) string {
	return fmt.Sprintf("attachment://%s", name)
}

// InputMediaConfig input media data.
type InputMediaConfig interface {
	GetInputMediaConfig() (string, map[string]*os.File)
}

// MediaGroupConfig input media data.
type MediaGroupConfig interface {
	GetMediaGroupConfig() (string, map[string]*os.File)
}

// File file.
type File struct {
	Local     *os.File
	Reference string
}

// FileID sets file id.
func FileID(id string) *File {
	return &File{nil, id}
}

// FileURL sets file url.
func FileURL(url string) *File {
	return &File{nil, url}
}

// LocalFile sets LocalFile file.
func LocalFile(file *os.File) *File {
	return &File{file, attachment(file.Name())}
}

// PhotoConfig Photo
type PhotoConfig struct {
	Local     *os.File `json:"-"`
	Reference string   `json:"media"`
	Caption   string   `json:"caption"`
	ParseMode string   `json:"parse_mode"`
	Type      string   `json:"type"`
}

// GetInputMediaConfig returns json-serialized string of PhotoConfig and attached files.
func (pc *PhotoConfig) GetInputMediaConfig() (string, map[string]*os.File) {
	res, _ := json.Marshal(pc)
	stringified := string(res)
	files := make(map[string]*os.File)
	files["file"] = pc.Local
	return stringified, files
}

// GetMediaGroupConfig returns json-serialized string of PhotoConfig and attached files.
func (pc *PhotoConfig) GetMediaGroupConfig() (string, map[string]*os.File) {
	res, _ := json.Marshal(pc)
	stringified := string(res)
	files := make(map[string]*os.File)
	files["file"] = pc.Local
	return stringified, files
}

// Photo photo.
func Photo(file *File, options ...PhotoConfigOption) *PhotoConfig {
	photo := new(PhotoConfig)
	photo.Type = "photo"
	photo.Local = file.Local
	photo.Reference = file.Reference
	for _, option := range options {
		option.ApplyPhotoConfigOption(photo)
	}
	return photo
}

// PhotoConfigOption its an functional parameters interface.
type PhotoConfigOption interface {
	ApplyPhotoConfigOption(*PhotoConfig)
}

// AudioConfig its an photo parameter.
type AudioConfig struct {
	Local     *os.File `json:"-"`
	Reference string   `json:"media"`
	Caption   string   `json:"caption"`
	ParseMode string   `json:"parse_mode"`
	Duration  int      `json:"duration"`
	Performer string   `json:"performer"`
	Title     string   `json:"title"`
	Thumb     string   `json:"thumb"`
	ThumbFile *os.File `json:"-"`
	Type      string   `json:"type"`
}

// GetInputMediaConfig returns json-serialized string of PhotoData.
func (ad *AudioConfig) GetInputMediaConfig() (string, map[string]*os.File) {
	res, _ := json.Marshal(ad)
	stringified := string(res)
	files := make(map[string]*os.File)
	files["file"] = ad.Local
	files["thumb"] = ad.ThumbFile
	return stringified, files
}

// Audio its an audio parameter setter.
func Audio(file *File, options ...AudioConfigOption) *AudioConfig {
	audio := new(AudioConfig)
	audio.Type = "audio"
	audio.Local = file.Local
	audio.Reference = file.Reference
	for _, option := range options {
		option.ApplyAudioConfigOption(audio)
	}
	return audio
}

// AudioConfigOption its an functional parameters interface.
type AudioConfigOption interface {
	ApplyAudioConfigOption(*AudioConfig)
}

// DocumentConfig its an photo parameter.
type DocumentConfig struct {
	Local     *os.File `json:"-"`
	Reference string   `json:"media"`
	Caption   string   `json:"caption"`
	ParseMode string   `json:"parse_mode"`
	Thumb     string   `json:"thumb"`
	ThumbFile *os.File `json:"-"`
	Type      string   `json:"type"`
}

// GetInputMediaConfig returns json-serialized string of PhotoData.
func (dd *DocumentConfig) GetInputMediaConfig() (string, map[string]*os.File) {
	res, _ := json.Marshal(dd)
	stringified := string(res)
	files := make(map[string]*os.File)
	files["file"] = dd.Local
	files["thumb"] = dd.ThumbFile
	return stringified, files
}

// Document its an document parameter setter.
func Document(file *File, options ...DocumentConfigOption) *DocumentConfig {
	doc := new(DocumentConfig)
	doc.Type = "document"
	doc.Local = file.Local
	doc.Reference = file.Reference
	for _, option := range options {
		option.ApplyDocumentConfigOption(doc)
	}
	return doc
}

// DocumentConfigOption its an functional parameters interface.
type DocumentConfigOption interface {
	ApplyDocumentConfigOption(*DocumentConfig)
}

// VideoConfig its an photo parameter.
type VideoConfig struct {
	Local             *os.File `json:"-"`
	Reference         string   `json:"media"`
	Caption           string   `json:"caption"`
	ParseMode         string   `json:"parse_mode"`
	Duration          int      `json:"duration"`
	Width             int      `json:"width"`
	Height            int      `json:"height"`
	SupportsStreaming bool     `json:"supports_streaming"`
	Thumb             string   `json:"thumb"`
	ThumbFile         *os.File `json:"-"`
	Type              string   `json:"type"`
}

// GetInputMediaConfig returns json-serialized string of PhotoData.
func (vd *VideoConfig) GetInputMediaConfig() (string, map[string]*os.File) {
	res, _ := json.Marshal(vd)
	stringified := string(res)
	files := make(map[string]*os.File)
	files["file"] = vd.Local
	files["thumb"] = vd.ThumbFile
	return stringified, files
}

// GetMediaGroupConfig returns json-serialized string of PhotoData.
func (vd *VideoConfig) GetMediaGroupConfig() (string, map[string]*os.File) {
	res, _ := json.Marshal(vd)
	stringified := string(res)
	files := make(map[string]*os.File)
	files["file"] = vd.Local
	files["thumb"] = vd.ThumbFile
	return stringified, files
}

// Video its an video parameter setter.
func Video(file *File, options ...VideoConfigOption) *VideoConfig {
	video := new(VideoConfig)
	video.Type = "video"
	video.Local = file.Local
	video.Reference = file.Reference
	for _, option := range options {
		option.ApplyVideoConfigOption(video)
	}
	return video
}

// VideoConfigOption its an functional parameters interface.
type VideoConfigOption interface {
	ApplyVideoConfigOption(*VideoConfig)
}

// AnimationConfig its an photo parameter.
type AnimationConfig struct {
	Local     *os.File `json:"-"`
	Reference string   `json:"media"`
	Caption   string   `json:"caption"`
	ParseMode string   `json:"parse_mode"`
	Duration  int      `json:"duration"`
	Width     int      `json:"width"`
	Height    int      `json:"height"`
	Thumb     string   `json:"thumb"`
	ThumbFile *os.File `json:"-"`
	Type      string   `json:"type"`
}

// GetInputMediaConfig returns json-serialized string of PhotoData.
func (ad *AnimationConfig) GetInputMediaConfig() (string, map[string]*os.File) {
	res, _ := json.Marshal(ad)
	stringified := string(res)
	files := make(map[string]*os.File)
	files["file"] = ad.Local
	files["thumb"] = ad.ThumbFile
	return stringified, files
}

// AnimationConfigOption its an functional parameters interface.
type AnimationConfigOption interface {
	ApplyAnimationConfigOption(*AnimationConfig)
}

// Animation its an video parameter setter.
func Animation(file *File, options ...AnimationConfigOption) *AnimationConfig {
	animation := new(AnimationConfig)
	animation.Type = "animation"
	animation.Local = file.Local
	animation.Reference = file.Reference
	for _, option := range options {
		option.ApplyAnimationConfigOption(animation)
	}
	return animation
}

// VoiceConfig its an photo parameter.
type VoiceConfig struct {
	Local     *os.File `json:"-"`
	Reference string   `json:"media"`
	Caption   string   `json:"caption"`
	ParseMode string   `json:"parse_mode"`
	Duration  int      `json:"duration"`
	Type      string   `json:"type"`
}

// Voice its an audio parameter setter.
func Voice(file *File, options ...VoiceConfigOption) *VoiceConfig {
	voice := new(VoiceConfig)
	voice.Local = file.Local
	voice.Reference = file.Reference
	for _, option := range options {
		option.ApplyVoiceConfigOption(voice)
	}
	return voice
}

// VoiceConfigOption its an functional parameters interface.
type VoiceConfigOption interface {
	ApplyVoiceConfigOption(*VoiceConfig)
}

// VideoNoteConfig its an photo parameter.
type VideoNoteConfig struct {
	Local     *os.File `json:"-"`
	Reference string   `json:"media"`
	Duration  int      `json:"duration"`
	Length    int      `json:"length"`
	Thumb     string   `json:"thumb"`
	ThumbFile *os.File `json:"-"`
}

// VideoNote its an video parameter setter.
func VideoNote(file *File, options ...VideoNoteConfigOption) *VideoNoteConfig {
	videoNote := new(VideoNoteConfig)
	videoNote.Local = file.Local
	videoNote.Reference = file.Reference
	for _, option := range options {
		option.ApplyVideoNoteConfigOption(videoNote)
	}
	return videoNote
}

// VideoNoteConfigOption its an functional parameters interface.
type VideoNoteConfigOption interface {
	ApplyVideoNoteConfigOption(*VideoNoteConfig)
}
