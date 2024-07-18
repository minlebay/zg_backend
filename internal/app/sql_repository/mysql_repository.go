package sql_repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
	"hash/crc32"
	"strconv"
	"sync"
	"zg_backend/internal/app/cache"
	"zg_backend/internal/app/sql_kv_db"
	"zg_backend/internal/model"
)

type MySQLRepository struct {
	Config *Config
	Logger *zap.Logger
	wg     sync.WaitGroup
	dbs    []*gorm.DB
	kvDb   sql_kv_db.SqlKvDb
	cache  cache.Cache
}

func NewMySQLRepository(
	logger *zap.Logger,
	config *Config,
	kvDb sql_kv_db.SqlKvDb,
	cache cache.Cache,
) *MySQLRepository {
	return &MySQLRepository{
		Config: config,
		Logger: logger,
		kvDb:   kvDb,
		cache:  cache,
	}
}

func (r *MySQLRepository) Start() {
	go func() {
		for _, db := range r.Config.Dbs {
			args := fmt.Sprintf(
				"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
				db.User, db.Password, db.Host, db.Port, db.DB,
			)
			gdb, err := gorm.Open("mysql", args)

			if err != nil {
				r.Logger.Error("Failed to connect to db", zap.Error(err))
				return
			}
			r.dbs = append(r.dbs, gdb)
		}
	}()
}

func (r *MySQLRepository) Stop() {
	r.wg.Wait()
	for _, db := range r.dbs {
		err := db.Close()
		if err != nil {
			r.Logger.Error("Failed to disconnect from db", zap.Error(err))
		}
	}
	r.Logger.Info("Repo started")
}

func (r *MySQLRepository) GetAll() ([]*model.Message, error) {

	dbs := r.dbs
	var allMessages []*model.Message

	for _, db := range dbs {
		var messages []*model.Message
		err := db.Find(&messages)
		if err.Error != nil {
			return nil, err.Error
		}
		allMessages = append(allMessages, messages...)
	}

	return allMessages, nil

}

func (r *MySQLRepository) GetById(uuid string) (*model.Message, error) {
	var dbNumber int
	dbBytes, err := r.kvDb.Get(uuid)
	if err != nil {
		r.Logger.Error("Failed to get shard index, use crc32 to target shard number", zap.Error(err))
		dbNumber, _ = r.getShardIndex(uuid, len(r.dbs)) // default shard index
	} else {
		var dbNumber64 int64
		if dbNumber64, err = strconv.ParseInt(string(dbBytes), 10, 0); err != nil {
			return nil, err
		}
		dbNumber = int(dbNumber64)
	}

	message, err := r.getFromCacheOrWarmUp(uuid, dbNumber)
	if err != nil || (message != nil && message.Uuid == "") {
		r.Logger.Error("cache miss", zap.Error(err))
	}

	return message, err
}

func (r *MySQLRepository) getShardIndex(uuid string, dbsCount int) (int, error) {
	uuidBytes := []byte(uuid)
	hash := crc32.ChecksumIEEE(uuidBytes)
	shardNumber := int(hash) % dbsCount
	return shardNumber, nil
}

func (r *MySQLRepository) getFromCacheOrWarmUp(uuid string, dbNumber int) (*model.Message, error) {
	var message model.Message

	// firstly we try to get message from cache
	bytes, err := r.cache.Get(uuid)
	if err == nil {
		err = message.Unmarshal(bytes)
		if err != nil {
			r.Logger.Error("Failed to unmarshal cached message", zap.Error(err))
		}
		if message.Uuid != "" { // empty message, cache miss
			return &message, nil
		}
	}

	r.Logger.Info("Cache miss, try to get message from db", zap.String("uuid", uuid))

	db := r.dbs[dbNumber]
	err = db.Where("uuid=?", uuid).First(&message).Error
	if err != nil {
		return nil, err
	}

	// warm up cache
	bytes, err = message.Marshal()
	if err != nil {
		r.Logger.Error("Failed to marshal message", zap.Error(err))
	}
	err = r.cache.Put(uuid, bytes)
	if err != nil {
		r.Logger.Error("Failed to put message to cache", zap.Error(err))
	}

	return &message, err
}
