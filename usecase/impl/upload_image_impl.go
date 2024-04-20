package impl

import (
	"log"
	"mime/multipart"
	"vita-message-service/data/entity"
	"vita-message-service/data/entity/image"
	"vita-message-service/repository"
	"vita-message-service/usecase"
)

type uploadImage struct {
	imageRepository   repository.ImageRepository
	settingRepository repository.SettingRepository
}

func NewUploadImage(imageRepository repository.ImageRepository, settingRepository repository.SettingRepository) usecase.UploadImage {
	return &uploadImage{
		imageRepository,
		settingRepository,
	}
}

func (sm *uploadImage) Invoke(email string, file multipart.File, header *multipart.FileHeader) (image.Scan, error) {
	message, err := sm.imageRepository.Insert(email, file, header)
	if err != nil {
		log.Fatalf("error insert image: %v", err)
		return image.Scan{}, err
	}

	setting, err := sm.settingRepository.Read()
	if err != nil {
		return image.Scan{}, err
	}

	result := sm.imageRepository.Scan(message, setting)
	return image.Scan{
		Messages: []entity.Message{message}, Possibilities: result,
	}, nil
}
