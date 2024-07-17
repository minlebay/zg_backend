package services

import (
	"go.uber.org/zap"
	"sort"
	"zg_backend/internal/app/nosql_repository"
	"zg_backend/internal/model"
)

type NoSqlService interface {
	GetAll(page int, size int) ([]*model.Message, error)
	GetMessageByID(string) (*model.Message, error)
}

type noSqlService struct {
	logger *zap.Logger
	repo   nosql_repository.NoSqlRepository
}

func NewNoSqlService(l *zap.Logger, r nosql_repository.NoSqlRepository) NoSqlService {
	return &noSqlService{
		logger: l,
		repo:   r,
	}
}

func (s *noSqlService) GetAll(page int, size int) ([]*model.Message, error) {

	messages, err := s.repo.GetMessages(nil)
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

func (s *noSqlService) GetMessageByID(id string) (*model.Message, error) {
	return s.repo.GetById(id)
}
