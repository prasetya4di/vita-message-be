package impl

import (
	"log"
	"vita-message-service/data/entity"
	"vita-message-service/repository"
	"vita-message-service/usecase"
	constant "vita-message-service/util/const"
	time2 "vita-message-service/util/local_time"
)

type sendMessage struct {
	messageRepository repository.MessageRepository
	settingRepository repository.SettingRepository
}

func NewSendMessage(messageRepository repository.MessageRepository, settingRepository repository.SettingRepository) usecase.SendMessage {
	return &sendMessage{
		messageRepository: messageRepository,
		settingRepository: settingRepository,
	}
}

func (sm *sendMessage) Invoke(user *entity.User, message entity.Message) ([]entity.Message, error) {
	createdDate := time2.CurrentTime()
	setting, err := sm.settingRepository.Read()
	if err != nil {
		return nil, err
	}
	prevMessage, err := sm.messageRepository.ReadByDate(message.Email, createdDate)
	if err != nil {
		return nil, err
	}
	response, err := sm.messageRepository.SendMessages(user, prevMessage, []entity.Message{message}, setting)
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
			Message:     choice.Message.Content,
			CreatedDate: createdDate,
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

	for _, newMessage := range newMessages {
		prevMessages = append(prevMessages, newMessage.Message)
	}

	return messages, nil
}
