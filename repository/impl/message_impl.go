package impl

import (
	"github.com/PullRequestInc/go-gpt3"
	"vita-message-service/data/entity"
	"vita-message-service/data/local"
	"vita-message-service/data/network"
	"vita-message-service/repository"
)

type messageRepository struct {
	dao     local.MessageDao
	network network.MessageService
}

func NewMessageRepository(dao local.MessageDao, network network.MessageService) repository.MessageRepository {
	return &messageRepository{
		dao:     dao,
		network: network,
	}
}

func (mr *messageRepository) Read(email string) ([]entity.Message, error) {
	return mr.dao.Read(email)
}

func (mr *messageRepository) Insert(message entity.Message) (entity.Message, error) {
	return mr.dao.Insert(message)
}

func (mr *messageRepository) Inserts(messages []entity.Message) ([]entity.Message, error) {
	return mr.dao.Inserts(messages)
}

func (mr *messageRepository) SendMessage(message entity.Message) (*gpt3.CompletionResponse, error) {
	return mr.network.SendMessage(message)
}
