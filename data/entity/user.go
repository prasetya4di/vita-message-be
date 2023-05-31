package entity

import (
	"time"
)

type User struct {
	Email     string    `json:"email" gorm:"primaryKey"`
	FirstName string    `json:"first_name" gorm:"size:50; not null"`
	LastName  string    `json:"last_name" gorm:"size:50; not null"`
	Nickname  string    `json:"nickname" gorm:"size:50; not null"`
	Password  string    `json:"password" gorm:"size:128; not null"`
	BirthDate time.Time `json:"birth_date" gorm:"not null"`
	Message   []Message `json:"messages" gorm:"foreignKey:email; references:email"`
	Energy    Energy    `json:"energy" gorm:"foreignKey:email; references:email"`
}

func (User) TableName() string {
	return "users"
}
