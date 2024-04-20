package usecase

import (
	"mime/multipart"
	"vita-message-service/data/entity"
)

type UploadImage interface {
	Invoke(email string, file multipart.File, header *multipart.FileHeader, message string) ([]entity.Message, error)
}
