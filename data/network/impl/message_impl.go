package impl

import (
	"context"
	"github.com/PullRequestInc/go-gpt3"
	"vita-message-service/data/entity"
	"vita-message-service/data/network"
	constant "vita-message-service/util/const"
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

func (ms *messageService) SendMessage(user *entity.User, message entity.Message) (*gpt3.CompletionResponse, error) {
	reqMessage := "Vita is an AI that help user to answer their question. " + user.Nickname + ": " + message.Message + " Vita: "
	return ms.client.Completion(ms.ctx, gpt3.CompletionRequest{
		Prompt:      []string{reqMessage},
		MaxTokens:   gpt3.IntPtr(256),
		Temperature: gpt3.Float32Ptr(0.8),
		Stop:        []string{user.Nickname + ":", "Vita:"},
	})
}

func (ms *messageService) SendMessages(user *entity.User, prevMessages []entity.Message, newMessage entity.Message) (*gpt3.ChatCompletionResponse, error) {
	var reqMessage []gpt3.ChatCompletionRequestMessage
	reqMessage = append(reqMessage, gpt3.ChatCompletionRequestMessage{
		Role:    "system",
		Content: "Vita is an AI that help user to answer their question.",
	})
	for _, message := range prevMessages {
		if message.MessageType == constant.Reply {
			reqMessage = append(reqMessage, gpt3.ChatCompletionRequestMessage{
				Role:    "assistant",
				Content: message.Message,
			})
		} else {
			reqMessage = append(reqMessage, gpt3.ChatCompletionRequestMessage{
				Role:    "user",
				Content: message.Message,
			})
		}
	}
	reqMessage = append(reqMessage, gpt3.ChatCompletionRequestMessage{
		Role:    "user",
		Content: newMessage.Message,
	})
	return ms.client.ChatCompletion(ms.ctx, gpt3.ChatCompletionRequest{
		Model:       "gpt-3.5-turbo",
		Messages:    reqMessage,
		MaxTokens:   256,
		Temperature: 0.8,
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
