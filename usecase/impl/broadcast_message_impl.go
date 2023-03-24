package impl

import (
	"vita-message-service/data/entity"
	"vita-message-service/repository"
	"vita-message-service/usecase"
)

type broadcastMessage struct {
	repo repository.MessageRepository
}

func NewBroadcastMessage(repo repository.MessageRepository) usecase.BroadcastMessage {
	return &broadcastMessage{repo: repo}
}

func (bm *broadcastMessage) Invoke(user *entity.User, messages []entity.Message) error {
	return bm.repo.BroadcastMessage(user, messages)
}
