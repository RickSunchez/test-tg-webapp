// t.me/DunkerLikeWebApp_bot
// https://shsq.ru/durgerwebapp
// https://habr.com/ru/post/666278/?ysclid=l6ynrep1vy299391856

package main

import (
	"fmt"
	"log"
	"os"
	"telegram/internal/api"
	"telegram/internal/api/keyboard"
)

func main() {
	file, _ := os.Open("config/token.txt")
	data := make([]byte, 46)
	file.Read(data)

	token := string(data)
	bot := api.NewBot(token)

	for {
		answer, err := bot.GetUpdates()
		if err != nil {
			log.Fatal(err)
		}

		if len(answer.Result) == 0 {
			continue
		}

		fmt.Println(answer)
		l := len(answer.Result) - 1
		bot.Offset = answer.Result[l].UpdateID + 1

		kb := keyboard.NewInlineKeyboard()
		kb.NewRow(
			keyboard.WebAppButton("app", "https://shsq.ru/durgerwebapp"),
		)
		view := kb.View()

		for _, result := range answer.Result {
			bot.SendMessage(result.Message.Chat.ID, "hi!", &view)
		}
	}
}
