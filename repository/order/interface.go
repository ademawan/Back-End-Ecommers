package order

import "Back-End-Ecommers/entities"

type Order interface {
	Get() ([]entities.Order, error)
	GetById(orderId int) (entities.Order, error)
	Create(userId int, order entities.Order) (entities.Order, error)
	Update(orderId int, newOrder entities.Order) (entities.Order, error)
	Delete(orderId int) error
}
