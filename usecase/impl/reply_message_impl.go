package impl

import (
	"log"
	"strings"
	"vita-message-service/data/entity"
	"vita-message-service/repository"
	"vita-message-service/usecase"
	constant "vita-message-service/util/const"
	time2 "vita-message-service/util/local_time"
)

type replyMessage struct {
	messageRepository      repository.MessageRepository
	settingRepository      repository.SettingRepository
	cacheMessageRepository repository.CacheMessageRepository
}

func NewReplyMessage(messageRepository repository.MessageRepository, cacheMessageRepository repository.CacheMessageRepository, settingRepository repository.SettingRepository) usecase.ReplyMessage {
	return &replyMessage{
		messageRepository:      messageRepository,
		cacheMessageRepository: cacheMessageRepository,
		settingRepository:      settingRepository,
	}
}

func (sm *replyMessage) Invoke(user *entity.User, message entity.Message) ([]entity.Message, error) {
	setting, err := sm.settingRepository.Read()
	if err != nil {
		return nil, err
	}
	response, err := sm.messageRepository.SendMessages(user, []entity.Message{}, message, setting)
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
			Message:     choice.Message.Content,
			CreatedDate: time2.CurrentTime(),
			MessageType: constant.Reply,
			FileType:    constant.Text,
			EnergyUsage: uint(response.Usage.CompletionTokens),
		}
		newMessages = append(newMessages, newReply)
	}

	messages, err := sm.messageRepository.Inserts(newMessages)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var prevMessages []string
	delimiter := "\n" // choose your delimiter

	for _, newMessage := range newMessages {
		prevMessages = append(prevMessages, newMessage.Message)
	}

	cacheMessage := entity.CacheMessage{
		Message:     message.Message,
		PrevMessage: strings.Join(prevMessages, delimiter),
		Answer:      newMessages[len(newMessages)-1].Message,
		EnergyUsage: newMessages[len(newMessages)-1].EnergyUsage,
	}
	_, err = sm.cacheMessageRepository.Insert(cacheMessage)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
