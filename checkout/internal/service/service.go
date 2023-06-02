package service

import (
	"context"

	"route256/checkout/internal/clients/loms"
	"route256/checkout/internal/clients/ps"
	"route256/checkout/internal/models"
	"route256/checkout/internal/repository/cart"
)

var _ Service = (*cartService)(nil)

type Service interface {
	AddToCart(ctx context.Context, user int64, sku uint32, count uint16) error
	DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint16) error
	ListCart(ctx context.Context, user int64) (*models.CartInfo, error)
	Purchase(ctx context.Context, user int64) (int64, error)
}

type cartService struct {
	repo       cart.Repository
	lomsClient loms.Client
	psClient   ps.Client
}

func New(repo cart.Repository, cLoms loms.Client, cPS ps.Client) *cartService {
	return &cartService{
		repo:       repo,
		lomsClient: cLoms,
		psClient:   cPS,
	}
}
