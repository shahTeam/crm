package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	StudentServiceAddr string `envconfig:"STUDENT_SERVICE_ADDR"`
}

func Load() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
           return Config{}, err
	}
	return cfg, nil
}