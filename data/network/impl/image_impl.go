package impl

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"vita-message-service/data/entity"
	"vita-message-service/data/entity/image"
	"vita-message-service/data/network"
)

type imageService struct {
	ctx    context.Context
	client *openai.Client
}

func NewImageService(client *openai.Client) network.ImageService {
	return &imageService{
		ctx:    context.Background(),
		client: client,
	}
}

func (m *imageService) Scan(message entity.Message, setting *entity.Setting) []image.Possibility {
	return []image.Possibility{}
}
