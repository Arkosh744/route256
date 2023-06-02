package service

import (
	"context"

	"route256/loms/internal/models"
	"route256/loms/internal/repository/cart"
)

var _ Service = (*service)(nil)

type Service interface {
	Stocks(ctx context.Context, sku uint32) ([]models.StockItem, error)
	Create(ctx context.Context, user int64, items []models.Item) (int64, error)
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
