package method

import (
	"fmt"
	"os"

	"gogram/args"
	"gogram/types"
)

// GetUpdates contains information about getUpdates telegram api method.
type GetUpdates struct {
	Offset  int `json:"offset,omitempty"`
	Limit   int `json:"limit,omitempty"`
	Timeout int `json:"timeout,omitempty"`
	SimpleMethod
}

// SetOffset sets offset
func (u *GetUpdates) SetOffset(v int) {
	u.Offset = v
}

// SetLimit sets limit
func (u *GetUpdates) SetLimit(v int) {
	u.Limit = v
}

// SetTimeout sets timeout
func (u *GetUpdates) SetTimeout(v int) {
	u.Timeout = v
}

// NewGetUpdates creates GetUpdates.
func NewGetUpdates(opts ...Option[*GetUpdates]) *GetUpdates {
	m := &GetUpdates{}
	for _, o := range opts {
		o(m)
	}
	fmt.Println(m)
	return m
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

// NewSetWebhook creates SetWebhook.
func NewSetWebhook(url string, file *args.File, fn ...func(*SetWebhook)) *SetWebhook {
	m := &SetWebhook{
		URL: url,
	}
	if file != nil {
		m.File = file.Local
		m.Certificate = file.Reference
	}
	for _, o := range fn {
		o(m)
	}
	return m
}

// IsMultipart returns true if method is multipart and false otherwise.
func (sw *SetWebhook) IsMultipart() bool {
	return sw.File != nil
}

// Method returns telegram api method endpoint.
func (sw *SetWebhook) Method() string {
	return "setWebhook"
}
