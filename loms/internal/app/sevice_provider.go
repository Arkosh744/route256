package app

import (
	"context"

	"route256/loms/internal/repository/cart"
	"route256/loms/internal/service"
)

type serviceProvider struct {
	service service.Service

	repo cart.Repository
}

func newServiceProvider(ctx context.Context) *serviceProvider {
	sp := &serviceProvider{}
	sp.GetCartService(ctx)

	return sp
}

func (s *serviceProvider) GetRepository(_ context.Context) cart.Repository {
	if s.repo == nil {
		s.repo = cart.NewRepo()
	}

	return s.repo
}

func (s *serviceProvider) GetCartService(ctx context.Context) service.Service {
	if s.service == nil {
		s.service = service.New(s.GetRepository(ctx))
	}

	return s.service
}
