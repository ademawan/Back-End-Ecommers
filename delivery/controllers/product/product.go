package product

import (
	"Back-End-Ecommers/delivery/controllers/common"
	"Back-End-Ecommers/entities"
	"Back-End-Ecommers/repository/product"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	repo product.Product
}

func New(repository product.Product) *ProductController {
	return &ProductController{
		repo: repository,
	}
}

func (ac *ProductController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ac.repo.GetAll()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get All Product", res))
	}
}

func (ac *ProductController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		productId, _ := strconv.Atoi(c.Param("id"))

		res, err := ac.repo.GetById(productId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "Not Found", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get Product", res))
	}
}

func (ac *ProductController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		product := InsertProductRequestFormat{}

		if err := c.Bind(&product); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := ac.repo.Create(entities.Product{Name: product.Name, Price: product.Price, Qty: product.Qty, Description: product.Description, Category_ID: product.Category_ID})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create Product", res))
	}
}

func (ac *ProductController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newProduct = UpdateProductRequestFormat{}
		productId, _ := strconv.Atoi(c.Param("id"))

		if err := c.Bind(&newProduct); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := ac.repo.Update(productId, entities.Product{Name: newProduct.Name, Price: newProduct.Price, Qty: newProduct.Qty, Description: newProduct.Description, Category_ID: newProduct.Category_ID})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Update Product", res))
	}
}

func (ac *ProductController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		productId, _ := strconv.Atoi(c.Param("id"))

		_, err := ac.repo.Delete(productId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Delete Product", nil))
	}
}
