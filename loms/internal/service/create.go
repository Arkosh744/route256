package service

import (
	"context"

	"route256/loms/internal/models"
)

func (s *service) Create(ctx context.Context, user int64, items []models.Item) (int64, error) {
	orderID, err := s.repo.CreateOrder(ctx, user, items)
	if err != nil {
		return 0, err
	}

	return orderID, nil
}
