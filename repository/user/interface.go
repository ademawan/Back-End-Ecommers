package user

import "Back-End-Ecommers/entities/user"

type User interface {
	get() ([]user.User, error)
	GetById(userId int)
}
