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
	Token   string `yaml:"token"`
	Timeout string `yaml:"timeout"`

	GRPC struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"grpc"`

	HTTP struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"http"`

	Swagger struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"swagger"`

	Metrics struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"metrics"`

	Jaeger struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"jaeger"`

	Postgres struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"postgres"`

	Services struct {
		Loms           string `yaml:"loms"`
		ProductService string `yaml:"productService"`
	} `yaml:"services"`

	Workers int `yaml:"workers"`

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

func (c *Config) GetHTTPAddr() string {
	return net.JoinHostPort(c.HTTP.Host, c.HTTP.Port)
}

func (c *Config) GetSwaggerAddr() string {
	return net.JoinHostPort(c.Swagger.Host, c.Swagger.Port)
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

func (c *Config) GetMetricsAddr() string {
	return net.JoinHostPort(c.Metrics.Host, c.Metrics.Port)
}

func (c *Config) GetJaegerAddr() string {
	return net.JoinHostPort(c.Jaeger.Host, c.Jaeger.Port)
}
