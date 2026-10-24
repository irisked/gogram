package transport

import (
	"bytes"
	"context"
	"testing"
)

func TestResponseParametersStringFormatsAsNumbers(t *testing.T) {
	rp := &ResponseParameters{MigrateToChatID: -1001234567890, RetryAfter: 30}
	got := rp.String()
	want := "migrate_to_chat_id=-1001234567890 retry_after=30"
	if got != want {
		t.Fatalf("String() = %q, want %q", got, want)
	}
}

func TestClientCheckReturnsErrorForNonOkResponse(t *testing.T) {
	c := NewClient("test-token")
	res := Response{
		Ok:          false,
		Description: "Bad Request: chat not found",
		Parameters:  &ResponseParameters{RetryAfter: 5},
	}

	_, err := c.check(res)
	if err == nil {
		t.Fatalf("expected an error for a non-ok response")
	}
	const want = "Bad Request: chat not foundmigrate_to_chat_id=0 retry_after=5"
	if err.Error() != want {
		t.Fatalf("err = %q, want %q", err.Error(), want)
	}
}

func TestClientCheckReturnsNoErrorForOkResponse(t *testing.T) {
	c := NewClient("test-token")
	res := Response{Ok: true}

	if _, err := c.check(res); err != nil {
		t.Fatalf("expected no error for an ok response, got %v", err)
	}
}

func TestCreateHTTPRequestUsesGivenContext(t *testing.T) {
	c := NewClient("test-token")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req, err := c.createHTTPRequest(ctx, "getMe", bytes.NewBuffer(nil), "application/json")
	if err != nil {
		t.Fatalf("createHTTPRequest() error = %v", err)
	}
	if req.Context() != ctx {
		t.Fatalf("expected the request to carry the context passed to createHTTPRequest")
	}
}
