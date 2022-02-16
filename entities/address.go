package entities

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Street      string
	City        string
	Region      string
	Postal_code int
	User_ID     int `gorm:"column:user_id" json:"user_id"`
}
