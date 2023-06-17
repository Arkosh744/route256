package service

import (
	"context"

	"route256/checkout/internal/models"
)

func (s *cartService) ListCart(ctx context.Context, user int64) (*models.CartInfo, error) {
	userItems, err := s.repo.GetUserCart(ctx, user)
	if err != nil {
		return nil, err
	}

	items := make([]models.Item, 0, len(userItems))

	var totalPrice uint32

	for i := range userItems {
		var res *models.ItemInfo

		res, err = s.psClient.GetProduct(ctx, userItems[i].SKU)
		if err != nil {
			return nil, err
		}

		items = append(items, models.Item{
			ItemInfo: models.ItemInfo{
				Name:  res.Name,
				Price: res.Price,
			},
			ItemData: models.ItemData{
				SKU:   userItems[i].SKU,
				Count: userItems[i].Count,
			},
		})

		totalPrice += res.Price * uint32(userItems[i].Count)
	}

	return &models.CartInfo{
		Items:      items,
		TotalPrice: totalPrice,
	}, nil
}
