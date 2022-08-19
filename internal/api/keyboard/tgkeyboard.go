package keyboard

import (
	"encoding/json"
	"fmt"
	"log"
)

// reply
type ReplyMarkupT struct {
	Keyboard []ReplyRowT `json:"keyboard"`
}
type ReplyRowT []ReplyButtonT
type ReplyButtonT string

// Reply keyboard

func NewKeyboard() ReplyMarkupT {
	return ReplyMarkupT{}
}

func (k *ReplyMarkupT) NewRow(buttons ...ReplyButtonT) {
	row := ReplyRowT{}
	for _, button := range buttons {
		row = append(row, button)
	}
	k.Keyboard = append(k.Keyboard, row)
}

func (k *ReplyMarkupT) View() string {
	kb, err := json.Marshal(k)
	if err != nil {
		log.Fatalf("Error converting to string: %v", err)
		return ""
	}

	return string(kb)
}

func ReplyButton(text interface{}) ReplyButtonT {
	switch t := text.(type) {
	case int, int8, int16, int32, int64, uint8, uint16, uint32, uint64:
		return ReplyButtonT(fmt.Sprintf("%d", t))
	case float32, float64:
		return ReplyButtonT(fmt.Sprintf("%.2f", t))
	case string:
		return ReplyButtonT(t)
	}

	log.Fatalf("Undefined symbol: %v", text)
	return ReplyButtonT("err")
}

// Inline keyboard
type InlineMarkupT struct {
	InlineKeyboard []InlineRowT `json:"inline_keyboard"`
}
type InlineRowT []InlineButtonT
type InlineButtonT struct {
	Text         string  `json:"text"`
	Url          string  `json:"url,omitempty"`
	CallbackData string  `json:"callback_data,omitempty"`
	WebApp       WebAppT `json:"web_app,omitempty"`
}
type WebAppT struct {
	Url string `json:"url"`
}

func NewInlineKeyboard() InlineMarkupT {
	return InlineMarkupT{}
}

func (k *InlineMarkupT) NewRow(buttons ...InlineButtonT) {
	row := InlineRowT{}
	for _, button := range buttons {
		row = append(row, button)
	}
	k.InlineKeyboard = append(k.InlineKeyboard, row)
}

func (k *InlineMarkupT) View() string {
	kb, err := json.Marshal(k)
	if err != nil {
		log.Fatalf("Error converting to string: %v", err)
		return ""
	}

	return string(kb)
}

func WebAppButton(text string, url string) InlineButtonT {
	return InlineButtonT{
		Text: text,
		WebApp: WebAppT{
			Url: url,
		},
	}
}
