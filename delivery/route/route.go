package route

import (
	"Back-End-Ecommers/delivery/controllers/address"
	"Back-End-Ecommers/delivery/controllers/auth"
	"Back-End-Ecommers/delivery/controllers/payment"
	"Back-End-Ecommers/delivery/controllers/product"
	"Back-End-Ecommers/delivery/controllers/user"
	"Back-End-Ecommers/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uc *user.UserController, ac *address.AddressController, aa *auth.AuthController, pc *product.ProductController) {

	//=========================================================

	//CORS
	e.Use(middleware.CORS())

	//LOGGER
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	//ROUTE REGISTER - LOGIN USERS
	e.POST("users/register", uc.Register())
	e.POST("users/login", aa.Login())

	//ROUTE USERS
	e.GET("users", uc.GetAll(), middlewares.JwtMiddleware())
	e.GET("users/me", uc.GetById(), middlewares.JwtMiddleware())
	e.PUT("users/me", uc.Update(), middlewares.JwtMiddleware())
	e.DELETE("users/me", uc.Delete(), middlewares.JwtMiddleware())

	//===========================================================
	//ROUTE ADDRESS
	ea := e.Group("")

	ea.POST("address", ac.Register(), middlewares.JwtMiddleware())
	ea.GET("address", ac.Get())
	ea.GET("address/:id", ac.GetById(), middlewares.JwtMiddleware())
	ea.PUT("address/:id", ac.Update(), middlewares.JwtMiddleware())
	ea.DELETE("address/:id", ac.Delete(), middlewares.JwtMiddleware())

	ea.POST("products", pc.Register(), middlewares.JwtMiddleware())
	ea.GET("products", pc.Get())
	ea.GET("products/:id", pc.GetById(), middlewares.JwtMiddleware())
	ea.PUT("products/:id", pc.Update(), middlewares.JwtMiddleware())
	ea.DELETE("products/:id", pc.Delete(), middlewares.JwtMiddleware())
}

func PaymentMethodPath(e *echo.Echo, pc *payment.PaymentController) {

	//CORS
	e.Use(middleware.CORS())

	//LOGGER
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	//ROUTE PAYMENT
	e.POST("paymentmethods", pc.Create(), middlewares.JwtMiddleware())
	e.GET("paymentmethods", pc.GetAll(), middlewares.JwtMiddleware())
	e.GET("paymentmethods/:id", pc.GetById(), middlewares.JwtMiddleware())
	e.PUT("paymentmethods/:id", pc.UpdateById(), middlewares.JwtMiddleware())
	e.DELETE("paymentmethods/:id", pc.DeleteById(), middlewares.JwtMiddleware())

}
