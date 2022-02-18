package orderdetail

import (
	"Back-End-Ecommers/entities"

	"gorm.io/gorm"
)

type OrderDetailRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *OrderDetailRepository {
	return &OrderDetailRepository{
		database: db,
	}
}

// func (ur *OrderDetailRepository) Get() ([]entities.Order_detail, error) {
// 	arrOrderDetail := []entities.Order_detail{}

// 	if err := ur.database.Preload("User").Preload("Payment").Preload("User.Address").Find(&arrOrderDetail).Error; err != nil {
// 		return nil, err
// 	}

// 	return arrOrderDetail, nil
// }

// func (ur *OrderDetailRepository) GetById(orderId int) (entities.Order_detail, error) {
// 	arrOrderDetail := entities.Order_detail{}
// 	// var artikel models.Artikel

// 	// Conn.Preload("Komentar").Find(&artikle)
// 	result := ur.database.Preload("User").Preload("Payment").Preload("User.Address").Where("ID = ?", orderId).First(&arrOrderDetail)
// 	// if err := ur.database.Preload("Task").Find(&arrOrderDetail, orderId).Error; err != nil {

// 	if err := result.Error; err != nil {
// 		return arrOrderDetail, err
// 	}

// 	return arrOrderDetail, nil
// }

func (ur *OrderDetailRepository) Create(productId, orderId, qty int) (entities.OrderDetail, error) {
	order := entities.OrderDetail{Product_ID: productId, Order_ID: orderId, Qty: qty}

	if err := ur.database.Create(&order).Error; err != nil {
		return entities.OrderDetail{}, err
	}

	return order, nil
}

// func (ur *OrderDetailRepository) Update(orderId int, newOrder entities.Order_detail) (entities.Order_detail, error) {

// 	var order entities.Order_detail
// 	ur.database.First(&order, orderId)

// 	if err := ur.database.Model(&order).Updates(&newOrder).Error; err != nil {
// 		return order, err
// 	}

// 	return order, nil
// }

// func (ur *OrderDetailRepository) Delete(orderId int) error {

// 	var order entities.Order_detail

// 	if err := ur.database.First(&order, orderId).Error; err != nil {
// 		return err
// 	}
// 	ur.database.Delete(&order, orderId)
// 	return nil

// }
