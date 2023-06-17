//go:generate mockgen -package=service -destination=./service_mock_internal_test.go -source=${GOFILE}
package service

import (
	"context"

	"route256/checkout/internal/models"
)

type cartService struct {
	repo       Repository
	lomsClient LomsClient
	psClient   PSClient
}

func New(repo Repository, loms LomsClient, ps PSClient) *cartService {
	return &cartService{
		repo:       repo,
		lomsClient: loms,
		psClient:   ps,
	}
}

type Repository interface{}

type LomsClient interface {
	Stocks(ctx context.Context, sku uint32) ([]*models.Stock, error)
	CreateOrder(ctx context.Context, user int64, items []*models.ItemData) (int64, error)
}

type PSClient interface {
	GetProduct(ctx context.Context, sku uint32) (*models.ItemBase, error)
}
