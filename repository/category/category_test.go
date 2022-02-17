package category

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
	db.Migrator().DropTable(&entities.Category{})
	db.AutoMigrate(&entities.Category{})

	t.Run("success run Create", func(t *testing.T) {
		mockCategory := entities.Category{Name: "Action Figure"}
		res, err := repo.Create(mockCategory)
		assert.Nil(t, err)
		assert.Equal(t, "Action Figure", res.Name)

	})

	t.Run("fail run Create", func(t *testing.T) {
		mockCategoryP := entities.Category{Name: "Accessories"}
		if _, err := repo.Create(mockCategoryP); err != nil {
			t.Fatal()
		}
		mockCategory := entities.Category{Model: gorm.Model{ID: 1}, Name: "anonim123"}
		_, err := repo.Create(mockCategory)
		assert.NotNil(t, err)
	})
}

func TestGetById(t *testing.T) {
	config := config.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.Category{})
	db.AutoMigrate(&entities.Category{})

	t.Run("success run Get By Id", func(t *testing.T) {
		mockCategory := entities.Category{Name: "Action Figure"}
		if _, err := repo.Create(mockCategory); err != nil {
			t.Fatal()
		}

		res, err := repo.GetById(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(res.ID))
		assert.Equal(t, "Action Figure", res.Name)

	})

	t.Run("fail run Get By Id", func(t *testing.T) {
		mockCategoryP := entities.Category{Name: "Accessories"}
		if _, err := repo.Create(mockCategoryP); err != nil {
			t.Fatal()
		}

		res, err := repo.GetById(10)
		assert.NotNil(t, err)
		assert.NotEqual(t, 1, int(res.ID))
	})
}

func TestUpdateById(t *testing.T) {
	config := config.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.Category{})
	db.AutoMigrate(&entities.Category{})

	t.Run("success run Update By Id", func(t *testing.T) {
		mockCategory := entities.Category{Name: "Action Figure"}
		if _, err := repo.Create(mockCategory); err != nil {
			t.Fatal()
		}

		mockCategoryUpdate := entities.Category{Name: "Model Kit"}

		res, err := repo.UpdateById(1, mockCategoryUpdate)
		assert.Nil(t, err)
		assert.Equal(t, "Model Kit", res.Name)

	})

	t.Run("fail run Update By Id", func(t *testing.T) {
		mockCategoryP := entities.Category{Name: "Accessories"}
		if _, err := repo.Create(mockCategoryP); err != nil {
			t.Fatal()
		}

		res, err := repo.GetById(10)
		assert.NotNil(t, err)
		assert.NotEqual(t, 1, int(res.ID))
	})
}

func TestDeleteById(t *testing.T) {
	config := config.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.Category{})
	db.AutoMigrate(&entities.Category{})

	t.Run("success run Delete By Id", func(t *testing.T) {
		mockCategory := entities.Category{Name: "Action Figure"}
		if _, err := repo.Create(mockCategory); err != nil {
			t.Fatal()
		}

		err := repo.Delete(1)
		assert.Nil(t, err)
		// assert.Equal(t, true, res.Valid)

	})

	t.Run("fail run Delete By Id", func(t *testing.T) {
		mockCategoryP := entities.Category{Name: "Accessories"}
		if _, err := repo.Create(mockCategoryP); err != nil {
			t.Fatal()
		}

		err := repo.Delete(10)
		assert.NotNil(t, err)
	})
}

func TestGetAll(t *testing.T) {
	config := config.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&entities.Category{})
	db.AutoMigrate(&entities.Category{})

	t.Run("success run Get All", func(t *testing.T) {
		mockCategory := entities.Category{Name: "Action Figure"}
		if _, err := repo.Create(mockCategory); err != nil {
			t.Fatal()
		}

		res, err := repo.GetAll()
		assert.Nil(t, err)
		assert.NotNil(t, res)

	})

	t.Run("fail run Get All", func(t *testing.T) {
		if err := repo.Delete(1); err != nil {
			t.Fatal()
		}
		res, err := repo.GetAll()
		log.Info(err)
		log.Info(res)
		assert.NotNil(t, err)
	})
}
