package sql_repository

import (
	"context"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"sync"
	"zg_backend/internal/model"
)

var stubMessage = &model.Message{
	Uuid:        "stub",
	ContentType: "stub",
}

type Stub struct {
	Config *Config
	Logger *zap.Logger
	dbs    []*gorm.DB
}

func NewRepositoryStub(config *Config, logger *zap.Logger) *Stub {
	return &Stub{
		Config: config,
		Logger: logger,
		dbs: []*gorm.DB{
			{
				RWMutex:      sync.RWMutex{},
				Value:        nil,
				Error:        nil,
				RowsAffected: 0,
			},
			{
				RWMutex:      sync.RWMutex{},
				Value:        nil,
				Error:        nil,
				RowsAffected: 0,
			},
		},
	}
}

func (r *Stub) Start() {

}

func (r *Stub) Stop() {

}

func (r *Stub) GetAll() ([]*model.Message, error) {
	return []*model.Message{stubMessage}, nil
}

func (r *Stub) Create(shard int, entity *model.Message) error {
	return nil
}

func (r *Stub) GetById(id string) (*model.Message, error) {
	return stubMessage, nil
}

func (r *Stub) Update(ctx context.Context, id string, entity *model.Message) (*model.Message, error) {
	return stubMessage, nil
}

func (r *Stub) Delete(ctx context.Context, id string) error {
	return nil
}

func (r *Stub) GetDbs() []*gorm.DB {
	return r.dbs
}

func (r *Stub) DoMigrations(pathToMigrations string, db *gorm.DB) error {
	return nil
}
