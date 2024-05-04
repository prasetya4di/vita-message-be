package impl

import (
	"log"
	"mime/multipart"
	"vita-message-service/data/entity"
	"vita-message-service/repository"
	"vita-message-service/usecase"
	constant "vita-message-service/util/const"
	time2 "vita-message-service/util/local_time"
)

type uploadImage struct {
	imageRepository   repository.ImageRepository
	settingRepository repository.SettingRepository
	messageRepository repository.MessageRepository
}

func NewUploadImage(imageRepository repository.ImageRepository, settingRepository repository.SettingRepository, messageRepository repository.MessageRepository) usecase.UploadImage {
	return &uploadImage{
		imageRepository,
		settingRepository,
		messageRepository,
	}
}

func (sm *uploadImage) Invoke(user *entity.User, file multipart.File, header *multipart.FileHeader, message string) ([]entity.Message, error) {
	var newMessages []entity.Message
	createdDate := time2.CurrentTime()
	imgMessage, err := sm.imageRepository.Insert(user.Email, file, header)
	if err != nil {
		log.Fatalf("error insert image: %v", err)
		return nil, err
	}

	prompt, err := sm.messageRepository.Insert(entity.Message{
		Message:     message,
		Email:       user.Email,
		MessageType: constant.Send,
		FileType:    constant.Text,
		CreatedDate: time2.CurrentTime(),
	})

	setting, err := sm.settingRepository.Read()
	if err != nil {
		log.Fatalf("error read setting: %v", err)
		return nil, err
	}

	prevMessage, err := sm.messageRepository.ReadByDate(user.Email, createdDate)
	if err != nil {
		return nil, err
	}

	scanResult, err := sm.messageRepository.SendMessages(user, prevMessage, []entity.Message{
		imgMessage, prompt,
	}, setting)
	if err != nil {
		return nil, err
	}

	for _, choice := range scanResult.Choices {
		newReply := entity.Message{
			Email:       imgMessage.Email,
			Message:     choice.Message.Content,
			CreatedDate: createdDate,
			MessageType: constant.Reply,
			FileType:    constant.Text,
			EnergyUsage: uint(scanResult.Usage.CompletionTokens),
		}
		newMessages = append(newMessages, newReply)
	}

	newMessage, err := sm.messageRepository.Inserts(newMessages)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return append([]entity.Message{
		imgMessage, prompt,
	}, newMessage...), nil
}
