package category

import (
	"Back-End-Ecommers/entities/category"

	"gorm.io/gorm"
)

type Category interface {
	Create(newCategory category.Category) (category.Category, error)
	GetById(categoryId int) (category.Category, error)
	Update(categoryId int, newCategory category.Category) (category.Category, error)
	Delete(categoryId int) (gorm.DeletedAt, error)
	GetAll() ([]category.Category, error)
}
