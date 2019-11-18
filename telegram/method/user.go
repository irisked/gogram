package method

// GetUserProfilePhotos contains information about getUserProfilePhotos telegram api method.
type GetUserProfilePhotos struct {
	UserID int `json:"user_id"`
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
	SimpleMethod
}

// GetUserProfilePhotosOption interface for optional parameters.
type GetUserProfilePhotosOption interface {
	ApplyGetUserProfilePhotosOption(*GetUserProfilePhotos)
}

// NewGetUserProfilePhotos creates GetUserProfilePhotos.
func NewGetUserProfilePhotos(userID int, options ...GetUserProfilePhotosOption) *GetUserProfilePhotos {
	method := new(GetUserProfilePhotos)
	method.UserID = userID
	for _, opt := range options {
		opt.ApplyGetUserProfilePhotosOption(method)
	}
	return method
}

// Method returns telegram api method endpoint.
func (r *GetUserProfilePhotos) Method() string {
	return "getUserProfilePhotos"
}

// KickChatMember contains information about kickChatMember telegram api method.
type KickChatMember struct {
	ChatID    int64 `json:"chat_id"`
	UserID    int   `json:"user_id"`
	UntilDate int   `json:"until_date,omitempty"`
	SimpleMethod
}

// KickChatMemberOption interface for optional parameters.
type KickChatMemberOption interface {
	ApplyKickChatMemberOption(*KickChatMember)
}

// NewKickChatMember creates KickChatMember.
func NewKickChatMember(userID int, chatID int64, options ...KickChatMemberOption) *KickChatMember {
	method := new(KickChatMember)
	method.ChatID = chatID
	method.UserID = userID
	for _, opt := range options {
		opt.ApplyKickChatMemberOption(method)
	}
	return method
}

// Method returns telegram api method endpoint.
func (r *KickChatMember) Method() string {
	return "kickChatMember"
}

// UnbanChatMember contains information about unbanChatMember telegram api method.
type UnbanChatMember struct {
	ChatID int64 `json:"chat_id"`
	UserID int   `json:"user_id"`
	SimpleMethod
}

// NewUnbanChatMember creates UnbanChatMember.
func NewUnbanChatMember(userID int, chatID int64) *UnbanChatMember {
	method := new(UnbanChatMember)
	method.ChatID = chatID
	method.UserID = userID
	return method
}

// Method returns telegram api method endpoint.
func (r *UnbanChatMember) Method() string {
	return "unbanChatMember"
}

// RestrictChatMember contains information about restrictChatMember telegram api method.
type RestrictChatMember struct {
	ChatID                int64 `json:"chat_id"`
	UserID                int   `json:"user_id"`
	UntilDate             int   `json:"until_date,omitempty"`
	CanSendMessages       bool  `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  bool  `json:"can_send_media_messages,omitempty"`
	CanSendOtherMessages  bool  `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool  `json:"can_add_web_page_previews,omitempty"`
	SimpleMethod
}

// RestrictChatMemberOption interface for optional parameters.
type RestrictChatMemberOption interface {
	ApplyRestrictChatMemberOption(*RestrictChatMember)
}

// NewRestrictChatMember creates RestrictChatMember.
func NewRestrictChatMember(userID int, chatID int64, options ...RestrictChatMemberOption) *RestrictChatMember {
	method := new(RestrictChatMember)
	method.ChatID = chatID
	method.UserID = userID
	for _, opt := range options {
		opt.ApplyRestrictChatMemberOption(method)
	}
	return method
}

// Method returns telegram api method endpoint.
func (r *RestrictChatMember) Method() string {
	return "restrictChatMember"
}

// PromoteChatMember contains information about promoteChatMember telegram api method.
type PromoteChatMember struct {
	ChatID             int64 `json:"chat_id"`
	UserID             int   `json:"user_id"`
	CanChangeInfo      bool  `json:"can_change_info,omitempty"`
	CanPostMessages    bool  `json:"can_post_messages,omitempty"`
	CanEditMessages    bool  `json:"can_edit_messages,omitempty"`
	CanDeleteMessages  bool  `json:"can_delete_messages,omitempty"`
	CanInviteUsers     bool  `json:"can_invite_users,omitempty"`
	CanRestrictMembers bool  `json:"can_restrict_members,omitempty"`
	CanPinMessages     bool  `json:"can_pin_messages,omitempty"`
	CanPromoteMembers  bool  `json:"can_promote_members,omitempty"`
	SimpleMethod
}

// NewPromoteChatMember creates PromoteChatMember.
func NewPromoteChatMember(userID int, chatID int64, options ...PromoteChatMemberOption) *PromoteChatMember {
	method := new(PromoteChatMember)
	method.ChatID = chatID
	method.UserID = userID
	for _, opt := range options {
		opt.ApplyPromoteChatMemberOption(method)
	}
	return method
}

// Method returns telegram api method endpoint.
func (r *PromoteChatMember) Method() string {
	return "promoteChatMember"
}

// PromoteChatMemberOption interface for optional parameters.
type PromoteChatMemberOption interface {
	ApplyPromoteChatMemberOption(*PromoteChatMember)
}
