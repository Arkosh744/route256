package converter

import (
	"route256/checkout/internal/models"
	desc "route256/pkg/checkout_v1"
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
