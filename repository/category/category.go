package category

import (
	"Back-End-Ecommers/entities"
	"errors"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (cr *CategoryRepository) Create(newCategory entities.Category) (entities.Category, error) {
	if err := cr.db.Create(&newCategory).Error; err != nil {
		return newCategory, err
	}
	return newCategory, nil
}

func (cr *CategoryRepository) GetById(categoryId int) (entities.Category, error) {
	arrCategory := entities.Category{}

	result := cr.db.Where("ID = ?", categoryId).First(&arrCategory)

	if err := result.Error; err != nil {
		return arrCategory, err
	}

	return arrCategory, nil
}

func (cr *CategoryRepository) UpdateById(categoryId int, newCategory entities.Category) (entities.Category, error) {

	res := cr.db.Model(&entities.Category{Model: gorm.Model{ID: uint(categoryId)}}).Updates(entities.Category{Name: newCategory.Name})

	if res.RowsAffected == 0 {
		return entities.Category{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return newCategory, nil
}

func (cr *CategoryRepository) DeleteById(id int) (gorm.DeletedAt, error) {
	category := entities.Category{}

	res := cr.db.Model(&category).Where("id = ?", id).Delete(&category)
	if res.RowsAffected == 0 {
		return category.DeletedAt, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return category.DeletedAt, nil
}

func (cr *CategoryRepository) GetAll() ([]entities.Category, error) {
	arrCategory := []entities.Category{}

	res := cr.db.Model(entities.Category{}).Find(&arrCategory)
	if res.RowsAffected == 0 {
		return nil, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return arrCategory, nil
}
