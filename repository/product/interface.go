package product

import "Back-End-Ecommers/entities"

type Product interface {
	Create(product entities.Product) (entities.Product, error)
	GetById(productId int) (entities.Product, error)
	Update(productId int, newProduct entities.Product) (entities.Product, error)
	Delete(productId int) error
	GetAll() ([]entities.Product, error)
}
