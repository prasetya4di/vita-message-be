package impl

import (
	"database/sql"
	"fmt"
	"strings"
	"vita-message-service/data/entity"
	"vita-message-service/data/local"
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
		if err := rows.Scan(&msg.ID, &msg.Email, &msg.Message, &msg.CreatedDate, &msg.MessageType); err != nil {
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
		"Insert into message (email, message, created_date, message_type) VALUES (?, ?, ?, ?)",
		message.Email,
		message.Message,
		message.CreatedDate,
		message.MessageType)
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
			"INSERT INTO message (email, message, created_date, message_type) VALUES (?, ?, ?, ?)",
			msg.Email,
			msg.Message,
			msg.CreatedDate,
			msg.MessageType)

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
