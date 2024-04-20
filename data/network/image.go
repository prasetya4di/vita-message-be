package network

import (
	"vita-message-service/data/entity"
	"vita-message-service/data/entity/image"
)

type ImageService interface {
	Scan(message entity.Message, setting *entity.Setting) []image.Possibility
}
