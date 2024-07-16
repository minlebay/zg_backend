package services

import (
	"hash/crc32"
	"zg_backend/internal/app/sql_repository"
	"zg_backend/internal/model"
)

type SqlService interface {
	GetAll(filter interface{}, page int, size int) ([]*model.Message, error)
	GetMessageByID(string) (*model.Message, error)
}

type sqlService struct {
	repo sql_repository.SqlRepository
}

func NewSqlService(r sql_repository.SqlRepository) SqlService {
	return &sqlService{r}
}

func (s sqlService) GetAll(filter interface{}, page int, size int) ([]*model.Message, error) {

	// TODO implement pagination

	dbs := s.repo.GetDbs()
	var messages []*model.Message

	for _, db := range dbs {
		msgs, err := s.repo.GetAll(nil, db)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msgs...)
	}

	return messages, nil
}

func (s sqlService) GetMessageByID(uuid string) (*model.Message, error) {

	dbs := s.repo.GetDbs()
	shardIndex, err := s.getShardIndex(uuid, len(dbs)) // TODO use redis instead of crc32
	if err != nil {
		return nil, err
	}

	// TODO implement work with cache

	return s.repo.GetById(nil, uuid, dbs[shardIndex])

}

func (s *sqlService) getShardIndex(uuid string, dbsCount int) (int, error) {
	uuidBytes := []byte(uuid)
	hash := crc32.ChecksumIEEE(uuidBytes)
	shardNumber := int(hash) % dbsCount
	return shardNumber, nil
}
