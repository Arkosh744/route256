package tg

import (
	"context"
	api "github.com/go-telegram-bot-api/telegram-bot-api"
	"route256/notifications/internal/models"
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
