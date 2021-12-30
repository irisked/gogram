package option

import (
	"gogram/telegram/method"
)

// VCard returns struct for setting optional VCard field.
func VCard(vcard string) interface {
	hasVCard
} {
	return &vcardOption{vcard}
}

type hasVCard interface {
	method.SendContactOption
}

type vcardOption struct {
	vcard string
}

// ApplySendContactOption ApplySendContactOption
func (vo *vcardOption) ApplySendContactOption(sc *method.SendContact) {
	sc.VCard = vo.vcard
}
