package server

import (
	cfg "go.uber.org/config"
)

type Config struct {
	Port string `yaml:"port"`
}

func NewServerConfig(provider cfg.Provider) (*Config, error) {
	config := Config{}

	err := provider.Get("server").Populate(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
