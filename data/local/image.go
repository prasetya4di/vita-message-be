package local

import (
	"mime/multipart"
	"vita-message-service/data/entity"
)

type ImageDao interface {
	Insert(email string, file multipart.File, header *multipart.FileHeader) (entity.Message, error)
}
