package impl

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/sirupsen/logrus"
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

	path := fmt.Sprintf("upload/images/%v", filename)

	file.Seek(0, 0)
	x, err := exif.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	orient, _ := x.Get(exif.Orientation)

	file.Seek(0, 0)
	imageFile, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	img := reverseOrientation(imageFile, orient.String())

	err = imaging.Save(img, path)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

	return filename
}

func reverseOrientation(img image.Image, o string) *image.NRGBA {
	switch o {
	case "1":
		return imaging.Clone(img)
	case "2":
		return imaging.FlipV(img)
	case "3":
		return imaging.Rotate180(img)
	case "4":
		return imaging.Rotate180(imaging.FlipV(img))
	case "5":
		return imaging.Rotate270(imaging.FlipV(img))
	case "6":
		return imaging.Rotate270(img)
	case "7":
		return imaging.Rotate90(imaging.FlipV(img))
	case "8":
		return imaging.Rotate90(img)
	}
	logrus.Errorf("unknown orientation %s, expect 1-8", o)
	return imaging.Clone(img)
}
