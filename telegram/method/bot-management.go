package method

import (
)

// BotCommand contains information about bot command.
type BotCommand struct {
	Command     string `json:"command" form:"command"`
	Description string `json:"description" form:"description"`
}

// SetMyCommands contains information about setMyCommands telegram api method.
type SetMyCommands struct {
	Commands      []BotCommand `json:"commands" form:"commands"`
	Scope         string       `json:"scope,omitempty" form:"scope,omitempty"`
	LanguageCode  string       `json:"language_code,omitempty" form:"language_code,omitempty"`
	SimpleMethod
}

// NewSetMyCommands creates SetMyCommands.
func NewSetMyCommands(commands []BotCommand, opts ...Option[*SetMyCommands]) *SetMyCommands {
	m := &SetMyCommands{
		Commands: commands,
	}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *SetMyCommands) Method() string {
	return "setMyCommands"
}

// SetScope sets scope
func (r *SetMyCommands) SetScope(s string) {
	r.Scope = s
}

// SetLanguageCode sets language_code
func (r *SetMyCommands) SetLanguageCode(l string) {
	r.LanguageCode = l
}

// GetMyCommands contains information about getMyCommands telegram api method.
type GetMyCommands struct {
	Scope        string `json:"scope,omitempty" form:"scope,omitempty"`
	LanguageCode string `json:"language_code,omitempty" form:"language_code,omitempty"`
	SimpleMethod
}

// NewGetMyCommands creates GetMyCommands.
func NewGetMyCommands(opts ...Option[*GetMyCommands]) *GetMyCommands {
	m := &GetMyCommands{}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *GetMyCommands) Method() string {
	return "getMyCommands"
}

// SetScope sets scope
func (r *GetMyCommands) SetScope(s string) {
	r.Scope = s
}

// SetLanguageCode sets language_code
func (r *GetMyCommands) SetLanguageCode(l string) {
	r.LanguageCode = l
}

// DeleteMyCommands contains information about deleteMyCommands telegram api method.
type DeleteMyCommands struct {
	Scope        string `json:"scope,omitempty" form:"scope,omitempty"`
	LanguageCode string `json:"language_code,omitempty" form:"language_code,omitempty"`
	SimpleMethod
}

// NewDeleteMyCommands creates DeleteMyCommands.
func NewDeleteMyCommands(opts ...Option[*DeleteMyCommands]) *DeleteMyCommands {
	m := &DeleteMyCommands{}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *DeleteMyCommands) Method() string {
	return "deleteMyCommands"
}

// SetScope sets scope
func (r *DeleteMyCommands) SetScope(s string) {
	r.Scope = s
}

// SetLanguageCode sets language_code
func (r *DeleteMyCommands) SetLanguageCode(l string) {
	r.LanguageCode = l
}

// BotName contains information about bot name.
type BotName struct {
	Name string `json:"name" form:"name"`
}

// SetMyName contains information about setMyName telegram api method.
type SetMyName struct {
	Name         string `json:"name,omitempty" form:"name,omitempty"`
	LanguageCode string `json:"language_code,omitempty" form:"language_code,omitempty"`
	SimpleMethod
}

// NewSetMyName creates SetMyName.
func NewSetMyName(opts ...Option[*SetMyName]) *SetMyName {
	m := &SetMyName{}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *SetMyName) Method() string {
	return "setMyName"
}

// SetName sets name
func (r *SetMyName) SetName(n string) {
	r.Name = n
}

// SetLanguageCode sets language_code
func (r *SetMyName) SetLanguageCode(l string) {
	r.LanguageCode = l
}

// GetMyName contains information about getMyName telegram api method.
type GetMyName struct {
	LanguageCode string `json:"language_code,omitempty" form:"language_code,omitempty"`
	SimpleMethod
}

// NewGetMyName creates GetMyName.
func NewGetMyName(opts ...Option[*GetMyName]) *GetMyName {
	m := &GetMyName{}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *GetMyName) Method() string {
	return "getMyName"
}

// SetLanguageCode sets language_code
func (r *GetMyName) SetLanguageCode(l string) {
	r.LanguageCode = l
}

// BotDescription contains information about bot description.
type BotDescription struct {
	Description string `json:"description" form:"description"`
}

// SetMyDescription contains information about setMyDescription telegram api method.
type SetMyDescription struct {
	Description  string `json:"description,omitempty" form:"description,omitempty"`
	LanguageCode string `json:"language_code,omitempty" form:"language_code,omitempty"`
	SimpleMethod
}

// NewSetMyDescription creates SetMyDescription.
func NewSetMyDescription(opts ...Option[*SetMyDescription]) *SetMyDescription {
	m := &SetMyDescription{}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *SetMyDescription) Method() string {
	return "setMyDescription"
}

// SetDescription sets description
func (r *SetMyDescription) SetDescription(d string) {
	r.Description = d
}

// SetLanguageCode sets language_code
func (r *SetMyDescription) SetLanguageCode(l string) {
	r.LanguageCode = l
}

// GetMyDescription contains information about getMyDescription telegram api method.
type GetMyDescription struct {
	LanguageCode string `json:"language_code,omitempty" form:"language_code,omitempty"`
	SimpleMethod
}

// NewGetMyDescription creates GetMyDescription.
func NewGetMyDescription(opts ...Option[*GetMyDescription]) *GetMyDescription {
	m := &GetMyDescription{}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *GetMyDescription) Method() string {
	return "getMyDescription"
}

// SetLanguageCode sets language_code
func (r *GetMyDescription) SetLanguageCode(l string) {
	r.LanguageCode = l
}

// BotShortDescription contains information about bot short description.
type BotShortDescription struct {
	ShortDescription string `json:"short_description" form:"short_description"`
}

// SetMyShortDescription contains information about setMyShortDescription telegram api method.
type SetMyShortDescription struct {
	ShortDescription string `json:"short_description,omitempty" form:"short_description,omitempty"`
	LanguageCode     string `json:"language_code,omitempty" form:"language_code,omitempty"`
	SimpleMethod
}

// NewSetMyShortDescription creates SetMyShortDescription.
func NewSetMyShortDescription(opts ...Option[*SetMyShortDescription]) *SetMyShortDescription {
	m := &SetMyShortDescription{}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *SetMyShortDescription) Method() string {
	return "setMyShortDescription"
}

// SetShortDescription sets short_description
func (r *SetMyShortDescription) SetShortDescription(s string) {
	r.ShortDescription = s
}

// SetLanguageCode sets language_code
func (r *SetMyShortDescription) SetLanguageCode(l string) {
	r.LanguageCode = l
}

// GetMyShortDescription contains information about getMyShortDescription telegram api method.
type GetMyShortDescription struct {
	LanguageCode string `json:"language_code,omitempty" form:"language_code,omitempty"`
	SimpleMethod
}

// NewGetMyShortDescription creates GetMyShortDescription.
func NewGetMyShortDescription(opts ...Option[*GetMyShortDescription]) *GetMyShortDescription {
	m := &GetMyShortDescription{}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *GetMyShortDescription) Method() string {
	return "getMyShortDescription"
}

// SetLanguageCode sets language_code
func (r *GetMyShortDescription) SetLanguageCode(l string) {
	r.LanguageCode = l
}

// BotAdministratorRights contains information about bot administrator rights.
type BotAdministratorRights struct {
	IsAnonymous        bool `json:"is_anonymous" form:"is_anonymous"`
	CanManageChat      bool `json:"can_manage_chat" form:"can_manage_chat"`
	CanChangeInfo      bool `json:"can_change_info" form:"can_change_info"`
	CanPostMessages    bool `json:"can_post_messages" form:"can_post_messages"`
	CanEditMessages    bool `json:"can_edit_messages" form:"can_edit_messages"`
	CanDeleteMessages  bool `json:"can_delete_messages" form:"can_delete_messages"`
	CanInviteUsers     bool `json:"can_invite_users" form:"can_invite_users"`
	CanRestrictMembers bool `json:"can_restrict_members" form:"can_restrict_members"`
	CanPinMessages     bool `json:"can_pin_messages" form:"can_pin_messages"`
	CanManageTopics    bool `json:"can_manage_topics" form:"can_manage_topics"`
}

// SetMyDefaultAdministratorRights contains information about setMyDefaultAdministratorRights telegram api method.
type SetMyDefaultAdministratorRights struct {
	Rights       *BotAdministratorRights `json:"rights" form:"rights"`
	ForSupergroups bool                  `json:"for_supergroups,omitempty" form:"for_supergroups,omitempty"`
	ForChannels    bool                  `json:"for_channels,omitempty" form:"for_channels,omitempty"`
	SimpleMethod
}

// NewSetMyDefaultAdministratorRights creates SetMyDefaultAdministratorRights.
func NewSetMyDefaultAdministratorRights(rights *BotAdministratorRights, opts ...Option[*SetMyDefaultAdministratorRights]) *SetMyDefaultAdministratorRights {
	m := &SetMyDefaultAdministratorRights{
		Rights: rights,
	}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *SetMyDefaultAdministratorRights) Method() string {
	return "setMyDefaultAdministratorRights"
}

// SetForSupergroups sets for_supergroups
func (r *SetMyDefaultAdministratorRights) SetForSupergroups(f bool) {
	r.ForSupergroups = f
}

// SetForChannels sets for_channels
func (r *SetMyDefaultAdministratorRights) SetForChannels(f bool) {
	r.ForChannels = f
}

// GetMyDefaultAdministratorRights contains information about getMyDefaultAdministratorRights telegram api method.
type GetMyDefaultAdministratorRights struct {
	ForSupergroups bool `json:"for_supergroups,omitempty" form:"for_supergroups,omitempty"`
	ForChannels    bool `json:"for_channels,omitempty" form:"for_channels,omitempty"`
	SimpleMethod
}

// NewGetMyDefaultAdministratorRights creates GetMyDefaultAdministratorRights.
func NewGetMyDefaultAdministratorRights(opts ...Option[*GetMyDefaultAdministratorRights]) *GetMyDefaultAdministratorRights {
	m := &GetMyDefaultAdministratorRights{}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *GetMyDefaultAdministratorRights) Method() string {
	return "getMyDefaultAdministratorRights"
}

// SetForSupergroups sets for_supergroups
func (r *GetMyDefaultAdministratorRights) SetForSupergroups(f bool) {
	r.ForSupergroups = f
}

// SetForChannels sets for_channels
func (r *GetMyDefaultAdministratorRights) SetForChannels(f bool) {
	r.ForChannels = f
}

// MenuButton contains information about menu button.
type MenuButton struct {
	Type         string `json:"type" form:"type"`
	Text         string `json:"text,omitempty" form:"text,omitempty"`
	WebURL       string `json:"web_app,omitempty" form:"web_app,omitempty"`
	ShortName    string `json:"short_name,omitempty" form:"short_name,omitempty"`
}

// SetChatMenuButton contains information about setChatMenuButton telegram api method.
type SetChatMenuButton struct {
	ChatID       int64  `json:"chat_id,omitempty" form:"chat_id,omitempty"`
	MenuButton   string `json:"menu_button" form:"menu_button"`
	SimpleMethod
}

// NewSetChatMenuButton creates SetChatMenuButton.
func NewSetChatMenuButton(opts ...Option[*SetChatMenuButton]) *SetChatMenuButton {
	m := &SetChatMenuButton{}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *SetChatMenuButton) Method() string {
	return "setChatMenuButton"
}

// SetChatID sets chat_id
func (r *SetChatMenuButton) SetChatID(c int64) {
	r.ChatID = c
}

// SetMenuButton sets menu_button
func (r *SetChatMenuButton) SetMenuButton(m string) {
	r.MenuButton = m
}

// GetChatMenuButton contains information about getChatMenuButton telegram api method.
type GetChatMenuButton struct {
	ChatID int64 `json:"chat_id,omitempty" form:"chat_id,omitempty"`
	SimpleMethod
}

// NewGetChatMenuButton creates GetChatMenuButton.
func NewGetChatMenuButton(opts ...Option[*GetChatMenuButton]) *GetChatMenuButton {
	m := &GetChatMenuButton{}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *GetChatMenuButton) Method() string {
	return "getChatMenuButton"
}

// SetChatID sets chat_id
func (r *GetChatMenuButton) SetChatID(c int64) {
	r.ChatID = c
}
