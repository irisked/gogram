package option

import (
	"gogram/telegram/method"
)

// LivePeriod returns struct for setting optional LivePeriod field.
func LivePeriod(p int) interface {
	livePeriodSetter
} {
	return &livePeriodOption{p}
}

type livePeriodSetter interface {
	method.SendLocationOption
}

type livePeriodOption struct {
	p int
}

func (lpo *livePeriodOption) ApplySendLocationOption(sl *method.SendLocation) {
	sl.LivePeriod = lpo.p
}
