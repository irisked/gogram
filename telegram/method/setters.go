package method

import (
	"os"

	"github.com/irisked/gogram/telegram/keyboard"
	"github.com/irisked/gogram/types"
)

// ReplyToMessageIDSetter interface for setting reply_to_message_id
type ReplyToMessageIDSetter interface {
	SetReplyToMessageID(int)
}

// DisableNotificationSetter interface for setting disable_notification
type DisableNotificationSetter interface {
	SetDisableNotification(bool)
}

// ReplyMarkupSetter interface for setting reply_markup
type ReplyMarkupSetter interface {
	SetReplyMarkup(keyboard.Keyboard)
}

// CaptionSetter interface for setting caption and parse_mode
type CaptionSetter interface {
	SetCaption(string)
	SetParseMode(string)
}

// ThumbSetter interface for setting thumb
type ThumbSetter interface {
	SetThumb(string)
}

// DurationSetter interface for setting duration
type DurationSetter interface {
	SetDuration(int)
}

// ParseModeSetter interface for setting parse_mode
type ParseModeSetter interface {
	SetParseMode(string)
}

// SupportsStreamingSetter interface for setting supports_streaming
type SupportsStreamingSetter interface {
	SetSupportsStreaming(bool)
}

// WidthSetter interface for setting width
type WidthSetter interface {
	SetWidth(int)
}

// HeightSetter interface for setting height
type HeightSetter interface {
	SetHeight(int)
}

// LengthSetter interface for setting length
type LengthSetter interface {
	SetLength(int)
}

// FileSetter interface for setting file
type FileSetter interface {
	SetFile(*os.File)
}

// ReferenceSetter interface for setting reference string field
type ReferenceSetter interface {
	SetReference(string)
}

// OffsetSetter interface for setting offset
type OffsetSetter interface {
	SetOffset(int)
}

// LimitSetter interface for setting limit
type LimitSetter interface {
	SetLimit(int)
}

// TimeoutSetter interface for setting timeout
type TimeoutSetter interface {
	SetTimeout(int)
}

// UserIDSetter interface for setting user_id
type UserIDSetter interface {
	SetUserID(int)
}

// ChatIDSetter interface for setting chat_id
type ChatIDSetter interface {
	SetChatID(int64)
}

// UntilDateSetter interface for setting until_date
type UntilDateSetter interface {
	SetUntilDate(int)
}

// CanSendMessagesSetter interface for setting can_send_messages
type CanSendMessagesSetter interface {
	SetCanSendMessages(bool)
}

// CanSendMediaMessagesSetter interface for setting can_send_media_messages
type CanSendMediaMessagesSetter interface {
	SetCanSendMediaMessages(bool)
}

// CanSendOtherMessagesSetter interface for setting can_send_other_messages
type CanSendOtherMessagesSetter interface {
	SetCanSendOtherMessages(bool)
}

// CanAddWebPagePreviewsSetter interface for setting can_add_web_page_previews
type CanAddWebPagePreviewsSetter interface {
	SetCanAddWebPagePreviews(bool)
}

// CanChangeInfoSetter interface for setting can_change_info
type CanChangeInfoSetter interface {
	SetCanChangeInfo(bool)
}

// CanPostMessagesSetter interface for setting can_post_messages
type CanPostMessagesSetter interface {
	SetCanPostMessages(bool)
}

// CanEditMessagesSetter interface for setting can_edit_messages
type CanEditMessagesSetter interface {
	SetCanEditMessages(bool)
}

// CanDeleteMessagesSetter interface for setting can_delete_messages
type CanDeleteMessagesSetter interface {
	SetCanDeleteMessages(bool)
}

// CanInviteUsersSetter interface for setting can_invite_users
type CanInviteUsersSetter interface {
	SetCanInviteUsers(bool)
}

// CanRestrictMembersSetter interface for setting can_restrict_members
type CanRestrictMembersSetter interface {
	SetCanRestrictMembers(bool)
}

// CanPinMessagesSetter interface for setting can_pin_messages
type CanPinMessagesSetter interface {
	SetCanPinMessages(bool)
}

// CanPromoteMembersSetter interface for setting can_promote_members
type CanPromoteMembersSetter interface {
	SetCanPromoteMembers(bool)
}

// URLSetter interface for setting url
type URLSetter interface {
	SetURL(string)
}

// CertificateSetter interface for setting certificate
type CertificateSetter interface {
	SetCertificate(string)
}

// MaxConnectionsSetter interface for setting max_connections
type MaxConnectionsSetter interface {
	SetMaxConnections(int)
}

// AllowedUpdatesSetter interface for setting allowed_updates
type AllowedUpdatesSetter interface {
	SetAllowedUpdates([]types.UpdateType)
}

// ShowAlertSetter interface for setting show_alert
type ShowAlertSetter interface {
	SetShowAlert(bool)
}

// URLSetter interface for setting url in callback queries
type CallbackQueryURLSetter interface {
	SetURL(string)
}

// CacheTimeSetter interface for setting cache_time
type CacheTimeSetter interface {
	SetCacheTime(int)
}

// DisableEditMessageSetter interface for setting disable_edit_message
type DisableEditMessageSetter interface {
	SetDisableEditMessage(bool)
}

// ForceSetter interface for setting force in games
type ForceSetter interface {
	SetForce(bool)
}

// All setter interfaces for new Bot API 9.4 method types
type (
	CopyMessageSetter interface {
		SetCaption(string)
		SetParseMode(string)
		SetDisableNotification(bool)
		SetReplyToMessageID(int)
		SetAllowCopy(bool)
		SetReplyMarkup(keyboard.Keyboard)
	}

	CopyMessagesSetter interface {
		SetDisableNotification(bool)
		SetRemoveSender(bool)
	}

	DeleteMessagesSetter interface {
	}

	StopPollSetter interface {
		SetReplyMarkup(keyboard.Keyboard)
	}

	EditMessageLiveLocationSetter interface {
		SetReplyMarkup(keyboard.Keyboard)
	}

	StopMessageLiveLocationSetter interface {
		SetReplyMarkup(keyboard.Keyboard)
	}

	DeleteWebhookSetter interface {
		SetDropPendingUpdates(bool)
	}

	GetStickerSetSetter interface {
	}

	UploadStickerFileSetter interface {
		SetStickerFormat(string)
		SetStickerType(string)
		SetPNGSticker(string)
		SetTSVGSticker(string)
		SetWebmSticker(string)
		SetThumbnail(string)
		SetMaskPosition(string)
		SetEmojiList([]string)
	}

	CreateNewStickerSetSetter interface {
		SetStickerType(string)
		SetMaskPosition(string)
		SetNeedsRepaired(*bool)
	}

	DeleteStickerSetSetter interface {
	}

	AddStickerToSetSetter interface {
		SetMaskPosition(string)
	}

	SetStickerPositionInSetSetter interface {
	}

	DeleteStickerFromSetSetter interface {
	}

	ReplaceStickerInSetSetter interface {
	}

	SetStickerEmojiListSetter interface {
	}

	SetStickerKeywordsSetter interface {
	}

	SetStickerMaskPositionSetter interface {
	}

	SetStickerSetTitleSetter interface {
	}

	SetStickerSetDescriptionSetter interface {
	}

	SetStickerSetThumbnailSetter interface {
	}

	SetCustomEmojiStickerSetThumbnailSetter interface {
	}

	GetCustomEmojiStickersSetter interface {
	}

	ForumTopicSetter interface {
		SetIconColor(int)
		SetIconEmojiID(string)
	}

	EditForumTopicSetter interface {
		SetName(string)
		SetIconEmojiID(string)
	}

	DeleteForumTopicSetter interface {
	}

	GetForumTopicSetter interface {
	}

	CloseForumTopicSetter interface {
	}

	ReopenForumTopicSetter interface {
	}

	GetForumTopicIconStickersSetter interface {
	}

	UnpinAllForumTopicMessagesSetter interface {
	}

	EditGeneralForumTopicSetter interface {
	}

	CloseGeneralForumTopicSetter interface {
	}

	ReopenGeneralForumTopicSetter interface {
	}

	HideGeneralForumTopicSetter interface {
	}

	UnhideGeneralForumTopicSetter interface {
	}

	UnpinAllGeneralForumTopicMessagesSetter interface {
	}

	SetMyCommandsSetter interface {
		SetScope(string)
		SetLanguageCode(string)
	}

	GetMyCommandsSetter interface {
		SetScope(string)
		SetLanguageCode(string)
	}

	DeleteMyCommandsSetter interface {
		SetScope(string)
		SetLanguageCode(string)
	}

	SetMyNameSetter interface {
		SetName(string)
		SetLanguageCode(string)
	}

	GetMyNameSetter interface {
		SetLanguageCode(string)
	}

	SetMyDescriptionSetter interface {
		SetDescription(string)
		SetLanguageCode(string)
	}

	GetMyDescriptionSetter interface {
		SetLanguageCode(string)
	}

	SetMyShortDescriptionSetter interface {
		SetShortDescription(string)
		SetLanguageCode(string)
	}

	GetMyShortDescriptionSetter interface {
		SetLanguageCode(string)
	}

	SetMyDefaultAdministratorRightsSetter interface {
		SetForSupergroups(bool)
		SetForChannels(bool)
	}

	GetMyDefaultAdministratorRightsSetter interface {
		SetForSupergroups(bool)
		SetForChannels(bool)
	}

	SetChatMenuButtonSetter interface {
		SetChatID(int64)
		SetMenuButton(string)
	}

	GetChatMenuButtonSetter interface {
		SetChatID(int64)
	}

	AnswerInlineQuerySetter interface {
		SetCacheTime(int)
		SetIsPersonal(bool)
		SetNextOffset(string)
		SetButton(string)
	}

	AnswerWebAppQuerySetter interface {
		SetResult(string)
	}

	SendPaidMediaSetter interface {
		SetCaption(string)
		SetParseMode(string)
		SetShowCaptionAboveMedia(*bool)
		SetReplyMarkup(string)
	}

	AnswerShippingQuerySetter interface {
		SetShippingOptions(string)
		SetErrorMessage(string)
	}

	AnswerPreCheckoutQuerySetter interface {
		SetErrorMessage(string)
	}

	GetStarTransactionsSetter interface {
		SetOffset(int)
		SetLimit(int)
	}

	RefundStarPaymentSetter interface {
	}
)
