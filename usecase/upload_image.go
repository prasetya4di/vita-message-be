package usecase

import (
	"mime/multipart"
	"vita-message-service/data/entity"
)

type UploadImage interface {
	Invoke(user *entity.User, file multipart.File, header *multipart.FileHeader, message string) ([]entity.Message, error)
}
