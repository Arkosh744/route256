//go:generate mockgen -package=checkout_v1 -destination=./service_mock_internal_test.go -source=${GOFILE}
package checkout_v1

import (
	"context"

	"route256/checkout/internal/models"
	desc "route256/pkg/checkout_v1"
)

type Implementation struct {
	desc.UnimplementedCheckoutServer

	cartService Service
}

func NewImplementation(s Service) *Implementation {
	return &Implementation{
		cartService: s,
	}
}

type Service interface {
	AddToCart(ctx context.Context, user int64, sku uint32, count uint16) error
	DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint16) error
	ListCart(ctx context.Context, user int64) (*models.CartInfo, error)
	Purchase(ctx context.Context, user int64) (int64, error)
}
