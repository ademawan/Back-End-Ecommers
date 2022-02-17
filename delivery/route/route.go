package route

import (
	"Back-End-Ecommers/delivery/controllers/address"
	"Back-End-Ecommers/delivery/controllers/user"
	"Back-End-Ecommers/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, uc *user.UserController, ac *address.AddressController) {

	//=========================================================
	//ROUTE USERS
	e.POST("users/register", uc.Register())

	e.GET("users", uc.Get(), middlewares.JwtMiddleware())
	e.GET("users/:id", uc.GetById())
	e.PUT("users/:id", uc.Update())
	e.DELETE("users/:id", uc.Delete())

	//===========================================================
	//ROUTE ADDRESS
	ea := e.Group("")

	ea.POST("address", ac.Register())
	ea.GET("address", ac.Get())
	ea.GET("address/:id", ac.GetById())
	ea.PUT("address/:id", ac.Update())
	ea.DELETE("address/:id", ac.Delete())

}
