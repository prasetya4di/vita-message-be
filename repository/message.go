package repository

import (
	"github.com/sashabaranov/go-openai"
	"time"
	"vita-message-service/data/entity"
)

type MessageRepository interface {
	Read(email string) ([]entity.Message, error)
	ReadByDate(email string, time time.Time) ([]entity.Message, error)
	Insert(message entity.Message) (entity.Message, error)
	Inserts(messages []entity.Message) ([]entity.Message, error)
	SendMessages(user *entity.User, prevMessages []entity.Message, newMessage entity.Message, setting *entity.Setting) (openai.ChatCompletionResponse, error)
}
