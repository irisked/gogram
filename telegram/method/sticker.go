package method

import (
	"os"
)

// GetStickerSet contains information about getStickerSet telegram api method.
type GetStickerSet struct {
	Name string `json:"name" form:"name"`
	SimpleMethod
}

// NewGetStickerSet creates GetStickerSet.
func NewGetStickerSet(name string) *GetStickerSet {
	m := new(GetStickerSet)
	m.Name = name
	return m
}

// Method returns telegram api method endpoint.
func (r *GetStickerSet) Method() string {
	return "getStickerSet"
}

// UploadStickerFile contains information about uploadStickerFile telegram api method.
type UploadStickerFile struct {
	Sticker       *os.File `json:"-" form:"sticker"`
	StickerFormat string   `json:"sticker_format" form:"sticker_format"`
	StickerType   string   `json:"sticker_type,omitempty" form:"sticker_type,omitempty"`
	PNGSticker    string   `json:"png_sticker" form:"png_sticker"`
	TSVGSticker   string   `json:"tgs_sticker,omitempty" form:"tgs_sticker,omitempty"`
	WebmSticker   string   `json:"webm_sticker,omitempty" form:"webm_sticker,omitempty"`
	Thumbnail     string   `json:"thumbnail,omitempty" form:"thumbnail,omitempty"`
	MaskPosition  string   `json:"mask_position,omitempty" form:"mask_position,omitempty"`
	EmojiList     []string `json:"emoji_list,omitempty" form:"emoji_list,omitempty"`
	SimpleMethod
}

// NewUploadStickerFile creates UploadStickerFile.
func NewUploadStickerFile(sticker *os.File, opts ...Option[*UploadStickerFile]) *UploadStickerFile {
	m := new(UploadStickerFile)
	m.Sticker = sticker
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *UploadStickerFile) Method() string {
	return "uploadStickerFile"
}

// SetStickerFormat sets sticker_format
func (r *UploadStickerFile) SetStickerFormat(s string) {
	r.StickerFormat = s
}

// SetStickerType sets sticker_type
func (r *UploadStickerFile) SetStickerType(s string) {
	r.StickerType = s
}

// SetPNGSticker sets png_sticker
func (r *UploadStickerFile) SetPNGSticker(p string) {
	r.PNGSticker = p
}

// SetTSVGSticker sets tgs_sticker
func (r *UploadStickerFile) SetTSVGSticker(t string) {
	r.TSVGSticker = t
}

// SetWebmSticker sets webm_sticker
func (r *UploadStickerFile) SetWebmSticker(w string) {
	r.WebmSticker = w
}

// SetThumbnail sets thumbnail
func (r *UploadStickerFile) SetThumbnail(t string) {
	r.Thumbnail = t
}

// SetMaskPosition sets mask_position
func (r *UploadStickerFile) SetMaskPosition(m string) {
	r.MaskPosition = m
}

// SetEmojiList sets emoji_list
func (r *UploadStickerFile) SetEmojiList(e []string) {
	r.EmojiList = e
}

// CreateNewStickerSet contains information about createNewStickerSet telegram api method.
type CreateNewStickerSet struct {
	OwnerID      int64    `json:"owner_user_id" form:"owner_user_id"`
	Name         string   `json:"name" form:"name"`
	Title        string   `json:"title" form:"title"`
	Stickers     []string `json:"stickers" form:"stickers"`
	StickerType  string   `json:"sticker_type,omitempty" form:"sticker_type,omitempty"`
	MaskPosition string   `json:"mask_position,omitempty" form:"mask_position,omitempty"`
	NeedsRepairs *bool    `json:"needs_repairs,omitempty" form:"needs_repairs,omitempty"`
	SimpleMethod
}

// NewCreateNewStickerSet creates CreateNewStickerSet.
func NewCreateNewStickerSet(ownerID int64, name, title string, stickers []string, opts ...Option[*CreateNewStickerSet]) *CreateNewStickerSet {
	m := &CreateNewStickerSet{
		OwnerID:  ownerID,
		Name:     name,
		Title:    title,
		Stickers: stickers,
	}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *CreateNewStickerSet) Method() string {
	return "createNewStickerSet"
}

// SetStickerType sets sticker_type
func (r *CreateNewStickerSet) SetStickerType(s string) {
	r.StickerType = s
}

// SetMaskPosition sets mask_position
func (r *CreateNewStickerSet) SetMaskPosition(m string) {
	r.MaskPosition = m
}

// SetNeedsRepaired sets needs_repairs
func (r *CreateNewStickerSet) SetNeedsRepaired(n *bool) {
	r.NeedsRepairs = n
}

// DeleteStickerSet contains information about deleteStickerSet telegram api method.
type DeleteStickerSet struct {
	Name string `json:"name" form:"name"`
	SimpleMethod
}

// NewDeleteStickerSet creates DeleteStickerSet.
func NewDeleteStickerSet(name string) *DeleteStickerSet {
	m := new(DeleteStickerSet)
	m.Name = name
	return m
}

// Method returns telegram api method endpoint.
func (r *DeleteStickerSet) Method() string {
	return "deleteStickerSet"
}

// AddStickerToSet contains information about addStickerToSet telegram api method.
type AddStickerToSet struct {
	OwnerID      int64  `json:"owner_user_id" form:"owner_user_id"`
	Name         string `json:"name" form:"name"`
	Sticker      string `json:"sticker" form:"sticker"`
	MaskPosition string `json:"mask_position,omitempty" form:"mask_position,omitempty"`
	SimpleMethod
}

// NewAddStickerToSet creates AddStickerToSet.
func NewAddStickerToSet(ownerID int64, name, sticker string, opts ...Option[*AddStickerToSet]) *AddStickerToSet {
	m := &AddStickerToSet{
		OwnerID: ownerID,
		Name:    name,
		Sticker: sticker,
	}
	for _, o := range opts {
		o(m)
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *AddStickerToSet) Method() string {
	return "addStickerToSet"
}

// SetMaskPosition sets mask_position
func (r *AddStickerToSet) SetMaskPosition(m string) {
	r.MaskPosition = m
}

// SetStickerPositionInSet contains information about setStickerPositionInSet telegram api method.
type SetStickerPositionInSet struct {
	Sticker  string `json:"sticker" form:"sticker"`
	Position int    `json:"position" form:"position"`
	SimpleMethod
}

// NewSetStickerPositionInSet creates SetStickerPositionInSet.
func NewSetStickerPositionInSet(sticker string, position int) *SetStickerPositionInSet {
	m := new(SetStickerPositionInSet)
	m.Sticker = sticker
	m.Position = position
	return m
}

// Method returns telegram api method endpoint.
func (r *SetStickerPositionInSet) Method() string {
	return "setStickerPositionInSet"
}

// DeleteStickerFromSet contains information about deleteStickerFromSet telegram api method.
type DeleteStickerFromSet struct {
	Sticker string `json:"sticker" form:"sticker"`
	SimpleMethod
}

// NewDeleteStickerFromSet creates DeleteStickerFromSet.
func NewDeleteStickerFromSet(sticker string) *DeleteStickerFromSet {
	m := new(DeleteStickerFromSet)
	m.Sticker = sticker
	return m
}

// Method returns telegram api method endpoint.
func (r *DeleteStickerFromSet) Method() string {
	return "deleteStickerFromSet"
}

// ReplaceStickerInSet contains information about replaceStickerInSet telegram api method.
type ReplaceStickerInSet struct {
	OwnerID    int64  `json:"owner_user_id" form:"owner_user_id"`
	Name       string `json:"name" form:"name"`
	OldSticker string `json:"old_sticker" form:"old_sticker"`
	Sticker    string `json:"sticker" form:"sticker"`
	SimpleMethod
}

// NewReplaceStickerInSet creates ReplaceStickerInSet.
func NewReplaceStickerInSet(ownerID int64, name, oldSticker, sticker string) *ReplaceStickerInSet {
	m := &ReplaceStickerInSet{
		OwnerID:    ownerID,
		Name:       name,
		OldSticker: oldSticker,
		Sticker:    sticker,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *ReplaceStickerInSet) Method() string {
	return "replaceStickerInSet"
}

// SetStickerEmojiList contains information about setStickerEmojiList telegram api method.
type SetStickerEmojiList struct {
	Sticker   string   `json:"sticker" form:"sticker"`
	EmojiList []string `json:"emoji_list" form:"emoji_list"`
	SimpleMethod
}

// NewSetStickerEmojiList creates SetStickerEmojiList.
func NewSetStickerEmojiList(sticker string, emojiList []string) *SetStickerEmojiList {
	m := &SetStickerEmojiList{
		Sticker:   sticker,
		EmojiList: emojiList,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *SetStickerEmojiList) Method() string {
	return "setStickerEmojiList"
}

// SetStickerKeywords contains information about setStickerKeywords telegram api method.
type SetStickerKeywords struct {
	Sticker  string   `json:"sticker" form:"sticker"`
	Keywords []string `json:"keywords,omitempty" form:"keywords,omitempty"`
	SimpleMethod
}

// NewSetStickerKeywords creates SetStickerKeywords.
func NewSetStickerKeywords(sticker string, keywords []string) *SetStickerKeywords {
	m := &SetStickerKeywords{
		Sticker:  sticker,
		Keywords: keywords,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *SetStickerKeywords) Method() string {
	return "setStickerKeywords"
}

// SetStickerMaskPosition contains information about setStickerMaskPosition telegram api method.
type SetStickerMaskPosition struct {
	Sticker      string `json:"sticker" form:"sticker"`
	MaskPosition string `json:"mask_position,omitempty" form:"mask_position,omitempty"`
	SimpleMethod
}

// NewSetStickerMaskPosition creates SetStickerMaskPosition.
func NewSetStickerMaskPosition(sticker, maskPosition string) *SetStickerMaskPosition {
	m := &SetStickerMaskPosition{
		Sticker:      sticker,
		MaskPosition: maskPosition,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *SetStickerMaskPosition) Method() string {
	return "setStickerMaskPosition"
}

// SetStickerSetTitle contains information about setStickerSetTitle telegram api method.
type SetStickerSetTitle struct {
	Name  string `json:"name" form:"name"`
	Title string `json:"title" form:"title"`
	SimpleMethod
}

// NewSetStickerSetTitle creates SetStickerSetTitle.
func NewSetStickerSetTitle(name, title string) *SetStickerSetTitle {
	m := &SetStickerSetTitle{
		Name:  name,
		Title: title,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *SetStickerSetTitle) Method() string {
	return "setStickerSetTitle"
}

// SetStickerSetDescription contains information about setStickerSetDescription telegram api method.
type SetStickerSetDescription struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description,omitempty" form:"description,omitempty"`
	SimpleMethod
}

// NewSetStickerSetDescription creates SetStickerSetDescription.
func NewSetStickerSetDescription(name, description string) *SetStickerSetDescription {
	m := &SetStickerSetDescription{
		Name:        name,
		Description: description,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *SetStickerSetDescription) Method() string {
	return "setStickerSetDescription"
}

// SetStickerSetThumbnail contains information about setStickerSetThumbnail telegram api method.
type SetStickerSetThumbnail struct {
	Name      string `json:"name" form:"name"`
	Thumbnail string `json:"thumbnail,omitempty" form:"thumbnail,omitempty"`
	SimpleMethod
}

// NewSetStickerSetThumbnail creates SetStickerSetThumbnail.
func NewSetStickerSetThumbnail(name, thumbnail string) *SetStickerSetThumbnail {
	m := &SetStickerSetThumbnail{
		Name:      name,
		Thumbnail: thumbnail,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *SetStickerSetThumbnail) Method() string {
	return "setStickerSetThumbnail"
}

// SetCustomEmojiStickerSetThumbnail contains information about setCustomEmojiStickerSetThumbnail telegram api method.
type SetCustomEmojiStickerSetThumbnail struct {
	Name        string `json:"name" form:"name"`
	CustomEmoji string `json:"custom_emoji_id" form:"custom_emoji_id"`
	SimpleMethod
}

// NewSetCustomEmojiStickerSetThumbnail creates SetCustomEmojiStickerSetThumbnail.
func NewSetCustomEmojiStickerSetThumbnail(name, customEmoji string) *SetCustomEmojiStickerSetThumbnail {
	m := &SetCustomEmojiStickerSetThumbnail{
		Name:        name,
		CustomEmoji: customEmoji,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *SetCustomEmojiStickerSetThumbnail) Method() string {
	return "setCustomEmojiStickerSetThumbnail"
}

// GetCustomEmojiStickers contains information about getCustomEmojiStickers telegram api method.
type GetCustomEmojiStickers struct {
	CustomEmojiIDs []string `json:"custom_emoji_ids" form:"custom_emoji_ids"`
	SimpleMethod
}

// NewGetCustomEmojiStickers creates GetCustomEmojiStickers.
func NewGetCustomEmojiStickers(customEmojiIDs []string) *GetCustomEmojiStickers {
	m := &GetCustomEmojiStickers{
		CustomEmojiIDs: customEmojiIDs,
	}
	return m
}

// Method returns telegram api method endpoint.
func (r *GetCustomEmojiStickers) Method() string {
	return "getCustomEmojiStickers"
}
