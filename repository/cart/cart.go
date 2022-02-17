package cart

import (
	"Back-End-Ecommers/entities"
	"errors"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (cr *CartRepository) Create(user_id int, newCart entities.Cart) (entities.Cart, error) {
	newCart.User_ID = user_id

	if err := cr.db.Create(&newCart).Error; err != nil {
		return newCart, err
	}
	return newCart, nil
}

func (cr *CartRepository) GetByUserId(user_id int) ([]entities.Cart, error) {
	arrCart := []entities.Cart{}

	res := cr.db.Preload("User").Preload("User.Address").Preload("Product").Where("user_id = ?", user_id).First(&arrCart)

	if err := res.Error; err != nil {
		return arrCart, err
	}

	return arrCart, nil
}

func (cr *CartRepository) Update(user_id int, newCart entities.Cart) (entities.Cart, error) {
	res := cr.db.Model(&entities.Cart{Model: gorm.Model{ID: uint(user_id)}}).Updates(entities.Cart{Qty: newCart.Qty, Product_ID: newCart.Product_ID})

	if res.RowsAffected == 0 {
		return entities.Cart{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return newCart, nil
}

func (cr *CartRepository) Delete(user_id int, cart_id int) error {
	Cart := entities.Cart{}

	res := cr.db.Model(&Cart).Where("user_id = ? AND cart_id = ?", user_id, cart_id).Delete(&Cart)

	if res.RowsAffected == 0 {
		return errors.New(gorm.ErrRecordNotFound.Error())
	}

	return nil
}
