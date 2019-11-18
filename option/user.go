package option

import (
	"github.com/irisked/gogram/telegram/method"
)

// CanChangeInfo returns struct for setting optional CanChangeInfo field.
func CanChangeInfo() interface {
	canSendMessagesOptionSetter
} {
	return &canSendMessages{true}
}

// CanPostMessages returns struct for setting optional CanPostMessages field.
func CanPostMessages() interface {
	canSendMessagesOptionSetter
} {
	return &canSendMessages{true}
}

// CanDeleteMessages returns struct for setting optional CanDeleteMessages field.
func CanDeleteMessages() interface {
	canSendMessagesOptionSetter
} {
	return &canSendMessages{true}
}

// CanInviteUsers returns struct for setting optional CanInviteUsers field.
func CanInviteUsers() interface {
	canSendMessagesOptionSetter
} {
	return &canSendMessages{true}
}

// CanRestrictMembers returns struct for setting optional CanRestrictMembers field.
func CanRestrictMembers() interface {
	canSendMessagesOptionSetter
} {
	return &canSendMessages{true}
}

// CanPinMessages returns struct for setting optional CanPinMessages field.
func CanPinMessages() interface {
	canSendMessagesOptionSetter
} {
	return &canSendMessages{true}
}

// CanPromoteMembers returns struct for setting optional CanPromoteMembers field.
func CanPromoteMembers() interface {
	canSendMessagesOptionSetter
} {
	return &canSendMessages{true}
}

// CanSendMessages returns struct for setting optional CanSendMessages field.
func CanSendMessages() interface {
	canSendMessagesOptionSetter
} {
	return &canSendMessages{true}
}

// CanSendMediaMessages returns struct for setting optional CanSendMessages field.
func CanSendMediaMessages() interface {
	canSendMediaMessagesOptionSetter
} {
	return &canSendMediaMessages{true}
}

// CanSendOtherMessages returns struct for setting optional CanSendMessages field.
func CanSendOtherMessages() interface {
	canSendOtherMessagesOptionSetter
} {
	return &canSendMediaMessages{true}
}

// CanAddWebPagePreviews returns struct for setting optional CanSendMessages field.
func CanAddWebPagePreviews() interface {
	canAddWebPagePreviewsOptionSetter
} {
	return &canSendMediaMessages{true}
}

// canSendMessagesOptionSetter is an interface for setting showAlert option.
type canSendMessagesOptionSetter interface {
	method.RestrictChatMemberOption
}

type canSendMessages struct {
	flag bool
}

func (csm *canSendMessages) ApplyRestrictChatMemberOption(acq *method.RestrictChatMember) {
	acq.CanSendMessages = csm.flag
}

// canSendMediaMessagesOptionSetter is an interface for setting showAlert option.
type canSendMediaMessagesOptionSetter interface {
	method.RestrictChatMemberOption
}

type canSendMediaMessages struct {
	flag bool
}

func (csm *canSendMediaMessages) ApplyRestrictChatMemberOption(acq *method.RestrictChatMember) {
	acq.CanSendMediaMessages = csm.flag
}

// canSendOtherMessagesOptionSetter is an interface for setting showAlert option.
type canSendOtherMessagesOptionSetter interface {
	method.RestrictChatMemberOption
}

type canSendOtherMessages struct {
	flag bool
}

func (csm *canSendOtherMessages) ApplyRestrictChatMemberOption(acq *method.RestrictChatMember) {
	acq.CanSendOtherMessages = csm.flag
}

// canAddWebPagePreviewsOptionSetter is an interface for setting showAlert option.
type canAddWebPagePreviewsOptionSetter interface {
	method.RestrictChatMemberOption
}

type canAddWebPagePreviews struct {
	flag bool
}

func (csm *canAddWebPagePreviews) ApplyRestrictChatMemberOption(acq *method.RestrictChatMember) {
	acq.CanSendOtherMessages = csm.flag
}

// canChangeInfoSetter is an interface for setting showAlert option.
type canChangeInfoSetter interface {
	method.PromoteChatMemberOption
}

type canChangeInfoOption struct {
	flag bool
}

func (csm *canChangeInfoOption) ApplyPromoteChatMemberOption(acq *method.PromoteChatMember) {
	acq.CanChangeInfo = csm.flag
}

// canPostMessagesSetter is an interface for setting showAlert option.
type canPostMessagesSetter interface {
	method.PromoteChatMemberOption
}

type canPostMessagesOption struct {
	flag bool
}

func (csm *canPostMessagesOption) ApplyPromoteChatMemberOption(acq *method.PromoteChatMember) {
	acq.CanPostMessages = csm.flag
}

// canEditMessagesSetter is an interface for setting showAlert option.
type canEditMessagesSetter interface {
	method.PromoteChatMemberOption
}

type canEditMessagesOption struct {
	flag bool
}

func (csm *canEditMessagesOption) ApplyPromoteChatMemberOption(acq *method.PromoteChatMember) {
	acq.CanEditMessages = csm.flag
}

// canDeleteMessagesSetter is an interface for setting showAlert option.
type canDeleteMessagesSetter interface {
	method.PromoteChatMemberOption
}

type canDeleteMessagesOption struct {
	flag bool
}

func (csm *canDeleteMessagesOption) ApplyPromoteChatMemberOption(acq *method.PromoteChatMember) {
	acq.CanDeleteMessages = csm.flag
}

// canInviteUsersSetter is an interface for setting showAlert option.
type canInviteUsersSetter interface {
	method.PromoteChatMemberOption
}

type canInviteUsersOption struct {
	flag bool
}

func (csm *canInviteUsersOption) ApplyPromoteChatMemberOption(acq *method.PromoteChatMember) {
	acq.CanInviteUsers = csm.flag
}

// canRestrictMembersSetter is an interface for setting showAlert option.
type canRestrictMembersSetter interface {
	method.PromoteChatMemberOption
}

type canRestrictMembersOption struct {
	flag bool
}

func (csm *canRestrictMembersOption) ApplyPromoteChatMemberOption(acq *method.PromoteChatMember) {
	acq.CanRestrictMembers = csm.flag
}

// canPinMessagesSetter is an interface for setting showAlert option.
type canPinMessagesSetter interface {
	method.PromoteChatMemberOption
}

type canPinMessagesOption struct {
	flag bool
}

func (csm *canPinMessagesOption) ApplyPromoteChatMemberOption(acq *method.PromoteChatMember) {
	acq.CanPinMessages = csm.flag
}

// canPromoteMembersSetter is an interface for setting showAlert option.
type canPromoteMembersSetter interface {
	method.PromoteChatMemberOption
}

type canPromoteMembersOption struct {
	flag bool
}

func (csm *canPromoteMembersOption) ApplyPromoteChatMemberOption(acq *method.PromoteChatMember) {
	acq.CanPromoteMembers = csm.flag
}
