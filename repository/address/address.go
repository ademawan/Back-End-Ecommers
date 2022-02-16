package address

import (
	"Back-End-Ecommers/entities"

	"gorm.io/gorm"
)

type AddressRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *AddressRepository {
	return &AddressRepository{
		database: db,
	}
}

func (tr *AddressRepository) Get() ([]entities.Address, error) {
	arrAddress := []entities.Address{}

	if err := tr.database.Find(&arrAddress).Error; err != nil {
		return nil, err
	}

	return arrAddress, nil
}

func (tr *AddressRepository) GetById(addressId int) (entities.Address, error) {
	arrAddress := entities.Address{}

	if err := tr.database.Preload("Task").Find(&arrAddress, addressId).Error; err != nil {
		return arrAddress, err
	}

	return arrAddress, nil
}

func (tr *AddressRepository) Insert(t entities.Address) (entities.Address, error) {
	if err := tr.database.Create(&t).Error; err != nil {
		return t, err
	}

	return t, nil
}

func (tr *AddressRepository) Update(addressId int, newAddress entities.Address) (entities.Address, error) {

	var address entities.Address
	tr.database.First(&address, addressId)

	if err := tr.database.Model(&address).Updates(&newAddress).Error; err != nil {
		return address, err
	}

	return address, nil
}

func (tr *AddressRepository) Delete(addressId int) error {

	var address entities.Address

	if err := tr.database.First(&address, addressId).Error; err != nil {
		return err
	}
	tr.database.Delete(&address, addressId)
	return nil

}
