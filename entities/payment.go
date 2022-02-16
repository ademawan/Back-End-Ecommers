package entities

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Name     string
	Order_ID int `gorm:"column:order_id" json:"order_id"`
}
