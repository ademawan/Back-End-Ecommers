package category

import (
	"Back-End-Ecommers/entities"
)

type Category interface {
	Create(newCategory entities.Category) (entities.Category, error)
	GetById(categoryId int) (entities.Category, error)
	// Update(categoryId int, newCategory entities.Category) (entities.Category, error)
	// Delete(categoryId int) (gorm.DeletedAt, error)
	// GetAll() ([]entities.Category, error)
}
