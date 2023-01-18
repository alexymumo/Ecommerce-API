package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"firstname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_At"`
	UpdatedAt time.Time `json:"updated_At"`
}
