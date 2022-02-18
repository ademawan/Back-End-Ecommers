package payment

import (
	"Back-End-Ecommers/entities"

	"gorm.io/gorm"
)

type Payment interface {
	Create(newPayment entities.Payment) (entities.Payment, error)
	GetById(PaymentId int) (entities.Payment, error)
	UpdateById(PaymentId int, newPayment entities.Payment) (entities.Payment, error)
	DeleteById(PaymentId int) (gorm.DeletedAt, error)
	GetAll() ([]entities.Payment, error)
}
