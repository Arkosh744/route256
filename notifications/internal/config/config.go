package config

import (
	"context"
	"fmt"
	"net"
	"os"

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

	Redis struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"redis"`

	Metrics struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"metrics"`

	Jaeger struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"jaeger"`

	Kafka struct {
		Brokers []string `yaml:"brokers"`
		Topic   string   `yaml:"topic"`
	} `yaml:"kafka"`

	Tg struct {
		Token  string `yaml:"token"`
		ChatID int64  `yaml:"chatId"`
	} `yaml:"tg"`

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

	return nil
}

func (c *Config) GetPostgresDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Postgres.Host, c.Postgres.Port, c.Postgres.User, c.Postgres.Password, c.Postgres.Database)
}

func (c *Config) GetRedisAddr() string {
	return net.JoinHostPort(c.Redis.Host, c.Redis.Port)
}

func (c *Config) GetGRPCAddr() string {
	return net.JoinHostPort(c.GRPC.Host, c.GRPC.Port)
}

func (c *Config) GetMetricsAddr() string {
	return net.JoinHostPort(c.Metrics.Host, c.Metrics.Port)
}

func (c *Config) GetJaegerAddr() string {
	return net.JoinHostPort(c.Jaeger.Host, c.Jaeger.Port)
}