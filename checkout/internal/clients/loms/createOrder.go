package loms

import (
	"context"

	"route256/checkout/internal/converter"
	"route256/checkout/internal/models"
	lomsV1 "route256/pkg/loms_v1"
)

func (c *client) CreateOrder(ctx context.Context, user int64, items []*models.ItemData) (int64, error) {
	res, err := c.lomsClient.CreateOrder(
		ctx,
		&lomsV1.CreateOrderRequest{
			User:  user,
			Items: converter.ItemsDataToDesc(items),
		})
	if err != nil {
		return 0, err
	}

	return res.GetOrderId(), nil
}
