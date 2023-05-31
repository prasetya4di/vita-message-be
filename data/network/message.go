package network

import (
	"github.com/PullRequestInc/go-gpt3"
	"vita-message-service/data/entity"
)

type MessageService interface {
	SendMessages(user *entity.User, prevMessages []entity.Message, newMessage entity.Message, setting entity.Setting) (*gpt3.ChatCompletionResponse, error)
	StreamMessage(message entity.Message, onData func(response *gpt3.CompletionResponse)) error
	BroadcastMessage(user *entity.User, messages []entity.Message) error
}
