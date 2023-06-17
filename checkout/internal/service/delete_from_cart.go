package service

import (
	"context"

	"route256/checkout/internal/models"
)

func (s *cartService) DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint16) error {
	if err := s.repo.DeleteFromCart(ctx, user, &models.ItemData{SKU: sku, Count: count}); err != nil {
		return err
	}

	return nil
}
