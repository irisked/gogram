package gogram

import (
	"os"

	"github.com/irisked/gogram/params"
	"github.com/irisked/gogram/telegram/method"
	"github.com/irisked/gogram/telegram/keyboard"
	"github.com/irisked/gogram/telegram"
	"github.com/irisked/gogram/types"
)

// Context its gogram context.
type Context struct {
	update types.Update
	tg telegram.Telegram
	State map[string]interface{}
}

// NewContext creates context.
func NewContext(telegram telegram.Telegram, update types.Update) *Context {
	context := new(Context)
	context.tg = telegram
	context.update = update
	context.State = make(map[string]interface{})
	return context
}

// Update returns raw update.
func (c *Context) Update() types.Update {
	return c.update
}

// Telegram returns telegram.
func (c *Context) Telegram() telegram.Telegram {
	return c.tg
}

// Message returns message.
func (c *Context) Message() *types.Message {
	return c.update.Message
}

// EditedMessage returns edited message.
func (c *Context) EditedMessage() *types.Message {
	return c.update.EditedMessage
}

// InlineQuery returns edited message.
func (c *Context) InlineQuery() *types.InlineQuery {
	return c.update.InlineQuery
}

// ShippingQuery returns edited message.
func (c *Context) ShippingQuery() *types.ShippingQuery {
	return c.update.ShippingQuery
}

// PreCheckoutQuery returns edited message.
func (c *Context) PreCheckoutQuery() *types.PreCheckoutQuery {
	return c.update.PreCheckoutQuery
}

// ChosenInlineResult returns edited message.
func (c *Context) ChosenInlineResult() *types.ChosenInlineResult {
	return c.update.ChosenInlineResult
}

// ChannelPost returns edited message.
func (c *Context) ChannelPost() *types.Message {
	return c.update.ChannelPost
}

// EditedChannelPost returns edited message.
func (c *Context) EditedChannelPost() *types.Message {
	return c.update.EditedChannelPost
}

// CallbackQuery returns edited message.
func (c *Context) CallbackQuery() *types.CallbackQuery {
	return c.update.CallbackQuery
}

// Chat returns chat.
func (c *Context) Chat() *types.Chat {
	switch c.update.Type() {
	case types.IncomingMessage:
		return c.update.Message.Chat
	case types.IncomingEditedMessage:
		return c.update.EditedMessage.Chat
	case types.IncomingChannelPost:
		return c.update.ChannelPost.Chat
	case types.IncomingCallbackQuery:
		return c.update.CallbackQuery.Message.Chat
	case types.IncomingEditedChannelPost:
		return c.update.EditedChannelPost.Chat
	default:
		return nil
	}
}

// ChatID returns chat id.
func (c *Context) ChatID() int64 {
	return c.Chat().ID
}

// From returns user that sends an update.
func (c *Context) From() *types.User {
	switch c.update.Type() {
	case types.IncomingMessage:
		return c.update.Message.From
	case types.IncomingEditedMessage:
		return c.update.EditedMessage.From
	case types.IncomingChannelPost:
		return c.update.ChannelPost.From
	case types.IncomingEditedChannelPost:
		return c.update.EditedChannelPost.From
	case types.IncomingInlineQuery:
		return c.update.InlineQuery.From
	case types.IncomingChosenInlineResult:
		return c.update.ChosenInlineResult.From
	case types.IncomingCallbackQuery:
		return c.update.CallbackQuery.Message.From
	case types.IncomingShippingQuery:
		return c.update.ShippingQuery.From
	case types.IncomingPreCheckoutQuery:
		return c.update.PreCheckoutQuery.From
	default:
		return nil
	}
}

// Reply performs sendMessage telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.DisableNotification()
// option.ReplyToMessageID(id int), option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendmessage
func (c *Context) Reply(text params.TextData, options ...method.SendMessageOption) (types.Message, error) {
	return c.tg.SendMessage(c.ChatID(), text, options...)
}

// ReplyWithPhoto performs sendPhoto telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendphoto
func (c *Context) ReplyWithPhoto(photo *params.PhotoConfig, options ...method.SendPhotoOption) (types.Message, error) {
	return c.tg.SendPhoto(c.ChatID(), photo, options...)
}

// ReplyWithAudio performs sendAudio telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendaudio
func (c *Context) ReplyWithAudio(audio *params.AudioConfig, options ...method.SendAudioOption) (types.Message, error) {
	return c.tg.SendAudio(c.ChatID(), audio, options...)
}

// ReplyWithVideo performs sendVideo telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendvideo
func (c *Context) ReplyWithVideo(video *params.VideoConfig, options ...method.SendVideoOption) (types.Message, error) {
	return c.tg.SendVideo(c.ChatID(), video, options...)
}

// ReplyWithAnimation performs sendAnimation telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendanimation
func (c *Context) ReplyWithAnimation(animation *params.AnimationConfig, options ...method.SendAnimationOption) (types.Message, error) {
	return c.tg.SendAnimation(c.ChatID(), animation, options...)
}

// ReplyWithDocument performs sendDocument telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#senddocument
func (c *Context) ReplyWithDocument(document *params.DocumentConfig, options ...method.SendDocumentOption) (types.Message, error) {
	return c.tg.SendDocument(c.ChatID(), document, options...)
}

// ReplyWithVoice performs sendVoice telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendvoice
func (c *Context) ReplyWithVoice(voice *params.VoiceConfig, options ...method.SendVoiceOption) (types.Message, error) {
	return c.tg.SendVoice(c.ChatID(), voice, options...)
}

// ReplyWithVideoNote performs sendVideoNote telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendvideonote
func (c *Context) ReplyWithVideoNote(videoNote *params.VideoNoteConfig, options ...method.SendVideoNoteOption) (types.Message, error) {
	return c.tg.SendVideoNote(c.ChatID(), videoNote, options...)
}

// ReplyWithMediaGroup performs sendMediaGroup telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
//
// Reference: https://core.telegram.org/bots/api#sendmediagroup
func (c *Context) ReplyWithMediaGroup(mediaGroup []params.MediaGroupConfig, options ...method.SendMediaGroupOption) (types.Message, error) {
	return c.tg.SendMediaGroup(c.ChatID(), mediaGroup, options...)
}

// ReplyWithLocation performs sendLocation telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendlocation
func (c *Context) ReplyWithLocation(location types.Location, options ...method.SendLocationOption) (types.Message, error) {
	return c.tg.SendLocation(c.ChatID(), location, options...)
}

// TODO: https://core.telegram.org/bots/api#editmessagelivelocation
// TODO: https://core.telegram.org/bots/api#stopmessagelivelocation

// ReplyWithVenue performs sendVenue telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard), option.Foursquare(id, type string)
//
// Reference: https://core.telegram.org/bots/api#sendvenue
func (c *Context) ReplyWithVenue(venue types.Venue, options ...method.SendVenueOption) (types.Message, error) {
	return c.tg.SendVenue(c.ChatID(), venue, options...)
}

// ReplyWithContact performs sendContact telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendcontact
func (c *Context) ReplyWithContact(contact types.Contact, options ...method.SendContactOption) (types.Message, error) {
	return c.tg.SendContact(c.ChatID(), contact, options...)
}

// ReplyWithChatAction performs sendChatAction telegram api method.
//
// Reference: https://core.telegram.org/bots/api#sendchataction
func (c *Context) ReplyWithChatAction(action types.ChatAction) (types.Message, error) {
	return c.tg.SendChatAction(c.ChatID(), action)
}

// KickChatMember performs kickChatMember telegram api method.
//
// options:
//
// option.UntilDate(date int)
//
// Reference: https://core.telegram.org/bots/api#kickchatmember
// TODO: swap userID <-> ChatID in params
func (c *Context) KickChatMember(userID int, options ...method.KickChatMemberOption) (bool, error) {
	return c.tg.KickChatMember(userID, c.ChatID(), options...)
}

// UnbanChatMember performs unbanChatMember telegram api method.
//
// options:
//
// option.UntilDate(date int)
//
// Reference: https://core.telegram.org/bots/api#unbanchatmember
func (c *Context) UnbanChatMember(userID int) (bool, error) {
	return c.tg.UnbanChatMember(userID, c.ChatID())
}

// RestrictChatMember performs restrictChatMember telegram api method.
//
// options:
//
// option.CanSendMessages(), options.CanSendMediaMessages(),
// option.CanSendOtherMessages(), option.CanAddWebPagePreview()
//
// Reference: https://core.telegram.org/bots/api#restrictchatmember
func (c *Context) RestrictChatMember(userID int, options ...method.RestrictChatMemberOption) (bool, error) {
	return c.tg.RestrictChatMember(userID, c.ChatID(), options...)
}

// PromoteChatMember performs promoteChatMember telegram api method.
//
// options:
//
// option.CanChangeInfo(), option.CanPostMessages(), option.CanDeleteMessages(), option.CanEditMessages()
// option.CanInviteUsers(), option.CanRestrictUsers(), option.CanPinMessages(), option.CanPromoteMembers()
//
// Reference: https://core.telegram.org/bots/api#promotechatmember
func (c *Context) PromoteChatMember(userID int, options ...method.PromoteChatMemberOption) (bool, error) {
	return c.tg.PromoteChatMember(userID, c.ChatID(), options...)
}

// ExportChatInviteLink performs exportChatInviteLink telegram api method.
//
// Reference: https://core.telegram.org/bots/api#exportchatinvitelink
func (c *Context) ExportChatInviteLink() (string, error) {
	return c.tg.ExportChatInviteLink(c.ChatID())
}

// SetChatPhoto performs setChatPhoto telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setchatphoto
func (c *Context) SetChatPhoto(file *os.File) (bool, error) {
	return c.tg.SetChatPhoto(c.ChatID(), file)
}

// DeleteChatPhoto performs deleteChatPhoto telegram api method.
//
// Reference: https://core.telegram.org/bots/api#deletechatphoto
func (c *Context) DeleteChatPhoto() (bool, error) {
	return c.tg.DeleteChatPhoto(c.ChatID())
}

// SetChatTitle performs setChatTitle telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setchattitle
func (c *Context) SetChatTitle(title string) (bool, error) {
	return c.tg.SetChatTitle(c.ChatID(), title)
}

// SetChatDescription performs setChatDescription telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setchatdescription
func (c *Context) SetChatDescription(description string) (bool, error) {
	return c.tg.SetChatDescription(c.ChatID(), description)
}

// PinChatMessage performs pinChatMessage telegram api method.
//
// Reference: https://core.telegram.org/bots/api#pinchatmessage
func (c *Context) PinChatMessage(messageID int, options ...method.PinChatMessageOption) (bool, error) {
	return c.tg.PinChatMessage(c.ChatID(), messageID, options...)
}

// UnpinChatMessage performs unpinChatMessage telegram api method.
//
// Reference: https://core.telegram.org/bots/api#unpinchatmessage
func (c *Context) UnpinChatMessage() (bool, error) {
	return c.tg.UnpinChatMessage(c.ChatID())
}

// LeaveChat performs leaveChat telegram api method.
//
// Reference: https://core.telegram.org/bots/api#leavechat
func (c *Context) LeaveChat() (bool, error) {
	return c.tg.LeaveChat(c.ChatID())
}

// GetChat performs getChat telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getchat
func (c *Context) GetChat() (types.Chat, error) {
	return c.tg.GetChat(c.ChatID())
}


// GetChatAdministrators performs getChatAdministrators telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getchatadministrators
func (c *Context) GetChatAdministrators() (types.ChatMember, error) {
	return c.tg.GetChatAdministrators(c.ChatID())
}

// GetChatMembersCount performs getChatMembersCount telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getchatmemberscount
func (c *Context) GetChatMembersCount() (int, error) {
	return c.tg.GetChatMembersCount(c.ChatID())
}

// GetChatMember performs getChatMember telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getchatmember
func (c *Context) GetChatMember(userID int) (types.ChatMember, error) {
	return c.tg.GetChatMember(c.ChatID(), userID)
}

// SetChatStickerSet performs setChatStickerSet telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setchatstickerset
func (c *Context) SetChatStickerSet(stickerSetName string) (bool, error) {
	return c.tg.SetChatStickerSet(c.ChatID(), stickerSetName);
}

// DeleteChatStickerSet performs deleteChatStickerSet telegram api method.
//
// Reference: https://core.telegram.org/bots/api#deletechatstickerset
func (c *Context) DeleteChatStickerSet() (bool, error) {
	return c.tg.DeleteChatStickerSet(c.ChatID())
}

// AnswerCallbackQuery performs answerCallbackQuery telegram api method.
//
// options:
//
// option.Text(text string), option.ShowAlert(), option.URL(url string),
// option.CacheTime(time int)
//
// Reference: https://core.telegram.org/bots/api#answercallbackquery
func (c *Context) AnswerCallbackQuery(options ...method.AnswerCallbackQueryOption) (bool, error) {
	return c.tg.AnswerCallbackQuery(c.CallbackQuery().ID, options...);
}

// EditMessageText performs editMessageText telegram api method.
// Returns pointer to the message, if message was sent by the bot or nil-pointer otherwise.
//
// options:
//
// option.DisableWebPagePreview(), option.Keyboard(keyboard markup.Keyboard)
// 
// Reference: https://core.telegram.org/bots/api#editmessagetext
func (c *Context) EditMessageText(text params.TextData, options ...method.EditMessageTextOption) (*types.Message, error) {
	return c.tg.EditMessageText(c.CallbackQuery().MessageID(), text, options...)
}

// EditMessageCaption performs editMessageCaption telegram api method.
// Returns pointer to the message, if message was sent by the bot or nil-pointer otherwise.
//
// options:
//
// option.Keyboard(keyboard markup.Keyboard)
// 
// Reference: https://core.telegram.org/bots/api#editmessagecaption
func (c *Context) EditMessageCaption(caption params.TextData, options ...method.EditMessageCaptionOption) (*types.Message, error) {
	return c.tg.EditMessageCaption(c.CallbackQuery().MessageID(), caption, options...)
}

// EditMessageMedia performs editMessageMedia telegram api method.
// Returns pointer to the message, if message was sent by the bot or nil-pointer otherwise.
//
// options:
//
// option.Keyboard(keyboard markup.Keyboard)
// 
// Reference: https://core.telegram.org/bots/api#editmessagemedia
func (c *Context) EditMessageMedia(media params.InputMediaConfig, options ...method.EditMessageMediaOption) (*types.Message, error) {
	return c.tg.EditMessageMedia(c.CallbackQuery().MessageID(), media, options...)
}

// EditMessageReplyMarkup performs editMessageReplyMarkup telegram api method.
// Returns pointer to the message, if message was sent by the bot or nil-pointer otherwise.
//
// options:
//
// option.Keyboard(keyboard markup.Keyboard)
// 
// Reference: https://core.telegram.org/bots/api#editmessagereplymarkup
func (c *Context) EditMessageReplyMarkup(kb keyboard.Keyboard) (*types.Message, error) {
	return c.tg.EditMessageReplyMarkup(c.CallbackQuery().MessageID(), kb)
}

// DeleteMessage performs deleteMessage telegram api method.
// 
// https://core.telegram.org/bots/api#deletemessage
func (c *Context) DeleteMessage(messageID int) (bool, error) {
	return c.tg.DeleteMessage(c.ChatID(), messageID)
}

// TODO: stickers methods.
// TODO: inline mode methods.
// TODO: payments methods.
// TODO: telegram passport methods.
// TODO: games methods.