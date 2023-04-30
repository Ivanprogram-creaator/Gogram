package main

type LabeledPrice struct {
	// Portion label
	Label string `json:"label"`

	// Price of the product in the smallest units of the currency
	// (integer, not float/double). For example, for a price of US$ 1.45
	// pass amount = 145. See the exp parameter in currencies.json,
	// it shows the number of digits past the decimal point for each currency
	// (2 for the majority of currencies).
	Amount int `json:"amount"`
}

type Invoice struct {
	// Product name
	Title string `json:"title"`

	// Product description
	StartParameter string `json:"start_parameter"`

	// Three-letter ISO 4217 currency code
	Currency string `json:"currency"`

	// Total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter
	// in currencies.json, it shows the number
	// of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
}

type ShippingAddress struct {
	// Two-letter ISO 3166-1 alpha-2 country code
	CountryCode string `json:"country_code"`

	// State, if applicable
	State string `json:"state"`

	// City
	City string `json:"city"`

	// First line for the address
	StreetLine1 string `json:"street_line1"`

	// Second line for the address
	StreetLine2 string `json:"street_line2"`

	// Address post code
	PostCode string `json:"post_code"`
}

type OrderInfo struct {
	// Optional. User name
	Name string `json:"name"`

	// Optional. User's phone number
	PhoneNumber string `json:"phone_number"`

	// Optional. User email
	Email string `json:"email"`

	// Optional. User shipping address
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

type ShippingOption struct {
	// Shipping option identifier
	Id string `json:"id"`

	// Option title
	Title string `json:"title"`

	// List of price portions
	Prices string `json:"prices"`
}

type SuccessfulPayment struct {
	// Three-letter ISO 4217 currency code
	Currency string `json:"currency"`

	// Total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter
	// in currencies.json, it shows the number of digits
	// past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`

	// Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`

	// Optional. Identifier of the shipping option chosen by the user
	ShippingOptionId string `json:"shipping_option_id"`

	// Optional. Order information provided by the user
	OrderInfo *OrderInfo `json:"order_info"`

	// Telegram payment identifier
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`

	// Provider payment identifier
	ProviderPaymentChargeId string `json:"provider_payment_charge_id"`
}

type ShippingQuery struct {
	// Unique query identifier
	Id string `json:"id"`

	// User who sent the query
	From *User `json:"from"`

	// Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`

	// User specified shipping address
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

type PreCheckoutQuery struct {
	// Unique query identifier
	Id string `json:"id"`

	// User who sent the query
	From *User `json:"from"`

	// Three-letter ISO 4217 currency code
	Currency string `json:"currency"`

	// Total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter
	// in currencies.json, it shows the number of digits
	// past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`

	// Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`

	// Optional. Identifier of the shipping option chosen by the user
	ShippingOptionId string `json:"shipping_option_id"`

	// Optional. Order information provided by the user
	OrderInfo *OrderInfo `json:"order_info"`
}
