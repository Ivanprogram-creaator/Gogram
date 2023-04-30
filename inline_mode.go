package main

type InlineQuery struct {
	// Unique identifier for this query
	Id string `json:"id"`

	// Sender
	From *User `json:"from"`

	// Text of the query (up to 256 characters)
	Query string `json:"query"`

	// Offset of the results to be returned, can be controlled by the bot
	Offeset string `json:"offest"`

	// Optional. Type of the chat from which the inline query was sent.
	// Can be either “sender” for a private chat with the inline query sender,
	// “private”, “group”, “supergroup”, or “channel”. The chat type should be
	// always known for requests sent from official clients and most third-party clients,
	// unless the request was sent from a secret chat
	ChatType string `json:"chat_type"`

	// Optional. Sender location, only for bots that request user location
	Location *Location `json:"location"`
}

type InlineQueryResultsButton struct {
	// Label text on the button
	Text string `json:"text"`

	// Optional. Description of the Web App that will be launched when the user presses the button.
	// The Web App will be able to switch back to the inline mode using the method switchInlineQuery inside the Web App.
	WebApp *WebAppInfo `json:"web_app"`

	// Optional. Deep-linking parameter for the /start message sent to the bot when a user presses the button.
	// 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed.

	// Example: An inline bot that sends YouTube videos can ask the user to connect the bot to their YouTube account
	// to adapt search results accordingly. To do this, it displays a 'Connect your YouTube account'
	// button above the results, or even before showing any. The user presses the button, switches to a private chat with
	// the bot and, in doing so, passes a start parameter that instructs the bot to return an OAuth link. Once done,
	// the bot can offer a switch_inline button so that the user can easily return to the chat where they wanted
	// to use the bot's inline capabilities.
	StartParameter string `json:"start_parameter"`
}

type InlineQueryResultArticle struct {
	// Type of the result, must be article
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 Bytes
	Id string `json:"id"`

	// Title of the result
	Title string `json:"title"`

	// Content of the message to be sent
	InputMessageContent *InputMessageContent `json:"input_message_content"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`

	// Optional. URL of the result
	Url string `json:"url"`

	// Optional. Pass True if you don't want the URL to be shown in the message
	HideUrl bool `json:"hide_url"`

	// Optional. Short description of the result
	Description string `json:"description"`

	// Optional. Url of the thumbnail for the result
	ThumbnailUrl string `json:"thumbnail_url"`

	// Optional. Thumbnail width
	ThumbnailWidth int `json:"thumbnail_width"`

	// Optional. Thumbnail height
	ThumbnailHeight int `json:"thumbnail_height"`
}

type InlineQueryResultPhoto struct {
	// Type of the result, must be photo
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`

	// A valid URL of the photo. Photo must be in JPEG format.
	// Photo size must not exceed 5MB
	PhotoUrl string `json:"photo_url"`

	// URL of the thumbnail for the photo
	ThumbnailUrl string `json:"thumbnail_url"`

	// Optional. Width of the photo
	PhotoWidth int `json:"photo_width"`

	// Optional. Height of the photo
	PhotoHeight int `json:"photo_height"`

	// Optional. Title for the result
	Title string `json:"title"`

	// Optional. Short description of the result
	Description string `json:"description"`

	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption, which can be specified
	// instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`

	// Optional. Content of the message to be sent instead of the photo
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

type InlineQueryResultGif struct {
	// Type of the result, must be gif
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`

	// A valid URL for the GIF file. File size must not exceed 1MB
	GifUrl string `json:"gif_url"`

	// Optional. Width of the GIF
	GifWidth int `json:"gif_width"`

	// Optional. Height of the GIF
	GifHeight int `json:"gif_height"`

	// Optional. Duration of the GIF in seconds
	GifDuration int `json:"gif_duration"`

	// URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbnailUrl string `json:"thumbnail_url"`

	// Optional. MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”,
	// or “video/mp4”. Defaults to “image/jpeg”
	ThumbnailMimeType string `json:"thumbnail_mime_type"`

	// Optional. Title for the result
	Title string `json:"title"`

	// Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the caption. See formatting options for more details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`

	// Optional. Content of the message to be sent instead of the GIF animation
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

type InlineQueryResultMpeg4Gif struct {
	// Type of the result, must be mpeg4_gif
	//
	//
	//
}

type InlineQueryResultCachedAudio struct {
	// Type of the result, must be audio
	Type string `json:"type"`

	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`

	// A valid file identifier for the audio file
	AudioFileId string `json:"audio_file_id"`

	// Optional. Caption, 0-1024 characters after entities parsing
	Caption string `json:"caption"`

	// Optional. Mode for parsing entities in the audio caption.
	// See formatting options for more details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	CaptionEntities []MessageEntity

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`

	// Optional. Content of the message to be sent instead of the audio
	InputMessageContent interface{} `json:"input_message_content"`

	// Note: This will only work in Telegram versions released after 9 April, 2016.
	// Older clients will ignore them.
}

type InputTextMessageContent struct {
	// Text of the message to be sent, 1-4096 characters
	MessageText string `json:"message_text"`

	// Optional. Mode for parsing entities in the message text.
	// See formatting options for more details.
	ParseMode string `json:"parse_mode"`

	// Optional. List of special entities that appear in message text,
	// which can be specified instead of parse_mode
	Entities []MessageEntity `json:"entities"`

	// Optional. Disables link previews for links in the sent message
	DisableWebPagePreview bool `json:"disable_web_page_preview"`
}

type InputLocationMessageContent struct {
	// Latitude of the location in degrees
	Latitude float64 `json:"latitude"`

	// Longitude of the location in degrees
	Longitude float64 `json:"longitude"`

	// Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy float64 `json:"horizontal_accuracy"`

	// Optional. Period in seconds for which the location can be updated,
	// should be between 60 and 86400.
	LivePeriod int `json:"live_period"`

	// Optional. For live locations, a direction in which the user is moving,
	// in degrees. Must be between 1 and 360 if specified.
	Heading int `json:"heading"`

	// Optional. For live locations, a maximum distance for proximity alerts
	// about approaching another chat member,
	// in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius int `json:"proximity_alert_radius"`
}

type InputVenueMessageContent struct {
	// Latitude of the venue in degrees
	Latitude float64 `json:"latitude"`

	// Longitude of the venue in degrees
	Longitude float64 `json:"longitude"`

	// Name of the venue
	Title string `json:"title"`

	// Address of the venue
	Address string `json:"address"`

	// Optional. Foursquare identifier of the venue, if known
	FoursquareId string `json:"foursquare_id"`

	// Optional. Foursquare type of the venue, if known.
	// (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type"`

	// Optional. Google Places identifier of the venue
	GooglePlaceId string `json:"google_place_id"`

	// Optional. Google Places type of the venue
	GooglePlaceType string `json:"google_place_type"`
}

type InputContactMessageContent struct {
	// Contact's phone number
	PhoneNumber string `json:"phone_number"`

	// Contact's first name
	FirstName string `json:"first_name"`

	// Optional. Contact's last name
	LastName string `json:"last_name"`

	// Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	VCard string `json:"vcard"`
}

type InputInvoiceMessageContent struct {
	// Product name, 1-32 characters
	Title string `json:"title"`

	// Product description, 1-255 characters
	Description string `json:"description"`

	// Bot-defined invoice payload, 1-128 bytes.
	// This will not be displayed to the user, use for your internal processes.
	Payload string `json:"payload"`

	// Payment provider token, obtained via @BotFather
	ProviderToken string `json:"provider_token"`

	// Three-letter ISO 4217 currency code, see more on currencies
	Currency string `json:"currency"`

	// Price breakdown, a JSON-serialized list of components
	// (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
	Prices []LabeledPrice `json:"prices"`

	// Optional. The maximum accepted amount for tips in the smallest units of the currency
	// (integer, not float/double). For example, for a maximum tip of US$ 1.45
	// pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number
	// of digits past the decimal point for each currency
	// (2 for the majority of currencies). Defaults to 0
	MaxTipAmount int `json:"max_tip_amount"`

	// Optional. A JSON-serialized array of suggested amounts of tip in the
	// smallest units of the currency (integer, not float/double). At most 4 suggested
	// // tip amounts can be specified. The suggested tip amounts must be positive,
	// passed in a strictly increased order and must not exceed max_tip_amount.
	SuggestedTipAmounts []int `json:"suggested_tip_amounts"`

	// Optional. A JSON-serialized object for data about the invoice,
	// which will be shared with the payment provider.
	// A detailed description of the required fields should be provided by the payment provider
	ProviderData string `json:"provider_data"`

	// Optional. URL of the product photo for the invoice.
	// Can be a photo of the goods or a marketing image for a service.
	PhotoUrl string `json:"photo_url"`

	// Optional. Photo size in bytes
	PhotoSize int `json:"photo_size"`

	// Optional. Photo width
	PhotoWidth int `json:"photo_width"`

	// Optional. Photo height
	PhotoHeight int `json:"photo_height"`

	// Optional. Pass True if you require the user's full name to complete the order
	NeedName bool `json:"need_name"`

	// Optional. Pass True if you require the user's phone number to complete the order
	NeedPhoneNumber bool `json:"need_phone_number"`

	// Optional. Pass True if you require the user's email address to complete the order
	NeedEmail bool `json:"need_email"`

	// Optional. Pass True if you require the user's shipping address to complete the order
	NeedShippingAddress bool `json:"need_shipping_address"`

	// Optional. Pass True if the user's phone number should be sent to provider
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider"`

	// Optional. Pass True if the user's email address should be sent to provider
	SendEmailToProvider bool `json:"send_email_to_provider"`

	// Optional. Pass True if the final price depends on the shipping method
	IsFlexible bool `json:"is_flexible"`
}

type ChosenInlineResult struct {
	// The unique identifier for the result that was chosen
	ResultId string `json:"result_id"`

	// The user that chose the result
	From *User `json:"from"`

	// Optional. Sender location, only for bots that require user location
	Location *Location `json:"location"`

	// Optional. Identifier of the sent inline message.
	// Available only if there is an inline keyboard attached to the message.
	// Will be also received in callback queries and can be used to edit the message.
	InlineMessageId string `json:"inline_message_id"`

	// The query that was used to obtain the result
	Query string `json:"query"`

	// Note: It is necessary to enable inline feedback via
	// @BotFather in order to receive these objects in updates.
}

type SentWebAppMessage struct {
	// Optional. Identifier of the sent inline message.
	// Available only if there is an inline keyboard attached to the message.
	InlineMessageId string `json:"inline_message_id"`
}
