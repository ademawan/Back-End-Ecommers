package order

import (
	"Back-End-Ecommers/entities/payment"
	"Back-End-Ecommers/entities/user"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	User_id    int
	Payment_id int
	User       user.User
	Payment    payment.Payment
}
