package entity

import (
	"fmt"
	"time"
)

type Message struct {
	ID          int64     `json:"id"`
	Email       string    `json:"email"`
	Message     string    `json:"message"`
	CreatedDate time.Time `json:"created_date"`
	MessageType string    `json:"message_type"`
}

func (m Message) String() string {
	return fmt.Sprintf(
		"(%a, %b, %c, %d)",
		m.Email,
		m.Message,
		m.CreatedDate,
		m.MessageType,
	)
}
