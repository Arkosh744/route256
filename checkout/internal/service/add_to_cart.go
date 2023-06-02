package service

import (
	"context"
	"errors"
	"fmt"

	"route256/checkout/internal/log"
)

var ErrStockInsufficient = errors.New("stock insufficient")

func (s *cartService) AddToCart(ctx context.Context, user int64, sku uint32, count uint16) error {
	stocks, err := s.lomsClient.Stocks(ctx, sku)
	if err != nil {
		return fmt.Errorf("get stocks: %w", err)
	}

	log.Infof("stocks: %v", stocks)

	counter := int64(count)
	for _, stock := range stocks {
		counter -= int64(stock.Count)
		if counter <= 0 {
			return nil
		}
	}

	return ErrStockInsufficient
}
