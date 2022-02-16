package auth

import "Back-End-Ecommers/entities/user"

type Auth interface {
	Login(email, password string) (user.User, error)
}
