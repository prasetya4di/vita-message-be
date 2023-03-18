package entity

import (
	"time"
)

type Message struct {
	ID          int64     `json:"id" gorm:"primaryKey"`
	Email       string    `json:"email" gorm:"size:128; not null"`
	Message     string    `json:"message" gorm:"type:text"`
	CreatedDate time.Time `json:"created_date" gorm:"not null"`
	MessageType string    `json:"message_type" gorm:"size:5; not null"`
	FileType    string    `json:"file_type" gorm:"size:5; not null"`
}
