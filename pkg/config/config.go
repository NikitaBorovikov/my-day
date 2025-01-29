package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		Http     Http `yml:"http"`
		Postgres Postgres
	}
	Http struct {
		Port       string `yml:"port"`
		SessionKey string `env:"SESSION_KEY"`
	}
	Postgres struct {
		Host     string `env:"PG_HOST"`
		Port     int64  `env:"PG_PORT"`
		User     string `env:"PG_USER"`
		Password string `env:"PG_PASSWORD"`
		Name     string `env:"PG_NAME"`
	}
)

func Init() (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig("config/config.yml", &cfg); err != nil {
		return nil, err
	}
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	if err := cleanenv.ReadEnv(&cfg.Postgres); err != nil {
		return nil, err
	}
	if err := cleanenv.ReadEnv(&cfg.Http); err != nil {
		return nil, err
	}

	return &cfg, nil
}
