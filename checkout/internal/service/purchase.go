package service

import (
	"context"
	"math/rand"
)

func (s *cartService) Purchase(ctx context.Context, user int64) (int64, error) {
	return rand.Int63(), nil
}
