package entity

import (
	"time"
)

type Cart struct {
	UserID    uint       `gorm:"primary_key; auto_increment" json:"id"`
	CartItems []CartItem `json:"carts"`
}

type CartItem struct {
	CartID    uint      `gorm:"primary_key; auto_increment" json:"cartid"`
	ProductID uint      `gorm:"size:255" json:"productid"`
	Total     uint      `gorm:"size:255" json:"total"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (c *Cart) Prepare() {
	//c.ID = 0
	//c.Product = []Product{}
	//c.Total = html.EscapeString(strings.TrimSpace(c.Total))
}
