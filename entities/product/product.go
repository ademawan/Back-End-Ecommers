package product

import (
	"Back-End-Ecommers/entities/category"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name_product string `gorm:"unique"`
	Price        int
	Qty          int
	Description  string
	Category_id  []category.Category `gorm:"foreignKey:Product_id"`
}
