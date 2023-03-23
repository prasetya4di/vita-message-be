package impl

import (
	"vita-message-service/data/entity"
	"vita-message-service/repository"
	"vita-message-service/usecase"
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

func (rf *readFromCacheMessage) Invoke(message entity.Message) (string, error) {
	cacheMessage, err := rf.repository.Read(message.Message)
	if err != nil {
		return "", err
	}

	//Todo: Check previous message and detect the context before return message

	return cacheMessage.Answer, nil
}
