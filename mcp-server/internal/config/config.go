package config

import (
	"cmp"
	"os"
)

type Config struct {
	PostgreSQL PostgreSQL
}

type PostgreSQL struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

func NewConfig() *Config {
	cfg := &Config{
		PostgreSQL: PostgreSQL{
			Host:     cmp.Or(os.Getenv("DB_HOST"), "sublimex"),
			User:     cmp.Or(os.Getenv("DB_USER"), "sublimex-user"),
			Password: cmp.Or(os.Getenv("DB_PASSWORD"), "sublimex-password"),
			Name:     cmp.Or(os.Getenv("DB_NAME"), "sublimex-name"),
			Port:     cmp.Or(os.Getenv("DB_PORT"), "5432"),
		},
	}

	return cfg
}
