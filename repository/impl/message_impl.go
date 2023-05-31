package impl

import (
	"github.com/PullRequestInc/go-gpt3"
	"time"
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

func (mr *messageRepository) ReadByDate(email string, time time.Time) ([]entity.Message, error) {
	return mr.dao.ReadByDate(email, time)
}

func (mr *messageRepository) Insert(message entity.Message) (entity.Message, error) {
	return mr.dao.Insert(message)
}

func (mr *messageRepository) Inserts(messages []entity.Message) ([]entity.Message, error) {
	return mr.dao.Inserts(messages)
}

func (mr *messageRepository) SendMessages(user *entity.User, prevMessages []entity.Message, newMessage entity.Message, setting *entity.Setting) (*gpt3.ChatCompletionResponse, error) {
	return mr.network.SendMessages(user, prevMessages, newMessage, setting)
}

func (mr *messageRepository) StreamMessage(message entity.Message, onData func(response *gpt3.CompletionResponse)) error {
	return mr.network.StreamMessage(message, onData)
}

func (mr *messageRepository) BroadcastMessage(user *entity.User, messages []entity.Message) error {
	return mr.network.BroadcastMessage(user, messages)
}
