# Option Pattern Redesign

## Problem

The current functional option pattern requires writing N `Apply*` methods per option — one for each compatible method type. For example, `ReplyToMessageID` implements 12 separate `ApplySend*Option` methods.

## Proposed Design: Generic Setter Interfaces

Each option field has a corresponding setter interface. Options are generic functions constrained to that interface, accepting any struct that implements it.

### Step 1: Define setter interfaces

```go
type ReplyToMessageIDSetter interface {
    SetReplyToMessageID(int)
}

type DisableNotificationSetter interface {
    SetDisableNotification(bool)
}

type CaptionSetter interface {
    SetCaption(string)
    SetParseMode(string)
}
```

### Step 2: Implement setters on method structs

Use embedded structs to avoid repeating setter implementations:

```go
type SendOptions struct {
    ReplyToMessageID    int
    DisableNotification bool
}

func (o *SendOptions) SetReplyToMessageID(id int)    { o.ReplyToMessageID = id }
func (o *SendOptions) SetDisableNotification(v bool) { o.DisableNotification = v }

type SendMessage struct {
    ChatID int64
    Text   string
    SendOptions
}

type SendPhoto struct {
    ChatID  string
    Photo   string
    Caption string
    SendOptions
}
```

### Step 3: Options are generic functions

```go
type Option[T any] func(T)

func ReplyToMessageID[T ReplyToMessageIDSetter](id int) Option[T] {
    return func(t T) { t.SetReplyToMessageID(id) }
}

func DisableNotification[T DisableNotificationSetter]() Option[T] {
    return func(t T) { t.SetDisableNotification(true) }
}

func SimpleCaption[T CaptionSetter](text string) Option[T] {
    return func(t T) { t.SetCaption(text) }
}

func MarkdownCaption[T CaptionSetter](text string) Option[T] {
    return func(t T) { t.SetCaption(text); t.SetParseMode("markdown") }
}
```

### Step 4: Constructors use a single variadic

```go
func NewSendMessage(chatID int64, text string, opts ...Option[*SendMessage]) *SendMessage {
    m := &SendMessage{ChatID: chatID, Text: text}
    for _, o := range opts {
        o(m)
    }
    return m
}

func NewSendPhoto(chatID int64, photo string, opts ...Option[*SendPhoto]) *SendPhoto {
    m := &SendPhoto{ChatID: chatID, Photo: photo}
    for _, o := range opts {
        o(m)
    }
    return m
}
```

### Usage

```go
method.NewSendMessage(123, "hello",
    option.ReplyToMessageID[*SendMessage](42),
    option.DisableNotification[*SendMessage](),
)

method.NewSendPhoto(123, "photo.jpg",
    option.ReplyToMessageID[*SendPhoto](42),
    option.SimpleCaption[*SendPhoto]("hello"),
)
```

## Comparison

|                                   | Current        | Generic + typecast | Generic setter    |
|-----------------------------------|----------------|--------------------|-------------------|
| N `Apply*` methods per option     | yes            | no                 | **no**            |
| Runtime typecasts                 | no             | yes                | **no**            |
| Compile-time type safety          | yes            | yes                | **yes**           |
| Boilerplate for a new option      | N methods      | 1 func + union     | 1 func + 1 iface  |
| Boilerplate for a new method type | N option impls | N union entries    | **0**             |
| Single `opts ...Option` variadic  | yes            | yes                | **yes**           |

## Tradeoff

Go cannot infer the type parameter for options because `T` only appears in the return type, not the arguments. Callers must write the explicit `[*ConcreteType]` at every option instantiation.

Also note: setter methods must have pointer receivers, so the type argument must always be the pointer type (`[*SendPhoto]`, not `[SendPhoto]`). Passing the value type is a compile error.

## JSON Marshalling

Embedded structs work transparently with `encoding/json`. Fields from the embedded struct are promoted to the top level of the JSON output — no nesting, no extra wrapper key.

```go
type SendOptions struct {
    ReplyToMessageID    int  `json:"reply_to_message_id,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
}

type SendMessage struct {
    ChatID int64  `json:"chat_id"`
    Text   string `json:"text"`
    SendOptions
}
```

```json
{
  "chat_id": 123,
  "text": "hello",
  "reply_to_message_id": 42
}
```

**Gotchas:**

- If the outer struct and an embedded struct have a field with the same `json` tag, the outer struct field wins (shadowing). Unlikely to be an issue since `SendOptions` fields are shared across method types.
- Embed by value (`SendOptions`), not pointer (`*SendOptions`). With a pointer embed, a nil pointer causes the fields to marshal as `null` instead of being omitted.

## API Coverage Gap (Bot API 9.4)

The current implementation covers roughly Bot API 4.x–5.x. Everything below is missing.

### Deprecated / Renamed

| Old name | New name |
|---|---|
| `kickChatMember` | `banChatMember` |
| `getChatMembersCount` | `getChatMemberCount` |

---

### Missing Methods

#### Messaging
- `deleteWebhook`
- `copyMessage`, `copyMessages`
- `forwardMessages` (batch)
- `deleteMessages` (batch)
- `stopPoll`
- `editMessageLiveLocation`, `stopMessageLiveLocation`
- `sendMessageDraft`

#### Media / Content
- `sendSticker`
- `sendPoll`
- `sendDice`
- `sendPaidMedia`
- `sendInvoice`, `createInvoiceLink`, `createSubscriptionInvoice`

#### Stickers (entire category)
- `getStickerSet`, `getStickerSetThumbnail`
- `uploadStickerFile`
- `createNewStickerSet`, `deleteStickerSet`
- `addStickerToSet`, `setStickerPositionInSet`, `deleteStickerFromSet`, `replaceStickerInSet`
- `setStickerEmojiList`, `setStickerKeywords`, `setStickerMaskPosition`
- `setStickerSetTitle`, `setStickerSetDescription`, `setStickerSetThumbnail`
- `setCustomEmojiStickerSetThumbnail`, `getCustomEmojiStickers`
- `convertRegularStickerSetToMask`

#### Inline Mode (entire category)
- `answerInlineQuery`
- `answerWebAppQuery`

#### Payments
- `answerShippingQuery`, `answerPreCheckoutQuery`
- `getStarTransactions`, `refundStarPayment`
- `sendUpdatedPaymentReceipt`

#### Chat Administration
- `unpinAllChatMessages`
- `banChatSenderChat`, `unbanChatSenderChat`
- `setChatAdministratorCustomTitle`
- `setChatPermissions`
- `createChatInviteLink`, `editChatInviteLink`, `revokeChatInviteLink`
- `approveChatJoinRequest`, `declineChatJoinRequest`

#### Forum Topics (entire category)
- `createForumTopic`, `editForumTopic`, `deleteForumTopic`, `getForumTopic`
- `closeForumTopic`, `reopenForumTopic`
- `getForumTopicIconStickers`
- `unpinAllForumTopicMessages`
- `editGeneralForumTopic`, `closeGeneralForumTopic`, `reopenGeneralForumTopic`
- `hideGeneralForumTopic`, `unhideGeneralForumTopic`
- `unpinAllGeneralForumTopicMessages`

#### Bot Self-Management (entire category)
- `setMyCommands`, `getMyCommands`, `deleteMyCommands`
- `setMyName`, `getMyName`
- `setMyDescription`, `getMyDescription`
- `setMyShortDescription`, `getMyShortDescription`
- `setMyDefaultAdministratorRights`, `getMyDefaultAdministratorRights`
- `setChatMenuButton`, `getChatMenuButton`
- `setMyProfilePhoto`, `removeMyProfilePhoto`

#### Boosts / Business
- `getUserChatBoosts`, `editUserChatBoosts`
- `getBusinessConnection`, `deleteBusinessConnection`

#### Gifts / Stories
- `getUserGifts`, `getChatGifts`, `getBusinessAccountGifts`, `getAvailableGifts`
- `getUserProfileAudios`
- `repostStory`
- `approveSuggestedPost`, `declineSuggestedPost`

---

### Missing Types

#### Message Structure Overhaul
- `InaccessibleMessage`, `MaybeInaccessibleMessage` — callback messages can be inaccessible after 48h
- `MessageOrigin`, `MessageOriginUser`, `MessageOriginHiddenUser`, `MessageOriginChat`, `MessageOriginChannel` — replaces flat forward_* fields
- `ExternalReplyInfo`, `TextQuote`, `ReplyParameters` — new reply system
- `LinkPreviewOptions`
- `ChatFullInfo` — `getChat` returns this richer type instead of `Chat`

#### New Content Types
- `Dice`
- `Poll`, `PollOption`, `InputPollOption`, `PollAnswer`
- `Story`, `VideoQuality`
- `PaidMediaInfo`, `PaidMedia`, `PaidMediaPreview`, `PaidMediaPhoto`, `PaidMediaVideo`
- `Checklist`, `ChecklistTask`, `InputChecklist`, `InputChecklistTask`, `ChecklistTasksDone`, `ChecklistTasksAdded`

#### Message Service Event Types
- `ProximityAlertTriggered`, `MessageAutoDeleteTimerChanged`
- `VideoChatScheduled`, `VideoChatStarted`, `VideoChatEnded`, `VideoChatParticipantsInvited`
- `ChatBoostAdded`
- `SharedUser`, `UsersShared`, `ChatShared`
- `WriteAccessAllowed`
- `PaidMessagePriceChanged`, `DirectMessagePriceChanged`
- `SuggestedPostApproved`, `SuggestedPostApprovalFailed`, `SuggestedPostDeclined`, `SuggestedPostPaid`, `SuggestedPostRefunded`
- `SuggestedPostParameters`, `SuggestedPostPrice`, `SuggestedPostInfo`
- `ChatOwnerLeft`, `ChatOwnerChanged`
- `GiveawayCreated`, `Giveaway`, `GiveawayWinners`, `GiveawayCompleted`

#### Chat Backgrounds
- `BackgroundFill`, `BackgroundFillSolid`, `BackgroundFillGradient`, `BackgroundFillFreeformGradient`
- `BackgroundType`, `BackgroundTypeFill`, `BackgroundTypeWallpaper`, `BackgroundTypePattern`, `BackgroundTypeChatTheme`
- `ChatBackground`

#### Forum Topics
- `ForumTopicCreated`, `ForumTopicClosed`, `ForumTopicEdited`, `ForumTopicReopened`
- `GeneralForumTopicHidden`, `GeneralForumTopicUnhidden`

#### Chat Management
- `ChatPermissions`
- `ChatMemberUpdated`, `ChatJoinRequest`
- `ChatBoostUpdated`, `ChatBoostRemoved`
- `ReactionType`
- `DirectMessagesTopic`
- `Birthdate`
- `BusinessIntro`, `BusinessLocation`, `BusinessOpeningHours`

#### Bot Config
- `BotCommand` + `BotCommandScope` variants
- `MenuButton` variants
- `ChatAdministratorRights`

#### Inline Mode
- `InlineQueryResult*` (all variants)
- `InputMessageContent*` variants
- `SavedPreparedInlineMessage`

#### Payments
- `StarAmount`, `RefundedPayment`

#### Gifts
- `Gift`, `GiftInfo`, `GiftBackground`
- `OwnedGiftRegular`
- `UniqueGift`, `UniqueGiftModel`, `UniqueGiftInfo`, `UniqueGiftColors`
- `AcceptedGiftTypes`, `UserRating`, `UserProfileAudios`

## Relevant Go Issues

The explicit `[*ConcreteType]` requirement is a known language limitation, not a design flaw. Tracked upstream:

- [#47868 — spec: Type inference and functional options?](https://github.com/golang/go/issues/47868) — the exact pain point: functional options with generics force explicit type args at every call site.
- [#50285 — Generic should infer type from variable definition](https://github.com/golang/go/issues/50285) (`Proposal-Hold`) — canonical tracker for inferring `T` from return/assignment context.
- [#58650 — spec: a general approach to type inference](https://github.com/golang/go/issues/58650) — active proposal by @griesemer that would resolve both of the above by unifying result types with the expected context type.
- [#73068 — proposal: Allow type inference from return signature](https://github.com/golang/go/issues/73068) — closed as duplicate of #50285; notes TypeScript and Java already handle this.
- [#64197 — proposal: infer parameter type from arguments in call to result function](https://github.com/golang/go/issues/64197) — related: infer `T` from arguments passed to the returned function.
