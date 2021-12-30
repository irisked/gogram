package option

import (
	"gogram/telegram/method"
)

// DisableNotification returns struct for setting optional DisableNotification field.
func DisableNotification() interface {
	disableNotificationOptionSetter
} {
	return &disableNotificationOption{true}
}

// disableNotificationOptionSetter is an interface for setting disableNotification option.
type disableNotificationOptionSetter interface {
	method.SendMessageOption
	method.SendGameOption
	method.ForwardMessageOption
	method.SendPhotoOption
	method.SendAudioOption
	method.SendDocumentOption
	method.SendVideoOption
	method.SendAnimationOption
	method.SendVoiceOption
	method.SendVideoNoteOption
	method.PinChatMessageOption
	method.SendContactOption
	method.SendLocationOption
	method.SendVenueOption
}

type disableNotificationOption struct {
	disable bool
}

func (d *disableNotificationOption) ApplySendMessageOption(sm *method.SendMessage) {
	sm.DisableNotification = d.disable
}

func (d *disableNotificationOption) ApplySendGameOption(sm *method.SendGame) {
	sm.DisableNotification = d.disable
}

func (d *disableNotificationOption) ApplyForwardMessageOption(fm *method.ForwardMessage) {
	fm.DisableNotification = d.disable
}

func (d *disableNotificationOption) ApplySendPhotoOption(sf *method.SendPhoto) {
	sf.DisableNotification = d.disable
}

func (d *disableNotificationOption) ApplySendAudioOption(sa *method.SendAudio) {
	sa.DisableNotification = d.disable
}

func (d *disableNotificationOption) ApplySendDocumentOption(sd *method.SendDocument) {
	sd.DisableNotification = d.disable
}

func (d *disableNotificationOption) ApplySendVideoOption(sv *method.SendVideo) {
	sv.DisableNotification = d.disable
}

func (d *disableNotificationOption) ApplySendAnimationOption(sa *method.SendAnimation) {
	sa.DisableNotification = d.disable
}

func (d *disableNotificationOption) ApplySendVoiceOption(sv *method.SendVoice) {
	sv.DisableNotification = d.disable
}

func (d *disableNotificationOption) ApplySendVideoNoteOption(sv *method.SendVideoNote) {
	sv.DisableNotification = d.disable
}

func (d *disableNotificationOption) ApplyPinChatMessageOption(sv *method.PinChatMessage) {
	sv.DisableNotification = d.disable
}

func (d *disableNotificationOption) ApplySendContactOption(sv *method.SendContact) {
	sv.DisableNotification = d.disable
}

func (d *disableNotificationOption) ApplySendLocationOption(sv *method.SendLocation) {
	sv.DisableNotification = d.disable
}

func (d *disableNotificationOption) ApplySendVenueOption(sv *method.SendVenue) {
	sv.DisableNotification = d.disable
}
