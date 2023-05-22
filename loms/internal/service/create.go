package service

import (
	"context"
)

func (s *service) Create(ctx context.Context, user int64, sku uint32, count uint16) (int64, error) {
	return 1, nil
}
