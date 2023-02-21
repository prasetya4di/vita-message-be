package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email     string    `json:"email" gorm:"size:128; not null; unique"`
	FirstName string    `json:"first_name" gorm:"size:50; not null; unique"`
	LastName  string    `json:"last_name" gorm:"size:50; not null; unique"`
	Nickname  string    `json:"nickname" gorm:"size:50; not null; unique"`
	Password  string    `json:"password" gorm:"size:128; not null; unique"`
	BirthDate time.Time `json:"birth_date" gorm:"not null"`
}
