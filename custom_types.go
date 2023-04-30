package main

type GetMe struct {
	// True, if bot or user has been founded
	Ok bool `json:"ok"`

	// struct User
	Result *User `json:"result"`
}
