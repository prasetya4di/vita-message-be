package entity

import "time"

type User struct {
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Nickname  string    `json:"nickname"`
	Password  string    `json:"password"`
	BirthDate time.Time `json:"birth_date"`
}
