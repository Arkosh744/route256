package loms

import (
	"context"

	"route256/checkout/internal/converter"
	"route256/checkout/internal/models"
	lomsV1 "route256/pkg/loms_v1"
)

func (c *client) Stocks(ctx context.Context, sku uint32) ([]*models.Stock, error) {
	res, err := c.lomsClient.Stocks(ctx, &lomsV1.StocksRequest{Sku: sku})
	if err != nil {
		return nil, err
	}

	return converter.DescToStock(res), nil
}
