package entities

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Price       int
	Qty         int
	Description string
	Category_ID int
}
