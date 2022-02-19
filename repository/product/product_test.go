package product

import (
	"Back-End-Ecommers/configs"
	"Back-End-Ecommers/entities"

	"Back-End-Ecommers/utils"

	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run Create Product", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Product{})
		db.AutoMigrate(&entities.Product{})

		mockProduct := entities.Product{Name: "test", Price: 2000, Qty: 20, Description: "test", Category_ID: 1}

		resP, errP := repo.Create(mockProduct)

		assert.Nil(t, errP)
		assert.Equal(t, 1, int(resP.ID))
	})

	t.Run("fail run Create Product", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Product{})
		db.AutoMigrate(&entities.Product{})

		mocProduct := entities.Product{Name: "product1", Price: 2000, Qty: 20, Description: "test", Category_ID: 1}
		_, errP := repo.Create(mocProduct)
		if errP != nil {
			t.Fail()
		}

		mockProduct := entities.Product{Model: gorm.Model{ID: 1}, Name: "test", Price: 2000, Qty: 20, Description: "test", Category_ID: 1}
		_, errN := repo.Create(mockProduct)

		assert.NotNil(t, errN)
	})
}

func TestGetById(t *testing.T) {
	confg := configs.GetConfig()
	db := utils.InitDB(confg)
	repo := New(db)

	t.Run("success run GetById", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Product{})
		db.AutoMigrate(&entities.Product{})

		mockProduct := entities.Product{Name: "test", Price: 2000, Qty: 20, Description: "test", Category_ID: 1}
		if _, err := repo.Create(mockProduct); err != nil {
			t.Fatal()
		}

		res, err := repo.GetById(1)

		assert.Nil(t, err)
		assert.Equal(t, 1, int(res.ID))
		assert.Equal(t, 20, int(res.Qty))
	})

	t.Run("fail run GetById", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Product{})
		db.AutoMigrate(&entities.Product{})
		mockProduct := entities.Product{Name: "test", Price: 2000, Qty: 20, Description: "test", Category_ID: 1}
		if _, err := repo.Create(mockProduct); err != nil {
			t.Fatal()
		}

		_, err := repo.GetById(10)
		assert.NotNil(t, err)
	})
}

func TestUpdateById(t *testing.T) {
	confg := configs.GetConfig()
	db := utils.InitDB(confg)
	repo := New(db)

	t.Run("success run UpdateById", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Product{})
		db.AutoMigrate(&entities.Product{})

		mockProduct := entities.Product{Name: "Product", Price: 2000, Qty: 20, Description: "test", Category_ID: 1}
		if _, err := repo.Create(mockProduct); err != nil {
			t.Fatal()
		}

		mockNewP := entities.Product{Description: "update"}

		res, err := repo.Update(1, mockNewP)

		assert.Nil(t, err)
		assert.Equal(t, "update", res.Description)
	})

	t.Run("fail run UpdateById", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Product{})
		db.AutoMigrate(&entities.Product{})

		mockProduct := entities.Product{Name: "Product1", Price: 2000, Qty: 20, Description: "test", Category_ID: 1}
		if _, err := repo.Create(mockProduct); err != nil {
			t.Fatal()
		}

		mockNewP := entities.Product{Description: "update"}

		_, err := repo.Update(10, mockNewP)
		assert.NotNil(t, err)
	})
}

func TestDeleteById(t *testing.T) {
	confg := configs.GetConfig()
	db := utils.InitDB(confg)
	repo := New(db)

	t.Run("success run DeleteById", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Product{})
		db.AutoMigrate(&entities.Product{})
		mockProduct := entities.Product{Name: "Product1", Price: 2000, Qty: 20, Description: "test", Category_ID: 1}
		if _, err := repo.Create(mockProduct); err != nil {
			t.Fatal()
		}

		err := repo.Delete(1)
		assert.Nil(t, err)
	})

	t.Run("fail run DeleteById", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Product{})
		db.AutoMigrate(&entities.Product{})
		mockProduct := entities.Product{Name: "Product1", Price: 2000, Qty: 20, Description: "test", Category_ID: 1}
		if _, err := repo.Create(mockProduct); err != nil {
			t.Fatal()
		}
		err := repo.Delete(10)
		assert.NotNil(t, err)
	})
}

func TestGetAll(t *testing.T) {
	confg := configs.GetConfig()
	db := utils.InitDB(confg)
	repo := New(db)

	t.Run("success run GetAll", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Product{})
		db.AutoMigrate(&entities.Product{})
		mockProduct := entities.Product{Name: "Product1", Price: 2000, Qty: 20, Description: "test", Category_ID: 1}
		if _, err := repo.Create(mockProduct); err != nil {
			t.Fatal()
		}

		res, err := repo.GetAll()

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

	t.Run("fail run GetAll", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Product{})
		db.AutoMigrate(&entities.Product{})

		_, err := repo.GetAll()
		assert.NotNil(t, err)
	})

}
