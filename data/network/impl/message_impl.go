package impl

import (
	"context"
	"encoding/json"
	"firebase.google.com/go/messaging"
	"github.com/PullRequestInc/go-gpt3"
	"log"
	"strconv"
	"vita-message-service/data/entity"
	"vita-message-service/data/network"
	constant "vita-message-service/util/const"
)

type messageService struct {
	ctx      context.Context
	client   gpt3.Client
	firebase *messaging.Client
}

func NewMessageService(client gpt3.Client, firebase *messaging.Client) network.MessageService {
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
	setting *entity.Setting) (*gpt3.ChatCompletionResponse, error) {
	var reqMessage []gpt3.ChatCompletionRequestMessage
	reqMessage = append(reqMessage, gpt3.ChatCompletionRequestMessage{
		Role:    "system",
		Content: setting.SystemContent,
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
		Model:       setting.AiModel,
		Messages:    reqMessage,
		MaxTokens:   int(setting.MaxTokens),
		Temperature: gpt3.Float32Ptr(setting.Temperature),
		User:        user.Nickname,
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

func (ms *messageService) BroadcastMessage(user *entity.User, messages []entity.Message) error {
	chatMessage, err := json.Marshal(messages)
	if err != nil {
		log.Fatalf("error marshaling chat message: %v\n", err)
	}

	message := &messaging.Message{
		Data: map[string]string{
			"type": "chat",
			"data": string(chatMessage),
		},
		Notification: &messaging.Notification{
			Title: "New Message From Vita",
			Body:  messages[len(messages)-1].Message,
		},
		Topic: user.Nickname + strconv.Itoa(int(user.ID)),
	}

	_, err = ms.firebase.Send(ms.ctx, message)
	if err != nil {
		log.Fatalf("error broadcasting chat message: %v\n", err)
		return err
	}
	log.Println("Message sended")
	return nil
}
