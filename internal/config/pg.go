package config

type PG struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     uint16 `env:"PORT" envDefault:"5432"`
	User     string `env:"USER" envDefault:"postgres"`
	Password string `env:"PASSWORD" envDefault:"postgres"`
	Database string `env:"DATABASE" envDefault:"postgres"`
}
