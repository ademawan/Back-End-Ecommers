package cart

import "Back-End-Ecommers/entities"

type Cart interface {
	Create(user_id int, newCart entities.Cart) (entities.Cart, error)
	GetByUserId(user_id int) ([]entities.Cart, error)
	Update(user_id int, newCart entities.Cart) (entities.Cart, error)
	Delete(user_id int, cart_id int) error
}
