package order

import (
	"Back-End-Ecommers/configs"
	"Back-End-Ecommers/entities"
	repoPay "Back-End-Ecommers/repository/payment"
	repoUser "Back-End-Ecommers/repository/user"
	"Back-End-Ecommers/utils"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{}, &entities.Payment{}, &entities.Order{})
	db.AutoMigrate(&entities.User{}, &entities.Payment{}, &entities.Order{})

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

		res, err := repo.Create(int(resU.ID), int(resP.ID))
		assert.Nil(t, err)
		assert.Equal(t, 1, res.User_ID)

	})

	t.Run("failed run Create", func(t *testing.T) {

		_, err := repo.Create(2, 3)

		assert.NotNil(t, err)

	})
}

func TestGetById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{}, &entities.Payment{}, &entities.Order{})
	db.AutoMigrate(&entities.User{}, &entities.Payment{}, &entities.Order{})

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

		_, errC := repo.Create(int(resU.ID), int(resP.ID))
		if errC != nil {
			t.Fail()
		}

		res, err := repo.GetById(1)
		if errC != nil {
			t.Fail()
		}

		assert.Nil(t, err)
		assert.Equal(t, 1, res.User_ID)

	})

	t.Run("failed run Create", func(t *testing.T) {

		_, err := repo.GetById(10)

		assert.NotNil(t, err)

	})
}

func TestGet(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{}, &entities.Payment{}, &entities.Order{})
	db.AutoMigrate(&entities.User{}, &entities.Payment{}, &entities.Order{})

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

		_, errC := repo.Create(int(resU.ID), int(resP.ID))
		if errC != nil {
			t.Fail()
		}

		res, err := repo.Get(1)
		if errC != nil {
			t.Fail()
		}

		log.Info(res)

		assert.Nil(t, err)
		assert.Equal(t, 1, int(resU.ID))

	})

	t.Run("failed run Create", func(t *testing.T) {

		_, err := repo.Get(10)

		assert.Nil(t, err)

	})
}

func TestUpdate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{}, &entities.Payment{}, &entities.Order{})
	db.AutoMigrate(&entities.User{}, &entities.Payment{}, &entities.Order{})

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

		resC, errC := repo.Create(int(resU.ID), int(resP.ID))
		if errC != nil {
			t.Fail()
		}

		mockUp := entities.Payment{Model: gorm.Model{ID: 2}, Name: "GoPay"}
		resUp, errUp := repoPay.New(db).Create(mockUp)
		if errUp != nil {
			t.Fail()
		}

		res, err := repo.Update(int(resC.ID), entities.Order{Payment_ID: int(resUp.ID)})
		if errC != nil {
			t.Fail()
		}

		assert.Nil(t, err)
		assert.Equal(t, 2, int(res.Payment_ID))

	})

	t.Run("failed run Create", func(t *testing.T) {

		_, err := repo.Update(2, entities.Order{})

		assert.NotNil(t, err)

	})
}

func TestDelete(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.User{}, &entities.Payment{}, &entities.Order{})
	db.AutoMigrate(&entities.User{}, &entities.Payment{}, &entities.Order{})

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

		resC, errC := repo.Create(int(resU.ID), int(resP.ID))
		if errC != nil {
			t.Fail()
		}

		err := repo.Delete(int(resC.ID))
		if err != nil {
			t.Fail()
		}

		assert.Nil(t, err)

	})

	t.Run("failed run Create", func(t *testing.T) {

		err := repo.Delete(2)

		assert.NotNil(t, err)

	})
}
