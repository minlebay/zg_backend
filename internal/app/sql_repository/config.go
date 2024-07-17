package sql_repository

import (
	cfg "go.uber.org/config"
)

type Db struct {
	Host     string `yaml:"host"`
	DB       string `yaml:"database"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Config struct {
	Dbs []Db `yaml:"sql_dbs"`
}

func NewRepositoryConfig(provider cfg.Provider) (*Config, error) {
	config := Config{}
	dbs := []Db{}

	err := provider.Get("sql_dbs").Populate(&dbs)
	if err != nil {
		return nil, err
	}
	config.Dbs = dbs

	return &config, nil
}
