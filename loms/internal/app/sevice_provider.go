package app

import (
	"context"

	LomsV1 "route256/loms/internal/api/loms_v1"
	"route256/loms/internal/repository/cart"
	"route256/loms/internal/service"
)

type serviceProvider struct {
	service LomsV1.Service

	lomsImpl *LomsV1.Implementation

	repo service.Repository
}

func newServiceProvider(_ context.Context) *serviceProvider {
	sp := &serviceProvider{}

	return sp
}

func (s *serviceProvider) GetRepository(_ context.Context) service.Repository {
	if s.repo == nil {
		s.repo = cart.NewRepo()
	}

	return s.repo
}

func (s *serviceProvider) GetLomsService(ctx context.Context) LomsV1.Service {
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
