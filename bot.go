package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Creates new bot
func NewBot(Token string) (*Bot, error) {
	bot := Bot{Token, true}
	_, err := bot.GetMe()
	return &bot, err
}

type Bot struct {
	Token string
	Debug bool
}

// This func makes requests
func (bot *Bot) MakeRequest(Method string, data any) (result *Response) {
	json_data, err := json.Marshal(data)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.telegram.org/bot%s/%s", bot.Token, Method), bytes.NewBuffer(json_data))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	if bot.Debug {
		log.Printf("Connecting to %s", fmt.Sprintf("https://api.telegram.org/bot%s/%s\n", bot.Token, Method))
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	responseBody, err := io.ReadAll(response.Body)
	json.Unmarshal(responseBody, &result)
	if result.Ok {
		log.Println("Successfully")
	} else {
		log.Println("Not successfully")
	}
	return
}

// A simple method for testing your bot's authentication token. Requires no parameters. Returns basic information about the bot in form of a User object.
func (bot *Bot) GetMe() (*User, error) {
	response := bot.MakeRequest("getMe", nil)
	if !response.Ok {
		return nil, fmt.Errorf("function GetMe finished with error_code: %d, description: %s", response.ErrorCode, response.Description)
	}
	result, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	var user User
	err = json.Unmarshal(result, &user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

// Use this method to log out from the cloud Bot API server before launching the bot locally. You must log out the bot before running it locally,
// otherwise there is no guarantee that the bot will receive updates. After a successful call, you can immediately log in on a local server,
// but will not be able to log in back to the cloud Bot API server for 10 minutes. Returns nil on success. Requires no parameters.
func (bot *Bot) LogOut() (err error) {
	response := bot.MakeRequest("logQut", nil)
	if !response.Ok {
		return fmt.Errorf("function LogOut finished with error_code: %d, description: %s", response.ErrorCode, response.Description)
	}
	return
}

func (bot *Bot) Close() (err error) {
	response := bot.MakeRequest("close", nil)
	if !response.Ok {
		return fmt.Errorf("function Close finished with error_code: %d, description: %s", response.ErrorCode, response.Description)
	}
	return
}

func main() {
	bot, err := NewBot(token())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(bot)
	}
}
