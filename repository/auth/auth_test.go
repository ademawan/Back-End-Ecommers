package auth

import (
	config "Back-End-Ecommers/configs"
	"Back-End-Ecommers/entities/user"
	"Back-End-Ecommers/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	config := config.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&user.User{})
	db.AutoMigrate(&user.User{})

	// t.Run("success run login", func(t *testing.T) {
	// 	mockAddress :=
	// 	mockUser := user.User{Name: "test", Email: "test", Password: "test", Status: "admin", Address_id: 1}
	// 	_, err := _lib.New(db).UserRegister(mockUser)
	// 	if err != nil {
	// 		t.Fail()
	// 	}
	// 	mockLogin := user.User{Email: "test", Password: "test"}
	// 	res, err := repo.Login(mockLogin.Email, mockLogin.Password)

	// 	assert.Nil(t, err)
	// 	assert.Equal(t, "test", res.Email)
	// 	assert.Equal(t, "test", res.Password)
	// })

	t.Run("fail run login", func(t *testing.T) {
		mockLogin := user.User{Email: "anonim@123", Password: "anonim123"}
		_, err := repo.Login(mockLogin.Email, mockLogin.Password)
		assert.NotNil(t, err)
	})

}
