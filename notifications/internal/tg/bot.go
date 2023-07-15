package tg

import (
	"context"
	"route256/notifications/internal/cache"
	"route256/notifications/internal/models"
	"strconv"

	api "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	api   *api.BotAPI
	repo  Repository
	cache RedisCache
}

func NewBot(token string, repo Repository, cache RedisCache) (*Bot, error) {
	bot, err := api.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{api: bot, repo: repo, cache: cache}, nil
}

type Repository interface {
	GetUserIDByOrderID(ctx context.Context, orderID int64) (int64, error)
}

type RedisCache interface {
	Get(ctx context.Context, key string) (string, error)
	AddToUserHistoryDay(ctx context.Context, msg models.OrderMessage) error
}

func (b *Bot) SendMessage(ctx context.Context, msg models.TelegramMessage) error {
	tgMsg := api.NewMessage(msg.ChatID, msg.Text)

	_, err := b.api.Send(tgMsg)
	if err != nil {
		return err
	}

	userID, err := b.cache.Get(ctx, cache.BuildGetUserIDByOrderKey(msg.OrderID))
	if err != nil {
		return err
	}

	var userInt int64
	if userID == "" {
		userInt, err = b.repo.GetUserIDByOrderID(ctx, msg.OrderID)
		if err != nil {
			return err
		}
	} else {
		userInt, err = strconv.ParseInt(userID, 10, 64)
		if err != nil {
			return err
		}
	}

	msg.UserID = userInt

	err = b.cache.AddToUserHistoryDay(ctx, msg.OrderMessage)
	if err != nil {
		return err
	}

	//log.Printf("sent to chat %d message: %s", msg.ChatID, msg.Text)

	return nil
}
