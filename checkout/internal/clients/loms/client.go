package loms

import (
	"context"

	"route256/checkout/internal/models"
	lomsV1 "route256/pkg/loms_v1"
)

var _ Client = (*client)(nil)

type Client interface {
	Stocks(ctx context.Context, sku uint32) ([]*models.Stock, error)
	CreateOrder(ctx context.Context, user int64, items []*models.ItemData) (int64, error)
}

type client struct {
	lomsClient lomsV1.LomsClient
}

func New(loms lomsV1.LomsClient) *client {
	return &client{
		lomsClient: loms,
	}
}
