package entities

import (
	"gorm.io/gorm"
)

type OrderDetail struct {
	gorm.Model

	Order_ID    int `gorm:"column:order_id" json:"order_id"`
	Product_ID  int `gorm:"column:product_id" json:"product_id"`
	Qty         int
	Total_price int
}
