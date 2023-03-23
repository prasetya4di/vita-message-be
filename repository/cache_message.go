package repository

import "vita-message-service/data/entity"

type CacheMessageRepository interface {
	Read(message string) (entity.CacheMessage, error)
	Insert(message entity.CacheMessage) (entity.CacheMessage, error)
}
