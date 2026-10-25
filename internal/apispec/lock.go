// Package apispec holds the Bot API spec that package telegram is
// generated from, taken from the original source: the Bot API reference
// at https://core.telegram.org/bots/api, which Telegram publishes only as
// HTML.
//
// The pipeline:
//
//  1. internal/cmd/fetchspec downloads the reference page, keeps its
//     documentation part (Extract) and checks that it parses (Parse).
//     It writes that snapshot to api.html and its SHA-256 to api.lock.
//  2. internal/cmd/genapi checks api.html against api.lock, parses it
//     and generates the Go code.
//
// The page has no versions to pin, so the committed snapshot is the pin:
// the same snapshot always generates the same code. The Makefile at the
// repository root runs these steps.
package apispec

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// Lock records the source of api.html.
type Lock struct {
	URL    string // the page the snapshot was taken from
	SHA256 string // hex SHA-256 of the snapshot
}

// Verify checks that data is the locked snapshot.
func (l Lock) Verify(data []byte) error {
	if got := Sum(data); got != l.SHA256 {
		return fmt.Errorf("api.html has SHA-256 %s, but api.lock expects %s; it was edited by hand: restore it with git, or run `make update-api` to take a new snapshot of %s",
			got, l.SHA256, l.URL)
	}
	return nil
}

// Sum returns the hex SHA-256 of data.
func Sum(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

const header = `# The source of api.html, like go.sum for the Bot API spec.
# Written by ` + "`make update-api`" + `; checked by genapi.
`

// Marshal encodes the lock file.
func (l Lock) Marshal() []byte {
	return fmt.Appendf([]byte(header), "url %s\nsha256 %s\n", l.URL, l.SHA256)
}

// ParseLock decodes a lock file: "key value" lines, with # comments.
func ParseLock(data []byte) (Lock, error) {
	var l Lock
	sc := bufio.NewScanner(bytes.NewReader(data))
	for n := 1; sc.Scan(); n++ {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, ok := strings.Cut(line, " ")
		value = strings.TrimSpace(value)
		if !ok || value == "" {
			return l, fmt.Errorf("api.lock:%d: want \"key value\", got %q", n, line)
		}
		switch key {
		case "url":
			l.URL = value
		case "sha256":
			l.SHA256 = value
		default:
			return l, fmt.Errorf("api.lock:%d: unknown key %q", n, key)
		}
	}
	if l.URL == "" || l.SHA256 == "" {
		return l, fmt.Errorf("api.lock: url and sha256 are both required")
	}
	return l, nil
}

// ReadLock reads a lock file.
func ReadLock(path string) (Lock, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Lock{}, err
	}
	return ParseLock(data)
}

// WriteLock writes a lock file.
func WriteLock(path string, l Lock) error {
	return os.WriteFile(path, l.Marshal(), 0o644)
}
