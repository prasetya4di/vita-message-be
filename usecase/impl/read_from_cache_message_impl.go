package impl

import (
	"vita-message-service/data/entity"
	"vita-message-service/repository"
	"vita-message-service/usecase"
	constant "vita-message-service/util/const"
	"vita-message-service/util/local_time"
)

type readFromCacheMessage struct {
	repository  repository.CacheMessageRepository
	messageRepo repository.MessageRepository
}

func NewReadFromCacheMessage(repository repository.CacheMessageRepository, messageRepository repository.MessageRepository) usecase.ReadFromCacheMessage {
	return &readFromCacheMessage{
		repository:  repository,
		messageRepo: messageRepository,
	}
}

func (rf *readFromCacheMessage) Invoke(message entity.Message) (entity.Message, error) {
	cacheMessage, err := rf.repository.Read(message.Message)
	if err != nil {
		return entity.Message{}, err
	}

	//Todo: Check previous message and detect the context before return message

	return entity.Message{
		Email:       message.Email,
		Message:     cacheMessage.Answer,
		CreatedDate: local_time.CurrentTime(),
		MessageType: constant.Reply,
		FileType:    constant.Text,
		EnergyUsage: cacheMessage.EnergyUsage,
	}, nil
}
