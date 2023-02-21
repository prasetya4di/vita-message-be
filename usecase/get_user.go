package usecase

import "vita-message-service/data/entity"

type GetUser interface {
	Invoke(email string) *entity.User
}
