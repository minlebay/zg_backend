package nosql_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"hash/crc32"
	url2 "net/url"
	"strconv"
	"strings"
	"sync"
	"time"
	"zg_backend/internal/app/nosql_kv_db"
	"zg_backend/internal/model"
)

type MongoRepository struct {
	Config           *Config
	Logger           *zap.Logger
	wg               sync.WaitGroup
	DBs              []*mongo.Database
	Collection       *mongo.Collection
	CancelFunc       context.CancelFunc
	ClientDisconnect func()
	kvDb             nosql_kv_db.NosqlKvDb
}

func NewMongoRepository(logger *zap.Logger, config *Config, kv nosql_kv_db.NosqlKvDb) *MongoRepository {
	return &MongoRepository{
		Config: config,
		Logger: logger,
		kvDb:   kv,
	}
}

func (r *MongoRepository) Start() {
	go func() {
		for _, db := range r.Config.Dbs {
			url, err := url2.Parse(db)
			if err != nil {
				r.Logger.Fatal("Failed to parse MongoDB URL: %v", zap.Error(err))
			}
			databaseName := strings.TrimPrefix(url.Path, "/")

			dbctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			r.CancelFunc = cancel

			clientOptions := options.Client().ApplyURI(db)
			client, err := mongo.Connect(dbctx, clientOptions)
			if err != nil {
				r.Logger.Fatal("Failed to connect to MongoDB: %v", zap.Error(err))
			}
			r.ClientDisconnect = func() {
				if err = client.Disconnect(dbctx); err != nil {
					r.Logger.Fatal("Failed to disconnect from MongoDB: %v", zap.Error(err))
				}
			}
			r.DBs = append(r.DBs, client.Database(databaseName))
		}
	}()
}

func (r *MongoRepository) Stop() {
	r.wg.Wait()
	r.ClientDisconnect()
	r.CancelFunc()

	r.Logger.Info("Repo stopped")
}

func (r *MongoRepository) GetMessages(filter interface{}) ([]*model.Message, error) {
	if filter == nil {
		filter = bson.D{}
	}

	dbs := r.DBs
	var allMessages []*model.Message

	for _, db := range dbs {

		r.Collection = db.Collection("messages")
		ctx := context.Background()
		cursor, err := r.Collection.Find(ctx, filter)
		if err != nil {
			return nil, err
		}
		var messages []*model.Message

		if err := cursor.All(ctx, &messages); err != nil {
			return nil, err
		}

		allMessages = append(allMessages, messages...)
	}

	return allMessages, nil
}

func (r *MongoRepository) GetById(uuid string) (*model.Message, error) {

	var dbNumber int
	dbBytes, err := r.kvDb.Get(uuid)
	if err != nil {
		dbNumber, _ = r.getShardIndex(uuid, len(r.DBs)) // default shard index
	} else {
		var dbNumber64 int64
		if dbNumber64, err = strconv.ParseInt(string(dbBytes), 10, 0); err != nil {
			return nil, err
		}
		dbNumber = int(dbNumber64)
	}

	db := r.DBs[dbNumber]
	r.Collection = db.Collection("messages")
	ctx := context.Background()

	var message model.Message
	err = r.Collection.FindOne(ctx, bson.M{"uuid": uuid}).Decode(&message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *MongoRepository) getShardIndex(uuid string, dbsCount int) (int, error) {
	uuidBytes := []byte(uuid)
	hash := crc32.ChecksumIEEE(uuidBytes)
	shardNumber := int(hash) % dbsCount
	return shardNumber, nil
}
