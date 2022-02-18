package orderdetail

import "Back-End-Ecommers/entities"

type OrderDetail interface {
	// Get() ([]entities.Order, error)
	// GetById(orderId int) (entities.Order, error)
	Create(productId, orderId, qty int) (entities.OrderDetail, error)
	// Update(orderId int, newOrder entities.Order) (entities.Order, error)
	// Delete(orderId int) error
}
