// Package sandbox helps test bots built with gogram: Server is a fake
// Bot API server that records the requests a bot makes and answers them,
// and Message, Command, Callback and the other builders create the
// updates to feed it.
//
//	func TestStart(t *testing.T) {
//		srv := sandbox.NewServer(t)
//		bot := srv.Bot()
//		registerHandlers(bot)
//
//		if err := bot.HandleUpdate(t.Context(), sandbox.Command("/start")); err != nil {
//			t.Fatal(err)
//		}
//		if got := srv.Calls("sendMessage"); len(got) != 1 || got[0].String("text") != "Welcome!" {
//			t.Errorf("sent %v", got)
//		}
//	}
package sandbox

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/irisked/gogram"
	"github.com/irisked/gogram/telegram"
)

// Request is an API call received by the Server.
type Request struct {
	Method string
	// Params holds the parameters: JSON values as decoded by
	// encoding/json, with numbers as json.Number.
	Params map[string]any
	// Files holds the uploaded files of a multipart request, by part name.
	Files map[string]File
}

// File is an uploaded file.
type File struct {
	Name    string
	Content []byte
}

// String returns a parameter as a string: strings as they are, other
// values JSON-encoded. It returns "" for a missing parameter.
func (r Request) String(key string) string {
	switch v := r.Params[key].(type) {
	case nil:
		return ""
	case string:
		return v
	case json.Number:
		return v.String()
	default:
		b, _ := json.Marshal(v)
		return string(b)
	}
}

// Int returns a numeric parameter, or 0 if it is missing or not a number.
func (r Request) Int(key string) int64 {
	switch v := r.Params[key].(type) {
	case json.Number:
		n, _ := v.Int64()
		return n
	case string:
		n, _ := strconv.ParseInt(v, 10, 64)
		return n
	}
	return 0
}

// HandlerFunc answers a request: the result, or an error, which can be a
// *telegram.Error to answer with a specific error code.
type HandlerFunc func(Request) (result any, err error)

// Server is a fake Bot API server. By default it answers getMe with the
// test bot, getUpdates with the updates added by AddUpdates, the send,
// copy, forward and edit methods with a message built from the request,
// and every other method with an empty result of the right type. Change
// that per method with Handle.
type Server struct {
	t   testing.TB
	srv *httptest.Server

	mu        sync.Mutex
	requests  []Request
	handlers  map[string]HandlerFunc
	updates   []telegram.Update
	nextID    int
	nextMsgID int
	changed   chan struct{} // closed and replaced on every request and AddUpdates
}

// NewServer starts a Server that is closed when the test ends.
func NewServer(t testing.TB) *Server {
	s := &Server{t: t, handlers: map[string]HandlerFunc{}, nextID: 1, nextMsgID: 1, changed: make(chan struct{})}
	s.srv = httptest.NewServer(http.HandlerFunc(s.serve))
	t.Cleanup(s.srv.Close)
	return s
}

// URL returns the base URL of the server, for telegram.WithBaseURL.
func (s *Server) URL() string { return s.srv.URL }

// Client returns a client that talks to the server.
func (s *Server) Client() *telegram.Client {
	return telegram.New("test-token", telegram.WithBaseURL(s.srv.URL), telegram.WithHTTPClient(s.srv.Client()))
}

// Bot returns a bot that talks to the server, with the username of the
// test bot (BotUsername) and the given options.
func (s *Server) Bot(opts ...gogram.Option) *gogram.Bot {
	return gogram.NewWithClient(s.Client(), append([]gogram.Option{gogram.WithUsername(BotUsername)}, opts...)...)
}

// Handle sets how the server answers a method.
func (s *Server) Handle(method string, h HandlerFunc) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers[method] = h
}

// AddUpdates queues updates for getUpdates. Updates with a zero UpdateID
// get the next free one.
func (s *Server) AddUpdates(updates ...telegram.Update) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, u := range updates {
		if u.UpdateID == 0 {
			u.UpdateID = s.nextID
		}
		s.nextID = max(s.nextID, u.UpdateID+1)
		s.updates = append(s.updates, u)
	}
	s.notifyLocked()
}

// Requests returns every request received so far, except getUpdates.
func (s *Server) Requests() []Request {
	s.mu.Lock()
	defer s.mu.Unlock()
	var out []Request
	for _, r := range s.requests {
		if r.Method != "getUpdates" {
			out = append(out, r)
		}
	}
	return out
}

// Calls returns the requests received so far for a method.
func (s *Server) Calls(method string) []Request {
	s.mu.Lock()
	defer s.mu.Unlock()
	var out []Request
	for _, r := range s.requests {
		if r.Method == method {
			out = append(out, r)
		}
	}
	return out
}

// WaitFor waits until the server has received n requests for a method
// and returns them. It fails the test after 5 seconds.
func (s *Server) WaitFor(method string, n int) []Request {
	s.t.Helper()
	deadline := time.After(5 * time.Second)
	for {
		s.mu.Lock()
		changed := s.changed
		s.mu.Unlock()
		if calls := s.Calls(method); len(calls) >= n {
			return calls
		}
		select {
		case <-changed:
		case <-deadline:
			s.t.Fatalf("sandbox: timed out waiting for %d %s calls, got %d", n, method, len(s.Calls(method)))
			return nil
		}
	}
}

func (s *Server) notifyLocked() {
	close(s.changed)
	s.changed = make(chan struct{})
}

func (s *Server) serve(w http.ResponseWriter, r *http.Request) {
	req, err := parseRequest(r)
	if err != nil {
		s.t.Errorf("sandbox: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.mu.Lock()
	s.requests = append(s.requests, req)
	h := s.handlers[req.Method]
	s.notifyLocked()
	s.mu.Unlock()

	var result any
	switch {
	case h != nil:
		result, err = h(req)
	case req.Method == "getUpdates":
		result = s.getUpdates(r.Context(), req)
	default:
		result = s.defaultResult(req)
	}
	writeResponse(w, result, err)
}

// getUpdates answers with the queued updates from the requested offset,
// waiting briefly for new ones like a long poll.
func (s *Server) getUpdates(ctx context.Context, req Request) []telegram.Update {
	offset := int(req.Int("offset"))
	limit := int(req.Int("limit"))
	timeout := time.NewTimer(50 * time.Millisecond)
	defer timeout.Stop()
	for {
		s.mu.Lock()
		var pending []telegram.Update
		for _, u := range s.updates {
			if u.UpdateID >= offset {
				pending = append(pending, u)
			}
		}
		changed := s.changed
		s.mu.Unlock()
		if limit > 0 && len(pending) > limit {
			pending = pending[:limit]
		}
		if len(pending) > 0 {
			return pending
		}
		select {
		case <-changed:
		case <-timeout.C:
			return []telegram.Update{}
		case <-ctx.Done():
			return []telegram.Update{}
		}
	}
}

func (s *Server) defaultResult(req Request) any {
	switch {
	case req.Method == "getMe":
		return Bot
	case req.Method == "sendMediaGroup":
		var media []any
		json.Unmarshal([]byte(req.String("media")), &media)
		msgs := make([]telegram.Message, 0, len(media))
		for range media {
			msgs = append(msgs, s.message(req))
		}
		return msgs
	case req.Method == "sendChatAction", strings.HasSuffix(req.Method, "Draft"):
		return true
	case strings.HasPrefix(req.Method, "send"), req.Method == "copyMessage", req.Method == "forwardMessage":
		return s.message(req)
	case strings.HasPrefix(req.Method, "editMessage"), req.Method == "stopMessageLiveLocation":
		if req.String("inline_message_id") != "" {
			return true
		}
		m := s.message(req)
		m.MessageID = int(req.Int("message_id"))
		return m
	default:
		if r, ok := defaultResults[req.Method]; ok {
			return json.RawMessage(r)
		}
		return true
	}
}

// message builds the message a send method returns.
func (s *Server) message(req Request) telegram.Message {
	s.mu.Lock()
	id := s.nextMsgID
	s.nextMsgID++
	s.mu.Unlock()
	chatID := req.Int("chat_id")
	chatType := "private"
	if chatID < 0 {
		chatType = "supergroup"
	}
	return telegram.Message{
		MessageID: id,
		Date:      time.Now().Unix(),
		Chat:      &telegram.Chat{ID: chatID, Type: chatType},
		From:      Bot,
		Text:      req.String("text"),
		Caption:   req.String("caption"),
	}
}

func parseRequest(r *http.Request) (Request, error) {
	req := Request{
		Method: r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:],
		Params: map[string]any{},
	}
	mediaType, params, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
	switch mediaType {
	case "application/json":
		dec := json.NewDecoder(r.Body)
		dec.UseNumber()
		if err := dec.Decode(&req.Params); err != nil && err != io.EOF {
			return req, fmt.Errorf("decode %s: %w", req.Method, err)
		}
	case "multipart/form-data":
		req.Files = map[string]File{}
		mr := multipart.NewReader(r.Body, params["boundary"])
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				return req, fmt.Errorf("read %s: %w", req.Method, err)
			}
			b, err := io.ReadAll(p)
			if err != nil {
				return req, fmt.Errorf("read %s: %w", req.Method, err)
			}
			if p.FileName() != "" {
				req.Files[p.FormName()] = File{Name: p.FileName(), Content: b}
				continue
			}
			// Objects, arrays, numbers and booleans are sent as JSON.
			var v any
			dec := json.NewDecoder(bytes.NewReader(b))
			dec.UseNumber()
			if err := dec.Decode(&v); err != nil || dec.More() {
				v = string(b)
			}
			req.Params[p.FormName()] = v
		}
	default:
		return req, fmt.Errorf("%s: unexpected Content-Type %q", req.Method, mediaType)
	}
	return req, nil
}

func writeResponse(w http.ResponseWriter, result any, err error) {
	w.Header().Set("Content-Type", "application/json")
	if err == nil {
		json.NewEncoder(w).Encode(map[string]any{"ok": true, "result": result})
		return
	}
	resp := map[string]any{"ok": false, "error_code": http.StatusInternalServerError, "description": err.Error()}
	if apiErr, ok := errors.AsType[*telegram.Error](err); ok {
		resp["error_code"] = apiErr.Code
		resp["description"] = apiErr.Description
		params := map[string]any{}
		if apiErr.RetryAfter > 0 {
			params["retry_after"] = int(apiErr.RetryAfter / time.Second)
		}
		if apiErr.MigrateToChatID != 0 {
			params["migrate_to_chat_id"] = apiErr.MigrateToChatID
		}
		if len(params) > 0 {
			resp["parameters"] = params
		}
	}
	json.NewEncoder(w).Encode(resp)
}
