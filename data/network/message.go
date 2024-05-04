package network

import (
	"github.com/sashabaranov/go-openai"
	"vita-message-service/data/entity"
)

type MessageService interface {
	SendMessages(user *entity.User, prevMessages []entity.Message, newMessages []entity.Message, setting *entity.Setting) (openai.ChatCompletionResponse, error)
}
