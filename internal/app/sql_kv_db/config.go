package sql_kv_db

import cfg "go.uber.org/config"

type Config struct {
	Address string `yaml:"address"`
	DB      string `yaml:"db"`
}

func NewSqlKeyValueDbConfig(provider cfg.Provider) (*Config, error) {
	config := Config{}

	if err := provider.Get("sql_kv_db").Populate(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
