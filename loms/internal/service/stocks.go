package service

import (
	"context"
	"route256/loms/internal/models"
)

func (s *service) Stocks(ctx context.Context, sku int64) ([]*models.StockItem, error) {
	items := []*models.StockItem{
		{WarehouseID: 4678287, Count: 10},
		{WarehouseID: 123545, Count: 20},
	}

	return items, nil
}
