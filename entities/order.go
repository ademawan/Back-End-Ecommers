package entities

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model

	User            User    `gorm:"ForeignKey:Order_ID"`
	Payment         Payment `gorm:"ForeignKey:Order_ID"`
	Order_detail_ID int     `gorm:"column:order_detail_id" json:"order_detail_id"`
}
