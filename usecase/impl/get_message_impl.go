package impl

import (
	"vita-message-service/data/entity"
	"vita-message-service/repository"
	"vita-message-service/usecase"
)

type getMessage struct {
	repo repository.MessageRepository
}

func NewGetMessage(messageRepository repository.MessageRepository) usecase.GetMessage {
	return &getMessage{
		repo: messageRepository,
	}
}

func (gm *getMessage) Invoke(email string) ([]entity.Message, error) {
	return gm.repo.Read(email)
}
