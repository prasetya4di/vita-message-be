package impl

import (
	"log"
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

	message.CreatedDate = time.Now()
	message.MessageType = constant.Send
	newMessages = append(newMessages, message)

	for _, choice := range response.Choices {
		newReply := entity.Message{
			Email:       message.Email,
			Message:     choice.Text,
			CreatedDate: time.Now(),
			MessageType: constant.Reply,
		}
		newMessages = append(newMessages, newReply)
	}

	messages, err := sm.repo.Inserts(newMessages)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return messages, nil
}
