package telegram

import (
	"encoding/json"
	"gogram/args"
	"gogram/telegram/internal/net"
	"gogram/telegram/keyboard"
	"gogram/telegram/method"
	"gogram/types"
	"os"
)

// Telegram telegram.
type Telegram struct {
	client *net.Client
}

// New creates Telegram.
func New(token string) *Telegram {
	telegram := new(Telegram)
	telegram.client = net.NewClient(token)
	return telegram
}

func (t *Telegram) process(m method.TelegramMethod, result interface{}) error {
	response, err := t.client.Call(m)
	if err != nil {
		return err
	}
	json.Unmarshal(response.Result, result)
	return nil
}

// GetMe performs getMe telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getme
func (t *Telegram) GetMe() (types.User, error) {
	m := method.NewGetMe()
	user := types.User{}
	err := t.process(m, &user)
	return user, err
}

// GetUpdates performs getMe telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getupdates
func (t *Telegram) GetUpdates(options ...method.GetUpdatesOption) ([]types.Update, error) {
	m := method.NewGetUpdates(options...)
	var updates []types.Update
	err := t.process(m, &updates)
	return updates, err
}

// SendMessage performs sendMessage telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.DisableNotification()
// option.ReplyToMessageID(id int), option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendmessage
func (t *Telegram) SendMessage(chatID int64, text args.TextData, options ...method.SendMessageOption) (types.Message, error) {
	m := method.NewSendMessage(chatID, text, options...)
	message := types.Message{}
	err := t.process(m, &message)
	return message, err
}

// ForwardMessage performs forwardMessage telegram api method.
//
// options:
//
// option.DisableNotification()
//
// Reference: https://core.telegram.org/bots/api#forwardmessage
func (t *Telegram) ForwardMessage(chatID, fromChatID, messageID int, options ...method.ForwardMessageOption) (types.Message, error) {
	m := method.NewForwardMessage(chatID, fromChatID, messageID, options...)
	message := types.Message{}
	err := t.process(m, &message)
	return message, err
}

// SendPhoto performs sendPhoto telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendphoto
func (t *Telegram) SendPhoto(chatID int64, photo *args.PhotoConfig, options ...method.SendPhotoOption) (types.Message, error) {
	m := method.NewSendPhoto(chatID, photo, options...)
	message := types.Message{}
	err := t.process(m, &message)
	return message, err
}

// SendAudio performs sendAudio telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendaudio
func (t *Telegram) SendAudio(chatID int64, audio *args.AudioConfig, options ...method.SendAudioOption) (types.Message, error) {
	m := method.NewSendAudio(chatID, audio, options...)
	message := types.Message{}
	err := t.process(m, &message)
	return message, err
}

// SendDocument performs sendDocument telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#senddocument
func (t *Telegram) SendDocument(chatID int64, document *args.DocumentConfig, options ...method.SendDocumentOption) (types.Message, error) {
	m := method.NewSendDocument(chatID, document, options...)
	message := types.Message{}
	err := t.process(m, &message)
	return message, err
}

// SendVideo performs sendVideo telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendvideo
func (t *Telegram) SendVideo(chatID int64, video *args.VideoConfig, options ...method.SendVideoOption) (types.Message, error) {
	m := method.NewSendVideo(chatID, video, options...)
	message := types.Message{}
	err := t.process(m, &message)
	return message, err
}

// SendAnimation performs sendAnimation telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendanimation
func (t *Telegram) SendAnimation(chatID int64, animation *args.AnimationConfig, options ...method.SendAnimationOption) (types.Message, error) {
	m := method.NewSendAnimation(chatID, animation, options...)
	message := types.Message{}
	err := t.process(m, &message)
	return message, err
}

// SendVoice performs sendVoice telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendvoice
func (t *Telegram) SendVoice(chatID int64, voice *args.VoiceConfig, options ...method.SendVoiceOption) (types.Message, error) {
	m := method.NewSendVoice(chatID, voice, options...)
	message := types.Message{}
	err := t.process(m, &message)
	return message, err
}

// SendVideoNote performs sendVideoNote telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendvideonote
func (t *Telegram) SendVideoNote(chatID int64, videoNote *args.VideoNoteConfig, options ...method.SendVideoNoteOption) (types.Message, error) {
	m := method.NewSendVideoNote(chatID, videoNote, options...)
	message := types.Message{}
	err := t.process(m, &message)
	return message, err
}

// SendMediaGroup performs sendMediaGroup telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
//
// Reference: https://core.telegram.org/bots/api#sendmediagroup
// TODO: parse response to array.
func (t *Telegram) SendMediaGroup(chatID int64, mediaGroup []args.MediaGroupConfig, options ...method.SendMediaGroupOption) (types.Message, error) {
	m := method.NewSendMediaGroup(chatID, mediaGroup, options...)
	message := types.Message{}
	err := t.process(m, &message)
	return message, err
}

// SendLocation performs sendLocation telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendlocation
func (t *Telegram) SendLocation(chatID int64, location types.Location, options ...method.SendLocationOption) (types.Message, error) {
	m := method.NewSendLocation(chatID, location, options...)
	message := types.Message{}
	err := t.process(m, &message)
	return message, err
}

// TODO: https://core.telegram.org/bots/api#editmessagelivelocation
// TODO: https://core.telegram.org/bots/api#stopmessagelivelocation

// SendVenue performs sendVenue telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard), option.Foursquare(id, type string)
//
// Reference: https://core.telegram.org/bots/api#sendvenue
func (t *Telegram) SendVenue(chatID int64, venue types.Venue, options ...method.SendVenueOption) (types.Message, error) {
	m := method.NewSendVenue(chatID, venue, options...)
	message := types.Message{}
	err := t.process(m, &message)
	return message, err
}

// SendContact performs sendContact telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendcontact
func (t *Telegram) SendContact(chatID int64, contact types.Contact, options ...method.SendContactOption) (types.Message, error) {
	m := method.NewSendContact(chatID, contact, options...)
	message := types.Message{}
	err := t.process(m, &message)
	return message, err
}

// SendChatAction performs sendChatAction telegram api method.
//
// Reference: https://core.telegram.org/bots/api#sendchataction
func (t *Telegram) SendChatAction(chatID int64, action types.ChatAction) (types.Message, error) {
	m := method.NewSendChatAction(chatID, action)
	message := types.Message{}
	err := t.process(m, &message)
	return message, err
}

// GetUserProfilePhotos performs getUserProfilePhotos telegram api method.
//
// options:
//
// option.Offset(o int), option.Limit(l int)
//
// Reference: https://core.telegram.org/bots/api#getuserprofilephotos
func (t *Telegram) GetUserProfilePhotos(userID int, options ...method.GetUserProfilePhotosOption) (types.UserProfilePhotos, error) {
	m := method.NewGetUserProfilePhotos(userID, options...)
	photos := types.UserProfilePhotos{}
	err := t.process(m, &photos)
	return photos, err
}

// GetFile performs getFile telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getfile
func (t *Telegram) GetFile(fileID string) (types.File, error) {
	m := method.NewGetFile(fileID)
	file := types.File{}
	err := t.process(m, &file)
	return file, err
}

// KickChatMember performs kickChatMember telegram api method.
//
// options:
//
// option.UntilDate(date int)
//
// Reference: https://core.telegram.org/bots/api#kickchatmember
func (t *Telegram) KickChatMember(userID int, chatID int64, options ...method.KickChatMemberOption) (bool, error) {
	m := method.NewKickChatMember(userID, chatID, options...)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// UnbanChatMember performs unbanChatMember telegram api method.
//
// options:
//
// option.UntilDate(date int)
//
// Reference: https://core.telegram.org/bots/api#unbanchatmember
func (t *Telegram) UnbanChatMember(userID int, chatID int64) (bool, error) {
	m := method.NewUnbanChatMember(userID, chatID)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// RestrictChatMember performs restrictChatMember telegram api method.
//
// options:
//
// option.CanSendMessages(), options.CanSendMediaMessages(),
// option.CanSendOtherMessages(), option.CanAddWebPagePreview()
//
// Reference: https://core.telegram.org/bots/api#restrictchatmember
// TODO: rework true/false
func (t *Telegram) RestrictChatMember(userID int, chatID int64, options ...method.RestrictChatMemberOption) (bool, error) {
	m := method.NewRestrictChatMember(userID, chatID, options...)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// PromoteChatMember performs promoteChatMember telegram api method.
//
// options:
//
// option.CanChangeInfo(), option.CanPostMessages(), option.CanDeleteMessages(), option.CanEditMessages()
// option.CanInviteUsers(), option.CanRestrictUsers(), option.CanPinMessages(), option.CanPromoteMembers()
//
// Reference: https://core.telegram.org/bots/api#promotechatmember
// TODO: rework true/false
func (t *Telegram) PromoteChatMember(userID int, chatID int64, options ...method.PromoteChatMemberOption) (bool, error) {
	m := method.NewPromoteChatMember(userID, chatID, options...)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// ExportChatInviteLink performs exportChatInviteLink telegram api method.
//
// Reference: https://core.telegram.org/bots/api#exportchatinvitelink
func (t *Telegram) ExportChatInviteLink(chatID int64) (string, error) {
	m := method.NewExportChatInviteLink(chatID)
	var result string
	err := t.process(m, &result)
	return result, err
}

// SetChatPhoto performs setChatPhoto telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setchatphoto
func (t *Telegram) SetChatPhoto(chatID int64, file *os.File) (bool, error) {
	m := method.NewSetChatPhoto(chatID, file)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// DeleteChatPhoto performs deleteChatPhoto telegram api method.
//
// Reference: https://core.telegram.org/bots/api#deletechatphoto
func (t *Telegram) DeleteChatPhoto(chatID int64) (bool, error) {
	m := method.NewDeleteChatPhoto(chatID)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// SetChatTitle performs setChatTitle telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setchattitle
func (t *Telegram) SetChatTitle(chatID int64, title string) (bool, error) {
	m := method.NewSetChatTitle(chatID, title)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// SetChatDescription performs setChatDescription telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setchatdescription
func (t *Telegram) SetChatDescription(chatID int64, description string) (bool, error) {
	m := method.NewSetChatDescription(chatID, description)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// PinChatMessage performs pinChatMessage telegram api method.
//
// Reference: https://core.telegram.org/bots/api#pinchatmessage
func (t *Telegram) PinChatMessage(chatID int64, messageID int, options ...method.PinChatMessageOption) (bool, error) {
	m := method.NewPinChatMessage(chatID, messageID, options...)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// UnpinChatMessage performs unpinChatMessage telegram api method.
//
// Reference: https://core.telegram.org/bots/api#unpinchatmessage
func (t *Telegram) UnpinChatMessage(chatID int64) (bool, error) {
	m := method.NewUnpinChatMessage(chatID)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// LeaveChat performs leaveChat telegram api method.
//
// Reference: https://core.telegram.org/bots/api#leavechat
func (t *Telegram) LeaveChat(chatID int64) (bool, error) {
	m := method.NewLeaveChat(chatID)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// GetChat performs getChat telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getchat
func (t *Telegram) GetChat(chatID int64) (types.Chat, error) {
	m := method.NewGetChat(chatID)
	chat := types.Chat{}
	err := t.process(m, &chat)
	return chat, err
}

// GetChatAdministrators performs getChatAdministrators telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getchatadministrators
// TODO: results is array
func (t *Telegram) GetChatAdministrators(chatID int64) (types.ChatMember, error) {
	m := method.NewGetChatAdministrators(chatID)
	administrators := types.ChatMember{}
	err := t.process(m, &administrators)
	return administrators, err
}

// GetChatMembersCount performs getChatMembersCount telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getchatmemberscount
func (t *Telegram) GetChatMembersCount(chatID int64) (int, error) {
	m := method.NewGetChatMembersCount(chatID)
	var result int
	err := t.process(m, &result)
	return result, err
}

// GetChatMember performs getChatMember telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getchatmember
func (t *Telegram) GetChatMember(chatID int64, userID int) (types.ChatMember, error) {
	m := method.NewGetChatMember(chatID, userID)
	member := types.ChatMember{}
	err := t.process(m, &member)
	return member, err
}

// SetChatStickerSet performs setChatStickerSet telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setchatstickerset
func (t *Telegram) SetChatStickerSet(chatID int64, stickerSetName string) (bool, error) {
	m := method.NewSetChatStickerSet(chatID, stickerSetName)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// DeleteChatStickerSet performs deleteChatStickerSet telegram api method.
//
// Reference: https://core.telegram.org/bots/api#deletechatstickerset
func (t *Telegram) DeleteChatStickerSet(chatID int64) (bool, error) {
	m := method.NewDeleteChatStickerSet(chatID)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// AnswerCallbackQuery performs answerCallbackQuery telegram api method.
//
// options:
//
// option.Text(text string), option.ShowAlert(), option.URL(url string),
// option.CacheTime(time int)
//
// Reference: https://core.telegram.org/bots/api#answercallbackquery
func (t *Telegram) AnswerCallbackQuery(id string, options ...method.AnswerCallbackQueryOption) (bool, error) {
	m := method.NewAnswerCallbackQuery(id, options...)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// EditMessageText performs editMessageText telegram api method.
// Returns pointer to the message, if message was sent by the bot or nil-pointer.
//
// options:
//
// option.DisableWebPagePreview(), option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#editmessagetext
func (t *Telegram) EditMessageText(messageID *args.MessageID, text args.TextData, options ...method.EditMessageTextOption) (*types.Message, error) {
	m := method.NewEditMessageText(messageID, text, options...)
	message := types.Message{}
	err := t.process(m, &message)
	return &message, err
}

// EditMessageCaption performs editMessageCaption telegram api method.
// Returns pointer to the message, if message was sent by the bot or nil-pointer.
//
// options:
//
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#editmessagecaption
func (t *Telegram) EditMessageCaption(messageID *args.MessageID, caption args.TextData, options ...method.EditMessageCaptionOption) (*types.Message, error) {
	m := method.NewEditMessageCaption(messageID, caption, options...)
	message := types.Message{}
	err := t.process(m, &message)
	return &message, err
}

// EditMessageMedia performs editMessageMedia telegram api method.
// Returns pointer to the message, if message was sent by the bot or nil-pointer.
//
// options:
//
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#editmessagemedia
func (t *Telegram) EditMessageMedia(messageID *args.MessageID, media args.InputMediaConfig, options ...method.EditMessageMediaOption) (*types.Message, error) {
	m := method.NewEditMessageMedia(messageID, media, options...)
	message := types.Message{}
	err := t.process(m, &message)
	return &message, err
}

// EditMessageReplyMarkup performs editMessageReplyMarkup telegram api method.
// Returns pointer to the message, if message was sent by the bot or nil-pointer.
//
// Reference: https://core.telegram.org/bots/api#editmessagereplymarkup
func (t *Telegram) EditMessageReplyMarkup(messageID *args.MessageID, kb keyboard.Keyboard) (*types.Message, error) {
	m := method.NewEditMessageReplyMarkup(messageID, kb)
	message := types.Message{}
	err := t.process(m, &message)
	return &message, err
}

// DeleteMessage performs deleteMessage telegram api method.
//
// https://core.telegram.org/bots/api#deletemessage
func (t *Telegram) DeleteMessage(chatID int64, messageID int) (bool, error) {
	m := method.NewDeleteMessage(chatID, messageID)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// TODO: stickers methods.
// TODO: inline mode methods.
// TODO: payments methods.
// TODO: telegram passport methods.
// TODO: games methods.
