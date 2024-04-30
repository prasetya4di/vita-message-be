package repository

import (
	"github.com/sashabaranov/go-openai"
	"mime/multipart"
	"vita-message-service/data/entity"
)

type ImageRepository interface {
	Insert(email string, file multipart.File, header *multipart.FileHeader) (entity.Message, error)
	Scan(setting *entity.Setting, imgPath string, prompt string) (openai.ChatCompletionResponse, error)
}
