package service

import (
	"context"
	"fmt"

	"route256/checkout/internal/models"
)

var ErrCartEmpty = fmt.Errorf("cart is empty")

func (s *cartService) Purchase(ctx context.Context, user int64) (int64, error) {
	cart, err := s.ListCart(ctx, user)
	if err != nil {
		return 0, err
	}

	if len(cart.Items) == 0 {
		return 0, ErrCartEmpty
	}

	itemsData := make([]*models.ItemData, 0, len(cart.Items))
	for _, item := range cart.Items {
		itemsData = append(itemsData, &models.ItemData{
			SKU:   item.SKU,
			Count: item.Count,
		})
	}

	orderID, err := s.lomsClient.CreateOrder(ctx, user, itemsData)
	if err != nil {
		return 0, fmt.Errorf("create order: %w", err)
	}

	return orderID, nil
}
