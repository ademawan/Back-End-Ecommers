package order

// import (
// 	"Back-End-Ecommers/delivery/controllers/auth"
// 	"Back-End-Ecommers/delivery/controllers/common"
// 	"Back-End-Ecommers/entities"
// 	"Back-End-Ecommers/repository/order"
// 	"net/http"
// 	"strconv"

// 	"github.com/labstack/echo/v4"
// )

// type OrderController struct {
// 	repo order.Order
// }

// func New(repository order.Order) *OrderController {
// 	return &OrderController{
// 		repo: repository,
// 	}
// }

// func (tc *OrderController) Get() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		res, err := tc.repo.Get()

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))

// 		}

// 		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get All Task", res))
// 	}
// }

// func (tc *OrderController) GetById() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		orderId, _ := strconv.Atoi(c.Param("id"))

// 		res, err := tc.repo.GetById(orderId)

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
// 		}

// 		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get Taks Id", res))
// 	}
// }

// func (tc *OrderController) TaskRegister() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		task := RegisterTaskRequestFormat{}
// 		userId := int(auth.ExtractTokenUserId(c))
// 		if err := c.Bind(&task); err != nil {
// 			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
// 		}

// 		res, err := tc.repo.TaskRegister(entities.Order{Nama: task.Nama, Priority: task.Priority, User_ID: userId, Project_ID: task.Project_ID, Status: -1})

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))

// 		}

// 		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create Task", res))
// 	}
// }

// func (tc *OrderController) Update() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var newTask = UpdateTaskRequestFormat{}
// 		orderId, _ := strconv.Atoi(c.Param("id"))

// 		if err := c.Bind(&newTask); err != nil {
// 			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
// 		}

// 		res, err := tc.repo.Update(orderId, entities.Order{Nama: newTask.Nama, Priority: newTask.Priority})

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
// 		}

// 		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Update Task", res))
// 	}
// }

// func (tc *OrderController) Delete() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		orderId, _ := strconv.Atoi(c.Param("id"))

// 		err := tc.repo.Delete(orderId)

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server", nil))
// 		}

// 		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Delete Task", nil))
// 	}
// }
