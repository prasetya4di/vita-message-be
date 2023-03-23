package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Energy struct {
	gorm.Model
	Email       string    `json:"email" gorm:"size:128; not null; index"`
	Energy      uint      `json:"energy" gorm:"not null"`
	ExpiredDate time.Time `json:"expired_date" gorm:"not null"`
}
