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

type CategoryController struct {
	repo category.Category
}

func New(repository category.Category) *CategoryController {
	return &CategoryController{
		repo: repository,
	}
}

func (ac *CategoryController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ac.repo.GetAll()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get All Product", res))
	}
}

func (cc *CategoryController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		requestFormat := RequestCategory{}

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

func (cc *CategoryController) GetById() echo.HandlerFunc {

	return func(c echo.Context) error {
		categoryid := int(middlewares.ExtractTokenId(c))

		res, err := cc.repo.GetById(categoryid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(nil, "error in access Get Category By id", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get Category By Id", res))
	}
}
func (ac *CategoryController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newCategory = UpdateCategoryRequestFormat{}
		productId, _ := strconv.Atoi(c.Param("id"))

		if err := c.Bind(&newCategory); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := ac.repo.Update(productId, entities.Category{Name: newCategory.Name})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Update Product", res))
	}
}

func (ac *CategoryController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		productId, _ := strconv.Atoi(c.Param("id"))

		_, err := ac.repo.Delete(productId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Delete Product", nil))
	}
}
