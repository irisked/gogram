package telegram

import (
	"context"

	"github.com/irisked/gogram/content"
	"github.com/irisked/gogram/telegram/method"
	"github.com/irisked/gogram/types"
)

// SendPhoto performs sendPhoto telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendphoto
func (t *Client) SendPhoto(ctx context.Context, chatID int64, photo *content.PhotoConfig, opts ...method.Option[*method.SendPhoto]) (types.Message, error) {
	m := method.NewSendPhoto(chatID, photo, opts...)
	message := types.Message{}
	err := t.process(ctx, m, &message)
	return message, err
}

// SendAudio performs sendAudio telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendaudio
func (t *Client) SendAudio(ctx context.Context, chatID int64, audio *content.AudioConfig, opts ...method.Option[*method.SendAudio]) (types.Message, error) {
	m := method.NewSendAudio(chatID, audio, opts...)
	message := types.Message{}
	err := t.process(ctx, m, &message)
	return message, err
}

// SendDocument performs sendDocument telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#senddocument
func (t *Client) SendDocument(ctx context.Context, chatID int64, document *content.DocumentConfig, opts ...method.Option[*method.SendDocument]) (types.Message, error) {
	m := method.NewSendDocument(chatID, document, opts...)
	message := types.Message{}
	err := t.process(ctx, m, &message)
	return message, err
}

// SendVideo performs sendVideo telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendvideo
func (t *Client) SendVideo(ctx context.Context, chatID int64, video *content.VideoConfig, opts ...method.Option[*method.SendVideo]) (types.Message, error) {
	m := method.NewSendVideo(chatID, video, opts...)
	message := types.Message{}
	err := t.process(ctx, m, &message)
	return message, err
}

// SendAnimation performs sendAnimation telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendanimation
func (t *Client) SendAnimation(ctx context.Context, chatID int64, animation *content.AnimationConfig, opts ...method.Option[*method.SendAnimation]) (types.Message, error) {
	m := method.NewSendAnimation(chatID, animation, opts...)
	message := types.Message{}
	err := t.process(ctx, m, &message)
	return message, err
}

// SendVoice performs sendVoice telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendvoice
func (t *Client) SendVoice(ctx context.Context, chatID int64, voice *content.VoiceConfig, opts ...method.Option[*method.SendVoice]) (types.Message, error) {
	m := method.NewSendVoice(chatID, voice, opts...)
	message := types.Message{}
	err := t.process(ctx, m, &message)
	return message, err
}

// SendVideoNote performs sendVideoNote telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#sendvideonote
func (t *Client) SendVideoNote(ctx context.Context, chatID int64, videoNote *content.VideoNoteConfig, opts ...method.Option[*method.SendVideoNote]) (types.Message, error) {
	m := method.NewSendVideoNote(chatID, videoNote, opts...)
	message := types.Message{}
	err := t.process(ctx, m, &message)
	return message, err
}

// SendMediaGroup performs sendMediaGroup telegram api method.
//
// options:
//
// option.DisableWebPagePreview(), option.ReplyToMessageID(id int),
//
// Reference: https://core.telegram.org/bots/api#sendmediagroup
// TODO: parse response to array.
func (t *Client) SendMediaGroup(ctx context.Context, chatID int64, mediaGroup []content.MediaGroupConfig, opts ...method.Option[*method.SendMediaGroup]) ([]types.Message, error) {
	m := method.NewSendMediaGroup(chatID, mediaGroup, opts...)
	var messages []types.Message
	err := t.process(ctx, m, &messages)
	return messages, err
}

// EditMessageMedia performs editMessageMedia telegram api method.
// Returns pointer to the message, if message was sent by the bot or nil-pointer.
//
// options:
//
// option.Keyboard(keyboard markup.Keyboard)
//
// Reference: https://core.telegram.org/bots/api#editmessagemedia
func (t *Client) EditMessageMedia(ctx context.Context, messageID *content.MessageID, media content.InputMediaConfig, opts ...method.Option[*method.EditMessageMedia]) (*types.Message, error) {
	m := method.NewEditMessageMedia(messageID, media, opts...)
	message := types.Message{}
	err := t.process(ctx, m, &message)
	return &message, err
}

// GetFile performs getFile telegram api method.
//
// Reference: https://core.telegram.org/bots/api#getfile
func (t *Client) GetFile(ctx context.Context, fileID string) (types.File, error) {
	m := method.NewGetFile(fileID)
	file := types.File{}
	err := t.process(ctx, m, &file)
	return file, err
}
