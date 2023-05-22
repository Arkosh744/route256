package loms

import (
	"context"
	"route256/checkout/internal/models"
)

const (
	stocksPath      = "stocks"
	createOrderPath = "createOrder"
)

type Client interface {
	Stocks(ctx context.Context, sku uint32) ([]*models.Stock, error)
	CreateOrder(ctx context.Context, user int64, items []*models.ItemData) (int64, error)
}

type client struct {
	host string
}

func New(host string) *client {
	return &client{host: host}
}
