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

	results := s.psClient.GetProducts(ctx, userItems)

	items := make([]models.Item, 0, len(userItems))
	var totalPrice uint32

	for i := range results {
		if results[i].Err != nil {
			// if we get any error from PS, we return error
			return nil, results[i].Err
		}

		items = append(items, results[i].Value)

		totalPrice += results[i].Value.Price * uint32(results[i].Value.Count)
	}

	return &models.CartInfo{
		Items:      items,
		TotalPrice: totalPrice,
	}, nil
}

