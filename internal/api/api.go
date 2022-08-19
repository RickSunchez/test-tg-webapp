package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	tgtypes "telegram/internal/api/types"
)

const (
	baseUrl     = "https://api.telegram.org/bot"
	getUpdates  = "getUpdates"
	getMe       = "getMe"
	sendMessage = "sendMessage"
)

type Bot struct {
	Token  string
	Url    string
	Offset int
}

func NewBot(token string) Bot {
	return Bot{
		Token:  token,
		Url:    baseUrl + token + "/",
		Offset: 0,
	}
}

func (b Bot) GetMe() (tgtypes.GetMeT, error) {
	url := urlConstructor(b.Url, getMe, nil)

	resp, err := request(url)
	if err != nil {
		return tgtypes.GetMeT{}, err
	}

	answer := tgtypes.GetMeT{}
	err = json.Unmarshal(resp, &answer)
	if err != nil {
		return tgtypes.GetMeT{}, err
	}

	return answer, nil
}

func (b Bot) GetUpdates() (tgtypes.GetUpdatesT, error) {
	url := urlConstructor(b.Url, getUpdates,
		&tgtypes.GetBody{
			"offset": strconv.Itoa(b.Offset),
		},
	)

	resp, err := request(url)
	if err != nil {
		return tgtypes.GetUpdatesT{}, nil
	}

	answer := tgtypes.GetUpdatesT{}
	err = json.Unmarshal(resp, &answer)
	if err != nil {
		return tgtypes.GetUpdatesT{}, nil
	}

	return answer, nil
}

func (b Bot) SendMessage(chatId int, text string, replyMarkup *string) error {
	get := tgtypes.GetBody{
		"chat_id": strconv.Itoa(chatId),
		"text":    text,
	}
	if replyMarkup != nil {
		get["reply_markup"] = (*replyMarkup)
	}

	url := urlConstructor(b.Url, sendMessage, &get)

	fmt.Println(url)
	_, err := request(url)
	if err != nil {
		return err
	}

	return nil
}

func request(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func urlConstructor(path string, method tgtypes.TGMethod, params *tgtypes.GetBody) string {
	url := path + string(method)

	if params != nil {
		url += "?"
		p := []string{}

		for key, value := range *params {
			p = append(p, key+"="+value)
		}

		url += strings.Join(p, "&")
	}

	return url
}
