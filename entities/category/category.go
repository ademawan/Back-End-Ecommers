package category

import (
	"Back-End-Ecommers/entities/product"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name    string
	Product []product.Product `gorm:"foreignKey:Category_id"`
}
