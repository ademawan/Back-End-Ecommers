package entities

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name            string `gorm:"unique"`
	Price           int
	Qty             int
	Description     string
	Category        Category `gorm:"ForeignKey:Product_ID"`
	Cart_ID         int      `gorm:"column:cart_id" json:"cart_id"`
	Order_detail_ID int      `gorm:"column:order_detail_id" json:"order_detail_id"`
}
