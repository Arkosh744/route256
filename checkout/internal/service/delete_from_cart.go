package service

import (
	"context"
)

func (s *cartService) DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint16) error {
	return nil
}
