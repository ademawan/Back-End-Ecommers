package category

import (
	"Back-End-Ecommers/entities/category"
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

func (cr *CategoryRepository) Create(newCategory category.Category) (category.Category, error) {
	if err := cr.db.Create(&newCategory).Error; err != nil {
		return newCategory, err
	}
	return newCategory, nil
}

func (cr *CategoryRepository) GetById(categoryId int) (category.Category, error) {
	arrCategory := category.Category{}

	result := cr.db.Where("ID = ?", categoryId).First(&arrCategory)

	if err := result.Error; err != nil {
		return arrCategory, err
	}

	return arrCategory, nil
}

func (cr *CategoryRepository) UpdateById(categoryId int, newCategory category.Category) (category.Category, error) {

	res := cr.db.Model(&category.Category{Model: gorm.Model{ID: uint(categoryId)}}).Updates(category.Category{Name: newCategory.Name})

	if res.RowsAffected == 0 {
		return category.Category{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return newCategory, nil
}

func (cr *CategoryRepository) DeleteById(id int) (gorm.DeletedAt, error) {
	category := category.Category{}

	res := cr.db.Model(&category).Where("id = ?", id).Delete(&category)
	if res.RowsAffected == 0 {
		return category.DeletedAt, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return category.DeletedAt, nil
}

func (cr *CategoryRepository) GetAll() ([]category.Category, error) {
	arrCategory := []category.Category{}

	res := cr.db.Model(category.Category{}).Find(&arrCategory)
	if res.RowsAffected == 0 {
		return nil, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return arrCategory, nil
}
