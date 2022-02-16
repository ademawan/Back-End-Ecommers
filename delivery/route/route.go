package route

import (
	"Back-End-Ecommers/configs"
	"Back-End-Ecommers/delivery/controllers/address"
	"Back-End-Ecommers/delivery/controllers/user"

	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uc *user.UserController, ac *address.AddressController) {

	//=========================================================
	//ROUT USERS
	e.POST("users/register", uc.Register())

	e.GET("users", uc.Get())
	e.GET("users/:id", uc.GetById())
	e.PUT("users/:id", uc.Update())
	e.DELETE("users/:id", uc.Delete())

	//===========================================================
	//ROUTE TASK
	ea := e.Group("")
	ea.Use(m.JWT([]byte(configs.JWT_SECRET)))
	ea.POST("address", ac.Register())
	ea.GET("address", ac.Get())
	ea.GET("address/:id", ac.GetById())
	ea.PUT("address/:id", ac.Update())
	ea.DELETE("address/:id", ac.Delete())

}
