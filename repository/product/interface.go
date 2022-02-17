package product

import (
	"Back-End-Ecommers/entities"

	"gorm.io/gorm"
)

type Product interface {
	Create(newProduct entities.Product) (entities.Product, error)
	GetById(productId int) (entities.Product, error)
	Update(productId int, newProduct entities.Product) (entities.Product, error)
	Delete(productId int) (gorm.DeletedAt, error)
	GetAll() ([]entities.Product, error)
}
