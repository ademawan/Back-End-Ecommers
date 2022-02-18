package entities

import (
	"gorm.io/gorm"
)

type OrderDetail struct {
	gorm.Model

	Order_ID    int `gorm:"column:order_id" json:"order_id"`
	Order       Order
	Product_ID  int `gorm:"column:product_id" json:"product_id"`
	Product     Product
	Qty         int
	Total_price int
}
