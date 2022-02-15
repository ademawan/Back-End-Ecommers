package product

import (
	"Back-End-Ecommers/entities/category"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Price       int
	Qty         int
	Description string
	Category_id int
	Category    category.Category
}
