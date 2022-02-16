package entities

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model

	User_ID    int `gorm:"column:user_id" json:"user_id"`
	User       User
	Payment_ID int `gorm:"column:payment_id" json:"payment_id"`
	Payment    Payment
}
