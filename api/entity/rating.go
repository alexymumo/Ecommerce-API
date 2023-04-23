package entity

type Rating struct {
	Rate  float32 `gorm:"size:100" json:"rate"`
	Count int     `gorm:"size:100 json:count"`
}
