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
		return nil, fmt.Errorf("Message for email %q: %v", email, err)
	}
	defer rows.Close()
	for rows.Next() {
		var msg entity.Message
		if err := rows.Scan(&msg.ID, &msg.Email, &msg.Message, &msg.CreatedDate, &msg.MessageType); err != nil {
			return nil, fmt.Errorf("Message for email %q: %v", email, err)
		}
		messages = append(messages, msg)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Message for email %q: %v", email, err)
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
		return message, fmt.Errorf("Add message: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return message, fmt.Errorf("Add message: %v", err)
	}
	message.ID = id
	return message, nil
}

func (md *messageDao) Inserts(messages []entity.Message) ([]entity.Message, error) {
	var newMessages []string

	for _, msg := range messages {
		newMessages = append(newMessages, msg.String())
	}

	rows, err := md.db.Query(
		"INSERT INTO message (email, message, created_date, message_type) VALUES ?",
		strings.Join(newMessages[:], ","))
	if err != nil {
		return messages, fmt.Errorf("Add message: %v", err)
	}

	defer rows.Close()

	var insertedMessages []entity.Message
	for rows.Next() {
		var msg entity.Message
		if err := rows.Scan(&msg.ID, &msg.Email, &msg.Message, &msg.CreatedDate, &msg.MessageType); err != nil {
			return messages, fmt.Errorf("Add message: %v", err)
		}
		insertedMessages = append(insertedMessages, msg)
	}
	if err := rows.Err(); err != nil {
		return messages, fmt.Errorf("Add message: %v", err)
	}

	return insertedMessages, nil
}
