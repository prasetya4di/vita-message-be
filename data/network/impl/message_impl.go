package impl

import (
	"context"
	"github.com/PullRequestInc/go-gpt3"
	"vita-message-service/data/entity"
	"vita-message-service/data/network"
)

type messageService struct {
	ctx    context.Context
	client gpt3.Client
}

func NewMessageService(client gpt3.Client) network.MessageService {
	return &messageService{
		ctx:    context.Background(),
		client: client,
	}
}

func (ms *messageService) SendMessage(message entity.Message) (*gpt3.CompletionResponse, error) {
	return ms.client.Completion(ms.ctx, gpt3.CompletionRequest{
		Prompt:      []string{message.Message},
		MaxTokens:   gpt3.IntPtr(256),
		Temperature: gpt3.Float32Ptr(0.8),
	})
}

func (ms *messageService) StreamMessage(message entity.Message, onData func(response *gpt3.CompletionResponse)) error {
	return ms.client.CompletionStream(
		ms.ctx,
		gpt3.CompletionRequest{
			Prompt:      []string{message.Message},
			MaxTokens:   gpt3.IntPtr(256),
			Temperature: gpt3.Float32Ptr(0.8),
		},
		onData)
}
