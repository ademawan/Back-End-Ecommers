package user

import (
	"Back-End-Ecommers/entities/user"

	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (ur *UserRepository) Get() ([]user.User, error) {
	arrUser := []user.User{}

	if err := ur.database.Find(&arrUser).Error; err != nil {
		return nil, err
	}

	return arrUser, nil
}

func (ur *UserRepository) GetById(userId int) (user.User, error) {
	arrUser := user.User{}
	// var artikel models.Artikel

	// Conn.Preload("Komentar").Find(&artikle)
	result := ur.database.Where("ID = ?", userId).First(&arrUser)
	// if err := ur.database.Preload("Task").Find(&arrUser, userId).Error; err != nil {

	if err := result.Error; err != nil {
		return arrUser, err
	}

	return arrUser, nil
}

func (ur *UserRepository) UserRegister(u user.User) (user.User, error) {
	if err := ur.database.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (ur *UserRepository) Update(userId int, newUser user.User) (user.User, error) {

	var user user.User
	ur.database.First(&user, userId)

	if err := ur.database.Model(&user).Updates(&newUser).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) Delete(userId int) error {

	var user user.User

	if err := ur.database.First(&user, userId).Error; err != nil {
		return err
	}
	ur.database.Delete(&user, userId)
	return nil

}
