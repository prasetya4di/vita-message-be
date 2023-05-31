package impl

import (
	"fmt"
	"gorm.io/gorm"
	"vita-message-service/data/entity"
	"vita-message-service/data/local"
)

type cacheMessageDao struct {
	db *gorm.DB
}

func NewCacheMessageDao(db *gorm.DB) local.CacheMessageDao {
	return &cacheMessageDao{
		db: db,
	}
}

func (cmd *cacheMessageDao) Read(message string) (entity.CacheMessage, error) {
	cacheMessage := entity.CacheMessage{}

	err := cmd.db.Where("message = ?", message).Take(&cacheMessage).Error
	if err != nil {
		return entity.CacheMessage{}, fmt.Errorf("error when get message %v from cache message", err)
	}

	return cacheMessage, nil
}

func (cmd *cacheMessageDao) Insert(message entity.CacheMessage) (entity.CacheMessage, error) {
	err := cmd.db.Create(&message).Error
	if err != nil {
		return entity.CacheMessage{}, fmt.Errorf("error when insert message %v from cache message", err)
	}
	return message, nil
}
