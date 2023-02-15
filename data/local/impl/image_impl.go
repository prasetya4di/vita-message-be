package impl

import (
	"database/sql"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"log"
	"mime/multipart"
	"path/filepath"
	"time"
	"vita-message-service/data/entity"
	"vita-message-service/data/local"
	constant "vita-message-service/util/const"
)

type imageDao struct {
	db *sql.DB
}

func NewImageDao(db *sql.DB) local.ImageDao {
	return &imageDao{
		db: db,
	}
}

func (md *imageDao) Insert(email string, file multipart.File, header *multipart.FileHeader) (entity.Message, error) {
	filename := saveImage(file, header)

	message := entity.Message{
		Email:       email,
		Message:     filename,
		CreatedDate: time.Now(),
		MessageType: constant.Send,
		FileType:    constant.IMAGE,
	}

	result, err := md.db.Exec(
		"Insert into message (email, message, created_date, message_type, file_type) VALUES (?, ?, ?, ?, ?)",
		message.Email,
		message.Message,
		message.CreatedDate,
		message.MessageType,
		message.FileType)

	if err != nil {
		return message, fmt.Errorf("add message: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return message, fmt.Errorf("add message: %v", err)
	}
	message.ID = id
	return message, nil
}

func saveImage(file multipart.File, header *multipart.FileHeader) string {
	fileExt := filepath.Ext(header.Filename)
	now := time.Now()
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
