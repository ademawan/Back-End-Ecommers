package order_detail

import (
	"Back-End-Ecommers/entities/order"
	"Back-End-Ecommers/entities/product"

	"gorm.io/gorm"
)

type Order_detail struct {
	gorm.Model
	Product_id  int
	Order_id    int
	Qty         int
	Total_price int
	Product     product.Product
	Order       order.Order
}
