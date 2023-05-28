package app

import (
	"context"

	checkoutV1 "route256/checkout/internal/api/checkout_v1"
	"route256/checkout/internal/clients/loms"
	"route256/checkout/internal/clients/ps"
	"route256/checkout/internal/config"
	"route256/checkout/internal/repository/cart"
	"route256/checkout/internal/service"
)

type serviceProvider struct {
	cartService service.Service

	repo cart.Repository

	cartImpl *checkoutV1.Implementation

	loms loms.Client
	ps   ps.Client
}

func newServiceProvider(_ context.Context) *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GetCartRepo(_ context.Context) cart.Repository {
	if s.repo == nil {
		s.repo = cart.NewRepo()
	}

	return s.repo
}

func (s *serviceProvider) GetLomsClient(_ context.Context) loms.Client {
	if s.loms == nil {
		s.loms = loms.New(config.AppConfig.Services.Loms)
	}

	return s.loms
}

func (s *serviceProvider) GetPSClient(_ context.Context) ps.Client {
	if s.ps == nil {
		s.ps = ps.New(config.AppConfig.Services.ProductService)
	}

	return s.ps
}

func (s *serviceProvider) GetCartService(ctx context.Context) service.Service {
	if s.cartService == nil {
		s.cartService = service.New(s.GetCartRepo(ctx), s.GetLomsClient(ctx), s.GetPSClient(ctx))
	}

	return s.cartService
}

func (s *serviceProvider) GetCheckoutImpl(ctx context.Context) *checkoutV1.Implementation {
	if s.cartImpl == nil {
		s.cartImpl = checkoutV1.NewImplementation(s.GetCartService(ctx))
	}

	return s.cartImpl
}