package main

import (
	config "Back-End-Ecommers/configs"
	addressController "Back-End-Ecommers/delivery/controllers/address"
	authController "Back-End-Ecommers/delivery/controllers/auth"
	orderController "Back-End-Ecommers/delivery/controllers/order"
	paymentController "Back-End-Ecommers/delivery/controllers/payment"
	productController "Back-End-Ecommers/delivery/controllers/product"
	userController "Back-End-Ecommers/delivery/controllers/user"
	"Back-End-Ecommers/delivery/route"
	"Back-End-Ecommers/dummy"
	addressRepo "Back-End-Ecommers/repository/address"
	authRepo "Back-End-Ecommers/repository/auth"
	orderRepo "Back-End-Ecommers/repository/order"
	paymentRepo "Back-End-Ecommers/repository/payment"
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

	//REPOSITORY-DATABASE
	userRepo := userRepo.New(db)
	orderRepo := orderRepo.New(db)
	productRepo := productRepo.New(db)
	addressRepo := addressRepo.New(db)
	authRepo := authRepo.New(db)
	paymentRepo := paymentRepo.New(db)

	//CONTROLLER
	orderController := orderController.New(orderRepo)
	userController := userController.New(userRepo)
	productController := productController.New(productRepo)
	addressController := addressController.New(addressRepo)
	authController := authController.New(authRepo)
	paymentController := paymentController.New(paymentRepo)

	e := echo.New()

	route.RegisterPath(
		e,
		userController,
		addressController,
		authController,
		productController,
		orderController,
	)

	route.PaymentMethodPath(
		e,
		paymentController,
	)

	dummy.Dummy()

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
