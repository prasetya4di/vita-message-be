package repository

import (
	"mime/multipart"
	"vita-message-service/data/entity"
	"vita-message-service/data/entity/image"
)

type ImageRepository interface {
	Insert(email string, file multipart.File, header *multipart.FileHeader) (entity.Message, error)
	Scan(message entity.Message, setting *entity.Setting) []image.Possibility
}
