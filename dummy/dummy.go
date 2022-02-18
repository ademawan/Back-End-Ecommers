package dummy

import (
	config "Back-End-Ecommers/configs"
	"Back-End-Ecommers/entities"
	addressRepo "Back-End-Ecommers/repository/address"
	cartRepo "Back-End-Ecommers/repository/cart"
	categoryRepo "Back-End-Ecommers/repository/category"

	// orderRepo "Back-End-Ecommers/repository/order"
	paymentRepo "Back-End-Ecommers/repository/payment"
	productRepo "Back-End-Ecommers/repository/product"
	userRepo "Back-End-Ecommers/repository/user"
	"Back-End-Ecommers/utils"
)

func Dummy() {
	config := config.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})
	db.Migrator().DropTable(&entities.Address{})
	db.AutoMigrate(&entities.Address{})
	db.Migrator().DropTable(&entities.Category{})
	db.AutoMigrate(&entities.Category{})
	db.Migrator().DropTable(&entities.Product{})
	db.AutoMigrate(&entities.Product{})
	db.Migrator().DropTable(&entities.Cart{})
	db.AutoMigrate(&entities.Cart{})
	db.Migrator().DropTable(&entities.Order{})
	db.AutoMigrate(&entities.Order{})
	db.Migrator().DropTable(&entities.OrderDetail{})
	db.AutoMigrate(&entities.OrderDetail{})

	repoUser := userRepo.New(db)
	repoAddress := addressRepo.New(db)
	repoCategory := categoryRepo.New(db)
	repoProduct := productRepo.New(db)
	repoPayment := paymentRepo.New(db)
	// repoOrder := orderRepo.New(db)
	repoCart := cartRepo.New(db)

	//user dummy
	mockUserAdmin := entities.User{Name: "Admin", Email: "admin@gmail.com", Password: "admin"}
	repoUser.Register(mockUserAdmin)
	mockUser1 := entities.User{Name: "testuser1", Email: "testuser1@gmail.com", Password: "xyz"}
	repoUser.Register(mockUser1)
	mockUser2 := entities.User{Name: "testuser2", Email: "testuser2@gmail.com", Password: "xyz"}
	repoUser.Register(mockUser2)
	mockUser3 := entities.User{Name: "testuser3", Email: "testuser3@gmail.com", Password: "xyz"}
	repoUser.Register(mockUser3)
	mockUser4 := entities.User{Name: "testuser4", Email: "testuser4@gmail.com", Password: "xyz"}
	repoUser.Register(mockUser4)

	//address
	mockAddress1 := entities.Address{Street: "teststreetaddres1", City: "city1", Region: "region1", Postal_code: "postal_code1"}
	repoAddress.Insert(1, mockAddress1)
	mockAddress2 := entities.Address{Street: "teststreetaddres2", City: "city2", Region: "region2", Postal_code: "postal_code2"}
	repoAddress.Insert(2, mockAddress2)
	mockAddress3 := entities.Address{Street: "teststreetaddres3", City: "city3", Region: "region3", Postal_code: "postal_code3"}
	repoAddress.Insert(3, mockAddress3)

	//Category Dummy
	mockCategory1 := entities.Category{Name: "category1"}
	repoCategory.Create(mockCategory1)
	mockCategory2 := entities.Category{Name: "category2"}
	repoCategory.Create(mockCategory2)
	mockCategory3 := entities.Category{Name: "category3"}
	repoCategory.Create(mockCategory3)
	mockCategory4 := entities.Category{Name: "category4"}
	repoCategory.Create(mockCategory4)

	//product dummy
	mockProduct1 := entities.Product{Name: "product1", Price: 30000, Qty: 30, Description: "bLABFJAEFKJHKJNEKFNKJWAEFAWNJNASNDJ", Category_ID: 1}
	repoProduct.Create(mockProduct1)
	mockProduct2 := entities.Product{Name: "product2", Price: 30000, Qty: 30, Description: "bLABFJAEFKJHKJNEKFNKJWAEFAWNJNASNDJ", Category_ID: 2}
	repoProduct.Create(mockProduct2)
	mockProduct3 := entities.Product{Name: "product3", Price: 30000, Qty: 30, Description: "bLABFJAEFKJHKJNEKFNKJWAEFAWNJNASNDJ", Category_ID: 3}
	repoProduct.Create(mockProduct3)
	mockProduct4 := entities.Product{Name: "product4", Price: 40000, Qty: 40, Description: "bLABFJAEFKJHKJNEKFNKJWAEFAWNJNASNDJ", Category_ID: 4}
	repoProduct.Create(mockProduct4)
	mockProduct5 := entities.Product{Name: "product5", Price: 40000, Qty: 40, Description: "bLABFJAEFKJHKJNEKFNKJWAEFAWNJNASNDJ", Category_ID: 1}
	repoProduct.Create(mockProduct5)
	mockProduct6 := entities.Product{Name: "product6", Price: 60000, Qty: 60, Description: "bLABFJAEFKJHKJNEKFNKJWAEFAWNJNASNDJ", Category_ID: 2}
	repoProduct.Create(mockProduct6)
	mockProduct7 := entities.Product{Name: "product7", Price: 20000, Qty: 35, Description: "bLABFJAEFKJHKJNEKFNKJWAEFAWNJNASNDJ", Category_ID: 3}
	repoProduct.Create(mockProduct7)
	mockProduct8 := entities.Product{Name: "product8", Price: 80000, Qty: 20, Description: "bLABFJAEFKJHKJNEKFNKJWAEFAWNJNASNDJ", Category_ID: 4}
	repoProduct.Create(mockProduct8)

	//payment dummy
	mockPayment1 := entities.Payment{Name: "payment1"}
	repoPayment.Create(mockPayment1)
	mockPayment2 := entities.Payment{Name: "payment2"}
	repoPayment.Create(mockPayment2)
	mockPayment3 := entities.Payment{Name: "payment3"}
	repoPayment.Create(mockPayment3)
	mockPayment4 := entities.Payment{Name: "payment4"}
	repoPayment.Create(mockPayment4)

	//order dummy
	// mockOrder1 := entities.Order{Payment_ID: 1}
	// repoOrder.Create(1, mockOrder1)

	//cart dummy
	mockCart1 := entities.Cart{Product_ID: 1, Qty: 2}
	repoCart.Create(2, mockCart1)
	mockCart2 := entities.Cart{Product_ID: 2, Qty: 1}
	repoCart.Create(2, mockCart2)

}
