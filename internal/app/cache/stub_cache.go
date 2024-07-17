package cache

import (
	"go.uber.org/zap"
)

type StubCache struct {
	Config *Config
	Logger *zap.Logger
}

func NewCacheStub(config *Config, logger *zap.Logger) *StubCache {
	return &StubCache{
		Config: config,
		Logger: logger,
	}
}

func (c *StubCache) Start() {

}

func (c *StubCache) Stop() {

}

func (c *StubCache) Get(key string) (out []byte, err error) {
	return []byte("stub"), nil
}

func (c *StubCache) Put(key string, value []byte) (err error) {
	return nil
}

func (c *StubCache) Iterate(filter string) (keys []string, err error) {
	return []string{"stub_key"}, nil
}

func (c *StubCache) Delete(key string) (err error) {
	return nil
}
