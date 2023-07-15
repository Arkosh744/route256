//go:generate mockgen -package=notifications_v1 -destination=./service_mock_internal_test.go -source=${GOFILE}
package notifications_v1

import (
	"context"
	"route256/notifications/internal/models"

	desc "route256/pkg/notifications_v1"
)

type Implementation struct {
	desc.UnimplementedNotificationsServer

	service Service
}

func NewImplementation(s Service) *Implementation {
	return &Implementation{
		service: s,
	}
}

type Service interface {
	ListUserHistoryDay(ctx context.Context, userID int64) ([]models.OrderMessage, error)
}
