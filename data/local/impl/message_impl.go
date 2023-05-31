package impl

import (
	"fmt"
	"github.com/disintegration/imaging"
	"gorm.io/gorm"
	"image"
	"log"
	"mime/multipart"
	"path/filepath"
	"time"
	"vita-message-service/data/entity"
	"vita-message-service/data/local"
	constant "vita-message-service/util/const"
	time2 "vita-message-service/util/local_time"
)

type messageDao struct {
	db *gorm.DB
}

func NewMessageDao(db *gorm.DB) local.MessageDao {
	return &messageDao{
		db: db,
	}
}

func (md *messageDao) Read(email string) ([]entity.Message, error) {
	var messages []entity.Message

	err := md.db.Where("email = ?", email).Find(&messages).Error
	if err != nil {
		return nil, fmt.Errorf("message for email %q: %v", email, err)
	}

	return messages, nil
}

func (md *messageDao) ReadByDate(email string, time2 time.Time) ([]entity.Message, error) {
	var messages []entity.Message

	err := md.db.Where("email = ? and created_date >= ? and file_type = ?", email, time2.Add(-time.Hour*1), constant.Text).Error
	if err != nil {
		return nil, fmt.Errorf("message for email %q: %v", email, err)
	}

	return messages, nil
}

func (md *messageDao) Insert(message entity.Message) (entity.Message, error) {
	err := md.db.Create(&message).Error
	if err != nil {
		return message, fmt.Errorf("add message: %v", err)
	}
	return message, nil
}

func (md *messageDao) Inserts(messages []entity.Message) ([]entity.Message, error) {
	tx := md.db.Begin()

	var newMessages []entity.Message
	for _, msg := range messages {
		if err := tx.Create(&msg).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		newMessages = append(newMessages, msg)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return newMessages, nil
}

func (md *messageDao) SaveImage(file multipart.File, header *multipart.FileHeader) string {
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
