package entities

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name    string
	Product []Product `gorm:"ForeignKey:Category_ID"`
}
