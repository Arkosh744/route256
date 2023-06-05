package service

import (
	"context"

	"route256/checkout/internal/models"
	"route256/libs/log"
)

func (s *cartService) ListCart(ctx context.Context, user int64) (*models.CartInfo, error) {
	userItems, err := s.repo.GetUserData(ctx, user)
	if err != nil {
		return nil, err
	}

	items := make([]models.ItemInfo, 0, len(userItems))

	var totalPrice uint32
	for i := range userItems {
		log.Infof("user item: %+v", userItems[i])
		res, err := s.psClient.GetProduct(ctx, userItems[i].SKU)
		if err != nil {
			return nil, err
		}

		items = append(items, models.ItemInfo{
			ItemBase: models.ItemBase{
				Name:  res.Name,
				Price: res.Price,
			},
			ItemData: models.ItemData{
				SKU:   userItems[i].SKU,
				Count: userItems[i].Count,
			},
		})

		totalPrice += res.Price * userItems[i].Count
	}

	log.Infof("items: %+v", items)
	return &models.CartInfo{
		Items:      items,
		TotalPrice: totalPrice,
	}, nil
}
