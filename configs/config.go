package configs

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		DB   `yaml:"db"`
		HTTP `yaml:"http"`
	}

	// DB -.
	DB struct {
		UserName string `env-required:"false" yaml:"username"`
		Host     string `env-required:"false" yaml:"host"`
		Port     string `env-required:"false" yaml:"port"`
		DbName   string `env-required:"false" yaml:"dbname"`
		SslMode  string `env-required:"false" yaml:"sslmode"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"false" yaml:"port"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./configs/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
