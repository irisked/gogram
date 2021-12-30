package option

import (
	"gogram/telegram/keyboard"
	"gogram/telegram/method"
)

// Keyboard returns struct for setting optional Keyboard field.
func Keyboard(kb keyboard.Keyboard) interface {
	keyboardOptionSetter
} {
	return &replyMarkupOption{kb}
}

// keyboardOptionSetter is an interface for setting replyMarkup option.
type keyboardOptionSetter interface {
	method.EditMessageTextOption
	method.SendMessageOption
	method.SendGameOption
	method.EditMessageCaptionOption
	method.EditMessageMediaOption
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

type replyMarkupOption struct {
	kb keyboard.Keyboard
}

func (r *replyMarkupOption) ApplySendMessageOption(sm *method.SendMessage) {
	sm.ReplyMarkup = r.kb
}

func (r *replyMarkupOption) ApplyEditMessageTextOption(emt *method.EditMessageText) {
	emt.ReplyMarkup = r.kb
}

func (r *replyMarkupOption) ApplySendGameOption(emt *method.SendGame) {
	emt.ReplyMarkup = r.kb
}

func (r *replyMarkupOption) ApplyEditMessageCaptionOption(emc *method.EditMessageCaption) {
	emc.ReplyMarkup = r.kb
}

func (r *replyMarkupOption) ApplyEditMessageMediaOption(emm *method.EditMessageMedia) {
	emm.ReplyMarkup = r.kb
}

func (r *replyMarkupOption) ApplySendPhotoOption(sf *method.SendPhoto) {
	sf.ReplyMarkup = r.kb
}

func (r *replyMarkupOption) ApplySendAudioOption(sa *method.SendAudio) {
	sa.ReplyMarkup = r.kb
}

func (r *replyMarkupOption) ApplySendDocumentOption(sd *method.SendDocument) {
	sd.ReplyMarkup = r.kb
}

func (r *replyMarkupOption) ApplySendVideoOption(sv *method.SendVideo) {
	sv.ReplyMarkup = r.kb
}

func (r *replyMarkupOption) ApplySendAnimationOption(sa *method.SendAnimation) {
	sa.ReplyMarkup = r.kb
}

func (r *replyMarkupOption) ApplySendVoiceOption(sv *method.SendVoice) {
	sv.ReplyMarkup = r.kb
}

func (r *replyMarkupOption) ApplySendVideoNoteOption(svn *method.SendVideoNote) {
	svn.ReplyMarkup = r.kb
}

func (r *replyMarkupOption) ApplySendContactOption(svn *method.SendContact) {
	svn.ReplyMarkup = r.kb
}

func (r *replyMarkupOption) ApplySendLocationOption(svn *method.SendLocation) {
	svn.ReplyMarkup = r.kb
}

func (r *replyMarkupOption) ApplySendVenueOption(svn *method.SendVenue) {
	svn.ReplyMarkup = r.kb
}
