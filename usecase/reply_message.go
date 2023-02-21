package usecase

import "vita-message-service/data/entity"

type ReplyMessage interface {
	Invoke(user *entity.User, message entity.Message) ([]entity.Message, error)
}
