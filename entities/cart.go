package entities

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model

	Qty        int
	User_ID    int
	User       User
	Product_ID int
	Product    Product
}
