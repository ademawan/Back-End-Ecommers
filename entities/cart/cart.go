package cart

import (
	"Back-End-Ecommers/entities/product"
	"os/user"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	User_id     int
	Product_id  int
	Qty         int
	Total_price int
	User        user.User
	Product     product.Product
}
