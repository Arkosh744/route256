package app

import (
	"context"

	LomsV1 "route256/loms/internal/api/loms_v1"
	"route256/loms/internal/repository/cart"
	"route256/loms/internal/service"
)

type serviceProvider struct {
	service service.Service

	lomsImpl *LomsV1.Implementation

	repo cart.Repository
}

func newServiceProvider(ctx context.Context) *serviceProvider {
	sp := &serviceProvider{}
	sp.GetLomsService(ctx)

	return sp
}

func (s *serviceProvider) GetRepository(_ context.Context) cart.Repository {
	if s.repo == nil {
		s.repo = cart.NewRepo()
	}

	return s.repo
}

func (s *serviceProvider) GetLomsService(ctx context.Context) service.Service {
	if s.service == nil {
		s.service = service.New(s.GetRepository(ctx))
	}

	return s.service
}

func (s *serviceProvider) GetLomsImpl(ctx context.Context) *LomsV1.Implementation {
	if s.lomsImpl == nil {
		s.lomsImpl = LomsV1.NewImplementation(s.GetLomsService(ctx))
	}

	return s.lomsImpl
}
