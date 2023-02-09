package network

import (
	"github.com/PullRequestInc/go-gpt3"
	"vita-message-service/data/entity"
)

type MessageService interface {
	SendMessage(message entity.Message) (*gpt3.CompletionResponse, error)
}