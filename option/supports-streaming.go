package option

import (
	"github.com/irisked/gogram/args"
)

// SupportsStreaming returns struct for setting optional SupportsStreaming field.
func SupportsStreaming() interface {
	supportsStreamingOptionSetter
} {
	return &supportsStreamingOption{true}
}

// supportsStreamingOptionSetter is an interface for setting duration option.
type supportsStreamingOptionSetter interface {
	args.VideoConfigOption
}

type supportsStreamingOption struct {
	support bool
}

func (sso *supportsStreamingOption) ApplyVideoConfigOption(parameter *args.VideoConfig) {
	parameter.SupportsStreaming = sso.support
}
