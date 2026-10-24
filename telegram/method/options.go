package method

import (
	"os"

	"gogram/telegram/keyboard"
)

// Type aliases for backward compatibility - now using generic Option[T]
type SendMessageOption = Option[*SendMessage]
type ForwardMessageOption = Option[*ForwardMessage]
type SendPhotoOption = Option[*SendPhoto]
type SendAudioOption = Option[*SendAudio]
type SendDocumentOption = Option[*SendDocument]
type SendVideoOption = Option[*SendVideo]
type SendAnimationOption = Option[*SendAnimation]
type SendVoiceOption = Option[*SendVoice]
type SendVideoNoteOption = Option[*SendVideoNote]
type SendMediaGroupOption = Option[*SendMediaGroup]
type SendLocationOption = Option[*SendLocation]
type SendVenueOption = Option[*SendVenue]
type SendContactOption = Option[*SendContact]
type GetUserProfilePhotosOption = Option[*GetUserProfilePhotos]
type KickChatMemberOption = Option[*KickChatMember]
type UnbanChatMemberOption = Option[*UnbanChatMember]
type RestrictChatMemberOption = Option[*RestrictChatMember]
type PromoteChatMemberOption = Option[*PromoteChatMember]
type PinChatMessageOption = Option[*PinChatMessage]
type AnswerCallbackQueryOption = Option[*AnswerCallbackQuery]
type EditMessageTextOption = Option[*EditMessageText]
type EditMessageCaptionOption = Option[*EditMessageCaption]
type EditMessageMediaOption = Option[*EditMessageMedia]
type EditMessageReplyMarkupOption = Option[*EditMessageReplyMarkup]

// New type aliases for Bot API 9.4 methods
type CopyMessageOption = Option[*CopyMessage]
type CopyMessagesOption = Option[*CopyMessages]
type DeleteMessagesOption = Option[*DeleteMessages]
type StopPollOption = Option[*StopPoll]
type EditMessageLiveLocationOption = Option[*EditMessageLiveLocation]
type StopMessageLiveLocationOption = Option[*StopMessageLiveLocation]
type DeleteWebhookOption = Option[*DeleteWebhook]
type GetStickerSetOption = Option[*GetStickerSet]
type UploadStickerFileOption = Option[*UploadStickerFile]
type CreateNewStickerSetOption = Option[*CreateNewStickerSet]
type DeleteStickerSetOption = Option[*DeleteStickerSet]
type AddStickerToSetOption = Option[*AddStickerToSet]
type SetStickerPositionInSetOption = Option[*SetStickerPositionInSet]
type DeleteStickerFromSetOption = Option[*DeleteStickerFromSet]
type ReplaceStickerInSetOption = Option[*ReplaceStickerInSet]
type SetStickerEmojiListOption = Option[*SetStickerEmojiList]
type SetStickerKeywordsOption = Option[*SetStickerKeywords]
type SetStickerMaskPositionOption = Option[*SetStickerMaskPosition]
type SetStickerSetTitleOption = Option[*SetStickerSetTitle]
type SetStickerSetDescriptionOption = Option[*SetStickerSetDescription]
type SetStickerSetThumbnailOption = Option[*SetStickerSetThumbnail]
type SetCustomEmojiStickerSetThumbnailOption = Option[*SetCustomEmojiStickerSetThumbnail]
type GetCustomEmojiStickersOption = Option[*GetCustomEmojiStickers]
type ForumTopicOption = Option[*CreateForumTopic]
type EditForumTopicOption = Option[*EditForumTopic]
type DeleteForumTopicOption = Option[*DeleteForumTopic]
type GetForumTopicOption = Option[*GetForumTopic]
type CloseForumTopicOption = Option[*CloseForumTopic]
type ReopenForumTopicOption = Option[*ReopenForumTopic]
type GetForumTopicIconStickersOption = Option[*GetForumTopicIconStickers]
type UnpinAllForumTopicMessagesOption = Option[*UnpinAllForumTopicMessages]
type EditGeneralForumTopicOption = Option[*EditGeneralForumTopic]
type CloseGeneralForumTopicOption = Option[*CloseGeneralForumTopic]
type ReopenGeneralForumTopicOption = Option[*ReopenGeneralForumTopic]
type HideGeneralForumTopicOption = Option[*HideGeneralForumTopic]
type UnhideGeneralForumTopicOption = Option[*UnhideGeneralForumTopic]
type UnpinAllGeneralForumTopicMessagesOption = Option[*UnpinAllGeneralForumTopicMessages]
type SetMyCommandsOption = Option[*SetMyCommands]
type GetMyCommandsOption = Option[*GetMyCommands]
type DeleteMyCommandsOption = Option[*DeleteMyCommands]
type SetMyNameOption = Option[*SetMyName]
type GetMyNameOption = Option[*GetMyName]
type SetMyDescriptionOption = Option[*SetMyDescription]
type GetMyDescriptionOption = Option[*GetMyDescription]
type SetMyShortDescriptionOption = Option[*SetMyShortDescription]
type GetMyShortDescriptionOption = Option[*GetMyShortDescription]
type SetMyDefaultAdministratorRightsOption = Option[*SetMyDefaultAdministratorRights]
type GetMyDefaultAdministratorRightsOption = Option[*GetMyDefaultAdministratorRights]
type SetChatMenuButtonOption = Option[*SetChatMenuButton]
type GetChatMenuButtonOption = Option[*GetChatMenuButton]
type AnswerInlineQueryOption = Option[*AnswerInlineQuery]
type AnswerWebAppQueryOption = Option[*AnswerWebAppQuery]
type SendPaidMediaOption = Option[*SendPaidMedia]
type AnswerShippingQueryOption = Option[*AnswerShippingQuery]
type AnswerPreCheckoutQueryOption = Option[*AnswerPreCheckoutQuery]
type GetStarTransactionsOption = Option[*GetStarTransactions]
type RefundStarPaymentOption = Option[*RefundStarPayment]

// SendOptions contains options common to send methods
type SendOptions struct {
	DisableWebPagePreview bool              `json:"disable_web_page_preview,omitempty" form:"disable_web_page_preview,omitempty"`
	DisableNotification   bool              `json:"disable_notification,omitempty" form:"disable_notification,omitempty"`
	ReplyToMessageID      int               `json:"reply_to_message_id,omitempty" form:"reply_to_message_id,omitempty"`
	ReplyMarkup           keyboard.Keyboard `json:"reply_markup,omitempty" form:"reply_markup,omitempty"`
	ParseMode             string            `json:"parse_mode,omitempty" form:"parse_mode,omitempty"`
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
func (o *SendOptions) SetReplyMarkup(kb keyboard.Keyboard) {
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
