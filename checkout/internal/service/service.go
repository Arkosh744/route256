//go:generate mockgen -package=service -destination=./service_mock_internal_test.go -source=${GOFILE}
package service

import (
	"context"

	wp "route256/libs/worker_pool"

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

type Repository interface {
	AddToCart(ctx context.Context, user int64, item *models.ItemData) error
	GetCount(ctx context.Context, user int64, sku uint32) (uint16, error)
	GetUserCart(ctx context.Context, user int64) ([]models.ItemData, error)
	DeleteFromCart(ctx context.Context, user int64, item *models.ItemData) error
	DeleteUserCart(ctx context.Context, user int64) error
}

type LomsClient interface {
	Stocks(ctx context.Context, sku uint32) ([]*models.Stock, error)
	CreateOrder(ctx context.Context, user int64, items []models.ItemData) (int64, error)
}

// ItemsResult is a workaround to use mockgen with generic variable.
type ItemsResult = []wp.Result[models.Item]

type PSClient interface {
	GetProducts(ctx context.Context, userItems []models.ItemData) ItemsResult
}
