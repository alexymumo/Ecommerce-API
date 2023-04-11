package entity

import "time"

type User struct {
	UserID    uint32    `json:"userId"`
	Username  string    `json:"firstname"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}
