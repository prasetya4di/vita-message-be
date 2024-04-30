package network

import (
	"github.com/sashabaranov/go-openai"
	"vita-message-service/data/entity"
)

type ImageService interface {
	Scan(setting *entity.Setting, imgPath string, prompt string) (openai.ChatCompletionResponse, error)
}
