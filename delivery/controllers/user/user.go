package user

import (
	"Back-End-Ecommers/delivery/controllers/common"
	"Back-End-Ecommers/entities"
	"Back-End-Ecommers/repository/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	repo user.User
}

func New(repository user.User) *UserController {
	return &UserController{
		repo: repository,
	}
}

func (ac *UserController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ac.repo.Get()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get All User", res))
	}
}

func (ac *UserController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, _ := strconv.Atoi(c.Param("id"))

		res, err := ac.repo.GetById(userId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "Not Found", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get User", res))
	}
}

func (ac *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := InsertUserRequestFormat{}

		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := ac.repo.Register(entities.User{Name: user.Name, Email: user.Email, Password: user.Password, Status: user.Status})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create User", res))
	}
}

func (ac *UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newUser = UpdateUserRequestFormat{}
		userId, _ := strconv.Atoi(c.Param("id"))

		if err := c.Bind(&newUser); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := ac.repo.Update(userId, entities.User{Name: newUser.Name, Email: newUser.Email, Status: newUser.Status})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Update User", res))
	}
}

func (ac *UserController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, _ := strconv.Atoi(c.Param("id"))

		err := ac.repo.Delete(userId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Delete User", nil))
	}
}
