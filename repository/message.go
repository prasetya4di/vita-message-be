package repository

import (
	"github.com/PullRequestInc/go-gpt3"
	"time"
	"vita-message-service/data/entity"
)

type MessageRepository interface {
	Read(email string) ([]entity.Message, error)
	ReadByDate(email string, time time.Time) ([]entity.Message, error)
	Insert(message entity.Message) (entity.Message, error)
	Inserts(messages []entity.Message) ([]entity.Message, error)
	SendMessage(user *entity.User, message entity.Message) (*gpt3.CompletionResponse, error)
	SendMessages(user *entity.User, prevMessages []entity.Message, newMessage entity.Message) (*gpt3.ChatCompletionResponse, error)
	StreamMessage(message entity.Message, onData func(response *gpt3.CompletionResponse)) error
}
