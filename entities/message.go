package entities

import "time"

type Message struct {
	Email       string    `json:"email"`
	Message     string    `json:"message"`
	CreatedDate time.Time `json:"created_date"`
	MessageType string    `json:"message_type"`
}
