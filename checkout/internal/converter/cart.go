package converter

import (
	"route256/checkout/internal/models"
	desc "route256/pkg/checkout_v1"
	descLoms "route256/pkg/loms_v1"
	descPS "route256/pkg/product_v1"
)

func ListToDesc(info *models.CartInfo) *desc.ListCartResponse {
	deskItems := make([]*desc.CartItem, 0, len(info.Items))
	for _, item := range info.Items {
		deskItems = append(deskItems, &desc.CartItem{
			Sku:   item.SKU,
			Count: item.Price,
			Price: item.Price,
			Name:  item.Name,
		})
	}

	return &desc.ListCartResponse{
		Items:      deskItems,
		TotalPrice: info.TotalPrice,
	}
}

func ItemsDataToDesc(items []*models.ItemData) []*descLoms.Item {
	deskItems := make([]*descLoms.Item, 0, len(items))

	for _, item := range items {
		deskItems = append(deskItems, &descLoms.Item{
			Sku:   item.SKU,
			Count: item.Count,
		})
	}

	return deskItems
}

func DescToStock(in *descLoms.StocksResponse) []*models.Stock {
	result := make([]*models.Stock, 0, len(in.Stocks))

	for _, v := range in.Stocks {
		result = append(result, &models.Stock{
			WarehouseID: v.GetWarehouseID(),
			Count:       v.GetCount(),
		})
	}

	return result
}

func DescToItemBase(in *descPS.GetProductResponse) *models.ItemBase {
	return &models.ItemBase{
		Name:  in.GetName(),
		Price: in.GetPrice(),
	}
}
