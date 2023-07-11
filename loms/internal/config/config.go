package config

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

const pathToConfig = "config.yaml"

type Config struct {
	GRPC struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"grpc"`

	Postgres struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"postgres"`

	Kafka struct {
		Brokers []string `yaml:"brokers"`
		Topic   string   `yaml:"topic"`
	} `yaml:"kafka"`

	RateLimit struct {
		Limit     int           `yaml:"limit"`
		PeriodRaw string        `yaml:"periodRaw"`
		Period    time.Duration `yaml:"-"`
	} `yaml:"rateLimit"`

	Log struct {
		Preset string `yaml:"preset"`
	} `yaml:"log"`
}

var AppConfig = Config{}

func Init(_ context.Context) error {
	rawYaml, err := os.ReadFile(pathToConfig)
	if err != nil {
		return fmt.Errorf("read config file: %w", err)
	}

	if err = yaml.Unmarshal(rawYaml, &AppConfig); err != nil {
		return fmt.Errorf("parse config file: %w", err)
	}

	if err = AppConfig.getReqLimitPeriod(); err != nil {
		return fmt.Errorf("get request limit period: %w", err)
	}

	return nil
}

func (c *Config) GetGRPCAddr() string {
	return net.JoinHostPort(c.GRPC.Host, c.GRPC.Port)
}

func (c *Config) GetPostgresDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Postgres.Host, c.Postgres.Port, c.Postgres.User, c.Postgres.Password, c.Postgres.Database)
}

func (c *Config) getReqLimitPeriod() error {
	dur, err := time.ParseDuration(c.RateLimit.PeriodRaw)
	if err != nil {
		return fmt.Errorf("parse request limit period: %w", err)
	}

	c.RateLimit.Period = dur

	return nil
}
