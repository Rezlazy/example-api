package config

import (
	"github.com/caarlos0/env/v10"
)

type Config struct {
	Auth   `envPrefix:"AUTH_"`
	Server `envPrefix:"SERVER_"`
	PG     `envPrefix:"PG_"`
}

func Parse() (Config, error) {
	config := Config{}
	if err := env.Parse(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
