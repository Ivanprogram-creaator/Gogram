package main

type Response struct {
	// True, if bot or user has been founded
	Ok bool `json:"ok"`

	// struct User
	Result map[string]interface{} `json:"result"`

	// Error code
	ErrorCode int `json:"error_code"`

	// Error code description
	Description string `json:"description"`
}
