package entity

import (
	"gorm.io/gorm"
	"time"
)

type Energy struct {
	gorm.Model
	Email       string    `json:"email" gorm:"size:128; not null"`
	Energy      uint      `json:"energy" gorm:"not null"`
	ExpiredDate time.Time `json:"expired_date" gorm:"not null"`
}

func (Energy) TableName() string {
	return "energies"
}
