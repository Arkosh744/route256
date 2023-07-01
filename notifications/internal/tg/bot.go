package tg

import (
	"log"

	api "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	*api.BotAPI
}

func NewBot(token string) (*Bot, error) {
	bot, err := api.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{bot}, nil
}

func (b *Bot) SendMessage(chatID int64, text string) error {
	msg := api.NewMessage(chatID, text)

	_, err := b.Send(msg)
	if err != nil {
		return err
	}

	log.Printf("sent to chat %d message: %s", chatID, text)

	return nil
}
