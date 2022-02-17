package main

import (
	config "Back-End-Ecommers/configs"
	addressController "Back-End-Ecommers/delivery/controllers/address"
	authController "Back-End-Ecommers/delivery/controllers/auth"
	categoryController "Back-End-Ecommers/delivery/controllers/category"
	productController "Back-End-Ecommers/delivery/controllers/product"
	userController "Back-End-Ecommers/delivery/controllers/user"
	"Back-End-Ecommers/delivery/route"
	addressRepo "Back-End-Ecommers/repository/address"
	authRepo "Back-End-Ecommers/repository/auth"
	categoryRepo "Back-End-Ecommers/repository/category"
	productRepo "Back-End-Ecommers/repository/product"
	userRepo "Back-End-Ecommers/repository/user"
	"Back-End-Ecommers/utils"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.GetConfig()
	db := utils.InitDB(config)

	userRepo := userRepo.New(db)
	authRepo := authRepo.New(db)
	productRepo := productRepo.New(db)
	addressRepo := addressRepo.New(db)
	categoryRepo := categoryRepo.New(db)

	authController := authController.New(authRepo)
	userController := userController.New(userRepo)
	productController := productController.New(productRepo)
	addressController := addressController.New(addressRepo)
	categoryController := categoryController.New(categoryRepo)
	e := echo.New()

	route.RegisterPath(e, authController, userController, addressController, productController, categoryController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
