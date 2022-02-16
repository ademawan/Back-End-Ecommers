package category

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name_category string
	Product_id    uint
}
