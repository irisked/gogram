package telegram

import (
	"encoding/json"
	"fmt"
)

// The Bot API has objects that come in several kinds, such as ChatMember
// (an owner, an administrator, a restricted member, ...) or InputMedia (a
// photo, a video, ...). Each is an interface implemented by one struct per
// kind: use a type switch to tell them apart.
//
//	switch m := member.(type) {
//	case *telegram.ChatMemberAdministrator:
//		// m.CanDeleteMessages, ...
//	case *telegram.ChatMemberBanned:
//		// m.UntilDate
//	}
//
// The field that tells the kinds apart on the wire ("type", "status" or
// "source") is not part of the structs: it is added when encoding and read
// when decoding.

// UnknownVariant is a value of one of the interface types, such as
// ChatMember, whose kind this version of the package doesn't know: one
// added in a newer Bot API version. Raw is the object as received.
type UnknownVariant struct {
	Raw json.RawMessage
}

// MarshalJSON returns Raw.
func (v *UnknownVariant) MarshalJSON() ([]byte, error) {
	if len(v.Raw) == 0 {
		return []byte("null"), nil
	}
	return v.Raw, nil
}

// RichTextString is a RichText that is plain text.
type RichTextString string

// RichTextArray is a RichText made of a sequence of rich texts.
type RichTextArray []RichText

// marshalTagged encodes v, a struct, with the discriminator field added
// as its first field.
func marshalTagged(field, value string, v any) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	k, _ := json.Marshal(field)
	val, _ := json.Marshal(value)
	head := append(append(append([]byte{'{'}, k...), ':'), val...)
	if len(b) == 2 { // "{}"
		return append(head, '}'), nil
	}
	return append(append(head, ','), b[1:]...), nil
}

// decodeResult decodes the result of a method returning an interface type.
func decodeResult[T any](req Request, raw json.RawMessage, decode func([]byte) (T, error)) (T, error) {
	v, err := decode(raw)
	if err != nil {
		var zero T
		return zero, fmt.Errorf("telegram: decode %s result: %w", req.Method(), err)
	}
	return v, nil
}

// decodeResults decodes the result of a method returning a list of an
// interface type.
func decodeResults[T any](req Request, raws []json.RawMessage, decode func([]byte) (T, error)) ([]T, error) {
	out := make([]T, 0, len(raws))
	for _, raw := range raws {
		v, err := decodeResult(req, raw, decode)
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return out, nil
}
