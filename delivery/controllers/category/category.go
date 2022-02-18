package category

import (
	"Back-End-Ecommers/delivery/controllers/common"
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

func (cc *CategoryController) Create() echo.HandlerFunc {
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
		categoryid, _ := strconv.Atoi(c.Param("id"))

		res, err := cc.repo.GetById(categoryid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(nil, "error in access Get Category By id", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get Category By Id", res))
	}
}
