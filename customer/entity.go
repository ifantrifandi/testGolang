package customer

import "time"

type Customer struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Occupation string `json:"occupation"`
	PasswordHash string `json:"password"`
	Role string `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}