package config

import (
	"github.com/shahTeam/crmconnect/postgres"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/joho/godotenv/autoload"
)


type Config struct {
	Host string `envconfig:"host" required:"true"`
	Port string `envconfig:"port" required:"true"`
	postgres.Config
}

func Load() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, err 
	}
	return cfg, nil
}
