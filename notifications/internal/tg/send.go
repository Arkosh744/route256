package tg

import (
	"context"
	api "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"route256/notifications/internal/models"
)

func (b *Bot) SendMessage(ctx context.Context, msg models.TelegramMessage) error {
	tgMsg := api.NewMessage(msg.ChatID, msg.Text)

	if _, err := b.api.Send(tgMsg); err != nil {
		return err
	}

	if err := b.addToCache(ctx, msg); err != nil {
		return err
	}

	log.Printf("sent to chat %d message: %s", msg.ChatID, msg.Text)

	return nil
}
