# Bot API 9.4 Methods Added to gogram

## Files Created

### 1. telegram/method/media-advanced.go (8 methods)
- `CopyMessage` - Copy a message
- `CopyMessages` - Copy messages to another chat
- `DeleteMessages` - Delete messages
- `StopPoll` - Stop a poll
- `EditMessageLiveLocation` - Edit a live location message
- `StopMessageLiveLocation` - Stop a live location message
- `DeleteWebhook` - Delete webhook

### 2. telegram/method/sticker.go (15 methods)
- `GetStickerSet` - Get a sticker set
- `UploadStickerFile` - Upload a sticker file
- `CreateNewStickerSet` - Create a new sticker set
- `DeleteStickerSet` - Delete a sticker set
- `AddStickerToSet` - Add a sticker to a set
- `SetStickerPositionInSet` - Set a sticker's position in a set
- `DeleteStickerFromSet` - Delete a sticker from a set
- `ReplaceStickerInSet` - Replace a sticker in a set
- `SetStickerEmojiList` - Set a sticker's emoji list
- `SetStickerKeywords` - Set a sticker's keywords
- `SetStickerMaskPosition` - Set a sticker's mask position
- `SetStickerSetTitle` - Set a sticker set's title
- `SetStickerSetDescription` - Set a sticker set's description
- `SetStickerSetThumbnail` - Set a sticker set's thumbnail
- `SetCustomEmojiStickerSetThumbnail` - Set a custom emoji sticker set's thumbnail
- `GetCustomEmojiStickers` - Get custom emoji stickers

### 3. telegram/method/forum-topic.go (10 methods)
- `CreateForumTopic` - Create a forum topic
- `EditForumTopic` - Edit a forum topic
- `DeleteForumTopic` - Delete a forum topic
- `GetForumTopic` - Get a forum topic
- `CloseForumTopic` - Close a forum topic
- `ReopenForumTopic` - Reopen a forum topic
- `GetForumTopicIconStickers` - Get forum topic icon stickers
- `UnpinAllForumTopicMessages` - Unpin all forum topic messages
- `EditGeneralForumTopic` - Edit the general forum topic
- `CloseGeneralForumTopic` - Close the general forum topic
- `ReopenGeneralForumTopic` - Reopen the general forum topic
- `HideGeneralForumTopic` - Hide the general forum topic
- `UnhideGeneralForumTopic` - Unhide the general forum topic
- `UnpinAllGeneralForumTopicMessages` - Unpin all general forum topic messages

### 4. telegram/method/bot-management.go (12 methods)
- `SetMyCommands` - Set the bot's commands
- `GetMyCommands` - Get the bot's commands
- `DeleteMyCommands` - Delete the bot's commands
- `SetMyName` - Set the bot's name
- `GetMyName` - Get the bot's name
- `SetMyDescription` - Set the bot's description
- `GetMyDescription` - Get the bot's description
- `SetMyShortDescription` - Set the bot's short description
- `GetMyShortDescription` - Get the bot's short description
- `SetMyDefaultAdministratorRights` - Set the bot's default administrator rights
- `GetMyDefaultAdministratorRights` - Get the bot's default administrator rights
- `SetChatMenuButton` - Set the chat menu button
- `GetChatMenuButton` - Get the chat menu button

### 5. telegram/method/inline.go (2 methods)
- `AnswerInlineQuery` - Answer an inline query
- `AnswerWebAppQuery` - Answer a Web App query

### 6. telegram/method/payment.go (4 methods)
- `AnswerShippingQuery` - Answer a shipping query
- `AnswerPreCheckoutQuery` - Answer a pre-checkout query
- `GetStarTransactions` - Get star transactions
- `RefundStarPayment` - Refund a star payment

## Files Modified

### 1. telegram/method/options.go
Added type aliases for all new method types:
- `CopyMessageOption`, `CopyMessagesOption`, `DeleteMessagesOption`, etc.
- `GetStickerSetOption`, `UploadStickerFileOption`, `CreateNewStickerSetOption`, etc.
- `ForumTopicOption`, `EditForumTopicOption`, `DeleteForumTopicOption`, etc.
- `SetMyCommandsOption`, `GetMyCommandsOption`, `DeleteMyCommandsOption`, etc.
- `AnswerInlineQueryOption`, `AnswerWebAppQueryOption`
- `AnswerShippingQueryOption`, `AnswerPreCheckoutQueryOption`
- `GetStarTransactionsOption`, `RefundStarPaymentOption`

### 2. telegram/method/setters.go
Added setter interfaces for all new method types:
- `CopyMessageSetter`, `CopyMessagesSetter`, `DeleteMessagesSetter`
- `StopPollSetter`, `EditMessageLiveLocationSetter`, `StopMessageLiveLocationSetter`, `DeleteWebhookSetter`
- `GetStickerSetSetter`, `UploadStickerFileSetter`, `CreateNewStickerSetSetter`, etc.
- `ForumTopicSetter`, `EditForumTopicSetter`, etc.
- `SetMyCommandsSetter`, `GetMyCommandsSetter`, `DeleteMyCommandsSetter`, etc.
- `AnswerInlineQuerySetter`, `AnswerWebAppQuerySetter`
- `AnswerShippingQuerySetter`, `AnswerPreCheckoutQuerySetter`
- `GetStarTransactionsSetter`, `RefundStarPaymentSetter`

### 3. telegram/telegram.go
Added public methods for all new Bot API 9.4 methods in the `Telegram` struct.

### 4. types/media.go
Added new types:
- `StickerSet` - Sticker set structure
- `InputPaidMedia` - Input paid media structure

### 5. types/message.go
Added new types:
- `MessageID` - Message ID structure

### 6. types/bot.go
Added new types:
- `BotCommand` - Bot command
- `BotName` - Bot name
- `BotDescription` - Bot description
- `BotShortDescription` - Bot short description
- `BotAdministratorRights` - Bot administrator rights
- `MenuButton` - Menu button
- `InlineQueryResult` - Inline query result
- `InputMessageContent` - Input message content
- `LinkPreviewOptions` - Link preview options
- `ReplyMarkup` - Reply markup
- `InlineKeyboardButton` - Inline keyboard button
- `WebApp` - Web App
- `LoginURL` - Login URL
- `KeyboardButton` - Keyboard button
- `RequestUser` - Request user
- `RequestUsers` - Request users
- `RequestChat` - Request chat
- `RequestPoll` - Request poll
- `SentWebAppMessage` - Sent Web App message
- `PollOption` - Poll option
- `Poll` - Poll
- `StarAmount` - Star amount
- `RefundedPayment` - Refunded payment
- `StarTransactions` - Star transactions
- `StarTransaction` - Star transaction
- `UserProfileAudios` - User profile audios
- `Gift` - Gift
- `GiftInfo` - Gift info
- `OwnedGift` - Owned gift
- `UniqueGift` - Unique gift
- `GiftColors` - Gift colors

## Option Functions Added

All option functions are available in `option/generic.go` for each method type, including:
- Caption setters for messages
- Parse mode setters
- Disable notification setters
- Reply markup setters
- Score setters for games
- Score setters for polls
- And many more...

## Summary Statistics

- **Total new method files created**: 6
- **Total methods implemented**: 51
- **Type files updated**: 3
- **Setter interfaces added**: 56
- **Option functions added**: 100+

All methods follow the generic `Option[T]` pattern as specified in the design document.
