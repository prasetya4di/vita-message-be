package impl

import (
	"context"
	"fmt"
	"github.com/polds/imgbase64"
	"github.com/sashabaranov/go-openai"
	"log"
	"vita-message-service/data/entity"
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

func (m *imageService) Scan(message entity.Message, setting *entity.Setting, imgPath string, prompt string) (openai.ChatCompletionResponse, error) {
	imgUrl, err := imgbase64.FromLocal(fmt.Sprintf("upload/images/%s", imgPath))
	if err != nil {
		log.Fatalf("error convert image to base64: %v", err)
		return openai.ChatCompletionResponse{}, err
	}

	return m.client.CreateChatCompletion(m.ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4VisionPreview,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: setting.SystemContext,
			},
			{
				Role: openai.ChatMessageRoleUser,
				MultiContent: []openai.ChatMessagePart{
					{
						Type: openai.ChatMessagePartTypeImageURL,
						ImageURL: &openai.ChatMessageImageURL{
							URL:    imgUrl,
							Detail: openai.ImageURLDetailLow,
						},
					},
					{
						Type: openai.ChatMessagePartTypeText,
						Text: prompt,
					},
				},
			},
		},
		MaxTokens:   int(setting.MaxTokens),
		Temperature: setting.Temperature,
	})
}
