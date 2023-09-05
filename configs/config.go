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
		UserName string `env-required:"true" yaml:"username"`
		Host     string `env-required:"true" yaml:"host"`
		Port     string `env-required:"true" yaml:"port"`
		DbName   string `env-required:"true" yaml:"dbname"`
		SslMode  string `env-required:"true" yaml:"sslmode"`
		Password string `env-required:"true" yaml:"password" env:"DB_PASSWORD"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port"`
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
