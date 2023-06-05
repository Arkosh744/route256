package service

import (
	"context"
	"fmt"
)

var ErrCartEmpty = fmt.Errorf("cart is empty")

func (s *cartService) Purchase(ctx context.Context, user int64) (int64, error) {
	items, err := s.repo.GetUserData(ctx, user)
	if err != nil {
		return 0, err
	}

	if len(items) == 0 {
		return 0, ErrCartEmpty
	}

	orderID, err := s.lomsClient.CreateOrder(ctx, user, items)
	if err != nil {
		return 0, fmt.Errorf("create order: %w", err)
	}

	return orderID, nil
}
