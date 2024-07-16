package services

import (
	"go.uber.org/zap"
	"hash/crc32"
	"zg_backend/internal/app/nosql_repository"
	"zg_backend/internal/model"
)

type NoSqlService interface {
	GetAll(filter interface{}, page int, size int) ([]*model.Message, error)
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

func (s *noSqlService) GetAll(filter interface{}, page int, size int) ([]*model.Message, error) {

	// TODO: Implement pagination
	dbs := s.repo.GetDbs()

	var messages []*model.Message
	for _, db := range dbs {
		msgs, err := s.repo.GetAll(nil, *db)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msgs...)
	}
	return messages, nil
}

func (s *noSqlService) GetMessageByID(id string) (*model.Message, error) {

	dbs := s.repo.GetDbs()
	dbsCount := len(dbs)
	shardIndex, err := s.getShardIndex(id, dbsCount) // TODO use redis instead of crc32
	if err != nil {
		s.logger.Error("Failed to get shard index", zap.Error(err))
	}

	return s.repo.GetById(nil, *dbs[shardIndex], id)
}

func (s *noSqlService) getShardIndex(uuid string, dbsCount int) (int, error) {
	uuidBytes := []byte(uuid)
	hash := crc32.ChecksumIEEE(uuidBytes)
	shardNumber := int(hash) % dbsCount
	return shardNumber, nil
}
