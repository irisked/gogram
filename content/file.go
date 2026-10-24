package content

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
	InputMediaConfig() (string, map[string]*os.File)
}

// MediaGroupConfig input media data.
type MediaGroupConfig interface {
	MediaGroupConfig() (string, map[string]*os.File)
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

// InputMediaConfig returns json-serialized string of PhotoConfig and attached files.
func (pc *PhotoConfig) InputMediaConfig() (string, map[string]*os.File) {
	res, _ := json.Marshal(pc)
	stringified := string(res)
	files := make(map[string]*os.File)
	files["file"] = pc.Local
	return stringified, files
}

// MediaGroupConfig returns json-serialized string of PhotoConfig and attached files.
func (pc *PhotoConfig) MediaGroupConfig() (string, map[string]*os.File) {
	res, _ := json.Marshal(pc)
	stringified := string(res)
	files := make(map[string]*os.File)
	files["file"] = pc.Local
	return stringified, files
}

// Photo photo.
func Photo(file *File) *PhotoConfig {
	photo := new(PhotoConfig)
	photo.Type = "photo"
	photo.Local = file.Local
	photo.Reference = file.Reference
	return photo
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

// InputMediaConfig returns json-serialized string of PhotoData.
func (ad *AudioConfig) InputMediaConfig() (string, map[string]*os.File) {
	res, _ := json.Marshal(ad)
	stringified := string(res)
	files := make(map[string]*os.File)
	files["file"] = ad.Local
	files["thumb"] = ad.ThumbFile
	return stringified, files
}

// Audio its an audio parameter setter.
func Audio(file *File) *AudioConfig {
	audio := new(AudioConfig)
	audio.Type = "audio"
	audio.Local = file.Local
	audio.Reference = file.Reference
	return audio
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

// InputMediaConfig returns json-serialized string of PhotoData.
func (dd *DocumentConfig) InputMediaConfig() (string, map[string]*os.File) {
	res, _ := json.Marshal(dd)
	stringified := string(res)
	files := make(map[string]*os.File)
	files["file"] = dd.Local
	files["thumb"] = dd.ThumbFile
	return stringified, files
}

// Document its an document parameter setter.
func Document(file *File) *DocumentConfig {
	doc := new(DocumentConfig)
	doc.Type = "document"
	doc.Local = file.Local
	doc.Reference = file.Reference
	return doc
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

// InputMediaConfig returns json-serialized string of PhotoData.
func (vd *VideoConfig) InputMediaConfig() (string, map[string]*os.File) {
	res, _ := json.Marshal(vd)
	stringified := string(res)
	files := make(map[string]*os.File)
	files["file"] = vd.Local
	files["thumb"] = vd.ThumbFile
	return stringified, files
}

// MediaGroupConfig returns json-serialized string of PhotoData.
func (vd *VideoConfig) MediaGroupConfig() (string, map[string]*os.File) {
	res, _ := json.Marshal(vd)
	stringified := string(res)
	files := make(map[string]*os.File)
	files["file"] = vd.Local
	files["thumb"] = vd.ThumbFile
	return stringified, files
}

// Video its an video parameter setter.
func Video(file *File) *VideoConfig {
	video := new(VideoConfig)
	video.Type = "video"
	video.Local = file.Local
	video.Reference = file.Reference
	return video
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

// InputMediaConfig returns json-serialized string of PhotoData.
func (ad *AnimationConfig) InputMediaConfig() (string, map[string]*os.File) {
	res, _ := json.Marshal(ad)
	stringified := string(res)
	files := make(map[string]*os.File)
	files["file"] = ad.Local
	files["thumb"] = ad.ThumbFile
	return stringified, files
}

// Animation its an video parameter setter.
func Animation(file *File) *AnimationConfig {
	animation := new(AnimationConfig)
	animation.Type = "animation"
	animation.Local = file.Local
	animation.Reference = file.Reference
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
func Voice(file *File) *VoiceConfig {
	voice := new(VoiceConfig)
	voice.Local = file.Local
	voice.Reference = file.Reference
	return voice
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
func VideoNote(file *File) *VideoNoteConfig {
	videoNote := new(VideoNoteConfig)
	videoNote.Local = file.Local
	videoNote.Reference = file.Reference
	return videoNote
}
