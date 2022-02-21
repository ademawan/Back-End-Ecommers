package main

import (
	config "Back-End-Ecommers/configs"
	addressController "Back-End-Ecommers/delivery/controllers/address"
	authController "Back-End-Ecommers/delivery/controllers/auth"
	cartController "Back-End-Ecommers/delivery/controllers/cart"
	categoryController "Back-End-Ecommers/delivery/controllers/category"
	orderController "Back-End-Ecommers/delivery/controllers/order"
	orderdetailController "Back-End-Ecommers/delivery/controllers/orderdetail"
	paymentController "Back-End-Ecommers/delivery/controllers/payment"
	productController "Back-End-Ecommers/delivery/controllers/product"
	userController "Back-End-Ecommers/delivery/controllers/user"
	"Back-End-Ecommers/delivery/route"
	addressRepo "Back-End-Ecommers/repository/address"
	authRepo "Back-End-Ecommers/repository/auth"
	cartRepo "Back-End-Ecommers/repository/cart"
	categoryRepo "Back-End-Ecommers/repository/category"
	orderRepo "Back-End-Ecommers/repository/order"
	orderdetailRepo "Back-End-Ecommers/repository/orderdetail"
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
	categoryRepo := categoryRepo.New(db)
	cartRepo := cartRepo.New(db)
	orderRepo := orderRepo.New(db)
	productRepo := productRepo.New(db)
	addressRepo := addressRepo.New(db)
	authRepo := authRepo.New(db)
	paymentRepo := paymentRepo.New(db)
	orderdetailRepo := orderdetailRepo.New(db)

	//CONTROLLER
	cartController := cartController.New(cartRepo)
	orderController := orderController.New(orderRepo)
	categoryController := categoryController.New(categoryRepo)
	userController := userController.New(userRepo)
	productController := productController.New(productRepo)
	addressController := addressController.New(addressRepo)
	authController := authController.New(authRepo)
	paymentController := paymentController.New(paymentRepo)
	orderdetailController := orderdetailController.New(orderdetailRepo)

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
	route.CartPath(
		e,
		cartController,
	)
	route.CategoryPath(
		e,
		categoryController,
	)
	route.OrderDetailPath(
		e,
		orderdetailController,
	)

	// dummy.Dummy()

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
