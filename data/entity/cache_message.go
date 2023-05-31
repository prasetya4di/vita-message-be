package entity

import "gorm.io/gorm"

type CacheMessage struct {
	gorm.Model
	Message     string `json:"message" gorm:"type:text; not null"`
	PrevMessage string `json:"prev_message" gorm:"type:text; not null"`
	Answer      string `json:"answer" gorm:"type:text; not null"`
	EnergyUsage uint   `json:"energy_usage" gorm:"not null"`
}

func (CacheMessage) TableName() string {
	return "cache_messages"
}
