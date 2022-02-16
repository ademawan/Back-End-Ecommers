package auth

import "Back-End-Ecommers/entities"

type Auth interface {
	Login(email, password string) (entities.User, error)
}
