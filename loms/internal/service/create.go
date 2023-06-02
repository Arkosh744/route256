package service

import (
	"context"

	"route256/loms/internal/models"
)

func (s *service) Create(ctx context.Context, user int64, items []models.Item) (int64, error) {
	return 1, nil
}
