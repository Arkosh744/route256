package app

import (
	"context"

	"route256/libs/rate_limiter"

	"route256/checkout/internal/clients/loms"
	"route256/checkout/internal/clients/ps"
	"route256/checkout/internal/config"
	"route256/checkout/internal/repository/cart"
	"route256/checkout/internal/service"
	"route256/libs/client/pg"
	"route256/libs/closer"
	"route256/libs/log"
	lomsV1 "route256/pkg/loms_v1"
	productV1 "route256/pkg/product_v1"

	checkoutV1 "route256/checkout/internal/api/checkout_v1"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type serviceProvider struct {
	cartService checkoutV1.Service

	repo service.Repository

	cartImpl *checkoutV1.Implementation

	pgClient pg.Client
	loms     service.LomsClient
	ps       service.PSClient
}

func newServiceProvider(_ context.Context) *serviceProvider {
	return &serviceProvider{}
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

func (s *serviceProvider) GetCartRepo(ctx context.Context) service.Repository {
	if s.repo == nil {
		s.repo = cart.NewRepo(s.GetPGClient(ctx))
	}

	return s.repo
}

func (s *serviceProvider) GetLomsClient(ctx context.Context) service.LomsClient {
	if s.loms == nil {
		conn, err := grpc.DialContext(
			ctx,
			config.AppConfig.Services.Loms,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalf("failed to connect %s: %s", config.AppConfig.Services.Loms, err)
		}

		closer.Add(conn.Close)

		lomsClient := lomsV1.NewLomsClient(conn)
		s.loms = loms.New(lomsClient)

		log.Infof("loms client created and connected %s", config.AppConfig.Services.Loms)
	}

	return s.loms
}

func (s *serviceProvider) GetRateLimiter() *rate_limiter.SlidingWindow {
	rl := rate_limiter.NewSlidingWindow(config.AppConfig.ReqLimit, config.AppConfig.ReqLimitPeriod)

	return rl
}

func (s *serviceProvider) GetPSClient(ctx context.Context) service.PSClient {
	if s.ps == nil {
		conn, err := grpc.DialContext(
			ctx,
			config.AppConfig.Services.ProductService,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalf("failed to connect %s: %s", config.AppConfig.Services.ProductService, err)
		}

		closer.Add(conn.Close)

		psClient := productV1.NewProductServiceClient(conn)

		s.ps = ps.New(psClient, s.GetRateLimiter())

		log.Infof("ps client created and connected %s", config.AppConfig.Services.ProductService)
	}

	return s.ps
}

func (s *serviceProvider) GetCartService(ctx context.Context) checkoutV1.Service {
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
