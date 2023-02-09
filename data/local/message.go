package local

import "vita-message-service/data/entity"

type MessageDao interface {
	Read(email string) ([]entity.Message, error)
	Insert(message entity.Message) (entity.Message, error)
	Inserts(messages []entity.Message) ([]entity.Message, error)
}
