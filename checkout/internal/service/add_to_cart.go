package service

import (
	"context"
	"errors"
	"fmt"

	"route256/checkout/internal/models"
)

var ErrStockInsufficient = errors.New("stock insufficient")

func (s *cartService) AddToCart(ctx context.Context, user int64, sku uint32, count uint16) error {
	stocks, err := s.lomsClient.Stocks(ctx, sku)
	if err != nil {
		return fmt.Errorf("get stocks: %w", err)
	}

	cartCount, err := s.repo.GetCount(ctx, user, sku)
	if err != nil {
		return err
	}

	var inStock bool
	counter := int64(cartCount) + int64(count)
	for _, stock := range stocks {
		counter -= int64(stock.Count)
		if counter <= 0 {
			inStock = true
			break
		}
	}

	if !inStock {
		return ErrStockInsufficient
	}

	if err = s.repo.AddToCart(ctx, user, &models.ItemData{SKU: sku, Count: count}); err != nil {
		return err
	}

	return nil
}
