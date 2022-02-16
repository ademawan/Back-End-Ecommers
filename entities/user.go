package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Status   string
	Order_ID int     `gorm:"column:order_id" json:"order_id"`
	Cart_ID  int     `gorm:"column:cart_id" json:"cart_id"`
	Address  Address `gorm:"ForeignKey:User_ID"`
}
