package address

import (
	"Back-End-Ecommers/delivery/controllers/common"
	"Back-End-Ecommers/entities"
	"Back-End-Ecommers/repository/address"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AddressController struct {
	repo address.Address
}

func New(repository address.Address) *AddressController {
	return &AddressController{
		repo: repository,
	}
}

func (ac *AddressController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ac.repo.Get()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get All Address", res))
	}
}

func (ac *AddressController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		addressId, _ := strconv.Atoi(c.Param("id"))

		res, err := ac.repo.GetById(addressId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "Not Found", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get Address", res))
	}
}

func (ac *AddressController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		address := InsertAddressRequestFormat{}

		if err := c.Bind(&address); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := ac.repo.Insert(entities.Address{Street: address.Street, City: address.City, Region: address.Region, Postal_code: address.Postal_code, User_ID: address.User_ID})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create Address", res))
	}
}

func (ac *AddressController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newAddress = UpdateAddressRequestFormat{}
		addressId, _ := strconv.Atoi(c.Param("id"))

		if err := c.Bind(&newAddress); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := ac.repo.Update(addressId, entities.Address{Street: newAddress.Street, City: newAddress.City, Region: newAddress.Region, Postal_code: newAddress.Postal_code})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Update Address", res))
	}
}

func (ac *AddressController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		addressId, _ := strconv.Atoi(c.Param("id"))

		err := ac.repo.Delete(addressId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Delete Address", nil))
	}
}
