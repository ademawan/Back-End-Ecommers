package entities

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model

	Qty         int
	Total_price int
	User        User    `gorm:"ForeignKey:Cart_ID"`
	Product     Product `gorm:"ForeignKey:Cart_ID"`
}
