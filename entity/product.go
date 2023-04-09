package entity

type Product struct {
	ID           int     `json:"id"`
	ProductName  string  `json:"productname"`
	ProductImage string  `json:"productimage"`
	Price        float64 `json:"price"`
	Category     string  `json:"category"`
	Description  string  `json:"description"`
}
