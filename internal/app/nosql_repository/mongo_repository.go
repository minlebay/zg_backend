package nosql_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	url2 "net/url"
	"strings"
	"sync"
	"time"
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
}

func NewMongoRepository(logger *zap.Logger, config *Config) *MongoRepository {
	return &MongoRepository{
		Config: config,
		Logger: logger,
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

func (r *MongoRepository) GetMessages(filter interface{}, db mongo.Database) ([]*model.Message, error) {
	if filter == nil {
		filter = bson.D{}
	}

	r.Collection = db.Collection("messages")
	ctx := context.Background()
	var entities []*model.Message

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &entities); err != nil {
		return nil, err
	}

	return entities, nil
}

func (r *MongoRepository) GetById(db mongo.Database, uuid string) (*model.Message, error) {
	r.Collection = db.Collection("messages")
	ctx := context.Background()

	var entity model.Message
	err := r.Collection.FindOne(ctx, bson.M{"uuid": uuid}).Decode(&entity)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *MongoRepository) GetDbs() []*mongo.Database {
	return r.DBs
}
