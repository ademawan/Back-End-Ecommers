package order

import (
	"Back-End-Ecommers/entities"
	"Back-End-Ecommers/entities/payment"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	User_id    int
	Payment_id int
	User       entities.User
	Payment    payment.Payment
}
