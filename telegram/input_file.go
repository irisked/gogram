package telegram

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
)

// InputFile is a file to send. It is either a reference to a file
// Telegram already has (a file_id or an HTTP URL) or new content to
// upload. Requests that carry an upload are sent as multipart/form-data
// automatically.
//
// An upload made with FileReader or LocalFile reads its io.Reader once,
// so it can be sent only once. One made with LocalPath opens the file
// each time it is sent.
type InputFile struct {
	ref    string // file_id or URL
	name   string // filename of an upload
	reader io.Reader
	path   string // file to open when sending

	attach string // multipart part name, assigned while encoding
}

// FileID refers to a file already stored on Telegram's servers.
func FileID(id string) *InputFile {
	return &InputFile{ref: id}
}

// FileURL lets Telegram download the file from an HTTP URL.
func FileURL(url string) *InputFile {
	return &InputFile{ref: url}
}

// FileReader uploads the contents of r under the given file name.
func FileReader(name string, r io.Reader) *InputFile {
	return &InputFile{name: name, reader: r}
}

// LocalFile uploads an open file. The caller still owns f and must
// close it after the request is sent; LocalPath does that for you.
func LocalFile(f *os.File) *InputFile {
	return FileReader(filepath.Base(f.Name()), f)
}

// LocalPath uploads the file at path. The file is opened when the
// request is sent and closed once it is uploaded.
func LocalPath(path string) *InputFile {
	return &InputFile{name: filepath.Base(path), path: path}
}

func (f *InputFile) isUpload() bool {
	return f != nil && (f.reader != nil || f.path != "")
}

// open returns the content of an upload.
func (f *InputFile) open() (io.ReadCloser, error) {
	if f.path != "" {
		return os.Open(f.path)
	}
	return io.NopCloser(f.reader), nil
}

// MarshalJSON encodes a reference as its file_id/URL and an upload as
// an "attach://<name>" pointer to its multipart part.
func (f *InputFile) MarshalJSON() ([]byte, error) {
	if !f.isUpload() {
		return json.Marshal(f.ref)
	}
	if f.attach == "" {
		return nil, errors.New("telegram: InputFile upload used outside a request")
	}
	return json.Marshal("attach://" + f.attach)
}
