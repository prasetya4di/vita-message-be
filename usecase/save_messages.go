package usecase

import "vita-message-service/data/entity"

type SaveMessages interface {
	Invoke(messages []entity.Message) ([]entity.Message, error)
}
