package order

import (
	"Back-End-Ecommers/entities/order_detail"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	User_id      int
	Payment_id   int
	Order_detail []order_detail.Order_detail `gorm:"foreignKey:Order_id"`
}
