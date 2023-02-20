package local

import (
	"mime/multipart"
	"time"
	"vita-message-service/data/entity"
)

type MessageDao interface {
	Read(email string) ([]entity.Message, error)
	ReadByDate(email string, time time.Time) ([]entity.Message, error)
	Insert(message entity.Message) (entity.Message, error)
	Inserts(messages []entity.Message) ([]entity.Message, error)
	SaveImage(file multipart.File, header *multipart.FileHeader) string
}
