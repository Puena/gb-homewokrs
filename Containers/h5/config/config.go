package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Address     string `env:"ADDRESS" envDefault:"localhost:8080"`
	DatabaseURI string `env:"DATABASE_URI" envDefault:"postgres://postgres:test@postgres:5432/postgres?sslmode=disable"`
}

func New() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
