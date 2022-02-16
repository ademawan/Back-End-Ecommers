package entities

import (
	"gorm.io/gorm"
)

type Order_detail struct {
	gorm.Model

	Qty         int
	Total_price int
	Product     Product `gorm:"ForeignKey:Order_detail_ID"`
	Order       Order   `gorm:"ForeignKey:Order_detail_ID"`
}
