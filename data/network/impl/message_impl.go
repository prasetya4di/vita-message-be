package impl

import (
	"context"
	"fmt"
	"github.com/polds/imgbase64"
	"github.com/sashabaranov/go-openai"
	"log"
	"vita-message-service/data/entity"
	"vita-message-service/data/network"
	constant "vita-message-service/util/const"
)

type messageService struct {
	ctx    context.Context
	client *openai.Client
}

func NewMessageService(client *openai.Client) network.MessageService {
	return &messageService{
		ctx:    context.Background(),
		client: client,
	}
}

func (ms *messageService) SendMessages(
	user *entity.User,
	prevMessages []entity.Message,
	newMessages []entity.Message,
	setting *entity.Setting) (openai.ChatCompletionResponse, error) {
	var reqMessage []openai.ChatCompletionMessage
	reqMessage = append(reqMessage, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: setting.SystemContext,
	})

	var prevMessage entity.Message

	for _, message := range prevMessages {
		if message.MessageType == constant.Reply {
			reqMessage = append(reqMessage, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleAssistant,
				Content: message.Message,
			})
		} else if message.MessageType == constant.Image {
			prevMessage = message
		} else {
			reqMessage = append(reqMessage, messageToOpenAiMessage(prevMessage, message))
			prevMessage = message
		}
	}

	for _, message := range newMessages {
		if message.MessageType == constant.Image {
			prevMessage = message
		} else {
			reqMessage = append(reqMessage, messageToOpenAiMessage(prevMessage, message))
		}
	}

	return ms.client.CreateChatCompletion(ms.ctx, openai.ChatCompletionRequest{
		Model:       setting.AiModel,
		Messages:    reqMessage,
		MaxTokens:   int(setting.MaxTokens),
		Temperature: setting.Temperature,
		User:        user.Nickname,
	})
}

func messageToOpenAiMessage(prevMessage entity.Message, newMessage entity.Message) openai.ChatCompletionMessage {
	if prevMessage.FileType == constant.Image {
		imgPath := prevMessage.Message
		imgData, err := imgbase64.FromLocal(fmt.Sprintf("upload/images/%s", imgPath))
		if err != nil {
			log.Fatalf("error convert image to base64: %v", err)
			return openai.ChatCompletionMessage{}
		}
		return openai.ChatCompletionMessage{
			Role: openai.ChatMessageRoleUser,
			MultiContent: []openai.ChatMessagePart{
				{
					Type: openai.ChatMessagePartTypeImageURL,
					ImageURL: &openai.ChatMessageImageURL{
						URL:    imgData,
						Detail: openai.ImageURLDetailLow,
					},
				},
				{
					Type: openai.ChatMessagePartTypeText,
					Text: newMessage.Message,
				},
			},
		}
	} else {
		return openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: newMessage.Message,
		}
	}
}
