package telegram

import (
	"fmt"
	"log"
	"strings"
	"unicode"

	"github.com/ashwinath/kubetelebot/pkg/shell"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram struct {
	bot         *tgbotapi.BotAPI
	allowedUser string
}

func New(
	apiKey string,
	isDebug bool,
	allowedUser string,
) (*Telegram, error) {
	bot, err := tgbotapi.NewBotAPI(apiKey)

	if err != nil {
		return nil, err
	}

	bot.Debug = isDebug
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &Telegram{
		bot:         bot,
		allowedUser: allowedUser,
	}, nil
}

func (t *Telegram) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := t.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message.From.UserName == t.allowedUser && update.Message != nil { // If we got a message
			log.Printf("[%s to bot] %s", update.Message.From.UserName, update.Message.Text)

			args := strings.Split(firstToLower(update.Message.Text), " ")
			reply, err := shell.RunShell("kubectl", args...)
			if err != nil {
				errString := err.Error()
				reply = &errString
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, wrapInMonospace(reply))
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ParseMode = "Markdown"

			_, err = t.bot.Send(msg)
			if err != nil {
				log.Printf("[bot to %s] error: %s", update.Message.From.UserName, err)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Message was too long to respond.")
				t.bot.Send(msg)
			}
		}
	}
}

func firstToLower(s string) string {
	if len(s) == 0 {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToLower(r[0])

	return string(r)
}

func wrapInMonospace(text *string) string {
	return fmt.Sprintf("```\n%s\n```", *text)
}
