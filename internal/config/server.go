package config

import "time"

type Server struct {
	Host           string        `env:"HOST" envDefault:"localhost"`
	Port           int           `env:"PORT" envDefault:"80"`
	DefaultTimeout time.Duration `env:"DEFAULT_TIMEOUT" envDefault:"30s"`
}
