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

func (s *Service) ListUserHistoryDay(ctx context.Context, userID int64) ([]models.OrderMessage, error) {
	msgs, err := s.cache.GetUserHistoryDay(ctx, userID)
	if err != nil {
		return nil, err
	}

	if len(msgs) > 0 {
		var lastMsgTime time.Time
		lastMsgTime, err = s.cache.GetLatestMessageTime(ctx, userID)
		if err != nil {
			return nil, err
		}

		var msgsNew []models.OrderMessage
		msgsNew, err = s.repo.ListUserHistoryDay(ctx, userID, &lastMsgTime)
		if err != nil {
			return nil, err
		}

		if len(msgsNew) > 0 {
			msgs = append(msgs, msgsNew...)
			for _, msg := range msgsNew {
				msg.UserID = userID
				if err = s.cache.AddToUserHistoryDay(ctx, msg); err != nil {
					return nil, err
				}
			}
		}

		return msgs, nil
	}

	msgs, err = s.repo.ListUserHistoryDay(ctx, userID, nil)
	if err != nil {
		return nil, err
	}

	for _, msg := range msgs {
		msg.UserID = userID
		if err = s.cache.AddToUserHistoryDay(ctx, msg); err != nil {
			return nil, err
		}
	}

	return msgs, nil
}
