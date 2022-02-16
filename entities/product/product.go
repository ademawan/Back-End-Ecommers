package product

import (
	carts "Back-End-Ecommers/entities/cart"
	"Back-End-Ecommers/entities/order_detail"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name         string `gorm:"unique"`
	Price        int
	Qty          int
	Description  string
	Category_id  int
	Cart         []carts.Cart                `gorm:"foreignKey:Product_id"`
	Order_detail []order_detail.Order_detail `gorm:"foreignKey:Product_id"`
}
