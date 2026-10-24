package method

// GetUserProfilePhotos contains information about getUserProfilePhotos telegram api method.
type GetUserProfilePhotos struct {
	UserID int `json:"user_id"`
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
	SimpleMethod
}

// SetUserID sets user_id
func (u *GetUserProfilePhotos) SetUserID(id int) {
	u.UserID = id
}

// SetOffset sets offset
func (u *GetUserProfilePhotos) SetOffset(v int) {
	u.Offset = v
}

// SetLimit sets limit
func (u *GetUserProfilePhotos) SetLimit(v int) {
	u.Limit = v
}

// NewGetUserProfilePhotos creates GetUserProfilePhotos.
func NewGetUserProfilePhotos(userID int, opts ...Option[*GetUserProfilePhotos]) *GetUserProfilePhotos {
	m := &GetUserProfilePhotos{
		UserID: userID,
	}
	for _, o := range opts {
		o(m)
	}
	return m
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

// SetChatID sets chat_id
func (k *KickChatMember) SetChatID(id int64) {
	k.ChatID = id
}

// SetUserID sets user_id
func (k *KickChatMember) SetUserID(id int) {
	k.UserID = id
}

// SetUntilDate sets until_date
func (k *KickChatMember) SetUntilDate(v int) {
	k.UntilDate = v
}

// NewKickChatMember creates KickChatMember.
func NewKickChatMember(userID int, chatID int64, opts ...Option[*KickChatMember]) *KickChatMember {
	m := &KickChatMember{
		ChatID: chatID,
		UserID: userID,
	}
	for _, o := range opts {
		o(m)
	}
	return m
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

// SetChatID sets chat_id
func (u *UnbanChatMember) SetChatID(id int64) {
	u.ChatID = id
}

// SetUserID sets user_id
func (u *UnbanChatMember) SetUserID(id int) {
	u.UserID = id
}

// NewUnbanChatMember creates UnbanChatMember.
func NewUnbanChatMember(userID int, chatID int64) *UnbanChatMember {
	m := &UnbanChatMember{
		ChatID: chatID,
		UserID: userID,
	}
	return m
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

// SetChatID sets chat_id
func (r *RestrictChatMember) SetChatID(id int64) {
	r.ChatID = id
}

// SetUserID sets user_id
func (r *RestrictChatMember) SetUserID(id int) {
	r.UserID = id
}

// SetUntilDate sets until_date
func (r *RestrictChatMember) SetUntilDate(v int) {
	r.UntilDate = v
}

// SetCanSendMessages sets can_send_messages
func (r *RestrictChatMember) SetCanSendMessages(v bool) {
	r.CanSendMessages = v
}

// SetCanSendMediaMessages sets can_send_media_messages
func (r *RestrictChatMember) SetCanSendMediaMessages(v bool) {
	r.CanSendMediaMessages = v
}

// SetCanSendOtherMessages sets can_send_other_messages
func (r *RestrictChatMember) SetCanSendOtherMessages(v bool) {
	r.CanSendOtherMessages = v
}

// SetCanAddWebPagePreviews sets can_add_web_page_previews
func (r *RestrictChatMember) SetCanAddWebPagePreviews(v bool) {
	r.CanAddWebPagePreviews = v
}

// NewRestrictChatMember creates RestrictChatMember.
func NewRestrictChatMember(userID int, chatID int64, opts ...Option[*RestrictChatMember]) *RestrictChatMember {
	m := &RestrictChatMember{
		ChatID: chatID,
		UserID: userID,
	}
	for _, o := range opts {
		o(m)
	}
	return m
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

// SetChatID sets chat_id
func (p *PromoteChatMember) SetChatID(id int64) {
	p.ChatID = id
}

// SetUserID sets user_id
func (p *PromoteChatMember) SetUserID(id int) {
	p.UserID = id
}

// SetCanChangeInfo sets can_change_info
func (p *PromoteChatMember) SetCanChangeInfo(v bool) {
	p.CanChangeInfo = v
}

// SetCanPostMessages sets can_post_messages
func (p *PromoteChatMember) SetCanPostMessages(v bool) {
	p.CanPostMessages = v
}

// SetCanEditMessages sets can_edit_messages
func (p *PromoteChatMember) SetCanEditMessages(v bool) {
	p.CanEditMessages = v
}

// SetCanDeleteMessages sets can_delete_messages
func (p *PromoteChatMember) SetCanDeleteMessages(v bool) {
	p.CanDeleteMessages = v
}

// SetCanInviteUsers sets can_invite_users
func (p *PromoteChatMember) SetCanInviteUsers(v bool) {
	p.CanInviteUsers = v
}

// SetCanRestrictMembers sets can_restrict_members
func (p *PromoteChatMember) SetCanRestrictMembers(v bool) {
	p.CanRestrictMembers = v
}

// SetCanPinMessages sets can_pin_messages
func (p *PromoteChatMember) SetCanPinMessages(v bool) {
	p.CanPinMessages = v
}

// SetCanPromoteMembers sets can_promote_members
func (p *PromoteChatMember) SetCanPromoteMembers(v bool) {
	p.CanPromoteMembers = v
}

// NewPromoteChatMember creates PromoteChatMember.
func NewPromoteChatMember(userID int, chatID int64, opts ...Option[*PromoteChatMember]) *PromoteChatMember {
	m := &PromoteChatMember{
		ChatID: chatID,
		UserID: userID,
	}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *PromoteChatMember) Method() string {
	return "promoteChatMember"
}
