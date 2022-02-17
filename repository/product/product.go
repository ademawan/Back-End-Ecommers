package product

import (
	"Back-End-Ecommers/entities"
	"errors"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (cr *ProductRepository) Create(newProduct entities.Product) (entities.Product, error) {
	if err := cr.db.Create(&newProduct).Error; err != nil {
		return newProduct, err
	}
	return newProduct, nil
}

func (cr *ProductRepository) GetById(productId int) (entities.Product, error) {
	arrProduct := entities.Product{}

	result := cr.db.Preload("Category").Where("ID = ?", productId).First(&arrProduct)

	if err := result.Error; err != nil {
		return arrProduct, err
	}

	return arrProduct, nil
}

func (cr *ProductRepository) Update(productId int, newProduct entities.Product) (entities.Product, error) {

	res := cr.db.Model(&entities.Product{Model: gorm.Model{ID: uint(productId)}}).Updates(entities.Product{Name: newProduct.Name})
	if res.RowsAffected == 0 {
		return entities.Product{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return newProduct, nil
}

func (cr *ProductRepository) Delete(id int) (gorm.DeletedAt, error) {
	product := entities.Product{}

	res := cr.db.Model(&product).Where("id = ?", id).Delete(&product)
	if res.RowsAffected == 0 {
		return product.DeletedAt, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return product.DeletedAt, nil
}

func (cr *ProductRepository) GetAll() ([]entities.Product, error) {
	arrProduct := []entities.Product{}
	// db.Model(&user).Association("Languages")
	// query.Model(articles).
	// 	Where("title like ?", "%"+title+"%").
	// 	Preload("Category").
	// 	Preload("Tag").Find(&articles).Error

	// ur.database.Preload("Task").Find(&arrUser, userId)

	res := cr.db.Preload("Category").Find(&arrProduct)
	if res.RowsAffected == 0 {
		return nil, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return arrProduct, nil
}
