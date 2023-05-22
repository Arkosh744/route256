package service

import (
	"context"
	"route256/loms/internal/models"
	"route256/loms/internal/repository/cart"
)

var _ Service = (*service)(nil)

type Service interface {
	Stocks(ctx context.Context, sku int64) ([]models.StockItem, error)
	Create(ctx context.Context, user int64, sku uint32, count uint16) (int64, error)
	Get(ctx context.Context, user int64) (*models.Order, error)
	Paid(ctx context.Context, orderID int64) error
	Cancel(ctx context.Context, orderID int64) error
}

type service struct {
	repo cart.Repository
}

func New(repo cart.Repository) *service {
	return &service{
		repo: repo,
	}
}
