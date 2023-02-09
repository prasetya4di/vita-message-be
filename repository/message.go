package repository

import (
	"github.com/PullRequestInc/go-gpt3"
	"vita-message-service/data/entity"
)

type MessageRepository interface {
	Read(email string) ([]entity.Message, error)
	Insert(message entity.Message) (entity.Message, error)
	Inserts(messages []entity.Message) ([]entity.Message, error)
	SendMessage(message entity.Message) (*gpt3.CompletionResponse, error)
}