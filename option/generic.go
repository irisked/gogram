// Package option provides functional options for configuring telegram
// API method calls.
package option

import (
	"github.com/irisked/gogram/telegram/keyboard"
	"github.com/irisked/gogram/telegram/method"
	"github.com/irisked/gogram/types"
)

// Option is an alias for method.Option[T], re-exported so callers of
// this package don't need to import telegram/method directly.
type Option[T any] = method.Option[T]

// ReplyToMessageID sets reply_to_message_id
func ReplyToMessageID[T method.ReplyToMessageIDSetter](id int) Option[T] {
	return func(t T) { t.SetReplyToMessageID(id) }
}

// DisableNotification sets disable_notification to true
func DisableNotification[T method.DisableNotificationSetter]() Option[T] {
	return func(t T) { t.SetDisableNotification(true) }
}

// Keyboard sets reply_markup
func Keyboard[T method.ReplyMarkupSetter](kb keyboard.Markup) Option[T] {
	return func(t T) { t.SetReplyMarkup(kb) }
}

// Caption sets caption
func Caption[T method.CaptionSetter](text string) Option[T] {
	return func(t T) { t.SetCaption(text) }
}

// CaptionWithParseMode sets caption and parse_mode
func CaptionWithParseMode[T method.CaptionSetter](text, parseMode string) Option[T] {
	return func(t T) {
		t.SetCaption(text)
		t.SetParseMode(parseMode)
	}
}

// ParseMode sets parse_mode
func ParseMode[T method.ParseModeSetter](mode string) Option[T] {
	return func(t T) { t.SetParseMode(mode) }
}

// Duration sets duration
func Duration[T method.DurationSetter](d int) Option[T] {
	return func(t T) { t.SetDuration(d) }
}

// Width sets width
func Width[T method.WidthSetter](w int) Option[T] {
	return func(t T) { t.SetWidth(w) }
}

// Height sets height
func Height[T method.HeightSetter](h int) Option[T] {
	return func(t T) { t.SetHeight(h) }
}

// Length sets length
func Length[T method.LengthSetter](l int) Option[T] {
	return func(t T) { t.SetLength(l) }
}

// Thumb sets thumb
func Thumb[T method.ThumbSetter](thumb string) Option[T] {
	return func(t T) { t.SetThumb(thumb) }
}

// SupportsStreaming sets supports_streaming
func SupportsStreaming[T method.SupportsStreamingSetter](s bool) Option[T] {
	return func(t T) { t.SetSupportsStreaming(s) }
}

// SimpleCaption is shorthand for Caption
func SimpleCaption[T method.CaptionSetter](text string) Option[T] {
	return Caption[T](text)
}

// MarkdownCaption sets caption with markdown parse mode
func MarkdownCaption[T method.CaptionSetter](text string) Option[T] {
	return CaptionWithParseMode[T](text, "markdown")
}

// HTMLCaption sets caption with HTML parse mode
func HTMLCaption[T method.CaptionSetter](text string) Option[T] {
	return CaptionWithParseMode[T](text, "HTML")
}

// Offset sets offset
func Offset[T method.OffsetSetter](o int) Option[T] {
	return func(t T) { t.SetOffset(o) }
}

// Limit sets limit
func Limit[T method.LimitSetter](l int) Option[T] {
	return func(t T) { t.SetLimit(l) }
}

// Timeout sets timeout
func Timeout[T method.TimeoutSetter](v int) Option[T] {
	return func(t T) { t.SetTimeout(v) }
}

// UserID sets user_id
func UserID[T method.UserIDSetter](id int) Option[T] {
	return func(t T) { t.SetUserID(id) }
}

// ChatID sets chat_id
func ChatID[T method.ChatIDSetter](id int64) Option[T] {
	return func(t T) { t.SetChatID(id) }
}

// UntilDate sets until_date
func UntilDate[T method.UntilDateSetter](d int) Option[T] {
	return func(t T) { t.SetUntilDate(d) }
}

// CanSendMessages sets can_send_messages
func CanSendMessages[T method.CanSendMessagesSetter](v bool) Option[T] {
	return func(t T) { t.SetCanSendMessages(v) }
}

// CanSendMediaMessages sets can_send_media_messages
func CanSendMediaMessages[T method.CanSendMediaMessagesSetter](v bool) Option[T] {
	return func(t T) { t.SetCanSendMediaMessages(v) }
}

// CanSendOtherMessages sets can_send_other_messages
func CanSendOtherMessages[T method.CanSendOtherMessagesSetter](v bool) Option[T] {
	return func(t T) { t.SetCanSendOtherMessages(v) }
}

// CanAddWebPagePreviews sets can_add_web_page_previews
func CanAddWebPagePreviews[T method.CanAddWebPagePreviewsSetter](v bool) Option[T] {
	return func(t T) { t.SetCanAddWebPagePreviews(v) }
}

// CanChangeInfo sets can_change_info
func CanChangeInfo[T method.CanChangeInfoSetter](v bool) Option[T] {
	return func(t T) { t.SetCanChangeInfo(v) }
}

// CanPostMessages sets can_post_messages
func CanPostMessages[T method.CanPostMessagesSetter](v bool) Option[T] {
	return func(t T) { t.SetCanPostMessages(v) }
}

// CanEditMessages sets can_edit_messages
func CanEditMessages[T method.CanEditMessagesSetter](v bool) Option[T] {
	return func(t T) { t.SetCanEditMessages(v) }
}

// CanDeleteMessages sets can_delete_messages
func CanDeleteMessages[T method.CanDeleteMessagesSetter](v bool) Option[T] {
	return func(t T) { t.SetCanDeleteMessages(v) }
}

// CanInviteUsers sets can_invite_users
func CanInviteUsers[T method.CanInviteUsersSetter](v bool) Option[T] {
	return func(t T) { t.SetCanInviteUsers(v) }
}

// CanRestrictMembers sets can_restrict_members
func CanRestrictMembers[T method.CanRestrictMembersSetter](v bool) Option[T] {
	return func(t T) { t.SetCanRestrictMembers(v) }
}

// CanPinMessages sets can_pin_messages
func CanPinMessages[T method.CanPinMessagesSetter](v bool) Option[T] {
	return func(t T) { t.SetCanPinMessages(v) }
}

// CanPromoteMembers sets can_promote_members
func CanPromoteMembers[T method.CanPromoteMembersSetter](v bool) Option[T] {
	return func(t T) { t.SetCanPromoteMembers(v) }
}

// URL sets url
func URL[T method.URLSetter](url string) Option[T] {
	return func(t T) { t.SetURL(url) }
}

// Certificate sets certificate
func Certificate[T method.CertificateSetter](cert string) Option[T] {
	return func(t T) { t.SetCertificate(cert) }
}

// MaxConnections sets max_connections
func MaxConnections[T method.MaxConnectionsSetter](c int) Option[T] {
	return func(t T) { t.SetMaxConnections(c) }
}

// AllowedUpdates sets allowed_updates
func AllowedUpdates[T method.AllowedUpdatesSetter](updates []types.UpdateType) Option[T] {
	return func(t T) { t.SetAllowedUpdates(updates) }
}

// ShowAlert sets show_alert
func ShowAlert[T method.ShowAlertSetter]() Option[T] {
	return func(t T) { t.SetShowAlert(true) }
}

// CacheTime sets cache_time
func CacheTime[T method.CacheTimeSetter](v int) Option[T] {
	return func(t T) { t.SetCacheTime(v) }
}

// DisableEditMessage sets disable_edit_message
func DisableEditMessage[T method.DisableEditMessageSetter]() Option[T] {
	return func(t T) { t.SetDisableEditMessage(true) }
}

// Force sets force
func Force[T method.ForceSetter]() Option[T] {
	return func(t T) { t.SetForce(true) }
}

// CopyMessageCaption sets caption.
func CopyMessageCaption[T method.CopyMessageSetter](text string) Option[T] {
	return func(t T) { t.SetCaption(text) }
}

// CopyMessageParseMode sets parse mode.
func CopyMessageParseMode[T method.CopyMessageSetter](mode string) Option[T] {
	return func(t T) { t.SetParseMode(mode) }
}

// CopyMessageDisableNotification sets disable notification to true.
func CopyMessageDisableNotification[T method.CopyMessageSetter]() Option[T] {
	return func(t T) { t.SetDisableNotification(true) }
}

// CopyMessageReplyToMessageID sets reply_to_message_id.
func CopyMessageReplyToMessageID[T method.CopyMessageSetter](id int) Option[T] {
	return func(t T) { t.SetReplyToMessageID(id) }
}

// CopyMessageAllowCopy sets allow copy.
func CopyMessageAllowCopy[T method.CopyMessageSetter](v bool) Option[T] {
	return func(t T) { t.SetAllowCopy(v) }
}

// CopyMessageReplyMarkup sets reply markup.
func CopyMessageReplyMarkup[T method.CopyMessageSetter](kb keyboard.Markup) Option[T] {
	return func(t T) { t.SetReplyMarkup(kb) }
}

// CopyMessagesDisableNotification sets disable notification to true.
func CopyMessagesDisableNotification[T method.CopyMessagesSetter]() Option[T] {
	return func(t T) { t.SetDisableNotification(true) }
}

// CopyMessagesRemoveSender sets remove sender.
func CopyMessagesRemoveSender[T method.CopyMessagesSetter](v bool) Option[T] {
	return func(t T) { t.SetRemoveSender(v) }
}

// DeleteMessages setters
// (no setters needed)

// StopPollReplyMarkup sets reply markup.
func StopPollReplyMarkup[T method.StopPollSetter](kb keyboard.Markup) Option[T] {
	return func(t T) { t.SetReplyMarkup(kb) }
}

// EditMessageLiveLocationReplyMarkup sets reply markup.
func EditMessageLiveLocationReplyMarkup[T method.EditMessageLiveLocationSetter](kb keyboard.Markup) Option[T] {
	return func(t T) { t.SetReplyMarkup(kb) }
}

// StopMessageLiveLocationReplyMarkup sets reply markup.
func StopMessageLiveLocationReplyMarkup[T method.StopMessageLiveLocationSetter](kb keyboard.Markup) Option[T] {
	return func(t T) { t.SetReplyMarkup(kb) }
}

// DeleteWebhookDropPendingUpdates sets drop pending updates.
func DeleteWebhookDropPendingUpdates[T method.DeleteWebhookSetter](v bool) Option[T] {
	return func(t T) { t.SetDropPendingUpdates(v) }
}

// GetStickerSet setters
// (no setters needed)

// UploadStickerFileFormat sets sticker format.
func UploadStickerFileFormat[T method.UploadStickerFileSetter](s string) Option[T] {
	return func(t T) { t.SetStickerFormat(s) }
}

// UploadStickerFileType sets sticker type.
func UploadStickerFileType[T method.UploadStickerFileSetter](s string) Option[T] {
	return func(t T) { t.SetStickerType(s) }
}

// UploadStickerFilePNGSticker sets the PNG sticker file.
func UploadStickerFilePNGSticker[T method.UploadStickerFileSetter](p string) Option[T] {
	return func(t T) { t.SetPNGSticker(p) }
}

// UploadStickerFileTSVGSticker sets the TGS (animated) sticker file.
func UploadStickerFileTSVGSticker[T method.UploadStickerFileSetter](t string) Option[T] {
	return func(s T) { s.SetTSVGSticker(t) }
}

// UploadStickerFileWebmSticker sets webm sticker.
func UploadStickerFileWebmSticker[T method.UploadStickerFileSetter](w string) Option[T] {
	return func(s T) { s.SetWebmSticker(w) }
}

// UploadStickerFileThumbnail sets thumbnail.
func UploadStickerFileThumbnail[T method.UploadStickerFileSetter](t string) Option[T] {
	return func(s T) { s.SetThumbnail(t) }
}

// UploadStickerFileMaskPosition sets mask position.
func UploadStickerFileMaskPosition[T method.UploadStickerFileSetter](m string) Option[T] {
	return func(t T) { t.SetMaskPosition(m) }
}

// UploadStickerFileEmojiList sets emoji list.
func UploadStickerFileEmojiList[T method.UploadStickerFileSetter](e []string) Option[T] {
	return func(t T) { t.SetEmojiList(e) }
}

// CreateNewStickerSetStickerType sets sticker type.
func CreateNewStickerSetStickerType[T method.CreateNewStickerSetSetter](s string) Option[T] {
	return func(t T) { t.SetStickerType(s) }
}

// CreateNewStickerSetMaskPosition sets mask position.
func CreateNewStickerSetMaskPosition[T method.CreateNewStickerSetSetter](m string) Option[T] {
	return func(t T) { t.SetMaskPosition(m) }
}

// CreateNewStickerSetNeedsRepaired sets needs repaired.
func CreateNewStickerSetNeedsRepaired[T method.CreateNewStickerSetSetter](n *bool) Option[T] {
	return func(t T) { t.SetNeedsRepaired(n) }
}

// DeleteStickerSet setters
// (no setters needed)

// AddStickerToSetMaskPosition sets mask position.
func AddStickerToSetMaskPosition[T method.AddStickerToSetSetter](m string) Option[T] {
	return func(t T) { t.SetMaskPosition(m) }
}

// SetStickerPositionInSet setters
// (no setters needed)

// DeleteStickerFromSet setters
// (no setters needed)

// ReplaceStickerInSet setters
// (no setters needed)

// SetStickerEmojiList setters
// (no setters needed)

// SetStickerKeywords setters
// (no setters needed)

// SetStickerMaskPosition setters
// (no setters needed)

// SetStickerSetTitle setters
// (no setters needed)

// SetStickerSetDescription setters
// (no setters needed)

// SetStickerSetThumbnail setters
// (no setters needed)

// SetCustomEmojiStickerSetThumbnail setters
// (no setters needed)

// GetCustomEmojiStickers setters
// (no setters needed)

// CreateForumTopicIconColor sets icon color.
func CreateForumTopicIconColor[T method.ForumTopicSetter](c int) Option[T] {
	return func(t T) { t.SetIconColor(c) }
}

// CreateForumTopicIconEmojiID sets icon emoji ID.
func CreateForumTopicIconEmojiID[T method.ForumTopicSetter](e string) Option[T] {
	return func(t T) { t.SetIconEmojiID(e) }
}

// EditForumTopicName sets name.
func EditForumTopicName[T method.EditForumTopicSetter](n string) Option[T] {
	return func(t T) { t.SetName(n) }
}

// EditForumTopicIconEmojiID sets icon emoji ID.
func EditForumTopicIconEmojiID[T method.EditForumTopicSetter](e string) Option[T] {
	return func(t T) { t.SetIconEmojiID(e) }
}

// DeleteForumTopic setters
// (no setters needed)

// GetForumTopic setters
// (no setters needed)

// CloseForumTopic setters
// (no setters needed)

// ReopenForumTopic setters
// (no setters needed)

// GetForumTopicIconStickers setters
// (no setters needed)

// UnpinAllForumTopicMessages setters
// (no setters needed)

// EditGeneralForumTopic setters
// (no setters needed)

// CloseGeneralForumTopic setters
// (no setters needed)

// ReopenGeneralForumTopic setters
// (no setters needed)

// HideGeneralForumTopic setters
// (no setters needed)

// UnhideGeneralForumTopic setters
// (no setters needed)

// UnpinAllGeneralForumTopicMessages setters
// (no setters needed)

// SetMyCommandsScope sets scope.
func SetMyCommandsScope[T method.SetMyCommandsSetter](s string) Option[T] {
	return func(t T) { t.SetScope(s) }
}

// SetMyCommandsLanguageCode sets language code.
func SetMyCommandsLanguageCode[T method.SetMyCommandsSetter](l string) Option[T] {
	return func(t T) { t.SetLanguageCode(l) }
}

// GetMyCommandsScope sets scope.
func GetMyCommandsScope[T method.GetMyCommandsSetter](s string) Option[T] {
	return func(t T) { t.SetScope(s) }
}

// GetMyCommandsLanguageCode sets language code.
func GetMyCommandsLanguageCode[T method.GetMyCommandsSetter](l string) Option[T] {
	return func(t T) { t.SetLanguageCode(l) }
}

// DeleteMyCommandsScope sets scope.
func DeleteMyCommandsScope[T method.DeleteMyCommandsSetter](s string) Option[T] {
	return func(t T) { t.SetScope(s) }
}

// DeleteMyCommandsLanguageCode sets language code.
func DeleteMyCommandsLanguageCode[T method.DeleteMyCommandsSetter](l string) Option[T] {
	return func(t T) { t.SetLanguageCode(l) }
}

// SetMyNameName sets name.
func SetMyNameName[T method.SetMyNameSetter](n string) Option[T] {
	return func(t T) { t.SetName(n) }
}

// SetMyNameLanguageCode sets language code.
func SetMyNameLanguageCode[T method.SetMyNameSetter](l string) Option[T] {
	return func(t T) { t.SetLanguageCode(l) }
}

// GetMyNameLanguageCode sets language code.
func GetMyNameLanguageCode[T method.GetMyNameSetter](l string) Option[T] {
	return func(t T) { t.SetLanguageCode(l) }
}

// SetMyDescriptionDescription sets description.
func SetMyDescriptionDescription[T method.SetMyDescriptionSetter](d string) Option[T] {
	return func(t T) { t.SetDescription(d) }
}

// SetMyDescriptionLanguageCode sets language code.
func SetMyDescriptionLanguageCode[T method.SetMyDescriptionSetter](l string) Option[T] {
	return func(t T) { t.SetLanguageCode(l) }
}

// GetMyDescriptionLanguageCode sets language code.
func GetMyDescriptionLanguageCode[T method.GetMyDescriptionSetter](l string) Option[T] {
	return func(t T) { t.SetLanguageCode(l) }
}

// SetMyShortDescriptionShortDescription sets short description.
func SetMyShortDescriptionShortDescription[T method.SetMyShortDescriptionSetter](s string) Option[T] {
	return func(t T) { t.SetShortDescription(s) }
}

// SetMyShortDescriptionLanguageCode sets language code.
func SetMyShortDescriptionLanguageCode[T method.SetMyShortDescriptionSetter](l string) Option[T] {
	return func(t T) { t.SetLanguageCode(l) }
}

// GetMyShortDescriptionLanguageCode sets language code.
func GetMyShortDescriptionLanguageCode[T method.GetMyShortDescriptionSetter](l string) Option[T] {
	return func(t T) { t.SetLanguageCode(l) }
}

// SetMyDefaultAdministratorRightsForSupergroups sets for supergroups.
func SetMyDefaultAdministratorRightsForSupergroups[T method.SetMyDefaultAdministratorRightsSetter](f bool) Option[T] {
	return func(t T) { t.SetForSupergroups(f) }
}

// SetMyDefaultAdministratorRightsForChannels sets for channels.
func SetMyDefaultAdministratorRightsForChannels[T method.SetMyDefaultAdministratorRightsSetter](f bool) Option[T] {
	return func(t T) { t.SetForChannels(f) }
}

// GetMyDefaultAdministratorRightsForSupergroups sets for supergroups.
func GetMyDefaultAdministratorRightsForSupergroups[T method.GetMyDefaultAdministratorRightsSetter](f bool) Option[T] {
	return func(t T) { t.SetForSupergroups(f) }
}

// GetMyDefaultAdministratorRightsForChannels sets for channels.
func GetMyDefaultAdministratorRightsForChannels[T method.GetMyDefaultAdministratorRightsSetter](f bool) Option[T] {
	return func(t T) { t.SetForChannels(f) }
}

// SetChatMenuButtonChatID sets chat_id.
func SetChatMenuButtonChatID[T method.SetChatMenuButtonSetter](c int64) Option[T] {
	return func(t T) { t.SetChatID(c) }
}

// SetChatMenuButtonMenuButton sets menu button.
func SetChatMenuButtonMenuButton[T method.SetChatMenuButtonSetter](m string) Option[T] {
	return func(t T) { t.SetMenuButton(m) }
}

// GetChatMenuButtonChatID sets chat_id.
func GetChatMenuButtonChatID[T method.GetChatMenuButtonSetter](c int64) Option[T] {
	return func(t T) { t.SetChatID(c) }
}

// AnswerInlineQueryCacheTime sets cache time.
func AnswerInlineQueryCacheTime[T method.AnswerInlineQuerySetter](c int) Option[T] {
	return func(t T) { t.SetCacheTime(c) }
}

// AnswerInlineQueryIsPersonal sets is personal.
func AnswerInlineQueryIsPersonal[T method.AnswerInlineQuerySetter](p bool) Option[T] {
	return func(t T) { t.SetIsPersonal(p) }
}

// AnswerInlineQueryNextOffset sets next offset.
func AnswerInlineQueryNextOffset[T method.AnswerInlineQuerySetter](n string) Option[T] {
	return func(t T) { t.SetNextOffset(n) }
}

// AnswerInlineQueryButton sets button.
func AnswerInlineQueryButton[T method.AnswerInlineQuerySetter](b string) Option[T] {
	return func(t T) { t.SetButton(b) }
}

// AnswerWebAppQueryResult sets result.
func AnswerWebAppQueryResult[T method.AnswerWebAppQuerySetter](r string) Option[T] {
	return func(t T) { t.SetResult(r) }
}

// SendPaidMediaCaption sets caption.
func SendPaidMediaCaption[T method.SendPaidMediaSetter](c string) Option[T] {
	return func(t T) { t.SetCaption(c) }
}

// SendPaidMediaParseMode sets parse mode.
func SendPaidMediaParseMode[T method.SendPaidMediaSetter](p string) Option[T] {
	return func(t T) { t.SetParseMode(p) }
}

// SendPaidMediaShowCaptionAboveMedia sets show caption above media.
func SendPaidMediaShowCaptionAboveMedia[T method.SendPaidMediaSetter](s *bool) Option[T] {
	return func(t T) { t.SetShowCaptionAboveMedia(s) }
}

// SendPaidMediaReplyMarkup sets reply markup.
func SendPaidMediaReplyMarkup[T method.SendPaidMediaSetter](r string) Option[T] {
	return func(t T) { t.SetReplyMarkup(r) }
}

// AnswerShippingQueryShippingOptions sets shipping options.
func AnswerShippingQueryShippingOptions[T method.AnswerShippingQuerySetter](s string) Option[T] {
	return func(t T) { t.SetShippingOptions(s) }
}

// AnswerShippingQueryErrorMessage sets error message.
func AnswerShippingQueryErrorMessage[T method.AnswerShippingQuerySetter](e string) Option[T] {
	return func(t T) { t.SetErrorMessage(e) }
}

// AnswerPreCheckoutQueryErrorMessage sets error message.
func AnswerPreCheckoutQueryErrorMessage[T method.AnswerPreCheckoutQuerySetter](e string) Option[T] {
	return func(t T) { t.SetErrorMessage(e) }
}

// GetStarTransactionsOffset sets offset.
func GetStarTransactionsOffset[T method.GetStarTransactionsSetter](o int) Option[T] {
	return func(t T) { t.SetOffset(o) }
}

// GetStarTransactionsLimit sets limit.
func GetStarTransactionsLimit[T method.GetStarTransactionsSetter](l int) Option[T] {
	return func(t T) { t.SetLimit(l) }
}

// RefundStarPayment setters
// (no setters needed)
