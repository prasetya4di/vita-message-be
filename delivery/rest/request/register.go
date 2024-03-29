package request

import "time"

type RegisterRequest struct {
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	Nickname  string    `json:"nickname" binding:"required"`
	BirthDate time.Time `json:"birth_date" binding:"required"`
}
