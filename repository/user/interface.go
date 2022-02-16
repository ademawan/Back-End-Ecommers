package user

import "Back-End-Ecommers/entities"

type User interface {
	Get() ([]entities.User, error)
	GetById(userId int) (entities.User, error)
	Register(user entities.User) (entities.User, error)
	Update(userId int, newUser entities.User) (entities.User, error)
	Delete(userId int) error
}
