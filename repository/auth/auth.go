package auth

import (
	"Back-End-Ecommers/entities"

	"gorm.io/gorm"
)

type AuthDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AuthDb {
	return &AuthDb{
		db: db,
	}
}

func (ad *AuthDb) Login(email, password string) (entities.User, error) {
	user := entities.User{}
	if err := ad.db.Model(&user).Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
