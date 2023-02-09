package impl

import (
	"time"
	"vita-message-service/data/entity"
	"vita-message-service/repository"
	"vita-message-service/usecase"
	constant "vita-message-service/util/const"
)

type sendMessage struct {
	repo repository.MessageRepository
}

func NewSendMessage(messageRepository repository.MessageRepository) usecase.SendMessage {
	return &sendMessage{
		repo: messageRepository,
	}
}

func (sm *sendMessage) Invoke(message entity.Message) ([]entity.Message, error) {
	response, err := sm.repo.SendMessage(message)
	if err != nil {
		return nil, err
	}

	var newMessages []entity.Message
	message.MessageType = constant.Send
	newMessage, err := sm.repo.Insert(message)
	if err != nil {
		return nil, err
	}

	for _, choice := range response.Choices {
		newReply := entity.Message{
			Email:       message.Email,
			Message:     choice.Text,
			CreatedDate: time.Now(),
			MessageType: constant.Send,
		}
		newMessages = append(newMessages, newReply)
	}
	messages, err := sm.repo.Inserts(newMessages)
	if err != nil {
		return nil, err
	}

	newMessages = append(newMessages, newMessage)
	return messages, nil
}
