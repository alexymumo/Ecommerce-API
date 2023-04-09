package entity

type User struct {
	ID       int    `json:"id"`
	Username string `json:"firstname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
