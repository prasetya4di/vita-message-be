package network

import (
	"vita-message-service/data/entity"
)

type MessageService interface {
	SendMessages(user *entity.User, prevMessages []entity.Message, newMessage entity.Message, setting *entity.Setting) (openai.ChatCompletionResponse, error)
}
