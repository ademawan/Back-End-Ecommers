package entities

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name       string
	Product_ID int `gorm:"column:product_id" json:"product_id"`
}
