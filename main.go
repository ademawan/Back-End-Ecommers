package main

import (
	config "Back-End-Ecommers/configs"
	addressController "Back-End-Ecommers/delivery/controllers/address"
	userController "Back-End-Ecommers/delivery/controllers/user"
	"Back-End-Ecommers/delivery/route"
	AddressRepo "Back-End-Ecommers/repository/address"
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
	addressRepo := AddressRepo.New(db)

	userController := userController.New(userRepo)
	addressController := addressController.New(addressRepo)
	e := echo.New()

	route.RegisterPath(e, userController, addressController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
