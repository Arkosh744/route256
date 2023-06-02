package converter

import (
	"route256/loms/internal/models"
	desc "route256/pkg/loms_v1"
)

func ToStockDesc(item models.StockItem) *desc.Stock {
	return &desc.Stock{
		WarehouseId: item.WarehouseID,
		Count:       item.Count,
	}
}

func ToStocksDesc(in []models.StockItem) []*desc.Stock {
	items := make([]*desc.Stock, len(in))

	for i, item := range in {
		items[i] = ToStockDesc(item)
	}

	return items
}
