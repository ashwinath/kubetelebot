package main

import (
	"flag"
	"log"

	"github.com/ashwinath/kubetelebot/pkg/telegram"
)

func main() {
	telegramApiKey := flag.String("apiKey", "", "Telegram API key, can be created via BotFather.")
	isDebug := flag.Bool("debug", false, "Debug mode")
	allowedUser := flag.String("allowedUser", "", "Telegram user of the allowed user.")
	flag.Parse()

	tg, err := telegram.New(*telegramApiKey, *isDebug, *allowedUser)
	if err != nil {
		log.Fatalf("error connecting to telegram: %v", err)
	}

	tg.Run()
}
