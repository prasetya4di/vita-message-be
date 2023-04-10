package response

import "time"

type User struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Nickname  string    `json:"nickname"`
	BirthDate time.Time `json:"birth_date"`
	Token     string    `json:"token"`
}
