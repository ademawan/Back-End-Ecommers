package payment

import (
	"Back-End-Ecommers/entities"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{
		database: db,
	}
}

func (ur *PaymentRepository) Register(newPayment entities.Payment) (entities.Payment, error) {
	if err := ur.database.Create(&newPayment).Error; err != nil {
		return newPayment, err
	}

	ur.database.Create(&newPayment)

	return newPayment, nil
}

func (ur *PaymentRepository) GetAll() ([]entities.User, error) {
	arrUser := []entities.User{}

	if err := ur.database.Preload("Address").Find(&arrUser).Error; err != nil {
		return nil, err
	}

	return arrUser, nil
}

func (ur *PaymentRepository) GetById(userId int) (entities.User, error) {
	arrUser := entities.User{}
	// var artikel models.Artikel

	// Conn.Preload("Komentar").Find(&artikle)
	result := ur.database.Preload("Address").Where("ID = ?", userId).First(&arrUser)
	// if err := ur.database.Preload("Task").Find(&arrUser, userId).Error; err != nil {

	if err := result.Error; err != nil {
		return arrUser, err
	}

	return arrUser, nil
}

func (ur *PaymentRepository) Update(userId int, newUser entities.User) (entities.User, error) {

	var user entities.User
	ur.database.First(&user, userId)

	if err := ur.database.Model(&user).Updates(&newUser).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ur *PaymentRepository) Delete(userId int) error {

	var user entities.User

	if err := ur.database.First(&user, userId).Error; err != nil {
		return err
	}
	ur.database.Delete(&user, userId)
	return nil

}
