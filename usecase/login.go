package usecase

import "vita-message-service/data/entity"

type Login interface {
	Invoke(email string, password string) (*entity.User, error)
}
