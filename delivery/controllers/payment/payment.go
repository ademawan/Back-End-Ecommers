package payment

import (
	"Back-End-Ecommers/delivery/controllers/common"
	"Back-End-Ecommers/delivery/middlewares"
	"Back-End-Ecommers/entities"
	"Back-End-Ecommers/repository/payment"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	repo payment.Payment
}

func New(repository payment.Payment) *PaymentController {
	return &PaymentController{
		repo: repository,
	}
}

func (pc *PaymentController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := middlewares.ExtractTokenAdmin(c)[0]
		password := middlewares.ExtractTokenAdmin(c)[1]

		if email != "admin@admin.com" && password != "admin" {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "invalid input", nil))
		}

		requestFormat := RequestPayment{}

		if err := c.Bind(&requestFormat); err != nil || requestFormat.Name == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest(nil, "Invalid in Input Data Payment", nil))
		}

		newPayment := entities.Payment{
			Name: requestFormat.Name,
		}

		res, err := pc.repo.Create(newPayment)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(nil, "error in apcess Create Payment", nil))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Supcess Create Payment", res.Name))
	}
}

func (pc *PaymentController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		// user_id := int(middlewares.ExtractTokenId(c))

		res, err := pc.repo.GetAll()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(
				http.StatusInternalServerError,
				"error in database process",
				nil,
			))
		}

		return c.JSON(http.StatusCreated, common.Success(
			http.StatusCreated,
			"success to get all Payment",
			res,
		))
	}
}

func (pc *PaymentController) GetById() echo.HandlerFunc {

	return func(c echo.Context) error {
		paymentId := int(middlewares.ExtractTokenId(c))

		res, err := pc.repo.GetById(paymentId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(nil, "error in apcess Get Payment By id", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Supcess Get Payment By Id", res))
	}
}

func (ac *PaymentController) UpdateById() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := middlewares.ExtractTokenAdmin(c)[0]
		password := middlewares.ExtractTokenAdmin(c)[1]

		if email != "admin@admin.com" && password != "admin" {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "invalid input", nil))
		}

		paymentId := int(middlewares.ExtractTokenId(c))
		var newPayment = UpdateUserRequestFormat{}

		if err := c.Bind(&newPayment); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := ac.repo.UpdateById(paymentId, entities.Payment{Name: newPayment.Name})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Update Payment", res))
	}
}

func (ac *PaymentController) DeleteById() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := middlewares.ExtractTokenAdmin(c)[0]
		password := middlewares.ExtractTokenAdmin(c)[1]

		if email != "admin@admin.com" && password != "admin" {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "invalid input", nil))
		}

		paymentId := int(middlewares.ExtractTokenId(c))

		res, err := ac.repo.DeleteById(paymentId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Delete Payment", res))
	}
}
