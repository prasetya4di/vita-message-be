package usecase

import "vita-message-service/data/entity"

type GetMessage interface {
	Invoke(email string) ([]entity.Message, error)
}
