package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

var config *Config

type Config struct {
	Database Database
}

func Get() *Config {
	if config == nil {
		log.Fatal("config was not successfully loaded")
	}
	return config
}

func Load() Config {
	config = new(Config)
	if err := env.Parse(config); err != nil {
		log.Fatal(err)
	}

	if err := config.Validate(); err != nil {
		log.Fatal(err)
	}

	return Config{}
}

// Validate validate data structure settings
func (s *Config) Validate() error {
	if err := s.Database.Validate(); err != nil {
		return err
	}
	return nil
}
