package network

import (
	"github.com/PullRequestInc/go-gpt3"
	"vita-message-service/data/entity"
)

type MessageService interface {
	SendMessage(user *entity.User, message entity.Message) (*gpt3.CompletionResponse, error)
	SendMessages(user *entity.User, prevMessages []entity.Message, newMessage entity.Message) (*gpt3.ChatCompletionResponse, error)
	StreamMessage(message entity.Message, onData func(response *gpt3.CompletionResponse)) error
	BroadcastMessage(user *entity.User, messages []entity.Message) error
}
