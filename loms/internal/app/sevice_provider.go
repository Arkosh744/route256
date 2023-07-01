package app

import (
	"context"

	"route256/libs/client/pg"
	"route256/libs/closer"
	"route256/libs/log"
	"route256/libs/rate_limiter"
	LomsV1 "route256/loms/internal/api/loms_v1"
	"route256/loms/internal/config"
	"route256/loms/internal/notifications/status"
	"route256/loms/internal/repository/cart"
	"route256/loms/internal/service"

	"github.com/Shopify/sarama"
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

func (s *serviceProvider) GetRateLimiter(_ context.Context) rate_limiter.RateLimiter {
	rl := rate_limiter.NewSlidingWindow(config.AppConfig.RateLimit.Limit, config.AppConfig.RateLimit.Period)

	return rl
}

func (s *serviceProvider) GetRateLimiterWithPG(ctx context.Context) rate_limiter.RateLimiter {
	rl, err := rate_limiter.NewSlidingWindowWithPG(
		ctx,
		config.AppConfig.RateLimit.Limit,
		config.AppConfig.RateLimit.Period,
		s.GetPGClient(ctx),
	)
	if err != nil {
		log.Fatalf("failed to create rate limiter with pg: %s", err)
	}

	return rl
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
func (s *serviceProvider) GetKafkaSyncProducer() (sarama.SyncProducer, error) {
	var cfg = sarama.NewConfig()

	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true
	cfg.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Retry.Max = 5
	cfg.Producer.Idempotent = true
	cfg.Net.MaxOpenRequests = 1

	producer, err := sarama.NewSyncProducer(config.AppConfig.Kafka.Brokers, cfg)
	if err != nil {
		return nil, err
	}

	log.Info("kafka sync producer created")

	admin, err := sarama.NewClusterAdmin(config.AppConfig.Kafka.Brokers, cfg)
	if err != nil {
		return nil, err
	}

	err = ensureTopicExists(admin, config.AppConfig.Kafka.Topic, 1, 1)
	if err != nil {
		return nil, err
	}

	closer.Add(producer.Close)

	return producer, nil
}

func ensureTopicExists(admin sarama.ClusterAdmin, topic string, numPartitions int32, replicationFactor int16) error {
	topics, err := admin.ListTopics()
	if err != nil {
		return err
	}

	if _, ok := topics[topic]; !ok {
		details := &sarama.TopicDetail{
			NumPartitions:     numPartitions,
			ReplicationFactor: replicationFactor,
		}

		if err = admin.CreateTopic(topic, details, false); err != nil {
			return err
		}
	}

	return nil
}

func (s *serviceProvider) GetLomsService(ctx context.Context) LomsV1.Service {
	if s.service == nil {
		producer, err := s.GetKafkaSyncProducer()
		if err != nil {
			log.Fatalf("failed to get kafka sync producer: %s", err)
		}

		kafka := status.NewOrderStatusSender(producer, config.AppConfig.Kafka.Topic)
		s.service = service.New(s.GetRepository(ctx), s.GetPGClient(ctx), kafka, s.GetRateLimiterWithPG(ctx))
	}

	return s.service
}

func (s *serviceProvider) GetLomsImpl(ctx context.Context) *LomsV1.Implementation {
	if s.lomsImpl == nil {
		s.lomsImpl = LomsV1.NewImplementation(s.GetLomsService(ctx))
	}

	return s.lomsImpl
}
