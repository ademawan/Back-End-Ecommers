package orderdetail

import (
	"Back-End-Ecommers/entities"
)

//----------------------------------------------------
//REQUEST FORMAT
//----------------------------------------------------
type RegisterOrderDetailRequestFormat struct {
	Order_ID   int `json:"order_id" form:"order_id"`
	Product_ID int `json:"product_id" form:"product_id"`
	Qty        int `json:"qty" form:"qty"`
}

// type UpdateOrderRequestFormat struct {
// 	Payment_ID int `json:"payment_id" form:"payment_id"`
// }

//-----------------------------------------------------
//RESPONSE FORMAT
//-----------------------------------------------------
type RegisterOrderDetailResponseFormat struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    entities.OrderDetail `json:"data"`
}

// type GetOrderrResponseFormat struct {
// 	Code    int              `json:"code"`
// 	Message string           `json:"message"`
// 	Data    []entities.Order `json:"data"`
// }

// type GetOrderResponseFormat struct {
// 	Code    int            `json:"code"`
// 	Message string         `json:"message"`
// 	Data    entities.Order `json:"data"`
// }

// type UpdateResponseFormat struct {
// 	Code    int            `json:"code"`
// 	Message string         `json:"message"`
// 	Data    entities.Order `json:"data"`
// }

// type DeleteResponseFormat struct {
// 	Code    int         `json:"code"`
// 	Message string      `json:"message"`
// 	Data    interface{} `json:"data"`
// }
