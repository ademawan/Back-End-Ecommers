package cart

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	User_id     int
	Product_id  int
	Qty         int
	Total_price int
}
