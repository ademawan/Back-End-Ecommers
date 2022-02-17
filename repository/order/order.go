package order

import (
	"Back-End-Ecommers/entities"

	"gorm.io/gorm"
)

type OrderRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		database: db,
	}
}

func (ur *OrderRepository) Get() ([]entities.Order, error) {
	arrOrder := []entities.Order{}

	if err := ur.database.Find(&arrOrder).Error; err != nil {
		return nil, err
	}

	return arrOrder, nil
}

func (ur *OrderRepository) GetById(orderId int) (entities.Order, error) {
	arrOrder := entities.Order{}
	// var artikel models.Artikel

	// Conn.Preload("Komentar").Find(&artikle)
	result := ur.database.Preload("User").Preload("Payment").Where("ID = ?", orderId).First(&arrOrder)
	// if err := ur.database.Preload("Task").Find(&arrOrder, orderId).Error; err != nil {

	if err := result.Error; err != nil {
		return arrOrder, err
	}

	return arrOrder, nil
}

func (ur *OrderRepository) Create(userId int, u entities.Order) (entities.Order, error) {
	order := entities.Order{User_ID: userId, Payment_ID: u.Payment_ID}

	if err := ur.database.Create(&order).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (ur *OrderRepository) Update(orderId int, newOrder entities.Order) (entities.Order, error) {

	var order entities.Order
	ur.database.First(&order, orderId)

	if err := ur.database.Model(&order).Updates(&newOrder).Error; err != nil {
		return order, err
	}

	return order, nil
}

func (ur *OrderRepository) Delete(orderId int) error {

	var order entities.Order

	if err := ur.database.First(&order, orderId).Error; err != nil {
		return err
	}
	ur.database.Delete(&order, orderId)
	return nil

}
