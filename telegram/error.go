package telegram

import (
	"fmt"
	"time"

	"github.com/irisked/gogram/telegram/internal/transport"
)

// Error is returned when the Bot API rejects a call. Use errors.As to
// inspect it:
//
//	var apiErr *telegram.Error
//	if errors.As(err, &apiErr) && apiErr.RetryAfter > 0 {
//		time.Sleep(apiErr.RetryAfter)
//	}
type Error struct {
	// Method is the API method that failed, e.g. "sendMessage".
	Method string
	// Code is the error code, which mirrors HTTP status codes (400, 403, 429...).
	Code int
	// Description is the human-readable message from Telegram.
	Description string
	// RetryAfter is set when the call was rate limited (Code 429).
	RetryAfter time.Duration
	// MigrateToChatID is set when a group was upgraded to a supergroup.
	MigrateToChatID int64
}

func (e *Error) Error() string {
	return fmt.Sprintf("telegram: %s: %d %s", e.Method, e.Code, e.Description)
}

func newError(method string, res transport.Response) *Error {
	e := &Error{Method: method, Code: res.ErrorCode, Description: res.Description}
	if p := res.Parameters; p != nil {
		e.RetryAfter = time.Duration(p.RetryAfter) * time.Second
		e.MigrateToChatID = p.MigrateToChatID
	}
	return e
}
