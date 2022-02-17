package order

import (
	"Back-End-Ecommers/delivery/controllers/common"
	"Back-End-Ecommers/delivery/middlewares"
	"Back-End-Ecommers/entities"
	"Back-End-Ecommers/repository/order"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	repo order.Order
}

func New(repository order.Order) *OrderController {
	return &OrderController{
		repo: repository,
	}
}

func (tc *OrderController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := tc.repo.Get()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))

		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get All Order", res))
	}
}

func (tc *OrderController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		orderId, _ := strconv.Atoi(c.Param("id"))

		res, err := tc.repo.GetById(orderId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get Order By Id", res))
	}
}

func (tc *OrderController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		order := RegisterOrderRequestFormat{}
		userId := int(middlewares.ExtractTokenId(c))
		if err := c.Bind(&order); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := tc.repo.Create(userId, entities.Order{Payment_ID: order.Payment_ID})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))

		}

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create Order", res))
	}
}

func (tc *OrderController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newOrder = UpdateOrderRequestFormat{}
		orderId, _ := strconv.Atoi(c.Param("id"))

		if err := c.Bind(&newOrder); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := tc.repo.Update(orderId, entities.Order{Payment_ID: newOrder.Payment_ID})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Update Order", res))
	}
}

func (tc *OrderController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		orderId, _ := strconv.Atoi(c.Param("id"))

		err := tc.repo.Delete(orderId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Delete Order", nil))
	}
}
