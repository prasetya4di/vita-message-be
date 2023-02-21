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

func (sm *sendMessage) Invoke(user *entity.User, message entity.Message) ([]entity.Message, error) {
	createdDate := time.Now()
	prevMessage, err := sm.repo.ReadByDate(message.Email, createdDate)
	if err != nil {
		return nil, err
	}
	response, err := sm.repo.SendMessages(user, prevMessage, message)
	if err != nil {
		return nil, err
	}

	var newMessages []entity.Message

	message.CreatedDate = createdDate
	message.MessageType = constant.Send
	message.FileType = constant.Text
	newMessages = append(newMessages, message)

	for _, choice := range response.Choices {
		newReply := entity.Message{
			Email:       message.Email,
			Message:     choice.Text,
			CreatedDate: time.Now(),
			MessageType: constant.Reply,
			FileType:    constant.Text,
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
