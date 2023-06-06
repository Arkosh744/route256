package service

import (
	"context"

	"route256/loms/internal/models"
)

type service struct {
	repo Repository
}

func New(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

type Repository interface{
	CreateOrder(ctx context.Context, user int64, items []models.Item) (int64, error)
	GetStocks(ctx context.Context, sku uint32) ([]models.StockItem, error)
}
