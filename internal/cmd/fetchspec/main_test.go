package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/irisked/gogram/internal/apispec"
)

// referencePage is a minimal Bot API reference page.
func referencePage(version string) string {
	return `<html><body><div id="dev_page_content">
<h4><a class="anchor" name="august-24-2026" href="#august-24-2026"></a>August 24, 2026</h4>
<p>` + version + `</p>
<h4><a class="anchor" name="user" href="#user"></a>User</h4>
<p>A user.</p>
<table><tbody><tr><td>id</td><td>Integer</td><td>Identifier</td></tr></tbody></table>
<h4><a class="anchor" name="getme" href="#getme"></a>getMe</h4>
<p>Returns basic information about the bot in form of a User object.</p>
</div><!-- page generated in 1ms --></body></html>`
}

func TestRun(t *testing.T) {
	snapshot := func(page string) []byte {
		s, err := apispec.Extract([]byte(page))
		if err != nil {
			t.Fatal(err)
		}
		return s
	}
	current := referencePage("Bot API 1.0")

	tests := []struct {
		name        string
		page        string // served by the fake Telegram
		status      int
		check       bool
		wantErr     string
		wantMsg     string
		wantChanged bool
		wantWritten bool // whether api.html and api.lock were rewritten
	}{
		{
			name:    "up to date",
			page:    current,
			wantMsg: "up to date",
		},
		{
			name:    "up to date even though the page footer changed",
			page:    strings.Replace(current, "1ms", "93ms", 1),
			wantMsg: "up to date",
		},
		{
			name:        "new version",
			page:        referencePage("Bot API 2.0"),
			wantMsg:     "new snapshot of ",
			wantChanged: true,
			wantWritten: true,
		},
		{
			name:        "check only reports a new version",
			page:        referencePage("Bot API 2.0"),
			check:       true,
			wantMsg:     "changed since the snapshot: Bot API 2.0",
			wantChanged: true,
		},
		{
			name:    "page that doesn't parse",
			page:    `<div id="dev_page_content"><h4><a name="x" href="#x"></a>broken</h4></div>`,
			wantErr: "doesn't parse",
		},
		{
			name:    "layout changed",
			page:    "<html>new design</html>",
			wantErr: "layout",
		},
		{
			name:    "server error",
			status:  http.StatusBadGateway,
			wantErr: "502",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			old := snapshot(current)
			os.WriteFile(filepath.Join(dir, "api.html"), old, 0o644)
			apispec.WriteLock(filepath.Join(dir, "api.lock"), apispec.Lock{URL: "x", SHA256: apispec.Sum(old)})

			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if tt.status != 0 {
					w.WriteHeader(tt.status)
				}
				w.Write([]byte(tt.page))
			}))
			defer srv.Close()
			f := fetcher{client: srv.Client(), url: srv.URL}

			msg, changed, err := f.run(t.Context(), dir, tt.check)
			if tt.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("error = %v, want one containing %q", err, tt.wantErr)
				}
			} else if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(msg, tt.wantMsg) || changed != tt.wantChanged {
				t.Errorf("run = %q, changed %v; want a message containing %q, changed %v", msg, changed, tt.wantMsg, tt.wantChanged)
			}

			html, _ := os.ReadFile(filepath.Join(dir, "api.html"))
			lock, _ := apispec.ReadLock(filepath.Join(dir, "api.lock"))
			if written := string(html) != string(old); written != tt.wantWritten {
				t.Errorf("api.html rewritten = %v, want %v", written, tt.wantWritten)
			}
			if lock.Verify(html) != nil {
				t.Error("api.lock doesn't match api.html")
			}
			if tt.wantWritten && lock.URL != srv.URL {
				t.Errorf("lock URL = %q, want %q", lock.URL, srv.URL)
			}
		})
	}
}
