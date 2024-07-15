package nosql_repository

import (
	cfg "go.uber.org/config"
)

type Config struct {
	Dbs []string `yaml:"mongodb"`
}

func NewNoSqlRepositoryConfig(provider cfg.Provider) (*Config, error) {
	config := Config{}

	err := provider.Get("nosql_dbs").Populate(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
