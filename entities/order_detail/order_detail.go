package order_detail

import "gorm.io/gorm"

type Order_detail struct {
	gorm.Model
	Product_id  int
	Order_id    int
	Total_price int
}
