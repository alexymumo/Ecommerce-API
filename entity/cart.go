package entity

type Cart struct {
	CartID  uint      `json:"id"`
	Product []Product `json:"products"`
}
