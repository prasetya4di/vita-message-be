package impl

import (
	"vita-message-service/data/entity"
	"vita-message-service/data/local"
	"vita-message-service/repository"
)

type cacheMessageRepository struct {
	dao local.CacheMessageDao
}

func NewCacheMessageRepository(dao local.CacheMessageDao) repository.CacheMessageRepository {
	return &cacheMessageRepository{dao: dao}
}

func (cmr *cacheMessageRepository) Read(message string) (entity.CacheMessage, error) {
	return cmr.dao.Read(message)
}

func (cmr *cacheMessageRepository) Insert(message entity.CacheMessage) (entity.CacheMessage, error) {
	return cmr.dao.Insert(message)
}
