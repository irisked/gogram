package telegram

import (
	"context"
	"os"

	"github.com/irisked/gogram/telegram/method"
	"github.com/irisked/gogram/types"
)

// GetStickerSet performs getStickerSet telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getstickerset
func (t *Client) GetStickerSet(ctx context.Context, name string) (types.StickerSet, error) {
	m := method.NewGetStickerSet(name)
	var result types.StickerSet
	err := t.process(ctx, m, &result)
	return result, err
}

// UploadStickerFile performs uploadStickerFile telegram api method.
//
// options:
//
// option.UploadStickerFileFormat(), option.UploadStickerFileType(),
// option.UploadStickerFilePNGSticker(), option.UploadStickerFileTSVGSticker(),
// option.UploadStickerFileWebmSticker(), option.UploadStickerFileThumbnail(),
// option.UploadStickerFileMaskPosition(), option.UploadStickerFileEmojiList()
//
// Reference: https://core.telegram.org/bots/api#uploadstickerfile
func (t *Client) UploadStickerFile(ctx context.Context, userID int64, sticker *os.File, opts ...method.Option[*method.UploadStickerFile]) (types.File, error) {
	m := method.NewUploadStickerFile(sticker, opts...)
	file := types.File{}
	err := t.process(ctx, m, &file)
	return file, err
}

// CreateNewStickerSet performs createNewStickerSet telegram api method.
//
// options:
//
// option.CreateNewStickerSetStickerType(), option.CreateNewStickerSetMaskPosition(),
// option.CreateNewStickerSetNeedsRepaired()
//
// Reference: https://core.telegram.org/bots/api#createnewstickerset
func (t *Client) CreateNewStickerSet(ctx context.Context, userID int64, name, title string, stickers []string, opts ...method.Option[*method.CreateNewStickerSet]) (bool, error) {
	m := method.NewCreateNewStickerSet(userID, name, title, stickers, opts...)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// DeleteStickerSet performs deleteStickerSet telegram api method.
//
// Reference: https://core.telegram.org/bots/api#deletestickerset
func (t *Client) DeleteStickerSet(ctx context.Context, name string) (bool, error) {
	m := method.NewDeleteStickerSet(name)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// AddStickerToSet performs addStickerToSet telegram api method.
//
// options:
//
// option.AddStickerToSetMaskPosition()
//
// Reference: https://core.telegram.org/bots/api#addstickertoset
func (t *Client) AddStickerToSet(ctx context.Context, userID int64, name, sticker string, opts ...method.Option[*method.AddStickerToSet]) (bool, error) {
	m := method.NewAddStickerToSet(userID, name, sticker, opts...)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// SetStickerPositionInSet performs setStickerPositionInSet telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setstickerpositioninset
func (t *Client) SetStickerPositionInSet(ctx context.Context, sticker string, position int) (bool, error) {
	m := method.NewSetStickerPositionInSet(sticker, position)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// DeleteStickerFromSet performs deleteStickerFromSet telegram api method.
//
// Reference: https://core.telegram.org/bots/api#deletestickerfromset
func (t *Client) DeleteStickerFromSet(ctx context.Context, sticker string) (bool, error) {
	m := method.NewDeleteStickerFromSet(sticker)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// ReplaceStickerInSet performs replaceStickerInSet telegram api method.
//
// Reference: https://core.telegram.org/bots/api#replacestickerinset
func (t *Client) ReplaceStickerInSet(ctx context.Context, userID int64, name, oldSticker, sticker string) (bool, error) {
	m := method.NewReplaceStickerInSet(userID, name, oldSticker, sticker)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// SetStickerEmojiList performs setStickerEmojiList telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setstickeremojilist
func (t *Client) SetStickerEmojiList(ctx context.Context, sticker string, emojiList []string) (bool, error) {
	m := method.NewSetStickerEmojiList(sticker, emojiList)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// SetStickerKeywords performs setStickerKeywords telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setstickerkeywords
func (t *Client) SetStickerKeywords(ctx context.Context, sticker string, keywords []string) (bool, error) {
	m := method.NewSetStickerKeywords(sticker, keywords)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// SetStickerMaskPosition performs setStickerMaskPosition telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setstickermaskposition
func (t *Client) SetStickerMaskPosition(ctx context.Context, sticker, maskPosition string) (bool, error) {
	m := method.NewSetStickerMaskPosition(sticker, maskPosition)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// SetStickerSetTitle performs setStickerSetTitle telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setstickersettitle
func (t *Client) SetStickerSetTitle(ctx context.Context, name, title string) (bool, error) {
	m := method.NewSetStickerSetTitle(name, title)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// SetStickerSetDescription performs setStickerSetDescription telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setstickersetdescription
func (t *Client) SetStickerSetDescription(ctx context.Context, name, description string) (bool, error) {
	m := method.NewSetStickerSetDescription(name, description)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// SetStickerSetThumbnail performs setStickerSetThumbnail telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setstickersetthumbnail
func (t *Client) SetStickerSetThumbnail(ctx context.Context, name, thumbnail string) (bool, error) {
	m := method.NewSetStickerSetThumbnail(name, thumbnail)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// SetCustomEmojiStickerSetThumbnail performs setCustomEmojiStickerSetThumbnail telegram api method.
//
// Reference: https://core.telegram.org/bots/api#setcustomemojistickersetthumbnail
func (t *Client) SetCustomEmojiStickerSetThumbnail(ctx context.Context, name, customEmoji string) (bool, error) {
	m := method.NewSetCustomEmojiStickerSetThumbnail(name, customEmoji)
	var result bool
	err := t.process(ctx, m, &result)
	return result, err
}

// GetCustomEmojiStickers performs getCustomEmojiStickers telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getcustomemojistickers
func (t *Client) GetCustomEmojiStickers(ctx context.Context, customEmojiIDs []string) ([]types.Sticker, error) {
	m := method.NewGetCustomEmojiStickers(customEmojiIDs)
	var result []types.Sticker
	err := t.process(ctx, m, &result)
	return result, err
}
