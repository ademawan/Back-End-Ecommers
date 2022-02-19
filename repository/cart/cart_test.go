package cart

import (
	"Back-End-Ecommers/configs"
	"Back-End-Ecommers/entities"
	repoCategory "Back-End-Ecommers/repository/category"
	repoProduct "Back-End-Ecommers/repository/product"
	repoUser "Back-End-Ecommers/repository/user"
	"Back-End-Ecommers/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Cart{}, &entities.User{}, &entities.Product{})
		db.AutoMigrate(&entities.Cart{}, &entities.User{}, &entities.Product{})

		mockUser := entities.User{Name: "anonim1", Email: "anonim@1", Password: "anonim1"}
		resU, errU := repoUser.New(db).Register(mockUser)
		if errU != nil {
			t.Fail()
		}

		mockCategory := entities.Category{Name: "product1"}
		resCa, errP := repoCategory.New(db).Create(mockCategory)
		if errP != nil {
			t.Fail()
		}

		mockProduct := entities.Product{Name: "product1", Qty: 1, Price: 2000, Category_ID: int(resCa.ID)}
		resP, errP := repoProduct.New(db).Create(mockProduct)
		if errP != nil {
			t.Fail()
		}

		mockCart := entities.Cart{Product_ID: int(resP.ID), Qty: 1}
		res, err := repo.Create(int(resU.ID), mockCart)
		if err != nil {
			t.Fail()
		}

		assert.Equal(t, 1, res.Product_ID)
		assert.Equal(t, 1, res.User_ID)
		assert.Nil(t, err)

	})

	t.Run("fail run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Cart{}, &entities.User{}, &entities.Product{})
		db.AutoMigrate(&entities.Cart{}, &entities.User{}, &entities.Product{})

		mockUser := entities.User{Name: "anonim1", Email: "anonim@1", Password: "anonim1"}
		resU, errU := repoUser.New(db).Register(mockUser)
		if errU != nil {
			t.Fail()
		}

		mockCategory := entities.Category{Name: "product1"}
		resCa, errP := repoCategory.New(db).Create(mockCategory)
		if errP != nil {
			t.Fail()
		}

		mockProduct := entities.Product{Name: "product1", Qty: 1, Price: 2000, Category_ID: int(resCa.ID)}
		resP, errP := repoProduct.New(db).Create(mockProduct)
		if errP != nil {
			t.Fail()
		}

		mockCart := entities.Cart{Model: gorm.Model{ID: 1}, Product_ID: int(resP.ID), Qty: 1}
		_, err := repo.Create(int(resU.ID), mockCart)
		if err != nil {
			t.Fail()
		}

		assert.Nil(t, err)

	})
}

func TestByUserId(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Cart{}, &entities.User{}, &entities.Product{})
		db.AutoMigrate(&entities.Cart{}, &entities.User{}, &entities.Product{})

		mockUser := entities.User{Name: "anonim1", Email: "anonim@1", Password: "anonim1"}
		resU, errU := repoUser.New(db).Register(mockUser)
		if errU != nil {
			t.Fail()
		}

		mockCategory := entities.Category{Name: "product1"}
		resCa, errP := repoCategory.New(db).Create(mockCategory)
		if errP != nil {
			t.Fail()
		}

		mockProduct := entities.Product{Name: "product1", Qty: 1, Price: 2000, Category_ID: int(resCa.ID)}
		resP, errP := repoProduct.New(db).Create(mockProduct)
		if errP != nil {
			t.Fail()
		}

		mockCart := entities.Cart{Product_ID: int(resP.ID), Qty: 1}
		res, errC := repo.Create(int(resU.ID), mockCart)
		if errC != nil {
			t.Fail()
		}

		_, err := repo.GetByUserId(res.User_ID)

		assert.Equal(t, 1, res.Product_ID)
		assert.Equal(t, 1, res.User_ID)
		assert.Nil(t, err)

	})

	t.Run("fail run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Cart{}, &entities.User{}, &entities.Product{})
		db.AutoMigrate(&entities.Cart{}, &entities.User{}, &entities.Product{})

		mockUser := entities.User{Name: "anonim1", Email: "anonim@1", Password: "anonim1"}
		resU, errU := repoUser.New(db).Register(mockUser)
		if errU != nil {
			t.Fail()
		}

		mockCategory := entities.Category{Name: "product1"}
		resCa, errP := repoCategory.New(db).Create(mockCategory)
		if errP != nil {
			t.Fail()
		}

		mockProduct := entities.Product{Name: "product1", Qty: 1, Price: 2000, Category_ID: int(resCa.ID)}
		resP, errP := repoProduct.New(db).Create(mockProduct)
		if errP != nil {
			t.Fail()
		}

		mockCart := entities.Cart{Product_ID: int(resP.ID), Qty: 1}
		_, errC := repo.Create(int(resU.ID), mockCart)
		if errC != nil {
			t.Fail()
		}

		_, err := repo.GetByUserId(10)

		assert.NotNil(t, err)

	})
}

func TestUpdate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Cart{}, &entities.User{}, &entities.Product{})
		db.AutoMigrate(&entities.Cart{}, &entities.User{}, &entities.Product{})

		mockUser := entities.User{Name: "anonim1", Email: "anonim@1", Password: "anonim1"}
		resU, errU := repoUser.New(db).Register(mockUser)
		if errU != nil {
			t.Fail()
		}

		mockCategory := entities.Category{Name: "product1"}
		resCa, errP := repoCategory.New(db).Create(mockCategory)
		if errP != nil {
			t.Fail()
		}

		mockProduct := entities.Product{Name: "product1", Qty: 1, Price: 2000, Category_ID: int(resCa.ID)}
		resP, errP := repoProduct.New(db).Create(mockProduct)
		if errP != nil {
			t.Fail()
		}

		mockCart := entities.Cart{Product_ID: int(resP.ID), Qty: 1}
		res, errC := repo.Create(int(resU.ID), mockCart)
		if errC != nil {
			t.Fail()
		}

		mockCartUp := entities.Cart{Product_ID: int(resP.ID), Qty: 2}

		resUp, err := repo.Update(res.User_ID, mockCartUp)

		assert.Equal(t, 1, res.Product_ID)
		assert.Equal(t, 1, res.User_ID)
		assert.Equal(t, 2, resUp.Qty)
		assert.Nil(t, err)

	})

	t.Run("fail run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Cart{}, &entities.User{}, &entities.Product{})
		db.AutoMigrate(&entities.Cart{}, &entities.User{}, &entities.Product{})

		mockUser := entities.User{Name: "anonim1", Email: "anonim@1", Password: "anonim1"}
		resU, errU := repoUser.New(db).Register(mockUser)
		if errU != nil {
			t.Fail()
		}

		mockCategory := entities.Category{Name: "product1"}
		resCa, errP := repoCategory.New(db).Create(mockCategory)
		if errP != nil {
			t.Fail()
		}

		mockProduct := entities.Product{Name: "product1", Qty: 1, Price: 2000, Category_ID: int(resCa.ID)}
		resP, errP := repoProduct.New(db).Create(mockProduct)
		if errP != nil {
			t.Fail()
		}

		mockCart := entities.Cart{Product_ID: int(resP.ID), Qty: 1}
		_, errC := repo.Create(int(resU.ID), mockCart)
		if errC != nil {
			t.Fail()
		}

		mockCartUp := entities.Cart{Product_ID: int(resP.ID), Qty: 2}

		_, err := repo.Update(10, mockCartUp)

		assert.NotNil(t, err)

	})
}

func TestDelete(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Cart{}, &entities.User{}, &entities.Product{})
		db.AutoMigrate(&entities.Cart{}, &entities.User{}, &entities.Product{})

		mockUser := entities.User{Name: "anonim1", Email: "anonim@1", Password: "anonim1"}
		resU, errU := repoUser.New(db).Register(mockUser)
		if errU != nil {
			t.Fail()
		}

		mockCategory := entities.Category{Name: "product1"}
		resCa, errP := repoCategory.New(db).Create(mockCategory)
		if errP != nil {
			t.Fail()
		}

		mockProduct := entities.Product{Name: "product1", Qty: 1, Price: 2000, Category_ID: int(resCa.ID)}
		resP, errP := repoProduct.New(db).Create(mockProduct)
		if errP != nil {
			t.Fail()
		}

		mockCart := entities.Cart{Product_ID: int(resP.ID), Qty: 1}
		resC, errC := repo.Create(int(resU.ID), mockCart)
		if errC != nil {
			t.Fail()
		}

		fmt.Println(resC)

		err := repo.Delete(1, 1)

		assert.NotNil(t, err)

	})

	t.Run("fail run Create", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Cart{}, &entities.User{}, &entities.Product{})
		db.AutoMigrate(&entities.Cart{}, &entities.User{}, &entities.Product{})

		mockUser := entities.User{Name: "anonim1", Email: "anonim@1", Password: "anonim1"}
		resU, errU := repoUser.New(db).Register(mockUser)
		if errU != nil {
			t.Fail()
		}

		mockCategory := entities.Category{Name: "product1"}
		resCa, errP := repoCategory.New(db).Create(mockCategory)
		if errP != nil {
			t.Fail()
		}

		mockProduct := entities.Product{Name: "product1", Qty: 1, Price: 2000, Category_ID: int(resCa.ID)}
		resP, errP := repoProduct.New(db).Create(mockProduct)
		if errP != nil {
			t.Fail()
		}

		mockCart := entities.Cart{Product_ID: int(resP.ID), Qty: 1}
		_, errC := repo.Create(int(resU.ID), mockCart)
		if errC != nil {
			t.Fail()
		}

		err := repo.Delete(10, 100)

		assert.NotNil(t, err)

	})
}
