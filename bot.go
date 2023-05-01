package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

)

func DoRequest(Token, Method string, data any) (responseBody []byte, err error) {
	json_data, err := json.Marshal(data)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.telegram.org/bot%s/%s", Token, Method), bytes.NewBuffer(json_data))
	fmt.Println(fmt.Sprintf("https://api.telegram.org/bot%s/%s", Token, Method))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	responseBody, err = ioutil.ReadAll(response.Body)
	return
}

type Bot struct {
	Token string
	Debug bool
}

func NewBot(Token string) (*Bot, error) {
	bot := Bot{Token, false}
	_, err := bot.GetMe()
	return &bot, err
}

func (bot *Bot) GetMe() (interface{}, error) {
	responseBody, err := DoRequest(bot.Token, "logOut", nil)
	if err != nil {
		return nil, err
	}
	var user GetMe
	err = json.Unmarshal(responseBody, &user)
	if err != nil {
		return nil, err
	}
	if !user.Ok {
		return nil, fmt.Errorf("error_code: %d, description: %s", user.ErrorCode, user.Description)
	}
	return user.Result, err
}

func main() {
	bot, err := NewBot("6047702549:AAF5jopKJhDc3nHNmKwhlrHms9hbkCnUUjg")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(bot)
	}
}
