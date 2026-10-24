package types

// BotCommand contains information about a bot command.
type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

// BotName contains information about bot name.
type BotName struct {
	Name string `json:"name"`
}

// BotDescription contains information about bot description.
type BotDescription struct {
	Description string `json:"description"`
}

// BotShortDescription contains information about bot short description.
type BotShortDescription struct {
	ShortDescription string `json:"short_description"`
}

// BotAdministratorRights contains information about bot administrator rights.
type BotAdministratorRights struct {
	IsAnonymous        bool `json:"is_anonymous"`
	CanManageChat      bool `json:"can_manage_chat"`
	CanChangeInfo      bool `json:"can_change_info"`
	CanPostMessages    bool `json:"can_post_messages"`
	CanEditMessages    bool `json:"can_edit_messages"`
	CanDeleteMessages  bool `json:"can_delete_messages"`
	CanInviteUsers     bool `json:"can_invite_users"`
	CanRestrictMembers bool `json:"can_restrict_members"`
	CanPinMessages     bool `json:"can_pin_messages"`
	CanManageTopics    bool `json:"can_manage_topics"`
}

// MenuButton contains information about menu button.
type MenuButton struct {
	Type      string `json:"type"`
	Text      string `json:"text,omitempty"`
	WebURL    string `json:"web_app,omitempty"`
	ShortName string `json:"short_name,omitempty"`
}

// InlineQueryResult contains information about inline query result.
type InlineQueryResult struct {
	Type                string               `json:"type"`
	ID                  string               `json:"id"`
	Title               string               `json:"title,omitempty"`
	Caption             string               `json:"caption,omitempty"`
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
	ReplyMarkup         *ReplyMarkup         `json:"reply_markup,omitempty"`
	URL                 string               `json:"url,omitempty"`
	ThumbURL            string               `json:"thumb_url,omitempty"`
	ThumbWidth          int                  `json:"thumb_width,omitempty"`
	ThumbHeight         int                  `json:"thumb_height,omitempty"`
}

// InputMessageContent contains information about input message content.
type InputMessageContent struct {
	Text               string              `json:"text,omitempty"`
	ParseMode          string              `json:"parse_mode,omitempty"`
	Entities           []MessageEntity     `json:"entities,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	Location           *Location           `json:"location,omitempty"`
	Venue              *Venue              `json:"venue,omitempty"`
	Contact            *Contact            `json:"contact,omitempty"`
}

// LinkPreviewOptions contains information about link preview options.
type LinkPreviewOptions struct {
	IsDisabled       bool   `json:"is_disabled,omitempty"`
	URL              string `json:"url,omitempty"`
	PreferSmallMedia bool   `json:"prefer_small_media,omitempty"`
	PreferLargeMedia bool   `json:"prefer_large_media,omitempty"`
	ShowAboveText    bool   `json:"show_above_text,omitempty"`
}

// ReplyMarkup contains information about reply markup.
type ReplyMarkup struct {
	InlineKeyboard  []InlineKeyboardButton `json:"inline_keyboard,omitempty"`
	Keyboard        []KeyboardButton       `json:"keyboard,omitempty"`
	ResizeKeyboard  bool                   `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool                   `json:"one_time_keyboard,omitempty"`
	Placeholder     string                 `json:"placeholder,omitempty"`
	Selective       bool                   `json:"selective,omitempty"`
}

// InlineKeyboardButton contains information about inline keyboard button.
type InlineKeyboardButton struct {
	Text         string        `json:"text"`
	URL          string        `json:"url,omitempty"`
	CallbackData string        `json:"callback_data,omitempty"`
	WebApp       *WebApp       `json:"web_app,omitempty"`
	LoginURL     *LoginURL     `json:"login_url,omitempty"`
	CallbackGame *CallbackGame `json:"callback_game,omitempty"`
	Pay          bool          `json:"pay,omitempty"`
}

// WebApp contains information about a Web App.
type WebApp struct {
	URL string `json:"url"`
}

// LoginURL contains information about a login URL.
type LoginURL struct {
	URL            string `json:"url"`
	ForwardText    string `json:"forward_text,omitempty"`
	ButtonName     string `json:"button_name,omitempty"`
	AskChannelName bool   `json:"ask_channel_name,omitempty"`
}

// CallbackGame contains information about a callback game.
type CallbackGame struct{}

// KeyboardButton contains information about keyboard button.
type KeyboardButton struct {
	Text            string        `json:"text"`
	RequestUser     *RequestUser  `json:"request_user,omitempty"`
	RequestUsers    *RequestUsers `json:"request_users,omitempty"`
	RequestChat     *RequestChat  `json:"request_chat,omitempty"`
	RequestContacts bool          `json:"request_contacts,omitempty"`
	RequestLocation bool          `json:"request_location,omitempty"`
	RequestPoll     *RequestPoll  `json:"request_poll,omitempty"`
	WebApp          *WebApp       `json:"web_app,omitempty"`
}

// RequestUser contains information about request user.
type RequestUser struct {
	RequestID     int   `json:"request_id"`
	UserIsBot     *bool `json:"user_is_bot,omitempty"`
	UserIsPremium *bool `json:"user_is_premium,omitempty"`
}

// RequestUsers contains information about request users.
type RequestUsers struct {
	RequestID     int   `json:"request_id"`
	TotalCount    int   `json:"total_count"`
	UserIsBot     *bool `json:"user_is_bot,omitempty"`
	UserIsPremium *bool `json:"user_is_premium,omitempty"`
}

// RequestChat contains information about request chat.
type RequestChat struct {
	RequestID               int                     `json:"request_id"`
	ChatIsChannel           bool                    `json:"chat_is_channel,omitempty"`
	ChatIsForum             *bool                   `json:"chat_is_forum,omitempty"`
	ChatHasUsername         *bool                   `json:"chat_has_username,omitempty"`
	ChatIsCreated           *bool                   `json:"chat_is_created,omitempty"`
	UserAdministratorRights *BotAdministratorRights `json:"user_administrator_rights,omitempty"`
	BotAdministratorRights  *BotAdministratorRights `json:"bot_administrator_rights,omitempty"`
	BotIsMember             *bool                   `json:"bot_is_member,omitempty"`
}

// RequestPoll contains information about request poll.
type RequestPoll struct {
	RequestID   int    `json:"request_id"`
	Formatter   string `json:"format,omitempty"`
	OptionCount int    `json:"option_count,omitempty"`
}

// SentWebAppMessage contains information about an inline message sent by a Web App on behalf of a user.
type SentWebAppMessage struct {
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

// PollOption contains information about a poll option.
type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

// Poll contains information about a poll.
type Poll struct {
	ID                    string          `json:"id"`
	Question              string          `json:"question"`
	Options               []PollOption    `json:"options"`
	TotalVoterCount       int             `json:"total_voter_count"`
	IsClosed              bool            `json:"is_closed"`
	IsAnonymous           bool            `json:"is_anonymous"`
	Type                  string          `json:"type"`
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"`
	CorrectOptionID       int             `json:"correct_option_id,omitempty"`
	Explanation           string          `json:"explanation,omitempty"`
	ExplanationEntities   []MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            int             `json:"open_period,omitempty"`
	CloseDate             int             `json:"close_date,omitempty"`
}

// StarAmount contains information about a star amount.
type StarAmount struct {
	Amount int `json:"amount"`
}

// RefundedPayment contains information about a refunded payment.
type RefundedPayment struct {
	Currency    string `json:"currency"`
	TotalAmount int    `json:"total_amount"`
}

// StarTransactions contains information about a list of star transactions.
type StarTransactions struct {
	Transactions []StarTransaction `json:"transactions"`
}

// StarTransaction contains information about a star transaction.
type StarTransaction struct {
	ID                           string `json:"id"`
	Amount                       int    `json:"amount"`
	ReverseAmount                int    `json:"reverse_amount"`
	Source                       string `json:"source,omitempty"`
	Destination                  string `json:"destination,omitempty"`
	TelegramPaymentChargeID      string `json:"telegram_payment_charge_id,omitempty"`
	ProviderPaymentChargeID      string `json:"provider_payment_charge_id,omitempty"`
	TelegramSubscriptionChargeID string `json:"telegram_subscription_charge_id,omitempty"`
	Satoshi                      int    `json:"satoshi,omitempty"`
}

// Gift contains information about a gift.
type Gift struct {
	ID             string  `json:"id"`
	Sticker        Sticker `json:"sticker"`
	StickerSetName string  `json:"sticker_set_name,omitempty"`
	Count          int     `json:"count"`
	IsDuplicate    bool    `json:"is_duplicate,omitempty"`
}

// GiftInfo contains information about a gift.
type GiftInfo struct {
	Gift        Gift  `json:"gift"`
	IsAnonymous bool  `json:"is_anonymous,omitempty"`
	Recipient   *User `json:"recipient,omitempty"`
	SentDate    int   `json:"sent_date,omitempty"`
}

// OwnedGift contains information about an owned gift.
type OwnedGift struct {
	Gift         Gift  `json:"gift"`
	Sender       *User `json:"sender,omitempty"`
	Receiver     *User `json:"receiver,omitempty"`
	SentDate     int   `json:"sent_date,omitempty"`
	IsIncoming   bool  `json:"is_incoming,omitempty"`
	IsDownloaded bool  `json:"is_downloaded,omitempty"`
	StarCount    int   `json:"star_count,omitempty"`
}

// UniqueGift contains information about a unique gift.
type UniqueGift struct {
	Gift         Gift        `json:"gift"`
	UniqueGiftID string      `json:"unique_gift_id"`
	Sender       *User       `json:"sender,omitempty"`
	Receiver     *User       `json:"receiver,omitempty"`
	SentDate     int         `json:"sent_date,omitempty"`
	IsIncoming   bool        `json:"is_incoming,omitempty"`
	IsDownloaded bool        `json:"is_downloaded,omitempty"`
	Color        *GiftColors `json:"color,omitempty"`
	Animated     bool        `json:"animated,omitempty"`
}

// GiftColors contains information about gift colors.
type GiftColors struct {
	Colors []int `json:"colors"`
}

// UserProfileAudios contains information about user profile audios.
type UserProfileAudios struct {
	Audios []Audio `json:"audios"`
}

// UserGift contains information about a user gift.
type UserGift struct {
	ID           string `json:"id"`
	From         *User  `json:"from"`
	Gift         Gift   `json:"gift"`
	IsAnonymous  bool   `json:"is_anonymous,omitempty"`
	IsDuplicate  bool   `json:"is_duplicate,omitempty"`
	SentDate     int    `json:"sent_date,omitempty"`
	IsDownloaded bool   `json:"is_downloaded,omitempty"`
	StarCount    int    `json:"star_count,omitempty"`
	IsFromGroup  bool   `json:"is_from_group,omitempty"`
	IsOutgoing   bool   `json:"is_outgoing,omitempty"`
	Author       string `json:"author,omitempty"`
	Title        string `json:"title,omitempty"`
}
