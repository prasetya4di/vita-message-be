package impl

import (
	"database/sql"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"log"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"
	"vita-message-service/data/entity"
	"vita-message-service/data/local"
	constant "vita-message-service/util/const"
)

type messageDao struct {
	db *sql.DB
}

func NewMessageDao(db *sql.DB) local.MessageDao {
	return &messageDao{
		db: db,
	}
}

func (md *messageDao) Read(email string) ([]entity.Message, error) {
	var messages []entity.Message

	rows, err := md.db.Query("SELECT * from message where email = ?", email)
	if err != nil {
		return nil, fmt.Errorf("message for email %q: %v", email, err)
	}
	defer rows.Close()
	for rows.Next() {
		var msg entity.Message
		if err := rows.Scan(&msg.ID, &msg.Email, &msg.Message, &msg.CreatedDate, &msg.MessageType, &msg.FileType); err != nil {
			return nil, fmt.Errorf("message for email %q: %v", email, err)
		}
		messages = append(messages, msg)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("message for email %q: %v", email, err)
	}
	return messages, nil
}

func (md *messageDao) ReadByDate(email string, time2 time.Time) ([]entity.Message, error) {
	var messages []entity.Message
	rows, err := md.db.Query("SELECT * from message where email = ? and created_date >= ? and file_type = ?", email, time2.Add(-time.Hour*1), constant.Text)
	if err != nil {
		return nil, fmt.Errorf("message for email %q: %v", email, err)
	}
	defer rows.Close()
	for rows.Next() {
		var msg entity.Message
		if err := rows.Scan(&msg.ID, &msg.Email, &msg.Message, &msg.CreatedDate, &msg.MessageType, &msg.FileType); err != nil {
			return nil, fmt.Errorf("message for email %q: %v", email, err)
		}
		messages = append(messages, msg)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("message for email %q: %v", email, err)
	}
	return messages, nil
}

func (md *messageDao) Insert(message entity.Message) (entity.Message, error) {
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

func (md *messageDao) Inserts(messages []entity.Message) ([]entity.Message, error) {
	var insertedMessages []entity.Message
	tx, _ := md.db.Begin()

	for _, msg := range messages {
		msg.Message = strings.TrimSpace(msg.Message)
		result, err := tx.Exec(
			"INSERT INTO message (email, message, created_date, message_type, file_type) VALUES (?, ?, ?, ?, ?)",
			msg.Email,
			msg.Message,
			msg.CreatedDate,
			msg.MessageType,
			msg.FileType)

		if err != nil {
			tx.Rollback()
			return nil, err
		}
		msg.ID, _ = result.LastInsertId()
		insertedMessages = append(insertedMessages, msg)
	}

	err := tx.Commit()
	if err != nil {
		return nil, err
	}
	return insertedMessages, nil
}

func (md *messageDao) SaveImage(file multipart.File, header *multipart.FileHeader) string {
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
