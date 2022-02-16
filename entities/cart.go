package entities

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model

	Qty         int
	Total_price int
	User_ID     int `gorm:"column:user_id" json:"user_id"`
	User        User
	Product_ID  int `gorm:"column:product_id" json:"product_id"`
	Product     Product
}
