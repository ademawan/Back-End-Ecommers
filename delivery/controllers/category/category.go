package category

import (
	"Back-End-Ecommers/delivery/controllers/common"
	"Back-End-Ecommers/delivery/middlewares"
	"Back-End-Ecommers/entities"
	"Back-End-Ecommers/repository/category"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryrController struct {
	repo category.Category
}

func New(repository category.Category) *CategoryrController {
	return &CategoryrController{
		repo: repository,
	}
}

func (cc *CategoryrController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		requestFormat := RequestCategory{}
		email := middlewares.ExtractTokenAdmin(c)[0]
		password := middlewares.ExtractTokenAdmin(c)[1]

		if email != "admin@admin.com" && password != "admin" {
			return c.JSON(http.StatusBadRequest, common.BadRequest(nil, "sorry you are not an admin", nil))
		}

		if err := c.Bind(&requestFormat); err != nil || requestFormat.Name == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest(nil, "Invalid in Input Data Category", nil))
		}

		newCategory := entities.Category{
			Name: requestFormat.Name,
		}

		res, err := cc.repo.Create(newCategory)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(nil, "error in access Create Category", nil))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create Category", res.Name))
	}
}

func (cc *CategoryrController) GetById() echo.HandlerFunc {

	return func(c echo.Context) error {
		categoryid := int(middlewares.ExtractTokenId(c))

		res, err := cc.repo.GetById(categoryid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(nil, "error in access Get Category By id", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get Category By Id", res))
	}
}

func (cc *CategoryrController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := cc.repo.GetAll()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get All Category", res))
	}
}

func (cc *CategoryrController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := middlewares.ExtractTokenAdmin(c)[0]
		password := middlewares.ExtractTokenAdmin(c)[1]

		if email != "admin@admin.com" && password != "admin" {
			return c.JSON(http.StatusBadRequest, common.BadRequest(nil, "sorry you are not an admin", nil))
		}

		categoryId, _ := strconv.Atoi(c.Param("id"))

		newCategory := RequestCategory{}

		if err := c.Bind(&newCategory); err != nil || newCategory.Name == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := cc.repo.Update(categoryId, entities.Category{Name: newCategory.Name})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Update Category", res))
	}
}

func (cc *CategoryrController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := middlewares.ExtractTokenAdmin(c)[0]
		password := middlewares.ExtractTokenAdmin(c)[1]

		if email != "admin@admin.com" && password != "admin" {
			return c.JSON(http.StatusBadRequest, common.BadRequest(nil, "sorry you are not an admin", nil))
		}
		categoryId, _ := strconv.Atoi(c.Param("id"))

		_, err := cc.repo.Delete(categoryId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Delete Category", nil))
	}
}
