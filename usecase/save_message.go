package usecase

import "vita-message-service/data/entity"

type SaveMessage interface {
	Invoke(message entity.Message) (entity.Message, error)
}
