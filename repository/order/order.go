package order

import (
	"Back-End-Ecommers/entities"
	"errors"
	"fmt"

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

func (ur *OrderRepository) Get(userId int) ([]entities.Order, error) {
	arrOrder := []entities.Order{}

	if err := ur.database.Preload("User").Preload("Payment").Preload("User.Address").Preload("OrderDetail").Where("user_id = ?", userId).Find(&arrOrder).Error; err != nil {
		return nil, err
	}

	return arrOrder, nil
}

func (ur *OrderRepository) GetById(orderId int) (entities.Order, error) {
	arrOrder := entities.Order{}

	result := ur.database.Preload("User").Preload("Payment").Preload("User.Address").Preload("OrderDetail").Where("ID = ?", orderId).First(&arrOrder)

	if err := result.Error; err != nil {
		return arrOrder, err
	}

	return arrOrder, nil
}

func (ur *OrderRepository) Create(userId, paymentId int) (entities.Order, error) {
	var orderRes = entities.Order{}
	err := ur.database.Transaction(func(tx *gorm.DB) error {

		order := entities.Order{User_ID: userId, Payment_ID: paymentId}
		arrCart := []entities.Cart{}

		if err := tx.Where("user_id = ?", userId).Find(&arrCart).Error; err != nil {
			return err
		}

		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Model(&orderRes).Preload("User").Preload("Payment").Preload("User.Address").Preload("OrderDetail").Create(&order).Error; err != nil {
			// return any error will rollback
			return err
		}
		for i := 0; i < len(arrCart); i++ {
			orderdetail := entities.OrderDetail{Order_ID: int(order.ID), Product_ID: arrCart[i].Product_ID, Qty: arrCart[i].Qty}

			//cek Qty product
			product := entities.Product{}
			if err := tx.Model(entities.Product{}).Where("ID = ?", arrCart[i].Product_ID).First(&product).Error; err != nil {
				return err
			}

			if product.Qty-arrCart[i].Qty < 0 {
				return errors.New("stock di gudang kurang broo")
			}
			newQty := product.Qty - arrCart[i].Qty

			if err := tx.Model(&entities.Product{}).Where("ID = ?", arrCart[i].Product_ID).Update("qty", newQty).Error; err != nil {
				return err
			}

			if err := tx.Model(&entities.OrderDetail{}).Create(&orderdetail).Error; err != nil {
				return err
			}
		}
		tx.Where("user_id = ?", userId).Delete(&entities.Cart{})

		// return nil will commit the whole transaction
		return nil
	})
	fmt.Println(orderRes)

	if err != nil {
		return orderRes, err
	}
	fmt.Println(orderRes)

	return orderRes, nil
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
