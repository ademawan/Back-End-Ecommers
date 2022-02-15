package user

import (
	"Back-End-Ecommers/entities/address"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string
	Email      string `gorm:"unique"`
	Password   string
	Status     string
	Address_id int
	Address    address.Address
}
