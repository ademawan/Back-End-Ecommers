package user

import (
	"Back-End-Ecommers/entities"
)

//----------------------------------------------------
//REQUEST FORMAT
//----------------------------------------------------
type InsertUserRequestFormat struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Status   int    `json:"status" form:"status"`
}
type UpdateUserRequestFormat struct {
	Name   string `json:"name" form:"name"`
	Email  string `json:"email" form:"email"`
	Status int    `json:"status" form:"status"`
}

//-----------------------------------------------------
//RESPONSE FORMAT
//-----------------------------------------------------
type InsertUserResponseFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.User `json:"data"`
}

type GetUsersResponseFormat struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    []entities.User `json:"data"`
}

type GetAllUserResponseFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.User `json:"data"`
}

type UpdateUserResponseFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.User `json:"data"`
}

type DeleteUserResponseFormat struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
