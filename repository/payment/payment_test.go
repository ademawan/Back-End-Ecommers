package payment

import (
	config "Back-End-Ecommers/configs"
	"Back-End-Ecommers/entities"
	"Back-End-Ecommers/utils"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	config := config.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.Payment{})
	db.AutoMigrate(&entities.Payment{})

	t.Run("success run Create", func(t *testing.T) {
		mockPayment := entities.Payment{Name: "BANK"}
		res, err := repo.Create(mockPayment)
		assert.Nil(t, err)
		assert.Equal(t, "BANK", res.Name)

	})

	t.Run("fail run Create", func(t *testing.T) {
		mockPaymentP := entities.Payment{Name: "OVO"}
		if _, err := repo.Create(mockPaymentP); err != nil {
			t.Fatal()
		}
		mockPayment := entities.Payment{Model: gorm.Model{ID: 1}, Name: "anonim123"}
		_, err := repo.Create(mockPayment)
		assert.NotNil(t, err)
	})
}

func TestGetById(t *testing.T) {
	config := config.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("fail run Get By Id", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Payment{})
		db.AutoMigrate(&entities.Payment{})

		mockPaymentP := entities.Payment{Name: "OVO"}
		if _, err := repo.Create(mockPaymentP); err != nil {
			t.Fatal()
		}

		res, err := repo.GetById(10)
		assert.NotNil(t, err)
		assert.NotEqual(t, 1, int(res.ID))
	})

	t.Run("success run Get By Id", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Payment{})
		db.AutoMigrate(&entities.Payment{})

		mockPayment := entities.Payment{Name: "BANK"}
		if _, err := repo.Create(mockPayment); err != nil {
			t.Fatal()
		}

		res, err := repo.GetById(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(res.ID))
		assert.Equal(t, "BANK", res.Name)

	})
}

func TestUpdateById(t *testing.T) {
	config := config.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run Update By Id", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Payment{})
		db.AutoMigrate(&entities.Payment{})

		mockPayment := entities.Payment{Name: "BANK"}
		if _, err := repo.Create(mockPayment); err != nil {
			t.Fatal()
		}

		mockPaymentUpdate := entities.Payment{Name: "GOPAY"}

		res, err := repo.UpdateById(1, mockPaymentUpdate)
		assert.Nil(t, err)
		assert.Equal(t, "GOPAY", res.Name)

	})

	t.Run("fail run Update By Id", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Payment{})
		db.AutoMigrate(&entities.Payment{})

		mockPaymentP := entities.Payment{Name: "OVO"}
		if _, err := repo.Create(mockPaymentP); err != nil {
			t.Fatal()
		}
		mockPaymentUpdate := entities.Payment{Name: "GOPAY"}

		res, err := repo.UpdateById(10, mockPaymentUpdate)
		assert.NotNil(t, err)
		assert.NotEqual(t, 1, int(res.ID))
	})
}

func TestDeleteById(t *testing.T) {
	config := config.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run Delete By Id", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Payment{})
		db.AutoMigrate(&entities.Payment{})

		mockPayment := entities.Payment{Name: "BANK"}
		if _, err := repo.Create(mockPayment); err != nil {
			t.Fatal()
		}

		_, err := repo.DeleteById(1)
		assert.Nil(t, err)

	})

	t.Run("fail run Delete By Id", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Payment{})
		db.AutoMigrate(&entities.Payment{})

		mockPaymentP := entities.Payment{Name: "OVO"}
		if _, err := repo.Create(mockPaymentP); err != nil {
			t.Fatal()
		}

		_, err := repo.DeleteById(10)
		assert.NotNil(t, err)
	})
}

func TestGetAll(t *testing.T) {
	config := config.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run Get All", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Payment{})
		db.AutoMigrate(&entities.Payment{})

		mockPayment := entities.Payment{Name: "BANK"}
		if _, err := repo.Create(mockPayment); err != nil {
			t.Fatal()
		}

		res, err := repo.GetAll()
		assert.Nil(t, err)
		assert.NotNil(t, res)

	})

	t.Run("fail run Get All", func(t *testing.T) {
		if _, err := repo.DeleteById(1); err != nil {
			t.Fatal()
		}
		res, err := repo.GetAll()
		log.Info(err)
		log.Info(res)
		assert.NotNil(t, err)
	})
}
