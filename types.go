package main

type User struct {
	// Unique identifier for this user or bot.
	// This number may have more than 32 significant bits and some programming languages
	// may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a 64-bit integer or double-precision float
	// type are safe for storing this identifier.
	Id int64 `json:"id"`

	// True, if this user is a bot
	Is_Bot bool `json:"is_bot"`

	// User's or bot's first name
	FirstName string `json:"first_name"`

	// Optional. User's or bot's last name
	LastName string `json:"last_name"`

	// Optional. User's or bot's username
	Username string `json:"username"`

	// Optional. IETF language tag of the user's language
	LanguageCode string `json:"language_code"`

	// Optional. True, if this user is a Telegram Premium user
	IsPremium bool `json:"is_premium"`

	// Optional. True, if this user added the bot to the attachment menu
	AddedToAttachmentMenu bool `json:"added_to_attachment_menu"`

	// Optional. True, if the bot can be invited to groups. Returned only in getMe.
	CanJoinGroups bool `json:"can_join_groups"`

	// Optional. True, if privacy mode is disabled for the bot. Returned only in getMe.
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages"`

	// Optional. True, if the bot supports inline queries. Returned only in getMe.
	SupportsInlineQueries bool `json:"supports_inline_queries"`
}

type Chat struct {
	// Unique identifier for this user or bot.
	// This number may have more than 32 significant bits and some programming languages
	// may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a 64-bit integer or double-precision float
	// type are safe for storing this identifier.
	Id int64 `json:"id"`

	// Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Type string `json:"type"`

	// Optional. Title, for supergroups, channels and group chats
	Title string `json:"title"`

	// Optional. Username, for private chats, supergroups and channels if available
	Username string `json:"username"`

	// Optional. First name of the other party in a private chat
	FirstName string `json:"first_name"`

	// Optional. True, if the supergroup chat is a forum (has topics enabled)
	IsForum bool `json:"is_forum"`

	// Optional. Chat photo. Returned only in getChat.
	Photo *ChatPhoto `json:"photo"`

	// Optional. If non-empty, the list of all active chat usernames;
	// for private chats, supergroups and channels. Returned only in getChat.
	ActiveUsernames []string `json:"active_usernames"`

	// Optional. Custom emoji identifier of emoji status of the other party in a private chat.
	// Returned only in getChat.
	EmojiStatusCustomEmojiId string `json:"emoji_status_custom_emoji_id"`

	// Optional. Bio of the other party in a private chat. Returned only in getChat.
	BIO string `json:"bio"`

	// Optional. True, if privacy settings of the other party in the private chat allows to use
	// tg://user?id=<user_id> links only in chats with the user. Returned only in getChat.
	HasPrivateForwards bool `json:"has_private_forwards"`

	// Optional. True, if the privacy settings of the other party restrict sending
	// voice and video note messages in the private chat. Returned only in getChat.
	HasRestrictedVoiceAndVideoMessages bool `json:"has_restricted_voice_and_video_messages"`

	// Optional. True, if users need to join the supergroup before they can send messages.
	// Returned only in getChat.
	JoinToSendMessages bool `json:"join_to_send_messages"`

	// Optional. True, if all users directly joining the supergroup need to be approved by
	// supergroup administrators. Returned only in getChat.
	JoinByRequest bool `json:"join_by_request"`

	// Optional. Description, for groups, supergroups and channel chats.
	// Returned only in getChat.
	Description string `json:"description"`

	// Optional. Primary invite link, for groups, supergroups and channel chats.
	// Returned only in getChat.
	InviteLink string `json:"invite_link"`

	// 	Optional. The most recent pinned message (by sending date).
	// Returned only in getChat.
	PinnedMessage *Message `json:"pinned_message"`

	// Optional. Default chat member permissions, for groups and supergroups.
	// Returned only in getChat.
	Permissions *ChatPermissions `json:"permissions"`

	// Optional. For supergroups, the minimum allowed delay between consecutive
	// messages sent by each unprivileged user; in seconds. Returned only in getChat.
	SlowModeDelay int `json:"slow_mode_delay"`

	// Optional. The time after which all messages sent to the chat will be automatically
	// deleted; in seconds. Returned only in getChat.
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`

	// Optional. True, if aggressive anti-spam checks are enabled in the supergroup.
	// The field is only available to chat administrators. Returned only in getChat.
	HasAggressiveAntiSpamEnabled bool `json:"has_aggressive_anti_spam_enabled"`

	// Optional. True, if non-administrators can only get the list of
	// bots and administrators in the chat. Returned only in getChat.
	HasHiddenMembers bool `json:"has_hidden_members"`

	// Optional. True, if messages from the chat can't be forwarded to other chats.
	// Returned only in getChat.
	HasProtectedContent bool `json:"has_protected_content"`

	// Optional. For supergroups, name of group sticker set. Returned only in getChat.
	StickerSetName string `json:"sticker_set_name"`

	// Optional. True, if the bot can change the group sticker set. Returned only in getChat.
	CanSetStickerSet bool `json:"can_set_sticker_set"`

	// Optional. Unique identifier for the linked chat, i.e. the discussion group identifier
	// for a channel and vice versa; for supergroups and channel chats.
	// This identifier may be greater than 32 bits and some programming languages may have
	// difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed
	// 64 bit integer or double-precision float type are safe for storing this identifier.
	// Returned only in getChat.
	LinkedChatId int `json:"linked_chat_id"`

	// Optional. For supergroups, the location to which the supergroup is connected.
	// Returned only in getChat.
	Location *ChatLocation `json:"location"`
}

type Message struct {
	// Unique message identifier inside this chat
	MessageId int `json:"message_id"`

	// Optional. Unique identifier of a message thread to which the message belongs;
	// for supergroups only
	MessageThreadId int `json:"message_thread_id"`

	// Optional. Sender of the message; empty for messages sent to channels.
	// For backward compatibility, the field contains a fake sender user in non-channel chats,
	// if the message was sent on behalf of a chat.
	From User `json:"From"`

	// Optional. Sender of the message, sent on behalf of a chat.
	// For example, the channel itself for channel posts,
	// the supergroup itself for messages from anonymous group administrators,
	// the linked channel for messages automatically forwarded to the discussion group.
	// For backward compatibility, the field from contains a fake sender user in non-channel chats,
	// if the message was sent on behalf of a chat.
	SenderChat *Chat `json:"sender_chat"`

	// Date the message was sent in Unix time
	Date int `json:"date"`

	// Conversation the message belongs to
	Chat *Chat `json:"chat"`

	// Optional. For forwarded messages, sender of the original message
	ForwardFrom *User `json:"forward_from"`

	// 	Optional. For messages forwarded from channels or from anonymous administrators,
	// information about the original sender chat
	ForwardFromChat *Chat `json:"forward_from_chat"`

	// Optional. For messages forwarded from channels, identifier of the original message in the channel
	ForwardFromMessageId int `json:"forward_from_message_id"`

	// Optional. For forwarded messages that were originally sent in channels or by an anonymous chat
	// administrator, signature of the message sender if present
	ForwardSignature string `json:"forward_signature"`

	// Optional. Sender's name for messages forwarded from users who disallow adding a link to their
	// account in forwarded messages
	ForwardSenderName string `json:"forward_sender_name"`

	// Optional. For forwarded messages, date the original message was sent in Unix time
	ForwardDate int `json:"forward_date"`

	// Optional. True, if the message is sent to a forum topic
	IsTopicMessage bool `json:"is_topic_message"`

	// Optional. True, if the message is a channel post that was automatically
	// forwarded to the connected discussion group
	IsAutomaticForward bool `json:"is_automatic_forward"`

	// Optional. For replies, the original message. Note that the Message object in this field
	// will not contain further reply_to_message fields even if it itself is a reply.
	ReplyToMessage *Message `json:"reply_to_message"`

	// Optional. Bot through which the message was sent
	VIABot User `json:"via_bot"`

	// Optional. Date the message was last edited in Unix time
	EditDate int `json:"edit_date"`

	// Optional. True, if the message can't be forwarded
	HasProtectedContent bool `json:"has_protected_content"`

	// Optional. The unique identifier of a media message group this message belongs to
	MediaGroupId string `json:"media_group_id"`

	// Optional. Signature of the post author for messages in channels, or the custom
	// title of an anonymous group administrator
	AuthorSignature string `json:"author_signature"`

	// Optional. For text messages, the actual UTF-8 text of the message
	Text string `json:"text"`

	// Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	Entities []MessageEntity `json:"entities"`

	// Optional. Message is an animation, information about the animation. For backward compatibility,
	// when this field is set, the document field will also be set
	Animation *Animation `json:"animation"`

	// Optional. Message is an audio file, information about the file
	Audio *Audio `json:"audio"`

	// Optional. Message is a general file, information about the file
	Document *Document `json:"document"`

	// Optional. Message is a photo, available sizes of the photo
	photo []PhotoSize `json:"photo"`

	// Optional. Message is a sticker, information about the sticker
	Sticker *Sticker `json:"sticker"`

	// Optional. Message is a video, information about the video
	VideoNote *VideoNote `json:"video_note"`

	// Optional. Message is a voice message, information about the file
	Voice *Voice `json:"voice"`

	// Optional. Caption for the animation, audio, document, photo, video or voice
	Caption string `json:"caption"`

	// 	Optional. For messages with a caption, special entities like usernames, URLs,
	// bot commands, etc. that appear in the caption
	CaptionEntities []MessageEntity `json:"caption_entities"`

	// Optional. True, if the message media is covered by a spoiler animation
	HasMediaSpoiler bool `json:"has_media_spoiler"`

	// 	Optional. Message is a shared contact, information about the contact
	Contact *Contact `json:"contact"`

	// Optional. Message is a dice with random value
	Dice *Dice `json:"dice"`

	// Optional. Message is a game, information about the game
	Poll *Poll `json:"poll"`

	// Optional. Message is a venue, information about the venue. For backward compatibility,
	// when this field is set, the location field will also be set
	Venue *Venue `json:"venue"`

	// Optional. Message is a shared location, information about the location
	Location *Location `json:"location"`

	// Optional. New members that were added to the group or supergroup and information about them
	// (the bot itself may be one of these members)
	NewChatMembers []User `json:"new_chat_members"`

	// Optional. A member was removed from the group, information about them (this member may be the bot itself)
	LeftChatMember *User `json:"left_chat_member"`

	// Optional. A chat title was changed to this value
	NewChatTitle string `json:"new_chat_title"`

	// Optional. A chat photo was change to this value
	NewChatPhoto []PhotoSize `json:"new_chat_photo"`

	// Optional. Service message: the chat photo was deleted
	DeleteChatPhoto bool `json:"delete_chat_photo"`

	// Optional. Service message: the group has been created
	GroupChatCreated bool `json:"group_chat_created"`

	// Optional. Service message: the supergroup has been created.
	// This field can't be received in a message coming through updates,
	// because bot can't be a member of a supergroup when it is created.
	// It can only be found in reply_to_message if someone replies to a very
	// first message in a directly created supergroup.
	SupergroupChatCreated bool `json:"supergroup_chat_created"`

	// Optional. Service message: the channel has been created.
	// This field can't be received in a message coming through updates,
	// because bot can't be a member of a channel when it is created.
	// It can only be found in reply_to_message if someone replies to a very first message in a channel.
	ChannelChatCreated bool `json:"channel_chat_created"`

	// Optional. Service message: auto-delete timer settings changed in the chat
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed"`

	// Optional. The group has been migrated to a supergroup with the specified identifier.
	// This number may have more than 32 significant bits and some programming
	// languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type
	// are safe for storing this identifier.
	MigrateToChatId int `json:"migrate_to_chat_id"`

	// 	Optional. The supergroup has been migrated from a group with the specified identifier.
	// This number may have more than 32 significant bits and some programming languages may have
	// difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type
	// are safe for storing this identifier.
	MigrateFromChatId int `json:"migrate_from_chat_id"`

	// Optional. Specified message was pinned.Note that the Message object in this field
	// will not contain further reply_to_message fields even if it is itself a reply.
	PinnedMessage *Message `json:"pinned_message"`

	// Optional. Message is an invoice for a payment, information about the invoice
	Invoice *Invoice `json:"invoice"`

	// Optional. Message is a service message about a successful payment, information about the payment
	SuccessfulPayment *SuccessfulPayment `json:"successful_payment"`

	// Optional. Service message: a user was shared with the bot
	UserShared *UserShared `json:"user_shared"`

	// Optional. Service message: a chat was shared with the bot
	ChatShared *ChatShared `json:"chat_shared"`

	// Optional. The domain name of the website on which the user has logged in
	ConnectedWebsite string `json:"connected_website"`

	// Optional. Service message: the user allowed the bot added to the attachment menu to write messages
	WriteAccessAllowed *WriteAccessAllowed `json:"write_access_allowed"`

	// Optional. Telegram Passport data
	PassportData *PassportData `json:"passport_data"`

	// Optional. Service message. A user in the chat triggered another user's proximity alert
	// while sharing Live Location.
	ProximityAlertTriggered *ProximityAlertTriggered `json:"proximity_alert_triggered"`

	// Optional. Service message: forum topic created
	ForumTopicCreated *ForumTopicCreated `json:"forum_topic_created"`

	// 	Optional. Service message: forum topic edited
	ForumTopicEdited *ForumTopicEdited `json:"forum_topic_edited"`

	// Optional. Service message: forum topic closed
	ForumTopicClosed *ForumTopicClosed `json:"forum_topic_closed"`

	// Optional. Service message: forum topic reopened
	ForumTopicReopened *ForumTopicReopened `json:"forum_topic_reopened"`

	// Optional. Service message: the 'General' forum topic hidden
	GeneralForumTopicHidden *GeneralForumTopicHidden `json:"general_forum_topic_hidden"`

	// Optional. Service message: the 'General' forum topic unhidden
	GeneralForumTopicUnhidden *GeneralForumTopicUnhidden `json:"general_forum_topic_unhidden"`

	// Optional. Service message: video chat scheduled
	VideoChatScheduled *VideoChatScheduled `json:"video_chat_scheduled"`

	// Optional. Service message: video chat started
	VideoChatStarted *VideoChatStarted `json:"video_chat_started"`

	// Optional. Service message: video chat ended
	VideoChatEnded *VideoChatEnded `json:"video_chat_ended"`

	// Optional. Service message: new participants invited to a video chat
	VideoChatParticipantsInvited *VideoChatParticipantsInvited `json:"video_chat_participants_invited"`

	// Optional. Service message: data sent by a Web App
	WebAppData *WebAppData `json:"web_app_data"`

	// Optional. Inline keyboard attached to the message. login_url buttons are represented as ordinary url buttons.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
}

type MessageId struct {
	// Unique message identifier
	MessageId int64 `json:"message_id"`
}

type MessageEntity struct {
	// Type of the entity. Currently, can be “mention” (@username), “hashtag” (#hashtag), “cashtag” ($USD),
	// “bot_command” (/start@jobs_bot), “url” (https://telegram.org), “email” (do-not-reply@telegram.org),
	// “phone_number” (+1-212-555-0123), “bold” (bold text), “italic” (italic text), “underline” (underlined text),
	// “strikethrough” (strikethrough text), “spoiler” (spoiler message), “code” (monowidth string),
	// “pre” (monowidth block), “text_link” (for clickable text URLs), “text_mention” (for users without usernames),
	// “custom_emoji” (for inline custom emoji stickers)
	Type string `json:"type"`

	// Offset in UTF-16 code units to the start of the entity
	Offset int `json:"offset"`

	// Length of the entity in UTF-16 code units
	Length int `json:"length"`

	// Optional. For “text_link” only, URL that will be opened after user taps on the text
	Url string `json:"url"`

	// Optional. For “text_mention” only, the mentioned user
	User *User `json:"user"`

	// Optional. For “pre” only, the programming language of the entity text
	Language string `json:"language"`

	// Optional. For “custom_emoji” only, unique identifier of the custom emoji.
	// Use getCustomEmojiStickers to get full information about the sticker
	CustomEmojiId string `json:"custom_emoji_id"`
}

type PhotoSize struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time
	// and for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`

	// Photo width
	Width int `json:"width"`

	// Photo height
	Height int `json:"height"`

	// Optional. File size in bytes
	FileSize int `json:"file_size"`
}

type Animation struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time
	// and for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`

	// Video width as defined by sender
	Width int `json:"width"`

	// Video height as defined by sender
	Height int `json:"height"`

	// Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`

	// Optional. Animation thumbnail as defined by sender
	Thumbnail *PhotoSize `json:"thumbnail"`

	// Optional. Original animation filename as defined by sender
	FileName string `json:"file_name"`

	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type"`

	// Optional. File size in bytes.It can be bigger than 2^31 and some programming languages
	// may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type
	// are safe for storing this value.
	FileSize int `json:"file_size"`
}

type Audio struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time
	// and for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`

	// Duration of the audio in seconds as defined by sender
	Duration int `json:"duration"`

	// Optional. Performer of the audio as defined by sender or by audio tags
	Performer string `json:"performer"`

	// Optional. Title of the audio as defined by sender or by audio tags
	Title string `json:"title"`

	// Optional. Original filename as defined by sender
	FileName string `json:"file_name"`

	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type"`

	// Optional. File size in bytes. It can be bigger than 2^31 and someprogramming languages
	// may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type
	// are safe for storing this value.
	FileSize int `json:"file_size"`

	// Optional. Thumbnail of the album cover to which the music file belongs
	Thumbnail *PhotoSize `json:"thumbnail"`
}

type Document struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time
	// and for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`

	// Optional. Document thumbnail as defined by sender
	Thumbnail *PhotoSize `json:"thumbnail"`

	// Optional. Original filename as defined by sender
	FileName string `json:"file_name"`

	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type"`

	// Optional. File size in bytes. It can be bigger than 2^31 and someprogramming languages
	// may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type
	// are safe for storing this value.
	FileSize int `json:"file_size"`
}

type Video struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time
	// and for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`

	// Video width as defined by sender
	Width int `json:"width"`

	// Video height as defined by sender
	Height int `json:"height"`

	// Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`

	// Optional. Video thumbnail
	Thumbnail *PhotoSize `json:"thumbnail"`

	// Optional. Original filename as defined by sender
	FileName string `json:"file_name"`

	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type"`

	// Optional. File size in bytes. It can be bigger than 2^31 and someprogramming languages
	// may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type
	// are safe for storing this value.
	FileSize int `json:"file_size"`
}

type VideoNote struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time
	// and for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`

	// Video width and height (diameter of the video message) as defined by sender
	Length int `json:"length"`

	// Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`

	// Optional. Video thumbnail
	Thumbnail *PhotoSize `json:"thumbnail"`

	// Optional. File size in bytes. It can be bigger than 2^31 and someprogramming languages
	// may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type
	// are safe for storing this value.
	FileSize int `json:"file_size"`
}

type Voice struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time
	// and for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`

	// Video width and height (diameter of the video message) as defined by sender
	Length int `json:"length"`

	// Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`

	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type"`

	// Optional. File size in bytes. It can be bigger than 2^31 and someprogramming languages
	// may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type
	// are safe for storing this value.
	FileSize int `json:"file_size"`
}

type Contact struct {
	// Contact's phone number
	PhoneNumber string `json:"phone_number"`

	// Contact's first name
	FirstName string `json:"first_name"`

	// Optional. Contact's last name
	LastName string `json:"last_name"`

	// Optional. Contact's user identifier in Telegram.This number may have more than 32 significant bits
	// and some programming languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a 64-bit integer or double-precision float type
	// are safe for storing this identifier.
	UserId int64 `json:"user_id"`

	// Optional. Additional data about the contact in the form of a vCard
	VCard string `json:"vcard"`
}

type Dice struct {
	// Emoji on which the dice throw animation is based
	Emoji string `json:"emoji"`

	// Value of the dice, 1-6 for “🎲”, “🎯” and “🎳” base emoji,
	// 1-5 for “🏀” and “⚽” base emoji, 1-64 for “🎰” base emoji
	Value int `json:"value"`
}

type PollOption struct {
	// Option text, 1-100 characters
	Text string `json:"text"`

	// Number of users that voted for this option
	VoterCount int `json:"voter_count"`
}

type PollAnswer struct {
	// Unique poll identifier
	PollId string `json:"poll_id"`

	// The user, who changed the answer to the poll
	User *User `json:"user"`

	// 0-based identifiers of answer options, chosen by the user.
	// May be empty if the user retracted their vote.
	OptionIds []int `json:"option_ids"`
}

type Poll struct {
	// Unique poll identifier
	Id string `json:"id"`

	// Poll question, 1-300 characters
	Question string `json:"question"`

	// List of poll options
	Options []PollOption `json:"options"`

	// Total number of users that voted in the poll
	TotalVoterСount int `json:"total_voter_count"`

	// True, if the poll is closed
	IsClosed bool `json:"is_closed"`

	// True, if the poll is anonymous
	IsAnonymous bool `json:"is_anonymous"`

	// Poll type, currently can be “regular” or “quiz”
	Type string `json:"type"`

	// 	True, if the poll allows multiple answers
	AllowsMultipleAnswers bool `json:"allows_multiple_answers"`

	// Optional. 0-based identifier of the correct answer option. Available only for polls in the quiz
	// mode, which are closed, or was sent (not forwarded) by the bot or to the private chat with the bot.
	CorrectOptionId int `json:"correct_option_id"`

	// 	Optional. Text that is shown when a user chooses an incorrect
	// answer or taps on the lamp icon in a quiz-style poll, 0-200 characters
	Explanation string `json:"explanation"`

	// Optional. Special entities like usernames, URLs, bot commands, etc. that appear in the explanation
	ExplanationEntities []MessageEntity `json:"explanation_entities"`

	// Optional. Amount of time in seconds the poll will be active after creation
	OpenPeriod int `json:"open_period"`

	// Optional. Point in time (Unix timestamp) when the poll will be automatically closed
	CloseDate int `json:"close_date"`
}

type Location struct {
	// Longitude as defined by sender
	Longitude float32 `json:"longitude"`

	// Latitude as defined by sender
	Latitude float32 `json:"latitude"`

	// Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy float64 `json:"horizontal_accuracy"`

	// Optional. Time relative to the message sending date, during which the location can be updated;
	// in seconds. For active live locations only.
	LivePeriod int `json:"live_period"`

	// Optional. The direction in which user is moving, in degrees; 1-360. For active live locations only.
	Heading int `json:"heading"`

	// Optional. The maximum distance for proximity alerts about approaching another chat member, in meters.
	// For sent live locations only.
	ProximityAlertRadius int `json:"proximity_alert_radius"`
}

type Venue struct {
	// Venue location. Can't be a live location
	Location *Location `json:"location"`

	// Name of the venue
	Title string `json:"title"`

	// Address of the venue
	Address string `json:"address"`

	// Optional. Foursquare identifier of the venue
	FoursquareId string `json:"foursquare_id"`

	// Optional. Foursquare type of the venue.
	// (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type"`

	// Optional. Google Places identifier of the venue
	GooglePlaceId string `json:"google_place_id"`

	// Optional. Google Places type of the venue
	GooglePlaceType string `json:"google_place_type"`
}

type WebAppData struct {
	// The data. Be aware that a bad client can send arbitrary data in this field.
	Data string `json:"data"`

	// Text of the web_app keyboard button from which the Web App was opened.
	// Be aware that a bad client can send arbitrary data in this field
	ButtonText string `json:"button_text"`
}

type ProximityAlertTriggered struct {
	// User that triggered the alert
	Traveler *User `json:"traveler"`

	// User that set the alert
	Watcher *User `json:"watcher"`

	// The distance between the users
	Distance int `json:"distance"`
}

type MessageAutoDeleteTimerChanged struct {
	// New auto-delete time for messages in the chat; in seconds
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

type ForumTopicCreated struct {
	// Name of the topic
	Name string `json:"name"`

	// Color of the topic icon in RGB format
	IconColor int `json:"icon_color"`

	// Optional. Unique identifier of the custom emoji shown as the topic icon
	IconCustomEmojiId string `json:"icon_custom_emoji_id"`
}

type ForumTopicClosed struct {
	// This object represents a service message about a forum topic closed in the chat.
	// Currently holds no information.
	_ interface{} `json:""`
}

type ForumTopicEdited struct {
	// Optional. New name of the topic, if it was edited
	Name string `json:"name"`

	// Optional. New identifier of the custom emoji shown as the topic icon, if it was edited;
	// an empty string if the icon was removed
	IconCustomEmojiId string `json:"icon_custom_emoji_id"`
}

type ForumTopicReopened struct {
	// This object represents a service message about a forum topic reopened in the chat.
	// Currently holds no information.
	_ interface{} `json:""`
}

type GeneralForumTopicHidden struct {
	// This object represents a service message about General forum topic hidden in the chat.
	// Currently holds no information.
	_ interface{} `json:""`
}

type GeneralForumTopicUnhidden struct {
	// This object represents a service message about General forum topic unhidden in the chat.
	// Currently holds no information.
	_ interface{} `json:""`
}

type UserShared struct {
	// Identifier of the request
	RequestId int `json:"request_id"`

	// Identifier of the shared user. This number may have more than 32 significant bits
	// and some programming languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a 64-bit integer or double-precision float
	// type are safe for storing this identifier. The bot may not have access to the user and
	// could be unable to use this identifier,
	// unless the user is already known to the bot by some other means.
	UserId int64 `json:"user_id"`
}

type ChatShared struct {
	// Identifier of the request
	RequestId int `json:"request_id"`

	// Identifier of the shared chat. This number may have more than 32 significant bits
	// and some programming languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a 64-bit integer or double-precision float
	// type are safe for storing this identifier. The bot may not have access to the user and
	// could be unable to use this identifier,
	// unless the user is already known to the bot by some other means.
	UserId int64 `json:"user_id"`
}

type WriteAccessAllowed struct {
	// Optional. Name of the Web App which was launched from a link
	WebAppName string `json:"web_app_name"`
}

type VideoChatScheduled struct {
	// Point in time (Unix timestamp) when the video chat is supposed
	// to be started by a chat administrator
	StartDate int `json:"start_date"`
}

type VideoChatStarted struct {
	// This object represents a service message about a video chat started in the chat.
	// Currently holds no information.
	_ interface{} `json:""`
}

type VideoChatEnded struct {
	// Video chat duration in seconds
	Duration int `json:"duration"`
}

type VideoChatParticipantsInvited struct {
	// New members that were invited to the video chat
	Users []User `json:"users"`
}

type UserProfilePhotos struct {
	// Total number of profile pictures the target user has
	TotalCount int `json:"total_count"`

	// Requested profile pictures (in up to 4 sizes each)
	Photos [][]PhotoSize `json:"photos"`
}

// The maximum file size to download is 20 MB
type File struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time
	// and for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`

	// Optional. File size in bytes. It can be bigger than 2^31 and some programming languages
	// may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits,
	// so a signed 64-bit integer or double-precision float type
	// are safe for storing this value.
	FileSize int `json:"file_size"`

	// Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
	FilePath string `json:"file_path"`
}

type WebAppInfo struct {
	// An HTTPS URL of a Web App to be opened with additional data as specified in Initializing Web Apps
	Url string `json:"url"`
}

type ReplyKeyboardMarkup struct {
	// Array of button rows, each represented by an Array of KeyboardButton objects
	Keyboard [][]KeyboardButton `json:"KeyboardButton"`

	// Optional. Requests clients to always show the keyboard
	// when the regular keyboard is hidden.
	// Defaults to false, in which case the custom keyboard
	// can be hidden and opened with a keyboard icon.
	IsPersistent bool `json:"is_persistent"`

	// Optional. Requests clients to resize the keyboard vertically for optimal
	// fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false,
	// in which case the custom keyboard is always of the same height as the app's standard keyboard.
	ResizeKeyboard bool `json:"resize_keyboard"`

	// Optional. Requests clients to hide the keyboard as soon as it's been used.
	// The keyboard will still be available, but clients will automatically
	// display the usual letter-keyboard in the chat - the user can press
	// a special button in the input field to see the custom keyboard again. Defaults to false.
	OneTimeKeyboard bool `json:"one_time_keyboard"`

	// Optional. The placeholder to be shown in the input
	// field when the keyboard is active; 1-64 characters
	InputFieldPlaceholder string `json:"input_field_placeholder"`

	// Optional. Use this parameter if you want to show the keyboard to specific users only.
	// Targets: 1) users that are @mentioned in the text of the Message object;
	// 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.

	// Example: A user requests to change the bot's language, bot replies to the request
	// with a keyboard to select the new language. Other users in the group don't see the keyboard.
	Selective bool `json:"selective"`
}

type KeyboardButton struct {
	// Text of the button.
	// If none of the optional fields are used,
	// it will be sent as a message when the button is pressed
	Text string `json:"text"`

	// Optional. If specified, pressing the button will open a list of suitable users.
	// Tapping on any user will send their
	// identifier to the bot in a “user_shared” service message. Available in private chats only.
	RequestUser *KeyboardButtonRequestUser `json:"request_user"`

	// Optional. If specified, pressing the button will open a list of suitable chats.
	// Tapping on a chat will send its
	// identifier to the bot in a “chat_shared” service message. Available in private chats only.
	RequestChat *KeyboardButtonRequestChat `json:"request_chat"`

	// Optional. If True, the user's phone number
	// will be sent as a contact when the button is pressed. Available in private chats only.
	RequestContact bool `json:"request_contact"`

	// Optional. If True, the user's current
	// location will be sent when the button is pressed. Available in private chats only.
	RequestLocation bool `json:"request_location"`

	// Optional. If specified, the user will be asked to create
	// a poll and send it to the bot when the button is pressed. Available in private chats only.
	RequestPoll *KeyboardButtonPollType `json:"request_poll"`

	// Optional. If specified, the described Web App
	// will be launched when the button is pressed. The Web App
	// will be able to send a “web_app_data” service message. Available in private chats only.
	WebApp *WebAppInfo `json:"web_app"`
}

type KeyboardButtonRequestUser struct {
	// Signed 32-bit identifier of the request,
	// which will be received back in the UserShared object. Must be unique within the message
	RequestId int `json:"request_id"`

	// Optional. Pass True to request a bot,
	// pass False to request a regular user.
	// If not specified, no additional restrictions are applied.
	UserIsBot bool `json:"user_is_bot"`

	// Optional. Pass True to request a premium user,
	// pass False to request a non-premium user.
	// If not specified, no additional restrictions are applied.
	UserIsPremium bool `json:"user_is_premium"`
}

type KeyboardButtonRequestChat struct {
	// Signed 32-bit identifier of the request,
	// which will be received back in the UserShared object. Must be unique within the message
	RequestId int `json:"request_id"`

	// Pass True to request a channel chat, pass False to request a group or a supergroup chat.
	ChatIsChannel bool `json:"chat_is_channel"`

	// Optional. Pass True to request a forum supergroup,
	// pass False to request a non-forum chat.
	// If not specified, no additional restrictions are applied.
	ChatIsForum bool `json:"chat_is_forum"`

	// Optional. Pass True to request a supergroup or a channel with a username,
	// pass False to request a chat without a username.
	// If not specified, no additional restrictions are applied.
	ChatHasUsername bool `json:"chat_has_username"`

	// Optional. Pass True to request a chat owned by the user.
	// Otherwise, no additional restrictions are applied.
	ChatIsCreated bool `json:"chat_is_created"`

	// Optional. A JSON-serialized object listing the required
	// administrator rights of the user in the chat.
	// The rights must be a superset of bot_administrator_rights.
	// If not specified, no additional restrictions are applied.
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights"`

	// Optional. A JSON-serialized object listing the required
	// administrator rights of the bot in the chat.
	// The rights must be a subset of user_administrator_rights.
	// If not specified, no additional restrictions are applied.
	BotAdministratorRights *ChatAdministratorRights `json:"bot_administrator_rights"`

	// Optional. Pass True to request a chat with the bot as a member.
	// Otherwise, no additional restrictions are applied.
	BotIsMember bool `json:"bot_is_member"`
}

type KeyboardButtonPollType struct {
	// Optional. If quiz is passed,
	// the user will be allowed to create only polls in the quiz mode.
	// If regular is passed, only regular polls will be allowed.
	// Otherwise, the user will be allowed to create a poll of any type.
	Type string `json:"type"`
}

type ReplyKeyboardRemove struct {
	// Requests clients to remove the custom keyboard
	// (user will not be able to summon this keyboard; if you want to hide
	// the keyboard from sight but keep it accessible,
	// use one_time_keyboard in ReplyKeyboardMarkup)
	RemoveKeyboard bool `json:"remove_keyboard"`

	// Optional. Use this parameter if you want to remove the keyboard for specific users only.
	// Targets: 1) users that are @mentioned in the text of the Message object;
	// 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.

	// Example: A user votes in a poll, bot returns confirmation message in reply to the vote
	// and removes the keyboard for that user, while still showing the keyboard
	// with poll options to users who haven't voted yet.
	Selective bool `json:"selective"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	// Label text on the button
	Text string `json:"text"`

	// Optional. HTTP or tg:// URL to be opened when the button is pressed.
	// Links tg://user?id=<user_id> can be used to mention
	// a user by their ID without using a username,
	// if this is allowed by their privacy settings.
	Url string `json:"url"`

	// Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
	CallbackData string `json:"callback_data"`

	// 	Optional. Description of the Web App that will be launched when the user presses the button.
	// The Web App will be able to send an arbitrary message on behalf of the user
	// using the method answerWebAppQuery.
	// Available only in private chats between a user and the bot.
	WebApp *WebAppInfo `json:"web_app"`

	// Optional. An HTTPS URL used to automatically authorize the user.
	// Can be used as a replacement for the Telegram Login Widget.
	LoginUrl *LoginUrl `json:"login_url"`

	// Optional. If set, pressing the button will prompt the user to select one of their chats,
	// open that chat and insert the bot's username and the specified inline
	// query in the input field. May be empty, in which case just the bot's
	// username will be inserted.

	// Note: This offers an easy way for users to start using your bot in inline mode
	// when they are currently in a private chat with it.
	// Especially useful when combined with switch_pm… actions - in this case the user
	// will be automatically returned to the chat they switched from,
	// skipping the chat selection screen.
	SwitchInlineQuery string `json:"switch_inline_query"`

	// Optional. If set, pressing the button will insert the bot's username
	// and the specified inline query in the current chat's input field.
	// May be empty, in which case only the bot's username will be inserted.

	// This offers a quick way for the user to open your bot in inline mode
	// in the same chat - good for selecting something from multiple options.
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat"`

	// Optional. If set, pressing the button will prompt the user to select one
	// of their chats of the specified type,
	// open that chat and insert the bot's username
	// and the specified inline query in the input field
	SwitchInlineQueryChosenChat *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat"`

	// Optional. Description of the game that will be launched when the user presses the button.

	// NOTE: This type of button must always be the first button in the first row.
	CallbackGame *CallbackGame `json:"callback_game"`

	// Optional. Specify True, to send a Pay button.

	// NOTE: This type of button must always be the first button
	// in the first row and can only be used in invoice messages.
	Pay bool `json:"pay"`
}

type LoginUrl struct {
	// An HTTPS URL to be opened with user authorization data added to the query
	// string when the button is pressed. If the user refuses to provide
	// authorization data, the original URL without information about the user will be opened.
	// The data added is the same as described in Receiving authorization data.

	// NOTE: You must always check the hash of the received data to verify the authentication
	// and the integrity of the data as described in Checking authorization.
	Url string `json:"url"`

	// Optional. New text of the button in forwarded messages.
	ForwardText string `json:"forward_text"`

	// Optional. Username of a bot, which will be used for user authorization.
	// See Setting up a bot for more details. If not specified,
	// the current bot's username will be assumed.
	// The url's domain must be the same as the domain linked with the bot.
	// See Linking your domain to the bot for more details.
	BotUsername string `json:"bot_username"`

	// Optional. Pass True to request the permission for your bot to send messages to the user.
	RequestWriteAccess bool `json:"request_write_access"`
}

type SwitchInlineQueryChosenChat struct {
	// Optional. The default inline query to be inserted in the input field.
	// If left empty, only the bot's username will be inserted
	Query string `json:"query"`

	// Optional. True, if private chats with users can be chosen
	AllowUserChats bool `json:"allow_user_chats"`

	// Optional. True, if private chats with bots can be chosen
	AllowBotChats bool `json:"allow_bot_chats"`

	// Optional. True, if group and supergroup chats can be chosen
	AllowGroupChats bool `json:"allow_group_chats"`

	// Optional. True, if channel chats can be chosen
	AllowChannelChats bool `json:"allow_channel_chats"`
}

type CallbackQuery struct {
	// Unique identifier for this query
	Id string `json:"id"`

	// Sender
	From *User `json:"from"`

	// Optional. Message with the callback button that originated the query.
	// Note that message content and message date will not be available if the message is too old
	Message *Message `json:"message"`

	// Optional. Identifier of the message sent via the bot in inline mode,
	// that originated the query.
	InlineMessageId string `json:"inline_message_id"`

	// Global identifier, uniquely corresponding to the chat to which the message
	// with the callback button was sent. Useful for high scores in games.
	ChatInstance string `json:"chat_instance"`

	// Optional. Data associated with the callback button.
	// Be aware that the message originated the query can
	// contain no callback buttons with this data.
	Data string `json:"data"`

	// Optional. Short name of a Game to be returned, serves as the unique identifier for the game
	GameShortName string `json:"game_short_name"`

	// NOTE: After the user presses a callback button,
	// Telegram clients will display a progress bar until you call answerCallbackQuery.
	// It is, therefore, necessary to react by calling answerCallbackQuery
	// even if no notification to the user is needed
	// (e.g., without specifying any of the optional parameters).
}

type ForceReply struct {
	// Shows reply interface to the user,
	// as if they manually selected the bot's message and tapped 'Reply'
	ForceReply bool `json:"force_reply"`

	// Optional. The placeholder to be shown in the input field when the reply is active;
	// 1-64 characters
	InputFieldPlaceholder string `json:"input_field_placeholder"`

	// Optional. Use this parameter if you want to force reply from specific users only.
	// Targets: 1) users that are @mentioned in the text of the Message object;
	// 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.

	Selective bool `json:"selective"`

	// Example: A poll bot for groups runs in privacy mode
	// (only receives commands, replies to its messages and mentions).
	// There could be two ways to create a new poll:

	// // Explain the user how to send a command with parameters
	// // (e.g. /newpoll question answer1 answer2).
	// // May be appealing for hardcore users but lacks modern day polish.

	// // Guide the user through a step-by-step process.
	// // 'Please send me your question', 'Cool, now let's add the first answer option',
	// // 'Great. Keep adding answer options, then send /done when you're ready'.

	// The last option is definitely more attractive.
	// And if you use ForceReply in your bot's questions,
	// it will receive the user's answers even if it only receives replies,
	// commands and mentions - without any extra work for the user.
}

type ChatPhoto struct {
	// File identifier of small (160x160) chat photo.
	// This file_id can be used only for photo download
	// and only for as long as the photo is not changed.
	SmallFileId string `json:"small_file_id"`

	// Unique file identifier of small (160x160) chat photo,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	SmallFileUniqueId string `json:"small_file_unique_id"`

	// File identifier of big (640x640) chat photo.
	// This file_id can be used only for photo download
	// and only for as long as the photo is not changed.
	BigFileId string `json:"big_file_id"`

	// Unique file identifier of big (640x640) chat photo,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	BigFileUniqueId string `json:"big_file_unique_id"`
}

type ChatInviteLink struct {
	// The invite link. If the link was created by another chat administrator,
	// then the second part of the link will be replaced with “…”.
	InviteLink string `json:"invite_link"`

	// Creator of the link
	Creator *User `json:"creator"`

	// True, if users joining the chat via the link need to be approved
	// by chat administrators
	CreatesJoinRequest bool `json:"creates_join_request"`

	// True, if the link is primary
	IsPrimary bool `json:"is_primary"`

	// True, if the link is revoked
	IsRevoked bool `json:"is_revoked"`

	// Optional. Invite link name
	Name string `json:"name"`

	// Optional. Point in time (Unix timestamp) when the link will
	// expire or has been expired
	ExpireDate int `json:"expire_date"`

	// Optional. The maximum number of users that can be members of the chat
	// simultaneously after joining the chat via this invite link; 1-99999
	MemberLimit int `json:"member_limit"`

	// Optional. Number of pending join requests created using this link
	PendingJoinRequestCount int `json:"pending_join_request_count"`
}

type ChatAdministratorRights struct {
	// True, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous"`

	// True, if the administrator can access the chat event log,
	// chat statistics, message statistics in channels, see channel members,
	// see anonymous administrators in supergroups and ignore slow mode.
	// Implied by any other administrator privilege
	CanManageChat bool `json:"can_manage_chat"`

	// True, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages"`

	// True, if the administrator can manage video chats
	CanManageVideoChats bool `json:"can_manage_video_chats"`

	// True, if the administrator can restrict, ban or unban chat members
	CanRestrictMembers bool `json:"can_restrict_members"`

	// True, if the administrator can add new administrators with a
	// subset of their own privileges or demote administrators that they
	// have promoted, directly or indirectly (promoted by administrators
	// that were appointed by the user)
	CanPromoteMembers bool `json:"can_promote_members"`

	// True, if the user is allowed to change the chat title,
	// photo and other settings
	CanChangeInfo bool `json:"can_change_info"`

	// True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`

	// Optional. True, if the administrator can post in the channel; channels only
	CanPostMessages bool `json:"can_post_messages"`

	// Optional. True, if the administrator can edit messages of other
	// users and can pin messages; channels only
	CanEditMessages bool `json:"can_edit_messages"`

	// Optional. True, if the user is allowed to pin messages; groups and supergroups only
	CanPinMessages bool `json:"can_pin_messages"`

	// Optional. True, if the user is allowed to create, rename, close, and
	// reopen forum topics; supergroups only
	CanManageTopics bool `json:"can_manage_topics"`
}

type ChatMemberOwner struct {
	// The member's status in the chat, always “creator”
	Status string `json:"status"`

	// Information about the user
	User *User `json:"user"`

	// True, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous"`

	// Optional. Custom title for this user
	CustomTitle string `json:"custom_title"`
}

type ChatMemberAdministrator struct {
	// The member's status in the chat, always “administrator”
	Status string `json:"status"`

	// Information about the user
	User *User `json:"user"`

	// True, if the bot is allowed to edit administrator privileges of that user
	CanBeEdited bool `json:"can_be_edited"`

	// True, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous"`

	// True, if the administrator can access the chat event log, chat statistics,
	// message statistics in channels, see channel members, see anonymous
	// administrators in supergroups and ignore slow mode. Implied by any other
	// administrator privilege
	CanManageChat bool `json:"can_manage_chat"`

	// True, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages"`

	// True, if the administrator can manage video chats
	CanManageVideoChats bool `json:"can_manage_video_chats"`

	// True, if the administrator can restrict, ban or unban chat members
	CanRestrictMembers bool `json:"can_restrict_members"`

	// True, if the administrator can add new administrators with a subset
	// of their own privileges or demote administrators that they have promoted,
	// directly or indirectly (promoted by administrators that were appointed by the user)
	CanPromoteMembers bool `json:"can_promote_members"`

	// True, if the user is allowed to change the chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info"`

	// True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`

	// Optional. True, if the administrator can post in the channel; channels only
	CanPostMessages bool `json:"can_post_messages"`

	// Optional. True, if the administrator can edit messages of other users
	// and can pin messages; channels only
	CanEditMessages bool `json:"can_edit_messages"`

	// Optional. True, if the user is allowed to pin messages; groups and supergroups only
	CanPinMessages bool `json:"can_pin_messages"`

	// Optional. True, if the user is allowed to create, rename, close, and
	// reopen forum topics; supergroups only
	CanManageTopics bool `json:"can_manage_topics"`

	// Optional. Custom title for this user
	CustomTitle string `json:"custom_title"`
}

type ChatMemberMember struct {
	// The member's status in the chat, always “member”
	Status string `json:"status"`

	// Information about the user
	User *User `json:"user"`
}

type ChatMemberRestricted struct {
	// The member's status in the chat, always “restricted”
	Status string `json:"status"`

	// Information about the user
	User *User `json:"user"`

	// True, if the user is a member of the chat at the moment of the request
	IsMember bool `json:"is_member"`

	// True, if the user is allowed to send text messages, contacts, invoices, locations and venues
	CanSendMessages bool `json:"can_send_messages"`

	// True, if the user is allowed to send audios
	CanSendAudios bool `json:"can_send_audios"`

	// True, if the user is allowed to send documents
	CanSendDocuments bool `json:"can_send_documents"`

	// True, if the user is allowed to send photos
	CanSendPhotos bool `json:"can_send_photos"`

	// True, if the user is allowed to send videos
	CanSendVideos bool `json:"can_send_videos"`

	// True, if the user is allowed to send video notes
	CanSendVideoNotes bool `json:"can_send_video_notes"`

	// True, if the user is allowed to send voice notes
	CanSendVoiceNotes bool `json:"can_send_voice_notes"`

	// True, if the user is allowed to send polls
	CanSendPolls bool `json:"can_send_polls"`

	// True, if the user is allowed to send animations, games, stickers and use inline bots
	CanSendOtherMessages bool `json:"can_send_other_messages"`

	// True, if the user is allowed to add web page previews to their messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`

	// True, if the user is allowed to change the chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info"`

	// True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`

	// True, if the user is allowed to pin messages
	CanPinMessages bool `json:"can_pin_messages"`

	// True, if the user is allowed to create forum topics
	CanManageTopics bool `json:"can_manage_topics"`

	// Date when restrictions will be lifted for this user; unix time. If 0,
	// then the user is restricted forever
	UntilDate int `json:"until_date"`
}

type ChatMemberLeft struct {
	// The member's status in the chat, always “left”
	Status string `json:"status"`

	// Information about the user
	User *User `json:"user"`
}

type ChatMemberBanned struct {
	// The member's status in the chat, always “kicked”
	Status string `json:"status"`

	// Information about the user
	User *User `json:"user"`

	// Date when restrictions will be lifted for this user; unix time. If 0,
	// then the user is banned forever
	UntilDate int `json:"until_date"`
}

type ChatMemberUpdated struct {
	// Chat the user belongs to
	Chat *Chat `json:"chat"`

	// Performer of the action, which resulted in the change
	From *User `json:"from"`

	// Date the change was done in Unix time
	Date int `json:"date"`

	// Previous information about the chat member
	OldChatMember interface{} `json:"old_chat_member"`

	// New information about the chat member
	NewChatMember interface{} `json:"new_chat_member"`

	// Optional. Chat invite link, which was used by the user to join the chat;
	// for joining by invite link events only.
	InviteLink *ChatInviteLink `json:"invite_link"`

	// Optional. True, if the user joined the chat via a chat folder invite link
	ViaChatFolderInviteLink bool `json:"via_chat_folder_invite_link"`
}

type ChatJoinRequest struct {
	// Chat to which the request was sent
	Chat *Chat `json:"chat"`

	// User that sent the join request
	From *User `json:"from"`

	// Identifier of a private chat with the user who sent the join request.
	// This number may have more than 32 significant bits and some programming
	// languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a 64-bit integer or
	// double-precision float type are safe for storing this identifier.
	// The bot can use this identifier for 24 hours to send messages until
	// the join request is processed, assuming no other administrator contacted the user.
	UserChatId int `json:"user_chat_id"`

	// Date the request was sent in Unix time
	Date int `json:"date"`

	// Optional. Bio of the user.
	Bio string `json:"bio"`

	// Optional. Chat invite link that was used by the user to send the join request
	InviteLink *ChatInviteLink `json:"invite_link`
}

type ChatPermissions struct {
	// Optional. True, if the user is allowed to send text messages, contacts,
	// invoices, locations and venues
	CanSendMessages bool `json:"can_send_messages"`

	// Optional. True, if the user is allowed to send audios
	CanSendAudios bool `json:"can_send_audios"`

	// Optional. True, if the user is allowed to send documents
	CanSendDocuments bool `json:"can_send_documents"`

	// Optional. True, if the user is allowed to send photos
	CanSendPhotos bool `json:"can_send_photos"`

	// Optional. True, if the user is allowed to send videos
	CanSendVideos bool `json:"can_send_videos"`

	// Optional. True, if the user is allowed to send video notes
	CanSendVideoNotes bool `json:"can_send_video_notes"`

	// Optional. True, if the user is allowed to send voice notes
	CanSendVoiceNotes bool `json:"can_send_voice_notes"`

	// Optional. True, if the user is allowed to send polls
	CanSendPolls bool `json:"can_send_polls"`

	// Optional. True, if the user is allowed to send animations, games,
	// stickers and use inline bots
	CanSendOtherMessages bool `json:"can_send_other_messages"`

	// Optional. True, if the user is allowed to add web page previews to their messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`

	// Optional. True, if the user is allowed to change the chat title,
	// photo and other settings. Ignored in public supergroups
	CanChangeInfo bool `json:"can_change_info"`

	// Optional. True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`

	// Optional. True, if the user is allowed to pin messages. Ignored in public supergroups
	CanPinMessages bool `json:"can_pin_messages"`

	// Optional. True, if the user is allowed to create forum topics.
	// If omitted defaults to the value of can_pin_messages
	CanManageTopics bool `json:"can_manage_topics"`
}

type ChatLocation struct {
	// The location to which the supergroup is connected. Can't be a live location.
	Location *Location `json:"locations"`

	// Location address; 1-64 characters, as defined by the chat owner
	Address string `json:"address"`
}

type ForumTopic struct {
	// Unique identifier of the forum topic
	MessageThreadId int `json:"message_thread_id"`

	// Name of the topic
	Name string `json:"name"`

	// Color of the topic icon in RGB format
	IconColor int `json:"icon_color"`

	// Optional. Unique identifier of the custom emoji shown as the topic icon
	IconCustomEmojiId string `json:"icon_custom_emoji_id"`
}

type BotCommand struct {
	// Text of the command; 1-32 characters. Can contain only lowercase English letters, digits and underscores.
	Command string `json:"command"`

	// Description of the command; 1-256 characters.
	Description string `json:"description"`
}

type BotCommandScopeDefault struct {
	// Scope type, must be default
	Type string `json:"type"`
}

type BotCommandScopeAllPrivateChats struct {
	// Scope type, must be all_private_chats
	Type string `json:"type"`
}

type BotCommandScopeAllGroupChats struct {
	// Scope type, must be all_group_chats
	Type string `json:"type"`
}

type BotCommandScopeAllChatAdministrators struct {
	// Scope type, must be all_chat_administrators
	Type string `json:"type"`
}

type BotCommandScopeChat struct {
	// Scope type, must be chat
	Type string `json:"type"`

	// Unique identifier for the target chat or username of the target supergroup
	// (in the format @supergroupusername)
	ChatId interface{} `json:"chat_id"`
}

type BotCommandScopeChatAdministrators struct {
	// Scope type, must be chat
	Type string `json:"type"`

	// Unique identifier for the target chat or username of the target supergroup
	// (in the format @supergroupusername)
	ChatId interface{} `json:"chat_id"`
}

type BotCommandScopeChatMember struct {
	// Scope type, must be chat_member
	Type string `json:"type"`

	// Unique identifier for the target chat or username of the target supergroup
	// (in the format @supergroupusername)
	ChatId interface{} `json:"chat_id"`

	// Unique identifier of the target user
	UserId int64 `json:"user_id"`
}

type BotName struct {
	// The bot's name
	Name string `json:"name"`
}

type BotDescription struct {
	// The bot's description
	Description string `json:"description"`
}

type BotShortDescription struct {
	// The bot's short description
	ShortDescription string `json:"short_description"`
}

type MenuButtonCommands struct {
	// Type of the button, must be commands
	Type string `json:"type"`
}

type MenuButtonWebApp struct {
	// Type of the button, must be web_app
	Type string `json:"type"`

	// Text on the button
	Text string `json:"text"`

	// Description of the Web App that will be launched when the user presses the button.
	// The Web App will be able to send an arbitrary message on behalf
	// of the user using the method answerWebAppQuery
	WebApp *WebAppInfo `json:"web_app"`
}

type MenuButtonDefault struct {
	// Type of the button, must be default
	Type string `json:"type"`
}

type ResponseParameters struct {
	// Optional. The group has been migrated to a supergroup with the specified identifier.
	// This number may have more than 32 significant bits and some programming
	// languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a signed 64-bit
	// integer or double-precision float type are safe for storing this identifier.
	MigrateToChatId int `json:"migrate_to_chat_id"`

	// Optional. In case of exceeding flood control,
	// the number of seconds left to wait before the request can be repeated
	RetryAfter int `json:"retry_after"`
}

type InputMediaPhoto struct {
	// Type of the result, must be video
	Type string `json:"type"`

	// File to send. Pass a file_id to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL for Telegram to get a file from the Internet,
	// or pass “attach://<file_attach_name>”
	// to upload a new one using multipart/form-data under <file_attach_name> name
	Media string `json:"media"`

	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the photo caption.
	// See formatting options for more details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities"`

	// Optional. Pass True if the photo needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler"`
}

type InputMediaVideo struct {
	// Type of the result, must be video
	Type string `json:"type"`

	// File to send. Pass a file_id to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL for Telegram to get a file from the Internet,
	// or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under
	// <file_attach_name> name
	Media string `json:"media"`

	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation
	// for the file is supported server-side. The thumbnail should be in JPEG format
	// and less than 200 kB in size. A thumbnail's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file, so you can pass
	// “attach://<file_attach_name>”
	// if the thumbnail was uploaded using multipart/form-data under <file_attach_name>
	Thumbnail interface{} `json:"thumbnail"`

	// Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the video caption.
	// See formatting options for more details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities"`

	// Optional. Animation width
	Width int `json:"width"`

	// Optional. Animation height
	Height int `json:"height"`

	// Optional. Animation duration in seconds
	Duration int `json:"duration"`

	// Optional. Pass True if the animation needs
	// to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler"`
}

type InputMediaAnimation struct {
	// Type of the result, must be animation
	Type string `json:"type"`

	// File to send. Pass a file_id to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL for Telegram to get a file from the Internet,
	// or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under
	// <file_attach_name> name
	Media string `json:"media"`

	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation
	// for the file is supported server-side. The thumbnail should be in JPEG format
	// and less than 200 kB in size. A thumbnail's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file,
	// so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using
	// multipart/form-data under <file_attach_name>
	Thumbnail interface{} `json:"thumbnail"`

	// Optional. Caption of the animation to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the animation caption.
	// See formatting options for more details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities"`

	// Optional. Animation width
	Width int `json:"width"`

	// Optional. Animation height
	Height int `json:"height"`

	// Optional. Animation duration in seconds
	Duration int `json"duration"`

	// Optional. Pass True if the animation needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler"`
}

type InputMediaAudio struct {
	// Type of the result, must be audio
	Type string `json:"type"`

	// File to send. Pass a file_id to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL for Telegram to get a file from the Internet,
	// or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data
	// under <file_attach_name> name
	Media string `json:"media"`

	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation
	// for the file is supported server-side. The thumbnail should be in JPEG format
	// and less than 200 kB in size. A thumbnail's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file, so you can pass
	// “attach://<file_attach_name>”
	// if the thumbnail was uploaded using multipart/form-data under <file_attach_name>
	Thumbnail interface{} `json:"thumbnail"`

	// Optional. Caption of the audio to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the audio caption.
	// See formatting options for more details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities"`

	// Optional. Duration of the audio in seconds
	Duration int `json"duration"`

	// Optional. Performer of the audio
	Performer string `json:"performer"`

	// Optional. Title of the audio
	Title string `json:"title"`
}

type InputMediaDocument struct {
	// Type of the result, must be document
	Type string `json:"type"`

	// File to send. Pass a file_id to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL for Telegram to get a file from the Internet,
	// or pass “attach://<file_attach_name>”
	// to upload a new one using multipart/form-data under <file_attach_name> name
	Media string `json:"media"`

	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation
	// for the file is supported server-side. The thumbnail should be in JPEG format
	// and less than 200 kB in size. A thumbnail's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file,
	// so you can pass “attach://<file_attach_name>”
	// if the thumbnail was uploaded using multipart/form-data under <file_attach_name>
	Thumbnail interface{} `json:"thumbnail"`

	// Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the document caption.
	// See formatting options for more details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities"`

	// Optional. Disables automatic server-side content type detection
	// for files uploaded using multipart/form-data.
	// Always True, if the document is sent as part of an album.
	DisableContentTypeDetection bool `json:"disable_content_type_detection"`
}

// type SetMyCommands setMyName {
// 	// A JSON-serialized list of bot commands to be set as the list of the bot's commands.
// 	// At most 100 commands can be specified.
// 	Commands [][]BotCommand `json:"commands"`

// 	// A JSON-serialized object, describing scope of users
// 	// for which the commands are relevant.
// 	// Defaults to BotCommandScopeDefault.
// 	Scope *BotCommandScope `json:"scope"`

// 	// A two-letter ISO 639-1 language code. If empty,
// 	// commands will be applied to all users from the given scope,
// 	// for whose language there are no dedicated commands
// 	LanguageCode string `json:"language_code"`
// }

// type DeleteMyCommands struct {
// 	// A JSON-serialized object, describing
// 	// scope of users for which the commands are relevant.
// 	// Defaults to BotCommandScopeDefault.
// 	Scope *BotCommandScope `json:"scope"`

// 	// A two-letter ISO 639-1 language code.
// 	// If empty, commands will be applied to all users from the given scope,
// 	// for whose language there are no dedicated commands
// 	LanguageCode string `json:"language_code"`
// }

// type GetMyCommands struct {
// 	// A JSON-serialized object, describing scope of users.
// 	// Defaults to BotCommandScopeDefault.
// 	Scope *BotCommandScope `json:"scope"`

// 	// A two-letter ISO 639-1 language code or an empty string
// 	LanguageCode string `json:"language_code"`
// }

// type setMyName struct {
// 	// New bot name; 0-64 characters. Pass an empty string to remove
// 	// the dedicated name for the given language.
// 	Name string `json:"name"`

// 	// A two-letter ISO 639-1 language code.
// 	// If empty, the name will be shown to all users for whose language there is no dedicated name.
// 	LanguageCode string `json:"language_code"`
// }

// type GetMyName struct {
// 	LanguageCode string `json:"language_code"`
// }

// type SetMyDescription struct {
// 	// New bot description; 0-512 characters.
// 	// Pass an empty string to remove the dedicated description for the given language.
// 	Description string `json:"description"`

// 	// A two-letter ISO 639-1 language code. If empty, the description
// 	// will be applied to all users for whose language there is no dedicated description.
// 	LanguageCode string `json:"language_code"`
// }

// type GetMyDescription struct {
// 	// A two-letter ISO 639-1 language code or an empty string
// 	LanguageCode string `json:"language_code"`
// }

// type SetMyShortDescription struct {
// 	// New short description for the bot; 0-120 characters.
// 	// Pass an empty string to remove the dedicated short description for the given language.
// 	ShortDescription int `json:"short_description"`

// 	// A two-letter ISO 639-1 language code. If empty, the short description
// 	// will be applied to all users for whose language there is no dedicated short description.
// 	LanguageCode string `json:"language_code"`
// }

// type GetMyShortDescription struct {
// 	// A two-letter ISO 639-1 language code or an empty string
// 	LanguageCode string `json:"language_code"`
// }

// type SetChatMenuButton struct {
// 	// Unique identifier for the target private chat.
// 	// If not specified, default bot's menu button will be changed
// 	ChatId int `json:"chat_id"`

// 	// A JSON-serialized object for the bot's new menu button.
// 	// Defaults to MenuButtonDefault
// 	MenuButton *MenuButton `json:"menu_button"`
// }

// type SetChatMenuButton struct {
// 	// Unique identifier for the target private chat.
// 	// If not specified, default bot's menu button will be changed
// 	ChatId int `json:"chat_id"`

// 	// A JSON-serialized object for the bot's new menu button.
// 	// Defaults to MenuButtonDefault
// 	MenuButton *MenuButton `json:"menu_button"`
// }

// type GetChatMenuButton struct {
// 	// Unique identifier for the target private chat.
// 	// If not specified, default bot's menu button will be returned
// 	ChatId int `json:"chat_id"`
// }

// type SetMyDefaultAdministratorRights struct {
// 	// A JSON-serialized object describing new default administrator rights.
// 	// If not specified, the default administrator rights will be cleared.
// 	Rights *ChatAdministratorRights `json:"rights"`

// 	// Pass True to change the default administrator rights of the bot in channels. Otherwise,
// 	// the default administrator rights of the bot for groups and supergroups will be changed.
// 	ForChannels bool `json:"for_channels"`
// }

// type GetMyDefaultAdministratorRights struct {
// 	// Pass True to get default administrator rights of the bot in channels.
// 	// Otherwise, default administrator rights of the bot for groups and supergroups will be returned.
// 	ForChannels bool `json:"for_channels"`
// }
