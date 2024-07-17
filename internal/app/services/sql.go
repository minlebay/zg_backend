package services

import (
	"sort"
	"zg_backend/internal/app/sql_repository"
	"zg_backend/internal/model"
)

type SqlService interface {
	GetAll(page int, size int) ([]*model.Message, error)
	GetMessageByID(string) (*model.Message, error)
}

type sqlService struct {
	repo sql_repository.SqlRepository
}

func NewSqlService(r sql_repository.SqlRepository) SqlService {
	return &sqlService{r}
}

func (s *sqlService) GetAll(page int, size int) ([]*model.Message, error) {

	messages, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	sort.Slice(messages, func(i, j int) bool {
		return messages[i].Uuid < messages[j].Uuid
	})

	totalMessages := len(messages)
	start := page * size
	end := start + size

	if start >= totalMessages {
		return []*model.Message{}, nil
	}

	if end > totalMessages {
		end = totalMessages
	}

	return messages[start:end], nil
}

func (s *sqlService) GetMessageByID(uuid string) (*model.Message, error) {
	return s.repo.GetById(uuid)
}
