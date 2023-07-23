package app

import (
	"context"
	"github.com/go-redis/redis/v8"
	redisClient "route256/libs/client/redis"
	"route256/notifications/internal/cache"
	"route256/notifications/internal/repo"
	"route256/notifications/internal/service"

	"route256/libs/client/pg"
	"route256/libs/closer"
	"route256/libs/log"
	NotifV1 "route256/notifications/internal/api/notifications_v1"
	"route256/notifications/internal/config"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type serviceProvider struct {
	service NotifV1.Service

	NotifImpl *NotifV1.Implementation

	pgClient    pg.Client
	cacheClient redisClient.Client
	repo        *repo.Repository
	cache       service.Cache
}

func newServiceProvider(_ context.Context) *serviceProvider {
	sp := &serviceProvider{}

	return sp
}

func (s *serviceProvider) GetPGClient(ctx context.Context) pg.Client {
	if s.pgClient == nil {
		pgCfg, err := pgxpool.ParseConfig(config.AppConfig.GetPostgresDSN())
		if err != nil {
			log.Fatal("failed to parse pg config", zap.Error(err))
		}

		cl, err := pg.NewClient(ctx, pgCfg)
		if err != nil {
			log.Fatal("failed to get pg client", zap.Error(err))
		}

		if cl.PG().Ping(ctx) != nil {
			log.Fatal("failed to ping pg", zap.Error(err))
		}

		closer.Add(cl.Close)

		s.pgClient = cl
	}

	return s.pgClient
}

func (s *serviceProvider) GetRepository(ctx context.Context) service.Repository {
	if s.repo == nil {
		s.repo = repo.NewRepo(s.GetPGClient(ctx))
	}

	return s.repo
}

func (s *serviceProvider) GetRedisCache(_ context.Context) service.Cache {
	if s.cache == nil {
		rdb := redis.NewClient(&redis.Options{
			Addr:     config.AppConfig.GetRedisAddr(),
			Password: "", // No password set
			DB:       0,  // Use default DB
		})

		s.cache = cache.NewRedis(redisClient.NewRedisService(rdb))
	}

	return s.cache
}

func (s *serviceProvider) GetNotificationsService(ctx context.Context) NotifV1.Service {
	if s.service == nil {
		s.service = service.NewService(s.GetRepository(ctx), s.GetRedisCache(ctx))
	}

	return s.service
}

func (s *serviceProvider) GetNotificationsImpl(ctx context.Context) *NotifV1.Implementation {
	if s.NotifImpl == nil {
		s.NotifImpl = NotifV1.NewImplementation(s.GetNotificationsService(ctx))
	}

	return s.NotifImpl
}
