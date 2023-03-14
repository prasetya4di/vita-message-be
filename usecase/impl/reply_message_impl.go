package impl

import (
	"log"
	"vita-message-service/data/entity"
	"vita-message-service/repository"
	"vita-message-service/usecase"
	constant "vita-message-service/util/const"
	time2 "vita-message-service/util/local_time"
)

type replyMessage struct {
	repo repository.MessageRepository
}

func NewReplyMessage(messageRepository repository.MessageRepository) usecase.ReplyMessage {
	return &replyMessage{
		repo: messageRepository,
	}
}

func (sm *replyMessage) Invoke(user *entity.User, message entity.Message) ([]entity.Message, error) {
	response, err := sm.repo.SendMessage(user, message)
	if err != nil {
		return nil, err
	}

	var newMessages []entity.Message

	message.CreatedDate = time2.CurrentTime()
	message.MessageType = constant.Reply
	message.FileType = constant.Text
	newMessages = append(newMessages, message)

	for _, choice := range response.Choices {
		newReply := entity.Message{
			Email:       message.Email,
			Message:     choice.Text,
			CreatedDate: time2.CurrentTime(),
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
