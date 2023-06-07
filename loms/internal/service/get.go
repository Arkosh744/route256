package service

import (
	"context"

	"route256/loms/internal/models"
)

func (s *service) Get(ctx context.Context, orderID int64) (*models.Order, error) {
	res, err := s.repo.GetOrder(ctx, orderID)
	if err != nil {
		return nil, err
	}

	items, err := s.repo.GetOrderItems(ctx, orderID)
	if err != nil {
		return nil, err
	}

	res.Items = items

	return res, nil
}
