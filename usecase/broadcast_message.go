package usecase

import "vita-message-service/data/entity"

type BroadcastMessage interface {
	Invoke(user *entity.User, messages []entity.Message) error
}
