package option

import (
	"github.com/irisked/gogram/telegram/method"
	"github.com/irisked/gogram/types"
)

// AllowedUpdates returns struct for setting optional AllowedUpdates field.
func AllowedUpdates(updateTypes ...types.UpdateType) interface {
	allowedUpdatesSetter
} {
	return &allowedUpdateOption{updateTypes}
}

// allowedUpdatesSetter is an interface for setting cacheTime option.
type allowedUpdatesSetter interface {
	method.SetWebhookOption
}

type allowedUpdateOption struct {
	allowedUpdates []types.UpdateType
}

func (auo *allowedUpdateOption) ApplySetWebhookOption(sw *method.SetWebhook) {
	sw.AllowedUpdates = auo.allowedUpdates
}
