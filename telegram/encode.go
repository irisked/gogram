package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

// encodeRequest serializes req as JSON, or as multipart/form-data when
// it carries at least one InputFile upload.
//
// Multipart requests reuse the JSON encoding for field names and
// omitempty: every top-level JSON field becomes a form field (objects
// and arrays stay JSON-encoded, as the Bot API expects), and every
// upload becomes a file part. A top-level upload is sent under its own
// field name; uploads nested inside other values (e.g. InputMedia in
// sendMediaGroup) are referenced with "attach://<part>".
//
// It also validates the values that implement validator, such as
// buttons with too much callback data.
func encodeRequest(req Request) (io.Reader, string, error) {
	uploads, err := walk(reflect.ValueOf(req), nil)
	if err != nil {
		return nil, "", err
	}
	if len(uploads) == 0 {
		b, err := json.Marshal(req)
		if err != nil {
			return nil, "", err
		}
		return bytes.NewReader(b), "application/json", nil
	}

	for i, f := range uploads {
		f.attach = "file" + strconv.Itoa(i)
	}
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(b, &fields); err != nil {
		return nil, "", err
	}

	// Open every upload first, so a missing file fails before sending.
	contents := make([]io.ReadCloser, 0, len(uploads))
	closeAll := func() {
		for _, c := range contents {
			c.Close()
		}
	}
	for _, f := range uploads {
		r, err := f.open()
		if err != nil {
			closeAll()
			return nil, "", fmt.Errorf("upload %s: %w", f.name, err)
		}
		contents = append(contents, r)
	}

	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)
	go func() {
		defer closeAll()
		pw.CloseWithError(writeMultipart(mw, fields, uploads, contents))
	}()
	return pr, mw.FormDataContentType(), nil
}

func writeMultipart(mw *multipart.Writer, fields map[string]json.RawMessage, uploads []*InputFile, contents []io.ReadCloser) error {
	partName := make(map[string]string, len(uploads))
	keys := make([]string, 0, len(fields))
	for k := range fields {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	for _, key := range keys {
		raw := fields[key]
		if string(raw) == "null" {
			continue
		}
		if raw[0] != '"' {
			if err := mw.WriteField(key, string(raw)); err != nil {
				return err
			}
			continue
		}
		var s string
		if err := json.Unmarshal(raw, &s); err != nil {
			return err
		}
		if attach, ok := strings.CutPrefix(s, "attach://"); ok && isUploadName(attach, uploads) {
			partName[attach] = key
			continue
		}
		if err := mw.WriteField(key, s); err != nil {
			return err
		}
	}

	for i, f := range uploads {
		name := f.attach
		if field, ok := partName[f.attach]; ok {
			name = field
		}
		part, err := mw.CreateFormFile(name, f.name)
		if err != nil {
			return err
		}
		if _, err := io.Copy(part, contents[i]); err != nil {
			return fmt.Errorf("upload %s: %w", f.name, err)
		}
	}
	return mw.Close()
}

func isUploadName(name string, uploads []*InputFile) bool {
	return slices.ContainsFunc(uploads, func(f *InputFile) bool { return f.attach == name })
}

var inputFileType = reflect.TypeFor[*InputFile]()

// validator is implemented by values the encoder checks before sending.
type validator interface {
	validate() error
}

// walk visits v and returns every distinct InputFile upload it contains,
// in field order. It fails on the first value whose validate fails.
func walk(v reflect.Value, out []*InputFile) ([]*InputFile, error) {
	if v.IsValid() && v.CanInterface() {
		if val, ok := v.Interface().(validator); ok && !(v.Kind() == reflect.Pointer && v.IsNil()) {
			if err := val.validate(); err != nil {
				return out, err
			}
		}
	}
	var err error
	switch v.Kind() {
	case reflect.Pointer:
		if v.IsNil() {
			return out, nil
		}
		if v.Type() == inputFileType {
			f := v.Interface().(*InputFile)
			if f.isUpload() && !slices.Contains(out, f) {
				out = append(out, f)
			}
			return out, nil
		}
		return walk(v.Elem(), out)
	case reflect.Interface:
		if !v.IsNil() {
			return walk(v.Elem(), out)
		}
	case reflect.Struct:
		t := v.Type()
		for i := range t.NumField() {
			if !t.Field(i).IsExported() {
				continue
			}
			if out, err = walk(v.Field(i), out); err != nil {
				return out, err
			}
		}
	case reflect.Slice, reflect.Array:
		for i := range v.Len() {
			if out, err = walk(v.Index(i), out); err != nil {
				return out, err
			}
		}
	}
	return out, nil
}
