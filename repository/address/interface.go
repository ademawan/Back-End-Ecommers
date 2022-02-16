package address

import "Back-End-Ecommers/entities"

type Address interface {
	Get() ([]entities.Address, error)
	GetById(addressId int) (entities.Address, error)
	Insert(address entities.Address) (entities.Address, error)
	Update(addressId int, newAddress entities.Address) (entities.Address, error)
	Delete(addressId int) error
}
