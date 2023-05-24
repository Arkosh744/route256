package service

import (
	"context"
	"route256/loms/internal/models"
)

func (s *service) Get(ctx context.Context, user int64) (*models.Order, error) {
	// for tests
	res := models.Order{
		Status: models.OrderStatusPaid,
		User:   user,
		Items: []models.Item{
			{SKU: 4678287, Count: 2},
		},
	}

	return &res, nil
}
