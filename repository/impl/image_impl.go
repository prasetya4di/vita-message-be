package impl

import (
	"mime/multipart"
	"vita-message-service/data/entity"
	"vita-message-service/data/entity/image"
	"vita-message-service/data/local"
	"vita-message-service/data/network"
	"vita-message-service/repository"
)

type imageRepository struct {
	imageDao     local.ImageDao
	imageService network.ImageService
}

func NewImageRepository(dao local.ImageDao, service network.ImageService) repository.ImageRepository {
	return &imageRepository{
		dao,
		service,
	}
}

func (mr *imageRepository) Insert(email string, file multipart.File, header *multipart.FileHeader) (entity.Message, error) {
	return mr.imageDao.Insert(email, file, header)
}

func (mr *imageRepository) Scan(message entity.Message) []image.Possibility {
	return mr.imageService.Scan(message)
}
