package impl

import (
	"log"
	"mime/multipart"
	"vita-message-service/data/entity/image"
	"vita-message-service/repository"
	"vita-message-service/usecase"
)

type uploadImage struct {
	repo repository.ImageRepository
}

func NewUploadImage(repo repository.ImageRepository) usecase.UploadImage {
	return &uploadImage{
		repo,
	}
}

func (sm *uploadImage) Invoke(email string, file multipart.File, header *multipart.FileHeader) ([]image.Possibility, error) {
	message, err := sm.repo.Insert(email, file, header)
	if err != nil {
		log.Fatalf("error insert image: %v", err)
		return nil, err
	}

	result := sm.repo.Scan(message)
	return result, nil
}
