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
func (t *Telegram) GetUpdates(opts ...method.Option[*method.GetUpdates]) ([]types.Update, error) {
	m := method.NewGetUpdates(opts...)
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
func (t *Telegram) SendMessage(chatID int64, text args.TextData, opts ...method.Option[*method.SendMessage]) (types.Message, error) {
	m := method.NewSendMessage(chatID, text, opts...)
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
func (t *Telegram) ForwardMessage(chatID, fromChatID, messageID int, opts ...method.Option[*method.ForwardMessage]) (types.Message, error) {
	m := method.NewForwardMessage(chatID, fromChatID, messageID, opts...)
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
func (t *Telegram) SendPhoto(chatID int64, photo *args.PhotoConfig, opts ...method.Option[*method.SendPhoto]) (types.Message, error) {
	m := method.NewSendPhoto(chatID, photo, opts...)
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
func (t *Telegram) SendAudio(chatID int64, audio *args.AudioConfig, opts ...method.Option[*method.SendAudio]) (types.Message, error) {
	m := method.NewSendAudio(chatID, audio, opts...)
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
func (t *Telegram) SendDocument(chatID int64, document *args.DocumentConfig, opts ...method.Option[*method.SendDocument]) (types.Message, error) {
	m := method.NewSendDocument(chatID, document, opts...)
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
func (t *Telegram) SendVideo(chatID int64, video *args.VideoConfig, opts ...method.Option[*method.SendVideo]) (types.Message, error) {
	m := method.NewSendVideo(chatID, video, opts...)
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
func (t *Telegram) SendAnimation(chatID int64, animation *args.AnimationConfig, opts ...method.Option[*method.SendAnimation]) (types.Message, error) {
	m := method.NewSendAnimation(chatID, animation, opts...)
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
func (t *Telegram) SendVoice(chatID int64, voice *args.VoiceConfig, opts ...method.Option[*method.SendVoice]) (types.Message, error) {
	m := method.NewSendVoice(chatID, voice, opts...)
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
func (t *Telegram) SendVideoNote(chatID int64, videoNote *args.VideoNoteConfig, opts ...method.Option[*method.SendVideoNote]) (types.Message, error) {
	m := method.NewSendVideoNote(chatID, videoNote, opts...)
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
func (t *Telegram) SendMediaGroup(chatID int64, mediaGroup []args.MediaGroupConfig, opts ...method.Option[*method.SendMediaGroup]) ([]types.Message, error) {
	m := method.NewSendMediaGroup(chatID, mediaGroup, opts...)
	var messages []types.Message
	err := t.process(m, &messages)
	return messages, err
}

// SendLocation performs sendLocation telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendlocation
func (t *Telegram) SendLocation(chatID int64, location types.Location, opts ...method.Option[*method.SendLocation]) (types.Message, error) {
	m := method.NewSendLocation(chatID, location, opts...)
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
func (t *Telegram) SendVenue(chatID int64, venue types.Venue, opts ...method.Option[*method.SendVenue]) (types.Message, error) {
	m := method.NewSendVenue(chatID, venue, opts...)
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
func (t *Telegram) SendContact(chatID int64, contact types.Contact, opts ...method.Option[*method.SendContact]) (types.Message, error) {
	m := method.NewSendContact(chatID, contact, opts...)
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
func (t *Telegram) GetUserProfilePhotos(userID int, opts ...method.Option[*method.GetUserProfilePhotos]) (types.UserProfilePhotos, error) {
	m := method.NewGetUserProfilePhotos(userID, opts...)
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
func (t *Telegram) KickChatMember(userID int, chatID int64, opts ...method.Option[*method.KickChatMember]) (bool, error) {
	m := method.NewKickChatMember(userID, chatID, opts...)
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
func (t *Telegram) RestrictChatMember(userID int, chatID int64, opts ...method.Option[*method.RestrictChatMember]) (bool, error) {
	m := method.NewRestrictChatMember(userID, chatID, opts...)
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
func (t *Telegram) PromoteChatMember(userID int, chatID int64, opts ...method.Option[*method.PromoteChatMember]) (bool, error) {
	m := method.NewPromoteChatMember(userID, chatID, opts...)
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
func (t *Telegram) PinChatMessage(chatID int64, messageID int, opts ...method.Option[*method.PinChatMessage]) (bool, error) {
	m := method.NewPinChatMessage(chatID, messageID, opts...)
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
func (t *Telegram) AnswerCallbackQuery(id string, opts ...method.Option[*method.AnswerCallbackQuery]) (bool, error) {
	m := method.NewAnswerCallbackQuery(id, opts...)
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
func (t *Telegram) EditMessageText(messageID *args.MessageID, text args.TextData, opts ...method.Option[*method.EditMessageText]) (*types.Message, error) {
	m := method.NewEditMessageText(messageID, text, opts...)
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
func (t *Telegram) EditMessageCaption(messageID *args.MessageID, caption args.TextData, opts ...method.Option[*method.EditMessageCaption]) (*types.Message, error) {
	m := method.NewEditMessageCaption(messageID, caption, opts...)
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
func (t *Telegram) EditMessageMedia(messageID *args.MessageID, media args.InputMediaConfig, opts ...method.Option[*method.EditMessageMedia]) (*types.Message, error) {
	m := method.NewEditMessageMedia(messageID, media, opts...)
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

// CopyMessage performs copyMessage telegram api method.
//
// options:
//
// option.CopyMessageCaption(), option.CopyMessageParseMode(),
// option.CopyMessageDisableNotification(), option.CopyMessageReplyToMessageID(id int),
// option.CopyMessageAllowCopy(), option.CopyMessageReplyMarkup()
//
// Reference: https://core.telegram.org/bots/api#copymessage
func (t *Telegram) CopyMessage(chatID, fromChatID int64, messageID int, opts ...method.Option[*method.CopyMessage]) (types.MessageID, error) {
	m := method.NewCopyMessage(chatID, fromChatID, messageID, opts...)
	var result types.MessageID
	err := t.process(m, &result)
	return result, err
}

// CopyMessages performs copyMessages telegram api method.
//
// options:
//
// option.CopyMessagesDisableNotification(), option.CopyMessagesRemoveSender()
//
// Reference: https://core.telegram.org/bots/api#copymessages
func (t *Telegram) CopyMessages(chatID, fromChatID int64, messageIDs []int, opts ...method.Option[*method.CopyMessages]) ([]types.MessageID, error) {
	m := method.NewCopyMessages(chatID, fromChatID, messageIDs, opts...)
	var result []types.MessageID
	err := t.process(m, &result)
	return result, err
}

// DeleteMessages performs deleteMessages telegram api method.
//
// Reference: https://core.telegram.org/bots/api#deletemessages
func (t *Telegram) DeleteMessages(chatID int64, messageIDs []int) (bool, error) {
	m := method.NewDeleteMessages(chatID, messageIDs)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// StopPoll performs stopPoll telegram api method.
//
// options:
//
// option.StopPollReplyMarkup()
//
// Reference: https://core.telegram.org/bots/api#stoppoll
func (t *Telegram) StopPoll(chatID int64, messageID int, opts ...method.Option[*method.StopPoll]) (types.Poll, error) {
	m := method.NewStopPoll(chatID, messageID, opts...)
	var result types.Poll
	err := t.process(m, &result)
	return result, err
}

// EditMessageLiveLocation performs editMessageLiveLocation telegram api method.
//
// options:
//
// option.EditMessageLiveLocationReplyMarkup()
//
// Reference: https://core.telegram.org/bots/api#editmessagelivelocation
func (t *Telegram) EditMessageLiveLocation(messageID *args.MessageID, location types.Location, opts ...method.Option[*method.EditMessageLiveLocation]) (*types.Message, error) {
	m := method.NewEditMessageLiveLocation(messageID, location, opts...)
	message := types.Message{}
	err := t.process(m, &message)
	return &message, err
}

// StopMessageLiveLocation performs stopMessageLiveLocation telegram api method.
//
// options:
//
// option.StopMessageLiveLocationReplyMarkup()
//
// Reference: https://core.telegram.org/bots/api#stopmessagelivelocation
func (t *Telegram) StopMessageLiveLocation(messageID *args.MessageID, opts ...method.Option[*method.StopMessageLiveLocation]) (*types.Message, error) {
	m := method.NewStopMessageLiveLocation(messageID, opts...)
	message := types.Message{}
	err := t.process(m, &message)
	return &message, err
}

// DeleteWebhook performs deleteWebhook telegram api method.
//
// options:
//
// option.DeleteWebhookDropPendingUpdates()
//
// Reference: https://core.telegram.org/bots/api#deletewebhook
func (t *Telegram) DeleteWebhook(opts ...method.Option[*method.DeleteWebhook]) (bool, error) {
	m := method.NewDeleteWebhook(opts...)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// GetStickerSet performs getStickerSet telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getstickerset
func (t *Telegram) GetStickerSet(name string) (types.StickerSet, error) {
	m := method.NewGetStickerSet(name)
	var result types.StickerSet
	err := t.process(m, &result)
	return result, err
}

// UploadStickerFile performs uploadStickerFile telegram api method.
//
// options:
//
// option.UploadStickerFileFormat(), option.UploadStickerFileType(),
// option.UploadStickerFilePNGSticker(), option.UploadStickerFileTSVGSticker(),
// option.UploadStickerFileWebmSticker(), option.UploadStickerFileThumbnail(),
// option.UploadStickerFileMaskPosition(), option.UploadStickerFileEmojiList()
//
// Reference: https://core.telegram.org/bots/api#uploadstickerfile
func (t *Telegram) UploadStickerFile(userID int64, sticker *os.File, opts ...method.Option[*method.UploadStickerFile]) (types.File, error) {
	m := method.NewUploadStickerFile(sticker, opts...)
	file := types.File{}
	err := t.process(m, &file)
	return file, err
}

// CreateNewStickerSet performs createNewStickerSet telegram api method.
//
// options:
//
// option.CreateNewStickerSetStickerType(), option.CreateNewStickerSetMaskPosition(),
// option.CreateNewStickerSetNeedsRepaired()
//
// Reference: https://core.telegram.org/bots/api#createnewstickerset
func (t *Telegram) CreateNewStickerSet(userID int64, name, title string, stickers []string, opts ...method.Option[*method.CreateNewStickerSet]) (bool, error) {
	m := method.NewCreateNewStickerSet(userID, name, title, stickers, opts...)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// DeleteStickerSet performs deleteStickerSet telegram api method.
//
// Reference: https://core.telegram.org/bots/api#deletestickerset
func (t *Telegram) DeleteStickerSet(name string) (bool, error) {
	m := method.NewDeleteStickerSet(name)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// AddStickerToSet performs addStickerToSet telegram api method.
//
// options:
//
// option.AddStickerToSetMaskPosition()
//
// Reference: https://core.telegram.org/bots/api#addstickertoset
func (t *Telegram) AddStickerToSet(userID int64, name, sticker string, opts ...method.Option[*method.AddStickerToSet]) (bool, error) {
	m := method.NewAddStickerToSet(userID, name, sticker, opts...)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// SetStickerPositionInSet performs setStickerPositionInSet telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setstickerpositioninset
func (t *Telegram) SetStickerPositionInSet(sticker string, position int) (bool, error) {
	m := method.NewSetStickerPositionInSet(sticker, position)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// DeleteStickerFromSet performs deleteStickerFromSet telegram api method.
//
// Reference: https://core.telegram.org/bots/api#deletestickerfromset
func (t *Telegram) DeleteStickerFromSet(sticker string) (bool, error) {
	m := method.NewDeleteStickerFromSet(sticker)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// ReplaceStickerInSet performs replaceStickerInSet telegram api method.
//
// Reference: https://core.telegram.org/bots/api#replacestickerinset
func (t *Telegram) ReplaceStickerInSet(userID int64, name, oldSticker, sticker string) (bool, error) {
	m := method.NewReplaceStickerInSet(userID, name, oldSticker, sticker)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// SetStickerEmojiList performs setStickerEmojiList telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setstickeremojilist
func (t *Telegram) SetStickerEmojiList(sticker string, emojiList []string) (bool, error) {
	m := method.NewSetStickerEmojiList(sticker, emojiList)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// SetStickerKeywords performs setStickerKeywords telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setstickerkeywords
func (t *Telegram) SetStickerKeywords(sticker string, keywords []string) (bool, error) {
	m := method.NewSetStickerKeywords(sticker, keywords)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// SetStickerMaskPosition performs setStickerMaskPosition telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setstickermaskposition
func (t *Telegram) SetStickerMaskPosition(sticker, maskPosition string) (bool, error) {
	m := method.NewSetStickerMaskPosition(sticker, maskPosition)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// SetStickerSetTitle performs setStickerSetTitle telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setstickersettitle
func (t *Telegram) SetStickerSetTitle(name, title string) (bool, error) {
	m := method.NewSetStickerSetTitle(name, title)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// SetStickerSetDescription performs setStickerSetDescription telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setstickersetdescription
func (t *Telegram) SetStickerSetDescription(name, description string) (bool, error) {
	m := method.NewSetStickerSetDescription(name, description)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// SetStickerSetThumbnail performs setStickerSetThumbnail telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setstickersetthumbnail
func (t *Telegram) SetStickerSetThumbnail(name, thumbnail string) (bool, error) {
	m := method.NewSetStickerSetThumbnail(name, thumbnail)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// SetCustomEmojiStickerSetThumbnail performs setCustomEmojiStickerSetThumbnail telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setcustomemojistickersetthumbnail
func (t *Telegram) SetCustomEmojiStickerSetThumbnail(name, customEmoji string) (bool, error) {
	m := method.NewSetCustomEmojiStickerSetThumbnail(name, customEmoji)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// GetCustomEmojiStickers performs getCustomEmojiStickers telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getcustomemojistickers
func (t *Telegram) GetCustomEmojiStickers(customEmojiIDs []string) ([]types.Sticker, error) {
	m := method.NewGetCustomEmojiStickers(customEmojiIDs)
	var result []types.Sticker
	err := t.process(m, &result)
	return result, err
}

// CreateForumTopic performs createForumTopic telegram api method.
//
// options:
//
// option.CreateForumTopicIconColor(), option.CreateForumTopicIconEmojiID()
//
// Reference: https://core.telegram.org/bots/api#createforumtopic
func (t *Telegram) CreateForumTopic(chatID int64, name string, opts ...method.Option[*method.CreateForumTopic]) (method.ForumTopic, error) {
	m := method.NewCreateForumTopic(chatID, name, opts...)
	var result method.ForumTopic
	err := t.process(m, &result)
	return result, err
}

// EditForumTopic performs editForumTopic telegram api method.
//
// options:
//
// option.EditForumTopicName(), option.EditForumTopicIconEmojiID()
//
// Reference: https://core.telegram.org/bots/api#editforumtopic
func (t *Telegram) EditForumTopic(chatID int64, messageID int, opts ...method.Option[*method.EditForumTopic]) (bool, error) {
	m := method.NewEditForumTopic(chatID, messageID, opts...)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// DeleteForumTopic performs deleteForumTopic telegram api method.
//
// Reference: https://core.telegram.org/bots/api#deleteforumtopic
func (t *Telegram) DeleteForumTopic(chatID int64, messageID int) (bool, error) {
	m := method.NewDeleteForumTopic(chatID, messageID)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// GetForumTopic performs getForumTopic telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getforumtopic
func (t *Telegram) GetForumTopic(chatID int64, messageID int) (method.ForumTopic, error) {
	m := method.NewGetForumTopic(chatID, messageID)
	var result method.ForumTopic
	err := t.process(m, &result)
	return result, err
}

// CloseForumTopic performs closeForumTopic telegram api method.
//
// Reference: https://core.telegram.org/bots/api#closeforumtopic
func (t *Telegram) CloseForumTopic(chatID int64, messageID int) (bool, error) {
	m := method.NewCloseForumTopic(chatID, messageID)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// ReopenForumTopic performs reopenForumTopic telegram api method.
//
// Reference: https://core.telegram.org/bots/api#reopenforumtopic
func (t *Telegram) ReopenForumTopic(chatID int64, messageID int) (bool, error) {
	m := method.NewReopenForumTopic(chatID, messageID)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// GetForumTopicIconStickers performs getForumTopicIconStickers telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getforumtopiciconstickers
func (t *Telegram) GetForumTopicIconStickers() ([]types.Sticker, error) {
	m := method.NewGetForumTopicIconStickers()
	var result []types.Sticker
	err := t.process(m, &result)
	return result, err
}

// UnpinAllForumTopicMessages performs unpinAllForumTopicMessages telegram api method.
//
// Reference: https://core.telegram.org/bots/api#unpinallforumtopicmessages
func (t *Telegram) UnpinAllForumTopicMessages(chatID int64, messageID int) (bool, error) {
	m := method.NewUnpinAllForumTopicMessages(chatID, messageID)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// EditGeneralForumTopic performs editGeneralForumTopic telegram api method.
//
// Reference: https://core.telegram.org/bots/api#editgeneralforumtopic
func (t *Telegram) EditGeneralTopic(chatID int64, name string) (bool, error) {
	m := method.NewEditGeneralForumTopic(chatID, name)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// CloseGeneralForumTopic performs closeGeneralForumTopic telegram api method.
//
// Reference: https://core.telegram.org/bots/api#closegeneralforumtopic
func (t *Telegram) CloseGeneralForumTopic(chatID int64) (bool, error) {
	m := method.NewCloseGeneralForumTopic(chatID)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// ReopenGeneralForumTopic performs reopenGeneralForumTopic telegram api method.
//
// Reference: https://core.telegram.org/bots/api#reopengeneralforumtopic
func (t *Telegram) ReopenGeneralForumTopic(chatID int64) (bool, error) {
	m := method.NewReopenGeneralForumTopic(chatID)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// HideGeneralForumTopic performs hideGeneralForumTopic telegram api method.
//
// Reference: https://core.telegram.org/bots/api#hidegeneralforumtopic
func (t *Telegram) HideGeneralForumTopic(chatID int64) (bool, error) {
	m := method.NewHideGeneralForumTopic(chatID)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// UnhideGeneralForumTopic performs unhideGeneralForumTopic telegram api method.
//
// Reference: https://core.telegram.org/bots/api#unhidegeneralforumtopic
func (t *Telegram) UnhideGeneralForumTopic(chatID int64) (bool, error) {
	m := method.NewUnhideGeneralForumTopic(chatID)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// UnpinAllGeneralForumTopicMessages performs unpinAllGeneralForumTopicMessages telegram api method.
//
// Reference: https://core.telegram.org/bots/api#unpinallgeneralforumtopicmessages
func (t *Telegram) UnpinAllGeneralForumTopicMessages(chatID int64) (bool, error) {
	m := method.NewUnpinAllGeneralForumTopicMessages(chatID)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// SetMyCommands performs setMyCommands telegram api method.
//
// options:
//
// option.SetMyCommandsScope(), option.SetMyCommandsLanguageCode()
//
// Reference: https://core.telegram.org/bots/api#setmycommands
func (t *Telegram) SetMyCommands(commands []method.BotCommand, opts ...method.Option[*method.SetMyCommands]) (bool, error) {
	m := method.NewSetMyCommands(commands, opts...)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// GetMyCommands performs getMyCommands telegram api method.
//
// options:
//
// option.GetMyCommandsScope(), option.GetMyCommandsLanguageCode()
//
// Reference: https://core.telegram.org/bots/api#getmycommands
func (t *Telegram) GetMyCommands(opts ...method.Option[*method.GetMyCommands]) ([]method.BotCommand, error) {
	m := method.NewGetMyCommands(opts...)
	var result []method.BotCommand
	err := t.process(m, &result)
	return result, err
}

// DeleteMyCommands performs deleteMyCommands telegram api method.
//
// options:
//
// option.DeleteMyCommandsScope(), option.DeleteMyCommandsLanguageCode()
//
// Reference: https://core.telegram.org/bots/api#deletemycommands
func (t *Telegram) DeleteMyCommands(opts ...method.Option[*method.DeleteMyCommands]) (bool, error) {
	m := method.NewDeleteMyCommands(opts...)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// SetMyName performs setMyName telegram api method.
//
// options:
//
// option.SetMyNameName(), option.SetMyNameLanguageCode()
//
// Reference: https://core.telegram.org/bots/api#setmyname
func (t *Telegram) SetMyName(opts ...method.Option[*method.SetMyName]) (bool, error) {
	m := method.NewSetMyName(opts...)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// GetMyName performs getMyName telegram api method.
//
// options:
//
// option.GetMyNameLanguageCode()
//
// Reference: https://core.telegram.org/bots/api#getmyname
func (t *Telegram) GetMyName(opts ...method.Option[*method.GetMyName]) (types.BotName, error) {
	m := method.NewGetMyName(opts...)
	var result types.BotName
	err := t.process(m, &result)
	return result, err
}

// SetMyDescription performs setMyDescription telegram api method.
//
// options:
//
// option.SetMyDescriptionDescription(), option.SetMyDescriptionLanguageCode()
//
// Reference: https://core.telegram.org/bots/api#setmydescription
func (t *Telegram) SetMyDescription(opts ...method.Option[*method.SetMyDescription]) (bool, error) {
	m := method.NewSetMyDescription(opts...)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// GetMyDescription performs getMyDescription telegram api method.
//
// options:
//
// option.GetMyDescriptionLanguageCode()
//
// Reference: https://core.telegram.org/bots/api#getmydescription
func (t *Telegram) GetMyDescription(opts ...method.Option[*method.GetMyDescription]) (types.BotDescription, error) {
	m := method.NewGetMyDescription(opts...)
	var result types.BotDescription
	err := t.process(m, &result)
	return result, err
}

// SetMyShortDescription performs setMyShortDescription telegram api method.
//
// options:
//
// option.SetMyShortDescriptionShortDescription(), option.SetMyShortDescriptionLanguageCode()
//
// Reference: https://core.telegram.org/bots/api#setmyshortdescription
func (t *Telegram) SetMyShortDescription(opts ...method.Option[*method.SetMyShortDescription]) (bool, error) {
	m := method.NewSetMyShortDescription(opts...)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// GetMyShortDescription performs getMyShortDescription telegram api method.
//
// options:
//
// option.GetMyShortDescriptionLanguageCode()
//
// Reference: https://core.telegram.org/bots/api#getmyshortdescription
func (t *Telegram) GetMyShortDescription(opts ...method.Option[*method.GetMyShortDescription]) (types.BotShortDescription, error) {
	m := method.NewGetMyShortDescription(opts...)
	var result types.BotShortDescription
	err := t.process(m, &result)
	return result, err
}

// SetMyDefaultAdministratorRights performs setMyDefaultAdministratorRights telegram api method.
//
// options:
//
// option.SetMyDefaultAdministratorRightsForSupergroups(), option.SetMyDefaultAdministratorRightsForChannels()
//
// Reference: https://core.telegram.org/bots/api#setmydefaultadministratorrights
func (t *Telegram) SetMyDefaultAdministratorRights(rights method.BotAdministratorRights, opts ...method.Option[*method.SetMyDefaultAdministratorRights]) (bool, error) {
	m := method.NewSetMyDefaultAdministratorRights(&rights, opts...)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// GetMyDefaultAdministratorRights performs getMyDefaultAdministratorRights telegram api method.
//
// options:
//
// option.GetMyDefaultAdministratorRightsForSupergroups(), option.GetMyDefaultAdministratorRightsForChannels()
//
// Reference: https://core.telegram.org/bots/api#getmydefaultadministratorrights
func (t *Telegram) GetMyDefaultAdministratorRights(opts ...method.Option[*method.GetMyDefaultAdministratorRights]) (method.BotAdministratorRights, error) {
	m := method.NewGetMyDefaultAdministratorRights(opts...)
	var result method.BotAdministratorRights
	err := t.process(m, &result)
	return result, err
}

// SetChatMenuButton performs setChatMenuButton telegram api method.
//
// options:
//
// option.SetChatMenuButtonChatID(), option.SetChatMenuButtonMenuButton()
//
// Reference: https://core.telegram.org/bots/api#setchatmenubutton
func (t *Telegram) SetChatMenuButton(opts ...method.Option[*method.SetChatMenuButton]) (bool, error) {
	m := method.NewSetChatMenuButton(opts...)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// GetChatMenuButton performs getChatMenuButton telegram api method.
//
// options:
//
// option.GetChatMenuButtonChatID()
//
// Reference: https://core.telegram.org/bots/api#getchatmenubutton
func (t *Telegram) GetChatMenuButton(opts ...method.Option[*method.GetChatMenuButton]) (types.MenuButton, error) {
	m := method.NewGetChatMenuButton(opts...)
	var result types.MenuButton
	err := t.process(m, &result)
	return result, err
}

// AnswerInlineQuery performs answerInlineQuery telegram api method.
//
// options:
//
// option.AnswerInlineQueryCacheTime(), option.AnswerInlineQueryIsPersonal(),
// option.AnswerInlineQueryNextOffset(), option.AnswerInlineQueryButton()
//
// Reference: https://core.telegram.org/bots/api#answerinlinequery
func (t *Telegram) AnswerInlineQuery(inlineQueryID string, results []method.InlineQueryResult, opts ...method.Option[*method.AnswerInlineQuery]) (bool, error) {
	m := method.NewAnswerInlineQuery(inlineQueryID, results, opts...)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// AnswerWebAppQuery performs answerWebAppQuery telegram api method.
//
// Reference: https://core.telegram.org/bots/api#answerwebappquery
func (t *Telegram) AnswerWebAppQuery(webAppQueryID, result string) (method.SentWebAppMessage, error) {
	m := method.NewAnswerWebAppQuery(webAppQueryID, result)
	var resultSent method.SentWebAppMessage
	err := t.process(m, &resultSent)
	return resultSent, err
}

// SendPaidMedia performs sendPaidMedia telegram api method.
//
// options:
//
// option.SendPaidMediaCaption(), option.SendPaidMediaParseMode(),
// option.SendPaidMediaShowCaptionAboveMedia(), option.SendPaidMediaReplyMarkup()
//
// Reference: https://core.telegram.org/bots/api#sendpaidmedia
func (t *Telegram) SendPaidMedia(chatID int64, starCount int, media []method.InputPaidMedia, opts ...method.Option[*method.SendPaidMedia]) (types.Message, error) {
	m := method.NewSendPaidMedia(chatID, starCount, media, opts...)
	message := types.Message{}
	err := t.process(m, &message)
	return message, err
}

// AnswerShippingQuery performs answerShippingQuery telegram api method.
//
// options:
//
// option.AnswerShippingQueryShippingOptions(), option.AnswerShippingQueryErrorMessage()
//
// Reference: https://core.telegram.org/bots/api#answershippingquery
func (t *Telegram) AnswerShippingQuery(shippingQueryID string, ok bool, opts ...method.Option[*method.AnswerShippingQuery]) (bool, error) {
	m := method.NewAnswerShippingQuery(shippingQueryID, ok, opts...)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// AnswerPreCheckoutQuery performs answerPreCheckoutQuery telegram api method.
//
// options:
//
// option.AnswerPreCheckoutQueryErrorMessage()
//
// Reference: https://core.telegram.org/bots/api#answerprecheckoutquery
func (t *Telegram) AnswerPreCheckoutQuery(preCheckoutQueryID string, ok bool, opts ...method.Option[*method.AnswerPreCheckoutQuery]) (bool, error) {
	m := method.NewAnswerPreCheckoutQuery(preCheckoutQueryID, ok, opts...)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// GetStarTransactions performs getStarTransactions telegram api method.
//
// options:
//
// option.GetStarTransactionsOffset(), option.GetStarTransactionsLimit()
//
// Reference: https://core.telegram.org/bots/api#getstartransactions
func (t *Telegram) GetStarTransactions(opts ...method.Option[*method.GetStarTransactions]) (types.StarTransactions, error) {
	m := method.NewGetStarTransactions(opts...)
	var result types.StarTransactions
	err := t.process(m, &result)
	return result, err
}

// RefundStarPayment performs refundStarPayment telegram api method.
//
// Reference: https://core.telegram.org/bots/api#refundstarpayment
func (t *Telegram) RefundStarPayment(userID int64, telegramPaymentCharge string) (bool, error) {
	m := method.NewRefundStarPayment(userID, telegramPaymentCharge)
	var result bool
	err := t.process(m, &result)
	return result, err
}

// TODO: telegram passport methods.
// TODO: games methods.
