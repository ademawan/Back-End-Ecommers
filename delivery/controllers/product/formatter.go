package product

import (
	"Back-End-Ecommers/entities"
)

//----------------------------------------------------
//REQUEST FORMAT
//----------------------------------------------------
type InsertProductRequestFormat struct {
	Name        string `json:"name" form:"name"`
	Price       int    `json:"price" form:"price"`
	Qty         int    `json:"qty" form:"qty"`
	Description string `json:"description" form:"description"`
	Category_ID int    `json:"category_id" form:"category_id"`
}
type UpdateProductRequestFormat struct {
	Name        string `json:"name" form:"name"`
	Price       int    `json:"price" form:"price"`
	Qty         int    `json:"qty" form:"qty"`
	Description string `json:"description" form:"description"`
	Category_ID int    `json:"category_id" form:"category_id"`
}

//-----------------------------------------------------
//RESPONSE FORMAT
//-----------------------------------------------------
type InsertProductResponseFormat struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    entities.Product `json:"data"`
}

type GetProductsResponseFormat struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    []entities.Product `json:"data"`
}

type GetAllProductResponseFormat struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    entities.Product `json:"data"`
}

type UpdateProductResponseFormat struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    entities.Product `json:"data"`
}

type DeleteProductResponseFormat struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
