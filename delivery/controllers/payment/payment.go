package category

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
		requestFormat := RequestPayment{}

		if err := c.Bind(&requestFormat); err != nil || requestFormat.Name == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest(nil, "Invalid in Input Data Category", nil))
		}

		newPayment := entities.Payment{
			Name: requestFormat.Name,
		}

		res, err := pc.repo.Create(newPayment)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(nil, "error in apcess Create Category", nil))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Supcess Create Category", res.Name))
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
			"success to get all project",
			res,
		))
	}
}

func (pc *PaymentController) GetById() echo.HandlerFunc {

	return func(c echo.Context) error {
		paymentId := int(middlewares.ExtractTokenId(c))

		res, err := pc.repo.GetById(paymentId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(nil, "error in apcess Get Category By id", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Supcess Get Category By Id", res))
	}
}

func (ac *PaymentController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := int(middlewares.ExtractTokenId(c))
		var newUser = UpdateUserRequestFormat{}

		if err := c.Bind(&newUser); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := ac.repo.Update(userId, entities.Payment{Name: newUser.Name})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Update User", res))
	}
}
