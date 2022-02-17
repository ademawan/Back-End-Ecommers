package order

import (
	"Back-End-Ecommers/entities"
)

//----------------------------------------------------
//REQUEST FORMAT
//----------------------------------------------------
type RegisterOrderRequestFormat struct {
	Payment_ID int `json:"payment_id" form:"payment_id"`
}

type UpdateOrderRequestFormat struct {
	Payment_ID int `json:"payment_id" form:"payment_id"`
}

//-----------------------------------------------------
//RESPONSE FORMAT
//-----------------------------------------------------
type RegisterOrderResponseFormat struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    entities.Order `json:"data"`
}

type GetOrderrResponseFormat struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    []entities.Order `json:"data"`
}

type GetOrderResponseFormat struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    entities.Order `json:"data"`
}

type UpdateResponseFormat struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    entities.Order `json:"data"`
}

type DeleteResponseFormat struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ReopenOrderResponFormat struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    entities.Order `json:"data"`
}

type CompleteOrderResponFormat struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    entities.Order `json:"data"`
}
