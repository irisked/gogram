package method

import (
	"os"

	"github.com/irisked/gogram/telegram/keyboard"
)

// Type aliases for backward compatibility - now using generic Option[T].

// SendMessageOption is an alias for Option[*SendMessage].
type SendMessageOption = Option[*SendMessage]

// ForwardMessageOption is an alias for Option[*ForwardMessage].
type ForwardMessageOption = Option[*ForwardMessage]

// SendPhotoOption is an alias for Option[*SendPhoto].
type SendPhotoOption = Option[*SendPhoto]

// SendAudioOption is an alias for Option[*SendAudio].
type SendAudioOption = Option[*SendAudio]

// SendDocumentOption is an alias for Option[*SendDocument].
type SendDocumentOption = Option[*SendDocument]

// SendVideoOption is an alias for Option[*SendVideo].
type SendVideoOption = Option[*SendVideo]

// SendAnimationOption is an alias for Option[*SendAnimation].
type SendAnimationOption = Option[*SendAnimation]

// SendVoiceOption is an alias for Option[*SendVoice].
type SendVoiceOption = Option[*SendVoice]

// SendVideoNoteOption is an alias for Option[*SendVideoNote].
type SendVideoNoteOption = Option[*SendVideoNote]

// SendMediaGroupOption is an alias for Option[*SendMediaGroup].
type SendMediaGroupOption = Option[*SendMediaGroup]

// SendLocationOption is an alias for Option[*SendLocation].
type SendLocationOption = Option[*SendLocation]

// SendVenueOption is an alias for Option[*SendVenue].
type SendVenueOption = Option[*SendVenue]

// SendContactOption is an alias for Option[*SendContact].
type SendContactOption = Option[*SendContact]

// GetUserProfilePhotosOption is an alias for Option[*GetUserProfilePhotos].
type GetUserProfilePhotosOption = Option[*GetUserProfilePhotos]

// KickChatMemberOption is an alias for Option[*KickChatMember].
type KickChatMemberOption = Option[*KickChatMember]

// UnbanChatMemberOption is an alias for Option[*UnbanChatMember].
type UnbanChatMemberOption = Option[*UnbanChatMember]

// RestrictChatMemberOption is an alias for Option[*RestrictChatMember].
type RestrictChatMemberOption = Option[*RestrictChatMember]

// PromoteChatMemberOption is an alias for Option[*PromoteChatMember].
type PromoteChatMemberOption = Option[*PromoteChatMember]

// PinChatMessageOption is an alias for Option[*PinChatMessage].
type PinChatMessageOption = Option[*PinChatMessage]

// AnswerCallbackQueryOption is an alias for Option[*AnswerCallbackQuery].
type AnswerCallbackQueryOption = Option[*AnswerCallbackQuery]

// EditMessageTextOption is an alias for Option[*EditMessageText].
type EditMessageTextOption = Option[*EditMessageText]

// EditMessageCaptionOption is an alias for Option[*EditMessageCaption].
type EditMessageCaptionOption = Option[*EditMessageCaption]

// EditMessageMediaOption is an alias for Option[*EditMessageMedia].
type EditMessageMediaOption = Option[*EditMessageMedia]

// EditMessageReplyMarkupOption is an alias for Option[*EditMessageReplyMarkup].
type EditMessageReplyMarkupOption = Option[*EditMessageReplyMarkup]

// New type aliases for Bot API 9.4 methods.

// CopyMessageOption is an alias for Option[*CopyMessage].
type CopyMessageOption = Option[*CopyMessage]

// CopyMessagesOption is an alias for Option[*CopyMessages].
type CopyMessagesOption = Option[*CopyMessages]

// DeleteMessagesOption is an alias for Option[*DeleteMessages].
type DeleteMessagesOption = Option[*DeleteMessages]

// StopPollOption is an alias for Option[*StopPoll].
type StopPollOption = Option[*StopPoll]

// EditMessageLiveLocationOption is an alias for Option[*EditMessageLiveLocation].
type EditMessageLiveLocationOption = Option[*EditMessageLiveLocation]

// StopMessageLiveLocationOption is an alias for Option[*StopMessageLiveLocation].
type StopMessageLiveLocationOption = Option[*StopMessageLiveLocation]

// DeleteWebhookOption is an alias for Option[*DeleteWebhook].
type DeleteWebhookOption = Option[*DeleteWebhook]

// GetStickerSetOption is an alias for Option[*GetStickerSet].
type GetStickerSetOption = Option[*GetStickerSet]

// UploadStickerFileOption is an alias for Option[*UploadStickerFile].
type UploadStickerFileOption = Option[*UploadStickerFile]

// CreateNewStickerSetOption is an alias for Option[*CreateNewStickerSet].
type CreateNewStickerSetOption = Option[*CreateNewStickerSet]

// DeleteStickerSetOption is an alias for Option[*DeleteStickerSet].
type DeleteStickerSetOption = Option[*DeleteStickerSet]

// AddStickerToSetOption is an alias for Option[*AddStickerToSet].
type AddStickerToSetOption = Option[*AddStickerToSet]

// SetStickerPositionInSetOption is an alias for Option[*SetStickerPositionInSet].
type SetStickerPositionInSetOption = Option[*SetStickerPositionInSet]

// DeleteStickerFromSetOption is an alias for Option[*DeleteStickerFromSet].
type DeleteStickerFromSetOption = Option[*DeleteStickerFromSet]

// ReplaceStickerInSetOption is an alias for Option[*ReplaceStickerInSet].
type ReplaceStickerInSetOption = Option[*ReplaceStickerInSet]

// SetStickerEmojiListOption is an alias for Option[*SetStickerEmojiList].
type SetStickerEmojiListOption = Option[*SetStickerEmojiList]

// SetStickerKeywordsOption is an alias for Option[*SetStickerKeywords].
type SetStickerKeywordsOption = Option[*SetStickerKeywords]

// SetStickerMaskPositionOption is an alias for Option[*SetStickerMaskPosition].
type SetStickerMaskPositionOption = Option[*SetStickerMaskPosition]

// SetStickerSetTitleOption is an alias for Option[*SetStickerSetTitle].
type SetStickerSetTitleOption = Option[*SetStickerSetTitle]

// SetStickerSetDescriptionOption is an alias for Option[*SetStickerSetDescription].
type SetStickerSetDescriptionOption = Option[*SetStickerSetDescription]

// SetStickerSetThumbnailOption is an alias for Option[*SetStickerSetThumbnail].
type SetStickerSetThumbnailOption = Option[*SetStickerSetThumbnail]

// SetCustomEmojiStickerSetThumbnailOption is an alias for Option[*SetCustomEmojiStickerSetThumbnail].
type SetCustomEmojiStickerSetThumbnailOption = Option[*SetCustomEmojiStickerSetThumbnail]

// GetCustomEmojiStickersOption is an alias for Option[*GetCustomEmojiStickers].
type GetCustomEmojiStickersOption = Option[*GetCustomEmojiStickers]

// ForumTopicOption is an alias for Option[*CreateForumTopic].
type ForumTopicOption = Option[*CreateForumTopic]

// EditForumTopicOption is an alias for Option[*EditForumTopic].
type EditForumTopicOption = Option[*EditForumTopic]

// DeleteForumTopicOption is an alias for Option[*DeleteForumTopic].
type DeleteForumTopicOption = Option[*DeleteForumTopic]

// GetForumTopicOption is an alias for Option[*GetForumTopic].
type GetForumTopicOption = Option[*GetForumTopic]

// CloseForumTopicOption is an alias for Option[*CloseForumTopic].
type CloseForumTopicOption = Option[*CloseForumTopic]

// ReopenForumTopicOption is an alias for Option[*ReopenForumTopic].
type ReopenForumTopicOption = Option[*ReopenForumTopic]

// GetForumTopicIconStickersOption is an alias for Option[*GetForumTopicIconStickers].
type GetForumTopicIconStickersOption = Option[*GetForumTopicIconStickers]

// UnpinAllForumTopicMessagesOption is an alias for Option[*UnpinAllForumTopicMessages].
type UnpinAllForumTopicMessagesOption = Option[*UnpinAllForumTopicMessages]

// EditGeneralForumTopicOption is an alias for Option[*EditGeneralForumTopic].
type EditGeneralForumTopicOption = Option[*EditGeneralForumTopic]

// CloseGeneralForumTopicOption is an alias for Option[*CloseGeneralForumTopic].
type CloseGeneralForumTopicOption = Option[*CloseGeneralForumTopic]

// ReopenGeneralForumTopicOption is an alias for Option[*ReopenGeneralForumTopic].
type ReopenGeneralForumTopicOption = Option[*ReopenGeneralForumTopic]

// HideGeneralForumTopicOption is an alias for Option[*HideGeneralForumTopic].
type HideGeneralForumTopicOption = Option[*HideGeneralForumTopic]

// UnhideGeneralForumTopicOption is an alias for Option[*UnhideGeneralForumTopic].
type UnhideGeneralForumTopicOption = Option[*UnhideGeneralForumTopic]

// UnpinAllGeneralForumTopicMessagesOption is an alias for Option[*UnpinAllGeneralForumTopicMessages].
type UnpinAllGeneralForumTopicMessagesOption = Option[*UnpinAllGeneralForumTopicMessages]

// SetMyCommandsOption is an alias for Option[*SetMyCommands].
type SetMyCommandsOption = Option[*SetMyCommands]

// GetMyCommandsOption is an alias for Option[*GetMyCommands].
type GetMyCommandsOption = Option[*GetMyCommands]

// DeleteMyCommandsOption is an alias for Option[*DeleteMyCommands].
type DeleteMyCommandsOption = Option[*DeleteMyCommands]

// SetMyNameOption is an alias for Option[*SetMyName].
type SetMyNameOption = Option[*SetMyName]

// GetMyNameOption is an alias for Option[*GetMyName].
type GetMyNameOption = Option[*GetMyName]

// SetMyDescriptionOption is an alias for Option[*SetMyDescription].
type SetMyDescriptionOption = Option[*SetMyDescription]

// GetMyDescriptionOption is an alias for Option[*GetMyDescription].
type GetMyDescriptionOption = Option[*GetMyDescription]

// SetMyShortDescriptionOption is an alias for Option[*SetMyShortDescription].
type SetMyShortDescriptionOption = Option[*SetMyShortDescription]

// GetMyShortDescriptionOption is an alias for Option[*GetMyShortDescription].
type GetMyShortDescriptionOption = Option[*GetMyShortDescription]

// SetMyDefaultAdministratorRightsOption is an alias for Option[*SetMyDefaultAdministratorRights].
type SetMyDefaultAdministratorRightsOption = Option[*SetMyDefaultAdministratorRights]

// GetMyDefaultAdministratorRightsOption is an alias for Option[*GetMyDefaultAdministratorRights].
type GetMyDefaultAdministratorRightsOption = Option[*GetMyDefaultAdministratorRights]

// SetChatMenuButtonOption is an alias for Option[*SetChatMenuButton].
type SetChatMenuButtonOption = Option[*SetChatMenuButton]

// GetChatMenuButtonOption is an alias for Option[*GetChatMenuButton].
type GetChatMenuButtonOption = Option[*GetChatMenuButton]

// AnswerInlineQueryOption is an alias for Option[*AnswerInlineQuery].
type AnswerInlineQueryOption = Option[*AnswerInlineQuery]

// AnswerWebAppQueryOption is an alias for Option[*AnswerWebAppQuery].
type AnswerWebAppQueryOption = Option[*AnswerWebAppQuery]

// SendPaidMediaOption is an alias for Option[*SendPaidMedia].
type SendPaidMediaOption = Option[*SendPaidMedia]

// AnswerShippingQueryOption is an alias for Option[*AnswerShippingQuery].
type AnswerShippingQueryOption = Option[*AnswerShippingQuery]

// AnswerPreCheckoutQueryOption is an alias for Option[*AnswerPreCheckoutQuery].
type AnswerPreCheckoutQueryOption = Option[*AnswerPreCheckoutQuery]

// GetStarTransactionsOption is an alias for Option[*GetStarTransactions].
type GetStarTransactionsOption = Option[*GetStarTransactions]

// RefundStarPaymentOption is an alias for Option[*RefundStarPayment].
type RefundStarPaymentOption = Option[*RefundStarPayment]

// SendOptions contains options common to send methods
type SendOptions struct {
	DisableWebPagePreview bool            `json:"disable_web_page_preview,omitempty" form:"disable_web_page_preview,omitempty"`
	DisableNotification   bool            `json:"disable_notification,omitempty" form:"disable_notification,omitempty"`
	ReplyToMessageID      int             `json:"reply_to_message_id,omitempty" form:"reply_to_message_id,omitempty"`
	ReplyMarkup           keyboard.Markup `json:"reply_markup,omitempty" form:"reply_markup,omitempty"`
	ParseMode             string          `json:"parse_mode,omitempty" form:"parse_mode,omitempty"`
}

// SetDisableWebPagePreview sets disable_web_page_preview
func (o *SendOptions) SetDisableWebPagePreview(v bool) {
	o.DisableWebPagePreview = v
}

// CaptionOptions contains caption-related fields
type CaptionOptions struct {
	Caption   string `json:"caption,omitempty" form:"caption,omitempty"`
	ParseMode string `json:"parse_mode,omitempty" form:"parse_mode,omitempty"`
}

// ThumbOptions contains thumbnail-related fields
type ThumbOptions struct {
	Thumb     string   `json:"thumb,omitempty" form:"thumb,omitempty"`
	ThumbFile *os.File `json:"-" form:"thumb"`
}

// DurationOptions contains duration-related fields
type DurationOptions struct {
	Duration int `json:"duration,omitempty" form:"duration,omitempty"`
}

// DimensionOptions contains width/height fields
type DimensionOptions struct {
	Width  int `json:"width,omitempty" form:"width,omitempty"`
	Height int `json:"height,omitempty" form:"height,omitempty"`
}

// LengthOptions contains length field
type LengthOptions struct {
	Length int `json:"length,omitempty" form:"length,omitempty"`
}

// SupportsStreamingOptions contains supports_streaming field
type SupportsStreamingOptions struct {
	SupportsStreaming bool `json:"supports_streaming,omitempty" form:"supports_streaming,omitempty"`
}

// SetReplyToMessageID sets reply_to_message_id
func (o *SendOptions) SetReplyToMessageID(id int) {
	o.ReplyToMessageID = id
}

// SetDisableNotification sets disable_notification
func (o *SendOptions) SetDisableNotification(v bool) {
	o.DisableNotification = v
}

// SetReplyMarkup sets reply_markup
func (o *SendOptions) SetReplyMarkup(kb keyboard.Markup) {
	o.ReplyMarkup = kb
}

// SetCaption sets caption
func (o *CaptionOptions) SetCaption(c string) {
	o.Caption = c
}

// SetParseMode sets parse_mode
func (o *CaptionOptions) SetParseMode(m string) {
	o.ParseMode = m
}

// SetThumb sets thumb
func (o *ThumbOptions) SetThumb(t string) {
	o.Thumb = t
}

// SetDuration sets duration
func (o *DurationOptions) SetDuration(d int) {
	o.Duration = d
}

// SetWidth sets width
func (o *DimensionOptions) SetWidth(w int) {
	o.Width = w
}

// SetHeight sets height
func (o *DimensionOptions) SetHeight(h int) {
	o.Height = h
}

// SetLength sets length
func (o *LengthOptions) SetLength(l int) {
	o.Length = l
}

// SetSupportsStreaming sets supports_streaming
func (o *SupportsStreamingOptions) SetSupportsStreaming(s bool) {
	o.SupportsStreaming = s
}
