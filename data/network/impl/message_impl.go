package impl

import (
	"context"
	"firebase.google.com/go/messaging"
	"github.com/sashabaranov/go-openai"
	"vita-message-service/data/entity"
	"vita-message-service/data/network"
	constant "vita-message-service/util/const"
)

type messageService struct {
	ctx      context.Context
	client   *openai.Client
	firebase *messaging.Client
}

func NewMessageService(client *openai.Client, firebase *messaging.Client) network.MessageService {
	return &messageService{
		ctx:      context.Background(),
		client:   client,
		firebase: firebase,
	}
}

func (ms *messageService) SendMessages(
	user *entity.User,
	prevMessages []entity.Message,
	newMessage entity.Message,
	setting *entity.Setting) (openai.ChatCompletionResponse, error) {
	var reqMessage []openai.ChatCompletionMessage
	reqMessage = append(reqMessage, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: setting.SystemContext,
	})
	for _, message := range prevMessages {
		if message.MessageType == constant.Reply {
			reqMessage = append(reqMessage, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleAssistant,
				Content: message.Message,
			})
		} else {
			reqMessage = append(reqMessage, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleUser,
				Content: message.Message,
			})
		}
	}
	reqMessage = append(reqMessage, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: newMessage.Message,
	})
	return ms.client.CreateChatCompletion(ms.ctx, openai.ChatCompletionRequest{
		Model:       setting.AiModel,
		Messages:    reqMessage,
		MaxTokens:   int(setting.MaxTokens),
		Temperature: setting.Temperature,
		User:        user.Nickname,
	})
}
