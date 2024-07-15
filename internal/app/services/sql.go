package services

import (
	"zg_backend/internal/app/sql_repository"
	"zg_backend/internal/model"
)

type SqlService interface {
	GetAll(filter interface{}, page int, size int) ([]model.Message, error)
	GetMessageByID(string) (*model.Message, error)
}

type sqlService struct {
	repo sql_repository.SqlRepository
}

func NewSqlService(r sql_repository.SqlRepository) SqlService {
	return &sqlService{r}
}

func (s sqlService) GetAll(filter interface{}, page int, size int) ([]model.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (s sqlService) GetMessageByID(id string) (*model.Message, error) {
	//TODO implement me
	panic("implement me")
}
