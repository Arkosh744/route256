package service

import (
	"context"

	"route256/checkout/internal/models"
)

func (s *cartService) ListCart(ctx context.Context, user int64) (*models.CartInfo, error) {
	// for testing purposes
	skus := []uint32{
		1076963,
		1148162,
		1625903,
		2618151,
	}
	counts := []uint32{1, 4, 2, 1}

	items := make([]models.ItemInfo, 0, len(skus))

	var totalPrice uint32

	for i, sku := range skus {
		res, err := s.psClient.GetProduct(ctx, sku)
		if err != nil {
			return nil, err
		}

		items = append(items, models.ItemInfo{
			ItemBase: models.ItemBase{
				Name:  res.Name,
				Price: res.Price,
			},
			ItemData: models.ItemData{
				SKU:   sku,
				Count: counts[i],
			},
		})

		totalPrice += res.Price * counts[i]
	}

	return &models.CartInfo{
		Items:      items,
		TotalPrice: totalPrice,
	}, nil
}
