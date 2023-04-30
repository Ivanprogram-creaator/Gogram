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

func NewBuffer(json_data []byte) {
	panic("unimplemented")
}

type Bot struct {
	Token string
}

func NewBot(Token string) (*Bot, error) {
	bot := Bot{Token}
	if is_bot, err := BotCheck(&bot); !is_bot && err != nil {
		return nil, err
	}
	return &Bot{Token}, nil
}

func BotCheck(bot *Bot) (bool, error) {
	user, err := bot.GetMe()
	if err != nil {
		return false, err
	}
	if !user.Is_Bot {
		return false, fmt.Errorf("token false")
	}
	return true, err
}

func (bot *Bot) GetMe() (*User, error) {
	responseBody, err := DoRequest(bot.Token, "getMe", nil)
	if err != nil {
		return nil, err
	}
	var user GetMe
	err = json.Unmarshal(responseBody, &user)
	if err != nil {
		return nil, err
	}
	if !user.Ok {
		return nil, fmt.Errorf("token false")
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
