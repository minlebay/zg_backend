package sql_repository

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
	"sync"
	"zg_backend/internal/model"
)

type MySQLRepository struct {
	Config *Config
	Logger *zap.Logger
	wg     sync.WaitGroup
	dbs    []*gorm.DB
}

func NewMySQLRepository(logger *zap.Logger, config *Config) *MySQLRepository {
	return &MySQLRepository{
		Config: config,
		Logger: logger,
	}
}

func (r *MySQLRepository) Start(ctx context.Context) {
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

func (r *MySQLRepository) Stop(ctx context.Context) {
	r.wg.Wait()
	for _, db := range r.dbs {
		err := db.Close()
		if err != nil {
			r.Logger.Error("Failed to disconnect from db", zap.Error(err))
		}
	}
	r.Logger.Info("Repo started")
}

func (r *MySQLRepository) GetAll(ctx context.Context, db *gorm.DB) ([]*model.Message, error) {
	var messages []*model.Message
	err := db.Find(&messages)
	if err.Error != nil {
		return nil, err.Error
	}
	return messages, nil
}

func (r *MySQLRepository) GetById(ctx context.Context, id string, db *gorm.DB) (*model.Message, error) {
	m := &model.Message{}
	err := db.Where("uuid=?", id).First(&m).Error
	return m, err
}

func (r *MySQLRepository) GetDbs() []*gorm.DB {
	return r.dbs
}
