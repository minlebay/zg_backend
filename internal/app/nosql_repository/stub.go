package nosql_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"zg_backend/internal/model"
)

type NoSqlStub struct {
}

func NewNoSqlStub() *NoSqlStub {
	return &NoSqlStub{}
}

func (s *NoSqlStub) Start() {}

func (s *NoSqlStub) Stop() {}

func (s *NoSqlStub) GetAll() ([]*model.Message, error) {
	return nil, nil
}

func (s *NoSqlStub) Create(ctx context.Context, db mongo.Database, entity *model.Message) (*model.Message, error) {
	return nil, nil
}

func (s *NoSqlStub) GetById(ctx context.Context, db mongo.Database, id string) (*model.Message, error) {
	return nil, nil
}

func (s *NoSqlStub) Update(ctx context.Context, db mongo.Database, id string, entity *model.Message) (*model.Message, error) {
	return nil, nil
}

func (s *NoSqlStub) Delete(ctx context.Context, db mongo.Database, id string) error {
	return nil
}
