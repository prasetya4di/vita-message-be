package impl

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/jinzhu/gorm"
	"image"
	"log"
	"mime/multipart"
	"path/filepath"
	"vita-message-service/data/entity"
	"vita-message-service/data/local"
	constant "vita-message-service/util/const"
	time2 "vita-message-service/util/local_time"
)

type imageDao struct {
	db *gorm.DB
}

func NewImageDao(db *gorm.DB) local.ImageDao {
	return &imageDao{
		db: db,
	}
}

func (md *imageDao) Insert(email string, file multipart.File, header *multipart.FileHeader) (entity.Message, error) {
	filename := saveImage(file, header)

	message := entity.Message{
		Email:       email,
		Message:     filename,
		CreatedDate: time2.CurrentTime(),
		MessageType: constant.Send,
		FileType:    constant.Image,
	}

	err := md.db.Create(&message).Error
	if err != nil {
		return message, fmt.Errorf("add message: %v", err)
	}

	return message, nil
}

func saveImage(file multipart.File, header *multipart.FileHeader) string {
	fileExt := filepath.Ext(header.Filename)
	now := time2.CurrentTime()
	filename := fmt.Sprintf("%v", now.Unix()) + fileExt

	file.Seek(0, 0)
	imageFile, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	src := imaging.Resize(imageFile, 1000, 0, imaging.Lanczos)
	err = imaging.Save(src, fmt.Sprintf("upload/images/%v", filename))
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

	return filename
}
