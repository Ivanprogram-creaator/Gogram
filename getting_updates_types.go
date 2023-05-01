package main

type Update struct {
	// The update's unique identifier. Update identifiers start from a certain positive number and increase sequentially.
	// This ID becomes especially handy if you're using webhooks, since it allows you to
	// ignore repeated updates or to restore the correct update sequence, should they get out of order.
	// If there are no new updates for at least a week, then identifier of
	// the next update will be chosen randomly instead of sequentially.
	UpdateId int `json:"update_id"`

	// Optional. New incoming message of any kind - text, photo, sticker, etc.
	Message *Message `json:"message"`

	// Optional. New version of a message that is known to the bot and was edited
	EditedMessage *Message `json:"edited_message"`

	// Optional. New incoming channel post of any kind - text, photo, sticker, etc.
	ChannelPost *Message `json:"channel_post"`

	// Optional. New version of a channel post that is known to the bot and was edited
	EditedChannelPost *Message `json:"edited_channel_post"`

	// Optional. New incoming inline query
	InlineQuery *InlineQuery `json:"inline_query"`

	// Optional. The result of an inline query that was chosen by a user and sent to their chat partner.
	// Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`

	// Optional. New incoming callback query
	CallbackQuery *CallbackQuery `json:"callback_query"`

	// Optional. New incoming shipping query. Only for invoices with flexible price
	ShippingQuery *ShippingQuery `json:"shipping_query"`

	// Optional. New incoming pre-checkout query. Contains full information about checkout
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query"`

	// Optional. New poll state. Bots receive only updates about stopped polls and polls, which are sent by the bot
	Poll *Poll `json:"poll"`

	// Optional. A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were sent by the bot itself.
	PollAnswer *PollAnswer `json:"poll_answer"`

	// Optional. The bot's chat member status was updated in a chat. For private chats,
	// this update is received only when the bot is blocked or unblocked by the user.
	MyChatMember *ChatMemberUpdated `json:"my_chat_member"`

	// Optional. A chat member's status was updated in a chat. The bot must be an administrator in the
	// chat and must explicitly specify “chat_member” in the list of allowed_updates to receive these updates.
	ChatMember *ChatMemberUpdated `json:"chat_member"`

	// Optional. A request to join the chat has been sent.
	// The bot must have the can_invite_users administrator right in the chat to receive these updates.
	ChatJoinRequest *ChatJoinRequest `json:"chat_join_request"`
}

type WebhookInfo struct {
	// Webhook URL, may be empty if webhook is not set up
	Url string `json:"url"`

	// True, if a custom certificate was provided for webhook certificate checks
	HasCustomCertificate bool `json:"has_custom_certificate"`

	// Number of updates awaiting delivery
	PendingUpdateCount int `json:"pending_update_count"`

	// Optional. Currently used webhook IP address
	IpAddress string `json:"ip_address"`

	// Optional. Unix time for the most recent error that happened when trying to deliver an update via webhook
	LastErrorDate int `json:"last_error_date"`

	// Optional. Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook
	LastErrorMessage string `json:"last_error_message"`

	// Optional. Unix time of the most recent error that happened when trying to synchronize available updates with Telegram datacenters
	LastSynchronizationErrorDate int `json:"last_synchronization_error_date"`

	// Optional. The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery
	MaxConnections int `json:"max_connections"`

	// Optional. A list of update types the bot is subscribed to. Defaults to all update types except chat_member
	AllowedUpdates []string `json:"allowed_updates"`
}
