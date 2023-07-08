package config

import (
	"fmt"
	"net"
	"os"

	"gopkg.in/yaml.v3"
)

const pathToConfig = "config.yaml"

type Config struct {
	Metrics struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"metrics"`

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

func Init() error {
	rawYaml, err := os.ReadFile(pathToConfig)
	if err != nil {
		return fmt.Errorf("read config file: %w", err)
	}

	if err = yaml.Unmarshal(rawYaml, &AppConfig); err != nil {
		return fmt.Errorf("parse config file: %w", err)
	}

	return nil
}

func (c *Config) GetMetricsAddr() string {
	return net.JoinHostPort(c.Metrics.Host, c.Metrics.Port)
}
