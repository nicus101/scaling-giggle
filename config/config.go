package config

import (
	"errors"
	"os"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Config struct {
	Db struct {
		Name string `env:"DB_NAME"`
		User string `env:"DB_USER"`
		Pass string `env:"DB_PASS"`
		Addr string `env:"DB_ADDR"`
	}
}

var config Config

func Load() error {
	err := godotenv.Load()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	return env.Parse(&config)
}

func Get() Config {
	return config
}
