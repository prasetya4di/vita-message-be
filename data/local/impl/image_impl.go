package impl

import (
	"fmt"
	"github.com/disintegration/imaging"
	"gorm.io/gorm"
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

// Todo: optimize image upload
func saveImage(file multipart.File, header *multipart.FileHeader) string {
	fileExt := filepath.Ext(header.Filename)
	now := time2.CurrentTime()
	filename := fmt.Sprintf("%v", now.Unix()) + fileExt

	_, err := file.Seek(0, 0)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
	imageFile, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	err = imaging.Save(imageFile, fmt.Sprintf("upload/images/%v", filename))
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

	return filename
}
