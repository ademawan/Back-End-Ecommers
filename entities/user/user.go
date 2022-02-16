package user

import (
	carts "Back-End-Ecommers/entities/cart"
	"Back-End-Ecommers/entities/order"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Status   string
	Address  string
	Order    []order.Order `gorm:"foreignKey:User_id"`
	Cart     []carts.Cart  `gorm:"foreignKey:User_id"`
}
