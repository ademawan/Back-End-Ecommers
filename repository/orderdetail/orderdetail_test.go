package orderdetail

import (
	"Back-End-Ecommers/configs"
	"Back-End-Ecommers/entities"
	repoPay "Back-End-Ecommers/repository/payment"
	repoUser "Back-End-Ecommers/repository/user"

	// repoProduct "Back-End-Ecommers/repository/product"
	repoO "Back-End-Ecommers/repository/order"
	"Back-End-Ecommers/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{}, &entities.Payment{}, &entities.Order{}, &entities.OrderDetail{})
	db.AutoMigrate(&entities.User{}, &entities.Payment{}, &entities.Order{}, &entities.OrderDetail{})

	t.Run("success run Create", func(t *testing.T) {
		mockUser := entities.User{Model: gorm.Model{ID: 1}, Name: "test"}
		resU, errU := repoUser.New(db).Register(mockUser)
		if errU != nil {
			t.Fail()
		}

		mockPayment := entities.Payment{Model: gorm.Model{ID: 1}, Name: "Ovo"}
		resP, errP := repoPay.New(db).Create(mockPayment)
		if errP != nil {
			t.Fail()
		}

		resO, errO := repoO.New(db).Create(int(resU.ID), int(resP.ID))
		if errO != nil {
			t.Fail()
		}

		res, err := repo.Create(1, int(resO.ID), 10)
		assert.Nil(t, err)
		assert.Equal(t, 10, res.Qty)

	})

	t.Run("failed run Create", func(t *testing.T) {

		_, err := repo.Create(1, 2, 10)

		assert.NotNil(t, err)

	})
}
