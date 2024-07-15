package services

import (
	"zg_backend/internal/app/nosql_repository"
	"zg_backend/internal/model"
)

type NoSqlService interface {
	GetAll(filter interface{}, page int, size int) ([]model.Message, error)
	GetMessageByID(string) (*model.Message, error)
}

type noSqlService struct {
	repo nosql_repository.NoSqlRepository
}

func NewNoSqlService(r nosql_repository.NoSqlRepository) NoSqlService {
	return &noSqlService{r}
}

func (s noSqlService) GetAll(filter interface{}, page int, size int) ([]model.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (s noSqlService) GetMessageByID(id string) (*model.Message, error) {
	//TODO implement me
	panic("implement me")
}
