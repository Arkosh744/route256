//go:generate mockgen -package=service -destination=./service_mock_internal_test.go -source=${GOFILE}
package service

import (
	"context"
	"route256/notifications/internal/models"
	"time"
)

type Service struct {
	repo  Repository
	cache Cache
}

func NewService(repo Repository, cache Cache) *Service {
	return &Service{
		repo:  repo,
		cache: cache,
	}
}

type Repository interface {
	ListUserHistoryDay(ctx context.Context, userID int64, lastMessageTime *time.Time) ([]models.OrderMessage, error)
}

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value any, ttl time.Duration) error

	AddToUserHistoryDay(ctx context.Context, msg models.OrderMessage) error
	GetUserHistoryDay(ctx context.Context, userID int64) ([]models.OrderMessage, error)
	GetLatestMessageTime(ctx context.Context, userID int64) (time.Time, error)
}
