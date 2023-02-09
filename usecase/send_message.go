package usecase

import "vita-message-service/data/entity"

type SendMessage interface {
	Invoke(message entity.Message) ([]entity.Message, error)
}
