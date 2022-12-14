package config

import (
	"github.com/shahTeam/crmconnect/postgres"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

// Config defines configurations values needed for the entire service
type Config struct {
	Host string `envconfig:"HOST" required:"true"`
	Port string `envconfig:"PORT" required:"true"`
	postgres.Config
}

func Load() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}