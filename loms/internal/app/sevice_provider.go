package app

import (
	"context"

	"route256/libs/client/pg"
	"route256/libs/closer"
	"route256/libs/log"
	LomsV1 "route256/loms/internal/api/loms_v1"
	"route256/loms/internal/config"
	"route256/loms/internal/repository/cart"
	"route256/loms/internal/service"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type serviceProvider struct {
	service LomsV1.Service

	lomsImpl *LomsV1.Implementation

	pgClient pg.Client
	repo     service.Repository
}

func newServiceProvider(_ context.Context) *serviceProvider {
	sp := &serviceProvider{}

	return sp
}

func (s *serviceProvider) GetPGClient(ctx context.Context) pg.Client {
	if s.pgClient == nil {
		pgCfg, err := pgxpool.ParseConfig(config.AppConfig.GetPostgresDSN())
		if err != nil {
			log.Fatalf("failed to parse pg config", zap.Error(err))
		}

		cl, err := pg.NewClient(ctx, pgCfg)
		if err != nil {
			log.Fatalf("failed to get pg client", zap.Error(err))
		}

		if cl.PG().Ping(ctx) != nil {
			log.Fatalf("failed to ping pg", zap.Error(err))
		}

		closer.Add(cl.Close)

		s.pgClient = cl
	}

	return s.pgClient
}

func (s *serviceProvider) GetRepository(ctx context.Context) service.Repository {
	if s.repo == nil {
		s.repo = cart.NewRepo(s.GetPGClient(ctx))
	}

	return s.repo
}

func (s *serviceProvider) GetLomsService(ctx context.Context) LomsV1.Service {
	if s.service == nil {
		s.service = service.New(s.GetRepository(ctx), s.GetPGClient(ctx))
	}

	return s.service
}

func (s *serviceProvider) GetLomsImpl(ctx context.Context) *LomsV1.Implementation {
	if s.lomsImpl == nil {
		s.lomsImpl = LomsV1.NewImplementation(s.GetLomsService(ctx))
	}

	return s.lomsImpl
}
