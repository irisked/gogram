// Command fetchspec takes a snapshot of the Bot API reference, the source
// that package telegram is generated from, into internal/apispec; see
// package apispec.
//
// Run it from the repository root, usually through the Makefile:
//
//	go run ./internal/cmd/fetchspec         # make update-api: snapshot the page into api.html and api.lock
//	go run ./internal/cmd/fetchspec -check  # make check-api: exit 1 if the page changed since the snapshot
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/irisked/gogram/internal/apispec"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("fetchspec: ")
	dir := flag.String("dir", "internal/apispec", "directory of api.html and api.lock")
	check := flag.Bool("check", false, "only report whether the page changed since the snapshot")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	f := fetcher{client: http.DefaultClient, url: apispec.URL}
	msg, changed, err := f.run(ctx, *dir, *check)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
	if *check && changed {
		os.Exit(1)
	}
}

type fetcher struct {
	client *http.Client
	url    string
}

// run downloads the reference page and, unless check is set, writes its
// snapshot and lock into dir. It reports whether the page differs from the
// current snapshot.
func (f fetcher) run(ctx context.Context, dir string, check bool) (msg string, changed bool, err error) {
	page, err := f.get(ctx)
	if err != nil {
		return "", false, fmt.Errorf("download the Bot API reference: %w", err)
	}
	snapshot, err := apispec.Extract(page)
	if err != nil {
		return "", false, err
	}
	spec, err := apispec.Parse(snapshot)
	if err != nil {
		return "", false, fmt.Errorf("the page doesn't parse as the Bot API reference:\n%w", err)
	}
	about := fmt.Sprintf("%s (%s): %d methods, %d types", spec.Version, spec.ReleaseDate, len(spec.Methods), len(spec.Types))

	lockPath := filepath.Join(dir, "api.lock")
	sum := apispec.Sum(snapshot)
	if old, err := apispec.ReadLock(lockPath); err == nil && old.SHA256 == sum {
		return "api.html is up to date with " + f.url + ": " + about, false, nil
	}
	if check {
		return "The Bot API reference changed since the snapshot: " + about + ". Run make update-api.", true, nil
	}

	if err := os.WriteFile(filepath.Join(dir, "api.html"), snapshot, 0o644); err != nil {
		return "", false, err
	}
	if err := apispec.WriteLock(lockPath, apispec.Lock{URL: f.url, SHA256: sum}); err != nil {
		return "", false, err
	}
	return "api.html: new snapshot of " + f.url + ": " + about, true, nil
}

// get downloads the page, retrying network errors a few times.
func (f fetcher) get(ctx context.Context) ([]byte, error) {
	var err error
	for attempt := range 3 {
		if attempt > 0 {
			select {
			case <-time.After(time.Duration(attempt) * time.Second):
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}
		var page []byte
		if page, err = f.getOnce(ctx); err == nil || errors.Is(err, errStatus) {
			return page, err
		}
	}
	return nil, err
}

var errStatus = errors.New("unexpected status")

func (f fetcher) getOnce(ctx context.Context) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, f.url, nil)
	if err != nil {
		return nil, err
	}
	res, err := f.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s: %w %s", f.url, errStatus, res.Status)
	}
	return io.ReadAll(res.Body)
}
