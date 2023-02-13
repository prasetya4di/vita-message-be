package usecase

import "vita-message-service/data/entity"

type ReplyMessage interface {
	Invoke(message entity.Message) ([]entity.Message, error)
}
