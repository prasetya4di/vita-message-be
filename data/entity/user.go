package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email     string    `json:"email" gorm:"size:128; not null; unique"`
	FirstName string    `json:"first_name" gorm:"size:50; not null"`
	LastName  string    `json:"last_name" gorm:"size:50; not null"`
	Nickname  string    `json:"nickname" gorm:"size:50; not null"`
	Password  string    `json:"password" gorm:"size:128; not null"`
	BirthDate time.Time `json:"birth_date" gorm:"not null"`
	Message   []Message `gorm:"foreignKey:email;references:email"`
}
