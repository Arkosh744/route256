//go:generate mockgen -package=loms_v1 -destination=./service_mock_internal_test.go -source=${GOFILE}
package loms_v1

import (
	"context"

	"route256/loms/internal/models"
	desc "route256/pkg/loms_v1"
)

type Implementation struct {
	desc.UnimplementedLomsServer

	lomsService Service
}

func NewImplementation(s Service) *Implementation {
	return &Implementation{
		lomsService: s,
	}
}

type Service interface {
	Stocks(ctx context.Context, sku uint32) ([]models.StockItem, error)
	Create(ctx context.Context, user int64, items []models.Item) (int64, error)
	Get(ctx context.Context, user int64) (*models.Order, error)
	Paid(ctx context.Context, orderID int64) error
	Cancel(ctx context.Context, orderID int64) error
}
