package impl

import (
	"vita-message-service/data/entity"
	"vita-message-service/repository"
	"vita-message-service/usecase"
)

type saveMessage struct {
	repository repository.MessageRepository
}

func NewSaveMessage(repository repository.MessageRepository) usecase.SaveMessage {
	return &saveMessage{repository: repository}
}

func (sm *saveMessage) Invoke(message entity.Message) (entity.Message, error) {
	return sm.repository.Insert(message)
}
