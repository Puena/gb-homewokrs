package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Address string `env:"ADDRESS" envDefault:"localhost:8080"`
}

func New() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
