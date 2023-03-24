package local

import "vita-message-service/data/entity"

type CacheMessageDao interface {
	Read(message string) (entity.CacheMessage, error)
	Insert(message entity.CacheMessage) (entity.CacheMessage, error)
}
