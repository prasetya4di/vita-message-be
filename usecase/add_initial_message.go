package usecase

import "vita-message-service/data/entity"

type AddInitialMessage interface {
	Invoke(email string) (*entity.Message, error)
}
