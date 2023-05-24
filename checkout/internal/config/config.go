package config

import (
	"context"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const pathToConfig = "config.yaml"

type Config struct {
	Token string `yaml:"token"`
	Host  string `yaml:"host"`
	Port  string `yaml:"port"`

	Services struct {
		Loms           string `yaml:"loms"`
		ProductService string `yaml:"productService"`
	} `yaml:"services"`

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

	err = yaml.Unmarshal(rawYaml, &AppConfig)
	if err != nil {
		return fmt.Errorf("parse config file: %w", err)
	}

	return nil
}
