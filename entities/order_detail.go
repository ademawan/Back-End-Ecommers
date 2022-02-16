package entities

import (
	"gorm.io/gorm"
)

type Order_detail struct {
	gorm.Model

	Qty         int
	Total_price int
	Product_ID  int `gorm:"column:product_id" json:"product_id"`
	Product     Product
	Order_ID    int `gorm:"column:order_id" json:"order_id"`
	Order       Order
}
