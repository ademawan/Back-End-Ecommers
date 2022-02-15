package main

import (
	config "Back-End-Ecommers/configs"
	"Back-End-Ecommers/utils"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.GetConfig()
	db := utils.InitDB(config)
	fmt.Println(db)

	e := echo.New()

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
