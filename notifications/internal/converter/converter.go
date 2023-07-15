package converter

import (
	"route256/notifications/internal/models"
	desc "route256/pkg/notifications_v1"
)

func ToOrderDesc(info []models.OrderMessage) []*desc.Message {
	resp := make([]*desc.Message, len(info))

	for i := range info {
		resp[i] = ToMessageDesc(info[i])
	}

	return resp
}

func ToMessageDesc(item models.OrderMessage) *desc.Message {
	return &desc.Message{
		OrderId: item.OrderID,
		Status:  item.Status,
	}
}
