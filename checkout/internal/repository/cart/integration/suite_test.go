//go:build integration
// +build integration

package integration

import (
	"context"
	"fmt"
	l "log"
	"os/exec"
	"testing"

	"route256/checkout/internal/repository/cart"
	"route256/libs/log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"route256/libs/client/pg"
)

type Suite struct {
	suite.Suite

	pool     *dockertest.Pool
	resource *dockertest.Resource
	repo     *cart.Repository
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	const (
		logPreset      = "dev"
		pgTestUser     = "test"
		pgTestPass     = "test"
		pgTestDB       = "test"
		migrationsPath = "../../../../migrations"
	)

	ctx := context.Background()
	if err := log.InitLogger(ctx, logPreset); err != nil {
		l.Fatalf("failed to init logger %v", err)
	}

	// Get postgres container

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatal("Could not connect to docker", zap.Error(err))
	}

	if err = pool.Client.Ping(); err != nil {
		log.Fatal("Could not connect to Docker", zap.Error(err))
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "15",
		Env:        []string{"POSTGRES_PASSWORD=" + pgTestPass, "POSTGRES_USER=" + pgTestUser, "POSTGRES_DB=" + pgTestDB},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatal("Could not start resource", zap.Error(err))
	}

	s.pool = pool
	s.resource = resource

	dsn := s.getPostgresDSN("localhost", resource.GetPort("5432/tcp"), pgTestUser, pgTestPass, pgTestDB)

	pgCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal("failed to parse pg config", zap.Error(err))
	}

	if err = pool.Retry(func() error {
		pgClient, clErr := pg.NewClient(ctx, pgCfg)
		if clErr != nil {
			return clErr
		}
		defer pgClient.Close()

		return pgClient.PG().Ping(ctx)
	}); err != nil {
		log.Fatal("Could not connect to pg docker", zap.Error(err))
	}

	migrationCmd := exec.Command("goose", "postgres", dsn, "up")
	migrationCmd.Dir = migrationsPath

	if err = migrationCmd.Run(); err != nil {
		log.Fatal("failed to run migrations", zap.Error(err))
	}

	log.Info("migrations applied successfully")

	db, err := pg.NewClient(ctx, pgCfg)
	if err != nil {
		log.Fatal("failed to create pg client", zap.Error(err))
	}

	s.repo = cart.NewRepo(db)
}

func (s *Suite) TearDownSuite() {
	if err := s.pool.Purge(s.resource); err != nil {
		log.Fatal("Could not purge resource", zap.Error(err))
	}
}

func (s *Suite) getPostgresDSN(host, port, user, password, db string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, db)
}
