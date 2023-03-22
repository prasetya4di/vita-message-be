package entity

import "github.com/jinzhu/gorm"

type CacheMessage struct {
	gorm.Model
	Message     string `gorm:"type:text; not null"`
	PrevMessage string `gorm:"type:text; not null"`
	Answer      string `gorm:"type:text; not null"`
	EnergyUsage int    `gorm:"type:uint; not null"`
}
