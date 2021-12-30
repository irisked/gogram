package method

import (
	"fmt"
	"os"

	"gogram/types"
)

// SendChatAction contains information about the sendChatAction telegram api method.
type SendChatAction struct {
	ChatID int64            `json:"chat_id"`
	Action types.ChatAction `json:"action"`
	SimpleMethod
}

// NewSendChatAction creates SendChatAction struct.
func NewSendChatAction(chatID int64, action types.ChatAction) *SendChatAction {
	method := new(SendChatAction)
	method.ChatID = chatID
	method.Action = action
	return method
}

// Method returns telegram api method endpoint.
func (r *SendChatAction) Method() string {
	return "sendChatAction"
}

// ExportChatInviteLink contains information about the exportChatInviteLink telegram api method.
type ExportChatInviteLink struct {
	ChatID int64 `json:"chat_id"`
	SimpleMethod
}

// NewExportChatInviteLink creates ExportChatInviteLink struct.
func NewExportChatInviteLink(chatID int64) *ExportChatInviteLink {
	method := new(ExportChatInviteLink)
	method.ChatID = chatID
	return method
}

// Method returns telegram api method endpoint.
func (r *ExportChatInviteLink) Method() string {
	return "exportChatInviteLink"
}

// SetChatPhoto contains information about the SetChatPhoto telegram api method.
type SetChatPhoto struct {
	File   *os.File `form:"photo"`
	ChatID int64    `json:"chat_id" form:"chat_id"`
	Photo  string   `json:"photo" form:"photo"`
	MultipartMethod
}

// SetChatPhotoOption interface for optional parameters.
type SetChatPhotoOption interface {
	ApplySetChatPhotoOption(*SetChatPhoto)
}

// NewSetChatPhoto creates SetChatPhoto struct.
func NewSetChatPhoto(chatID int64, file *os.File) *SetChatPhoto {
	method := new(SetChatPhoto)
	method.ChatID = chatID
	method.File = file
	method.Photo = fmt.Sprintf("attachment://%s", file.Name())
	return method
}

// Method returns telegram api method endpoint.
func (r *SetChatPhoto) Method() string {
	return "SetChatPhoto"
}

// DeleteChatPhoto contains information about the deleteChatPhoto telegram api method.
type DeleteChatPhoto struct {
	ChatID int64 `json:"chat_id"`
	SimpleMethod
}

// NewDeleteChatPhoto creates DeleteChatPhoto struct
func NewDeleteChatPhoto(chatID int64) *DeleteChatPhoto {
	method := new(DeleteChatPhoto)
	method.ChatID = chatID
	return method
}

// Method returns telegram api method endpoint.
func (r *DeleteChatPhoto) Method() string {
	return "deleteChatPhoto"
}

// SetChatTitle contains information about SetChatTitle telegram api method.
type SetChatTitle struct {
	ChatID int64  `json:"chat_id"`
	Title  string `json:"title"`
	SimpleMethod
}

// NewSetChatTitle creates SendMessage struct.
func NewSetChatTitle(chatID int64, title string) *SetChatTitle {
	method := new(SetChatTitle)
	method.ChatID = chatID
	method.Title = title
	return method
}

// Method returns telegram api method endpoint.
func (r *SetChatTitle) Method() string {
	return "SetChatTitle"
}

// SetChatDescription contains information about SetChatDescription telegram api method.
type SetChatDescription struct {
	ChatID      int64  `json:"chat_id"`
	Description string `json:"description"`
	SimpleMethod
}

// NewSetChatDescription creates SetChatDescription struct.
func NewSetChatDescription(chatID int64, desc string) *SetChatDescription {
	method := new(SetChatDescription)
	method.ChatID = chatID
	method.Description = desc
	return method
}

// Method returns telegram api method endpoint.
func (r *SetChatDescription) Method() string {
	return "SetChatDescription"
}

// PinChatMessage contains information about pinChatMessage telegram api method.
type PinChatMessage struct {
	ChatID              int64 `json:"chat_id"`
	MessageID           int   `json:"message_id"`
	DisableNotification bool  `json:"disable_notification,omitempty"`
	SimpleMethod
}

// PinChatMessageOption interface for optional parameters.
type PinChatMessageOption interface {
	ApplyPinChatMessageOption(*PinChatMessage)
}

// NewPinChatMessage creates PinChatMessage struct.
func NewPinChatMessage(chatID int64, messageID int, options ...PinChatMessageOption) *PinChatMessage {
	method := new(PinChatMessage)
	method.ChatID = chatID
	method.MessageID = messageID
	for _, option := range options {
		option.ApplyPinChatMessageOption(method)
	}
	return method
}

// Method returns telegram api method endpoint.
func (r *PinChatMessage) Method() string {
	return "pinChatMessage"
}

// UnpinChatMessage contains information about unpinChatMessage telegram api method.
type UnpinChatMessage struct {
	ChatID int64 `json:"chat_id"`
	SimpleMethod
}

// NewUnpinChatMessage creates UnpinChatMessage struct.
func NewUnpinChatMessage(chatID int64) *UnpinChatMessage {
	method := new(UnpinChatMessage)
	method.ChatID = chatID
	return method
}

// Method returns telegram api method endpoint.
func (r *UnpinChatMessage) Method() string {
	return "unpinChatMessage"
}

// LeaveChat contains information about leaveChat telegram api method.
type LeaveChat struct {
	ChatID int64 `json:"chat_id"`
	SimpleMethod
}

// NewLeaveChat creates LeaveChat struct.
func NewLeaveChat(chatID int64) *LeaveChat {
	method := new(LeaveChat)
	method.ChatID = chatID
	return method
}

// Method returns telegram api method endpoint.
func (r *LeaveChat) Method() string {
	return "leaveChat"
}

// GetChat contains information about getChat rtelegram api method.
type GetChat struct {
	ChatID int64 `json:"chat_id"`
	SimpleMethod
}

// NewGetChat creates GetChat struct.
func NewGetChat(chatID int64) *GetChat {
	method := new(GetChat)
	method.ChatID = chatID
	return method
}

// Method returns telegram api method endpoint.
func (r *GetChat) Method() string {
	return "getChat"
}

// GetChatAdministrators contains information about getChatAdministrators telegram api method.
type GetChatAdministrators struct {
	ChatID int64 `json:"chat_id"`
	SimpleMethod
}

// NewGetChatAdministrators creates GetChatAdministrators struct.
func NewGetChatAdministrators(chatID int64) *GetChatAdministrators {
	method := new(GetChatAdministrators)
	method.ChatID = chatID
	return method
}

// Method returns telegram api method endpoint.
func (r *GetChatAdministrators) Method() string {
	return "getChatAdministrators"
}

// GetChatMembersCount contains information about getChatMembersCount telegram api method.
type GetChatMembersCount struct {
	ChatID int64 `json:"chat_id"`
	SimpleMethod
}

// NewGetChatMembersCount creates GetChatMembersCount struct.
func NewGetChatMembersCount(chatID int64) *GetChatMembersCount {
	method := new(GetChatMembersCount)
	method.ChatID = chatID
	return method
}

// Method returns telegram api method endpoint.
func (r *GetChatMembersCount) Method() string {
	return "getChatMembersCount"
}

// GetChatMember contains information about getChatMember telegram api method.
type GetChatMember struct {
	ChatID int64 `json:"chat_id"`
	UserID int   `json:"user_id"`
	SimpleMethod
}

// NewGetChatMember creates GetChatMember struct method.
func NewGetChatMember(chatID int64, userID int) *GetChatMember {
	method := new(GetChatMember)
	method.ChatID = chatID
	method.UserID = userID
	return method
}

// Method returns telegram api method endpoint.
func (r *GetChatMember) Method() string {
	return "getChatMember"
}

// SetChatStickerSet contains information about setChatStickerSet telegram api method.
type SetChatStickerSet struct {
	ChatID         int64  `json:"chat_id"`
	StickerSetName string `json:"sticker_set_name"`
	SimpleMethod
}

// NewSetChatStickerSet creates SetChatStickerSet struct.
func NewSetChatStickerSet(chatID int64, stickerSetName string) *SetChatStickerSet {
	method := new(SetChatStickerSet)
	method.StickerSetName = stickerSetName
	method.ChatID = chatID
	return method
}

// Method returns telegram api method endpoint.
func (r *SetChatStickerSet) Method() string {
	return "SetChatStickerSet"
}

// DeleteChatStickerSet contains information about deleteChatStickerSet telegram api method.
type DeleteChatStickerSet struct {
	ChatID int64 `json:"chat_id"`
	SimpleMethod
}

// NewDeleteChatStickerSet creates DeleteChatStickerSet struct.
func NewDeleteChatStickerSet(chatID int64) *DeleteChatStickerSet {
	method := new(DeleteChatStickerSet)
	method.ChatID = chatID
	return method
}

// Method returns telegram api method endpoint.
func (r *DeleteChatStickerSet) Method() string {
	return "deleteChatStickerSet"
}
