package method

import (
	"fmt"
	"os"

	"github.com/irisked/gogram/args"
	"github.com/irisked/gogram/types"
)

// GetUpdates contains information about getUpdates telegram api method.
type GetUpdates struct {
	Offset  int `json:"offset,omitempty"`
	Limit   int `json:"limit,omitempty"`
	Timeout int `json:"timeout,omitempty"`
	SimpleMethod
}

// GetUpdatesOption interface for optional parameters.
type GetUpdatesOption interface {
	ApplyGetUpdatesOption(*GetUpdates)
}

// NewGetUpdates creates GetUpdates.
func NewGetUpdates(options ...GetUpdatesOption) *GetUpdates {
	method := new(GetUpdates)
	for _, option := range options {
		option.ApplyGetUpdatesOption(method)
	}
	fmt.Println(method)
	return method
}

// Method returns telegram api method endpoint.
func (r *GetUpdates) Method() string {
	return "getUpdates"
}

// GetMe contains information about getMe telegram api method.
type GetMe struct {
	SimpleMethod
}

// Method returns telegram api method endpoint.
func (r *GetMe) Method() string {
	return "getMe"
}

// NewGetMe creates GetMe.
func NewGetMe() *GetMe {
	return new(GetMe)
}

// GetWebhookInfo contains information about getWebhookInfo telegram api method.
type GetWebhookInfo struct {
	SimpleMethod
}

// Method returns telegram api method endpoint.
func (r *GetWebhookInfo) Method() string {
	return "getWebhookInfo"
}

// NewGetWebhookInfo creates GetWebhookInfo.
func NewGetWebhookInfo() *GetWebhookInfo {
	return new(GetWebhookInfo)
}

// SetWebhook contains information about SetWebhook telegram api method.
type SetWebhook struct {
	File           *os.File           `json:"-"`
	URL            string             `json:"url" form:"url"`
	Certificate    string             `json:"certificate,omitempty" form:"certificate,omitempty"`
	MaxConnections int                `json:"max_connections,omitempty" form:"max_connections,omitempty"`
	AllowedUpdates []types.UpdateType `json:"allowed_updates,omitempty" form:"allowed_updates,omitempty"`
}

// SetWebhookOption interface for optional parameters.
type SetWebhookOption interface {
	ApplySetWebhookOption(*SetWebhook)
}

// NewSetWebhook creates SetWebhook.
func NewSetWebhook(url string, file *args.File, options ...SetWebhookOption) *SetWebhook {
	method := new(SetWebhook)
	method.URL = url
	if file != nil {
		method.File = file.Local
		method.Certificate = file.Reference
	}
	for _, option := range options {
		option.ApplySetWebhookOption(method)
	}
	return method
}

// IsMultipart returns true if method is multipart and false otherwise.
func (sw *SetWebhook) IsMultipart() bool {
	return sw.File != nil
}

// Method returns telegram api method endpoint.
func (sw *SetWebhook) Method() string {
	return "setWebhook"
}
