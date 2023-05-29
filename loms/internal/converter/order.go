package converter

import (
	"route256/loms/internal/models"
	desc "route256/pkg/loms_v1"
)

func ToItem(in *desc.Item) models.Item {
	return models.Item{
		SKU:   in.GetSku(),
		Count: in.GetCount(),
	}
}

func ToItems(info []*desc.Item) []models.Item {
	items := make([]models.Item, len(info))

	for i, item := range info {
		items[i] = ToItem(item)
	}

	return items
}

func ToOrderDesc(info *models.Order) *desc.ListOrderResponse {
	return &desc.ListOrderResponse{
		Status: ToOrderStatusDesc(info.Status),
		User:    info.User,
		Items:   ToItemsDesc(info.Items),
	}
}

func ToItemDesc(item models.Item) *desc.Item {
	return &desc.Item{
		Sku:   item.SKU,
		Count: item.Count,
	}
}

func ToItemsDesc(in []models.Item) []*desc.Item {
	items := make([]*desc.Item, len(in))

	for i, item := range in {
		items[i] = ToItemDesc(item)
	}

	return items
}

func ToOrderStatusDesc(status string) desc.OrderStatus {
	switch status {
	case models.OrderStatusNew:
		return desc.OrderStatus_NEW
	case models.OrderStatusAwaitingPayment:
		return desc.OrderStatus_AWAITING_PAYMENT
	case models.OrderStatusFailed:
		return desc.OrderStatus_FAILED
	case models.OrderStatusPaid:
		return desc.OrderStatus_PAID
	case models.OrderStatusCanceled:
		return desc.OrderStatus_CANCELLED
	default:
		return desc.OrderStatus_UNKNOWN
	}

}