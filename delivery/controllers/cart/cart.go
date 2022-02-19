package cart

import (
	"Back-End-Ecommers/delivery/controllers/common"
	"Back-End-Ecommers/delivery/middlewares"
	"Back-End-Ecommers/entities"
	"Back-End-Ecommers/repository/cart"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

type CartController struct {
	repo cart.Cart
}

func New(repository cart.Cart) *CartController {
	return &CartController{
		repo: repository,
	}
}

func (cc *CartController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id := int(middlewares.ExtractTokenId(c))
		newCart := RequestCart{}

		if err := c.Bind(&newCart); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(
				http.StatusBadRequest,
				"There is some problem from input",
				nil))
		}

		res, err := cc.repo.Create(user_id, entities.Cart{Product_ID: newCart.Product_ID, Qty: newCart.Qty})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(
				http.StatusInternalServerError,
				"error in database process",
				nil,
			))
		}

		return c.JSON(http.StatusCreated, common.Success(
			http.StatusCreated,
			"success to create cart",
			res,
		))
	}
}

func (cc *CartController) GetByUserId() echo.HandlerFunc {
	return func(c echo.Context) error {
		// id, _ := strconv.Atoi(c.Param("id"))
		user_id, _ := strconv.Atoi(c.Param("id"))
		res, err := cc.repo.GetByUserId(user_id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(
				http.StatusInternalServerError,
				"error in database process",
				nil,
			))
		}

		return c.JSON(http.StatusOK, ResponseGetCart{
			Code:    http.StatusOK,
			Message: "Success Get Carts by user ID",
			Data:    res,
		})
	}
}

func (cc *CartController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		// id, _ := strconv.Atoi(c.Param("id"))
		user_id := int(middlewares.ExtractTokenId(c))
		upCart := RequestCart{}

		if err := c.Bind(&upCart); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(
				http.StatusBadRequest,
				"error in input cart",
				nil,
			))
		}

		res, err := cc.repo.Update(user_id, entities.Cart{Product_ID: upCart.Product_ID, Qty: upCart.Qty})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(
				http.StatusInternalServerError,
				"error in database process",
				nil,
			))
		}

		return c.JSON(http.StatusOK, common.Success(
			http.StatusOK,
			"success to update cart",
			res,
		))
	}
}

func (cc *CartController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		cart_id, _ := strconv.Atoi(c.Param("id"))

		user_id := int(middlewares.ExtractTokenId(c))

		err := cc.repo.Delete(user_id, cart_id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(
				http.StatusInternalServerError,
				"error in database process",
				nil,
			))
		}

		return c.JSON(http.StatusOK, common.Success(
			http.StatusOK,
			"success to delete cart",
			nil,
		))
	}
}
