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
	Address  Address `gorm:"ForeignKey:User_ID"`
}
