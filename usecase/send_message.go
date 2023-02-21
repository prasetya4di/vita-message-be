package usecase

import "vita-message-service/data/entity"

type SendMessage interface {
	Invoke(user *entity.User, message entity.Message) ([]entity.Message, error)
}
