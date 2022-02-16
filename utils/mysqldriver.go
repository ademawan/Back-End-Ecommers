package utils

import (
	config "Back-End-Ecommers/configs"
	"Back-End-Ecommers/entities"

	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *config.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name,
	)
	fmt.Println(connectionString)
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Info("error in connect database ", err)
		panic(err)
	}

	AutoMigrate(DB)
	return DB
}

func AutoMigrate(DB *gorm.DB) {

	DB.Migrator().DropTable(&entities.Address{})
	DB.Migrator().DropTable(&entities.User{})
	// DB.Migrator().DropTable(&category.Category{})
	// DB.Migrator().DropTable(&product.Product{})
	// DB.Migrator().DropTable(&payment.Payment{})
	// DB.Migrator().DropTable(&order.Order{})
	// DB.Migrator().DropTable(&order_detail.Order_detail{})
	// DB.Migrator().DropTable(&cart.Cart{})

	DB.AutoMigrate(&entities.User{})
	DB.AutoMigrate(&entities.Address{})
	// DB.AutoMigrate(&category.Category{})
	// DB.AutoMigrate(&product.Product{})
	// DB.AutoMigrate(&payment.Payment{})
	// DB.AutoMigrate(&order.Order{})
	// DB.AutoMigrate(&order_detail.Order_detail{})
	// DB.AutoMigrate(&cart.Cart{})
}
