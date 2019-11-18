package option

import (
	"github.com/irisked/gogram/telegram/method"
)

// ReplyToMessageID returns struct for setting optional ReplyToMessageID field.
func ReplyToMessageID(msgID int) interface {
	replyToMessageIDSetter
} {
	return &replyToMessageIDOption{msgID}
}

// replyToMessageIDSetter is an interface for setting replyToMessageId option.
type replyToMessageIDSetter interface {
	method.SendMessageOption
	method.SendGameOption
	method.SendPhotoOption
	method.SendAudioOption
	method.SendDocumentOption
	method.SendVideoOption
	method.SendAnimationOption
	method.SendVoiceOption
	method.SendVideoNoteOption
	method.SendContactOption
	method.SendLocationOption
	method.SendVenueOption
}

type replyToMessageIDOption struct {
	msgID int
}

func (r *replyToMessageIDOption) ApplySendMessageOption(sm *method.SendMessage) {
	sm.ReplyToMessageID = r.msgID
}

func (r *replyToMessageIDOption) ApplySendGameOption(sm *method.SendGame) {
	sm.ReplyToMessageID = r.msgID
}

func (r *replyToMessageIDOption) ApplySendPhotoOption(sf *method.SendPhoto) {
	sf.ReplyToMessageID = r.msgID
}

func (r *replyToMessageIDOption) ApplySendAudioOption(sa *method.SendAudio) {
	sa.ReplyToMessageID = r.msgID
}

func (r *replyToMessageIDOption) ApplySendDocumentOption(sd *method.SendDocument) {
	sd.ReplyToMessageID = r.msgID
}

func (r *replyToMessageIDOption) ApplySendVideoOption(sv *method.SendVideo) {
	sv.ReplyToMessageID = r.msgID
}

func (r *replyToMessageIDOption) ApplySendAnimationOption(sa *method.SendAnimation) {
	sa.ReplyToMessageID = r.msgID
}

func (r *replyToMessageIDOption) ApplySendVoiceOption(sv *method.SendVoice) {
	sv.ReplyToMessageID = r.msgID
}

func (r *replyToMessageIDOption) ApplySendVideoNoteOption(sv *method.SendVideoNote) {
	sv.ReplyToMessageID = r.msgID
}

func (r *replyToMessageIDOption) ApplySendContactOption(sv *method.SendContact) {
	sv.ReplyToMessageID = r.msgID
}

func (r *replyToMessageIDOption) ApplySendLocationOption(sv *method.SendLocation) {
	sv.ReplyToMessageID = r.msgID
}

func (r *replyToMessageIDOption) ApplySendVenueOption(sv *method.SendVenue) {
	sv.ReplyToMessageID = r.msgID
}
