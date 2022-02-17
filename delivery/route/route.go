package route

import (
	"Back-End-Ecommers/delivery/controllers/address"
	"Back-End-Ecommers/delivery/controllers/auth"
	"Back-End-Ecommers/delivery/controllers/category"
	"Back-End-Ecommers/delivery/controllers/product"
	"Back-End-Ecommers/delivery/controllers/user"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, auth *auth.AuthController, uc *user.UserController, ac *address.AddressController, pc *product.ProductController, cc *category.CategoryController) {

	//=========================================================
	//ROUTE USERS
	e.POST("users/register", uc.Register())
	e.POST("users/login", auth.Login())

	e.GET("users", uc.Get())
	e.GET("users/:id", uc.GetById())
	e.PUT("users/:id", uc.Update())
	e.DELETE("users/:id", uc.Delete())

	//===========================================================
	//ROUTE ADDRESS

	e.POST("address", ac.Register())
	e.GET("address", ac.Get())
	e.GET("address/:id", ac.GetById())
	e.PUT("address/:id", ac.Update())
	e.DELETE("address/:id", ac.Delete())

	e.POST("products", pc.Register())
	e.GET("products", pc.Get())
	e.GET("products/:id", pc.GetById())
	e.PUT("products/:id", pc.Update())
	e.DELETE("products/:id", pc.Delete())

	e.POST("category", cc.Register())
	e.GET("category", cc.Get())
	e.GET("category/:id", cc.GetById())
	e.PUT("category/:id", cc.Update())
	e.DELETE("category/:id", cc.Delete())
}
