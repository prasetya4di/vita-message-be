package impl

import (
	"vita-message-service/data/entity"
	"vita-message-service/repository"
	"vita-message-service/usecase"
)

type saveMessages struct {
	repository repository.MessageRepository
}

func NewSaveMessages(repository repository.MessageRepository) usecase.SaveMessages {
	return &saveMessages{repository: repository}
}

func (sm *saveMessages) Invoke(messages []entity.Message) ([]entity.Message, error) {
	return sm.repository.Inserts(messages)
}
