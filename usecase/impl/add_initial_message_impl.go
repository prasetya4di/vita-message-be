package impl

import (
	"vita-message-service/data/entity"
	"vita-message-service/repository"
	"vita-message-service/usecase"
	constant "vita-message-service/util/const"
	time2 "vita-message-service/util/local_time"
)

type addInitialMessage struct {
	repo repository.MessageRepository
}

func NewAddInitialMessage(messageRepository repository.MessageRepository) usecase.AddInitialMessage {
	return &addInitialMessage{messageRepository}
}

func (aim *addInitialMessage) Invoke(email string) (*entity.Message, error) {
	newMessage := entity.Message{
		Email: email,
		// Change with localization later
		Message:     "Hi my name is Vita, i'm an AI assistant, how can i help you ?",
		CreatedDate: time2.CurrentTime(),
		MessageType: constant.Reply,
		FileType:    constant.Text,
	}

	message, err := aim.repo.Insert(newMessage)
	if err != nil {
		return &entity.Message{}, err
	}

	return &message, nil
}
