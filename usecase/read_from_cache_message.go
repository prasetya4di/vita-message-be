package usecase

import "vita-message-service/data/entity"

type ReadFromCacheMessage interface {
	Invoke(message entity.Message) (entity.Message, error)
}
