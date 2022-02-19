package orderdetail

import (
	"Back-End-Ecommers/delivery/controllers/common"
	"Back-End-Ecommers/repository/orderdetail"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderDetailController struct {
	repo orderdetail.OrderDetail
}

func New(repository orderdetail.OrderDetail) *OrderDetailController {
	return &OrderDetailController{
		repo: repository,
	}
}

// func (tc *OrderController) Get() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		res, err := tc.repo.Get()

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))

// 		}

// 		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get All Order", res))
// 	}
// }

// func (tc *OrderController) GetById() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		orderId, _ := strconv.Atoi(c.Param("id"))

// 		res, err := tc.repo.GetById(orderId)

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
// 		}

// 		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get Order By Id", res))
// 	}
// }

func (tc *OrderDetailController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		orderdetail := RegisterOrderDetailRequestFormat{}
		// userId := int(middlewares.ExtractTokenId(c))
		if err := c.Bind(&orderdetail); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := tc.repo.Create(orderdetail.Product_ID, orderdetail.Order_ID, orderdetail.Qty)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))

		}

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create Order", res))
	}
}

// func (tc *OrderController) Update() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var newOrder = UpdateOrderRequestFormat{}
// 		orderId, _ := strconv.Atoi(c.Param("id"))

// 		if err := c.Bind(&newOrder); err != nil {
// 			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
// 		}

// 		res, err := tc.repo.Update(orderId, entities.Order{Payment_ID: newOrder.Payment_ID})

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
// 		}

// 		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Update Order", res))
// 	}
// }

// func (tc *OrderController) Delete() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		orderId, _ := strconv.Atoi(c.Param("id"))

// 		err := tc.repo.Delete(orderId)

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
// 		}

// 		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Delete Order", nil))
// 	}
// }
