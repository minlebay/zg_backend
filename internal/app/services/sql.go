package services

import (
	"hash/crc32"
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

	var messages []*model.Message
	dbs := s.repo.GetDbs()
	for _, db := range dbs {
		msgs, err := s.repo.GetAll(db)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msgs...)
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

	dbs := s.repo.GetDbs()
	shardIndex, err := s.getShardIndex(uuid, len(dbs)) // TODO use redis instead of crc32
	if err != nil {
		return nil, err
	}

	// TODO implement work with cache

	return s.repo.GetById(dbs[shardIndex], uuid)

}

func (s *sqlService) getShardIndex(uuid string, dbsCount int) (int, error) {
	uuidBytes := []byte(uuid)
	hash := crc32.ChecksumIEEE(uuidBytes)
	shardNumber := int(hash) % dbsCount
	return shardNumber, nil
}
