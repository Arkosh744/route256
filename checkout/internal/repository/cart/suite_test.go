package cart

import (
	"context"
	"fmt"
	l "log"
	"os/exec"
	"testing"

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
	repo     *repository
}

func TestSuite(t *testing.T) {
	s := new(Suite)
	s.SetupSuite()

	suite.Run(t, s)
}

func (s *Suite) SetupSuite() {
	const (
		logPreset      = "dev"
		pgUser         = "user"
		pgPass         = "secret"
		pgDB           = "test"
		migrationsPath = "../../../migrations"
	)

	// Get postgres container

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	if err = pool.Client.Ping(); err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "15",
		Env:        []string{"POSTGRES_PASSWORD=" + pgPass, "POSTGRES_USER=" + pgUser, "POSTGRES_DB=" + pgDB},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	s.pool = pool
	s.resource = resource

	ctx := context.Background()
	if err = log.InitLogger(ctx, logPreset); err != nil {
		l.Fatalf("failed to init logger %v", zap.Error(err))
	}

	dsn := s.getPostgresDSN("localhost", resource.GetPort("5432/tcp"), pgUser, pgPass, pgDB)

	pgCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("failed to parse pg config %v", zap.Error(err))
	}

	if err = pool.Retry(func() error {
		pgClient, clErr := pg.NewClient(ctx, pgCfg)
		if clErr != nil {
			return clErr
		}
		defer pgClient.Close()

		return pgClient.PG().Ping(ctx)
	}); err != nil {
		log.Fatalf("Could not connect to pg docker: %s", err)
	}

	migrationCmd := exec.Command("goose", "postgres", dsn, "up")
	migrationCmd.Dir = migrationsPath

	if err = migrationCmd.Run(); err != nil {
		log.Fatalf("failed to run migrations: %v", zap.Error(err))
	}

	log.Infof("migrations applied successfully")

	db, err := pg.NewClient(ctx, pgCfg)
	if err != nil {
		log.Fatalf("failed to create pg client %v", zap.Error(err))
	}

	s.repo = NewRepo(db)
}

func (s *Suite) TeardownSuite() {
	if err := s.pool.Purge(s.resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func (s *Suite) getPostgresDSN(host, port, user, password, db string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, db)
}
