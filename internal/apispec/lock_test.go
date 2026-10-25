package apispec

import (
	"os"
	"strings"
	"testing"
)

func TestParseLock(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Lock
		wantErr string
	}{
		{
			name:  "valid",
			input: "# comment\nurl https://example.com/api\n\nsha256 ff\n",
			want:  Lock{URL: "https://example.com/api", SHA256: "ff"},
		},
		{name: "missing key", input: "url https://example.com/api\n", wantErr: "required"},
		{name: "unknown key", input: "url x\nref main\n", wantErr: `unknown key "ref"`},
		{name: "no value", input: "url\n", wantErr: "api.lock:1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLock([]byte(tt.input))
			if tt.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("error = %v, want one containing %q", err, tt.wantErr)
				}
				return
			}
			if err != nil || got != tt.want {
				t.Fatalf("ParseLock = %+v, %v; want %+v", got, err, tt.want)
			}
			// Marshal and parse again.
			again, err := ParseLock(got.Marshal())
			if err != nil || again != got {
				t.Errorf("round trip = %+v, %v; want %+v", again, err, got)
			}
		})
	}
}

func TestLockVerify(t *testing.T) {
	l := Lock{URL: URL, SHA256: Sum([]byte("snapshot"))}
	tests := []struct {
		name    string
		data    string
		wantErr string
	}{
		{"matching snapshot", "snapshot", ""},
		{"edited snapshot", "snapshot!", "make update-api"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := l.Verify([]byte(tt.data))
			if tt.wantErr == "" && err != nil || tt.wantErr != "" && (err == nil || !strings.Contains(err.Error(), tt.wantErr)) {
				t.Errorf("Verify = %v, want error containing %q", err, tt.wantErr)
			}
		})
	}
}

// TestCommittedSnapshot checks that the committed api.html is the locked
// one and parses.
func TestCommittedSnapshot(t *testing.T) {
	l, err := ReadLock("api.lock")
	if err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile("api.html")
	if err != nil {
		t.Fatal(err)
	}
	if err := l.Verify(data); err != nil {
		t.Fatal(err)
	}
	s, err := Parse(data)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(s.Version, "Bot API ") || len(s.Methods) == 0 || len(s.Types) == 0 {
		t.Errorf("parsed %q with %d methods and %d types", s.Version, len(s.Methods), len(s.Types))
	}
}
