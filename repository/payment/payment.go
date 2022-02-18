package payment

import (
	"Back-End-Ecommers/entities"
	"errors"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

func (cr *PaymentRepository) Create(newPayment entities.Payment) (entities.Payment, error) {
	if err := cr.db.Create(&newPayment).Error; err != nil {
		return newPayment, err
	}
	return newPayment, nil
}

func (cr *PaymentRepository) GetById(PaymentId int) (entities.Payment, error) {
	arrPayment := entities.Payment{}

	result := cr.db.Where("ID = ?", PaymentId).First(&arrPayment)

	if err := result.Error; err != nil {
		return arrPayment, err
	}

	return arrPayment, nil
}

func (cr *PaymentRepository) UpdateById(PaymentId int, newPayment entities.Payment) (entities.Payment, error) {

	res := cr.db.Model(&entities.Payment{Model: gorm.Model{ID: uint(PaymentId)}}).Updates(entities.Payment{Name: newPayment.Name})

	if res.RowsAffected == 0 {
		return entities.Payment{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return newPayment, nil
}

func (cr *PaymentRepository) DeleteById(PaymentId int) (gorm.DeletedAt, error) {
	Payment := entities.Payment{}

	res := cr.db.Model(&Payment).Where("id = ?", PaymentId).Delete(&Payment)
	if res.RowsAffected == 0 {
		return Payment.DeletedAt, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return Payment.DeletedAt, nil
}

func (cr *PaymentRepository) GetAll() ([]entities.Payment, error) {
	arrPayment := []entities.Payment{}

	res := cr.db.Model(entities.Payment{}).Find(&arrPayment)
	if res.RowsAffected == 0 {
		return nil, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return arrPayment, nil
}
