package sql_kv_db

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"sort"
	"strconv"
	"sync"
)

type RedisSqlKvDb struct {
	Config *Config
	Logger *zap.Logger
	db     *redis.Client
	wg     sync.WaitGroup
}

func NewRedisSqlKvDb(logger *zap.Logger, config *Config) *RedisSqlKvDb {
	return &RedisSqlKvDb{
		Config: config,
		Logger: logger,
	}
}

func (r *RedisSqlKvDb) Start() {
	go func() {
		numdb, err := strconv.ParseInt(r.Config.DB, 10, 64)
		if err != nil {
			r.Logger.Error("Failed to parse DB", zap.Error(err))
		}

		if err != nil {
			r.Logger.Error("Failed to parse expiration time", zap.Error(err))
		}

		r.db = redis.NewClient(&redis.Options{
			Addr: r.Config.Address,
			DB:   int(numdb),
		})
	}()
}

func (r *RedisSqlKvDb) Stop() {
	r.wg.Wait()
	err := r.db.Close()
	if err != nil {
		r.Logger.Error("Failed to disconnect from RedisSqlKvDb", zap.Error(err))
	}
}

func (r *RedisSqlKvDb) Get(key string) (out []byte, err error) {
	out, err = r.db.Get(key).Bytes()
	return
}

func (r *RedisSqlKvDb) Put(key string, value []byte) (err error) {
	err = r.db.Set(key, string(value), 0).Err()
	return
}

func (r *RedisSqlKvDb) Delete(key string) (err error) {
	err = r.db.Del(key).Err()
	return
}

func (r *RedisSqlKvDb) Iterate(filter string) (out []string, err error) {
	if filter != "" {
		filter += "*"
	}
	iter := r.db.Scan(0, filter, 0).Iterator()
	for iter.Next() {
		key := iter.Val()
		out = append(out, key)
	}
	sort.Strings(out)
	err = iter.Err()
	return
}
