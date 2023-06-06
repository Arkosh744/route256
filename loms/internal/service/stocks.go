package service

import (
	"context"

	"route256/loms/internal/models"
)

func (s *service) Stocks(ctx context.Context, sku uint32) ([]models.StockItem, error) {
	items, err := s.repo.GetStocks(ctx, sku)
	if err != nil {
		return nil, err
	}

	return items, nil
}
