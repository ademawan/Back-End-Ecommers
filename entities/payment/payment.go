package payment

import (
	"Back-End-Ecommers/entities/order"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Name     string
	Order_id []order.Order `gorm:"foreignKey:Payment_id"`
}
