package category

import (
	"Back-End-Ecommers/entities"
	"time"
)

type RequestCategory struct {
	Name string `json:"name"`
}
type UpdateCategoryRequestFormat struct {
	Name string `json:"name" form:"name"`
}

type ResponseCategory struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Name string `json:"name"`
}
type UpdateProductResponseFormat struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    entities.Category `json:"data"`
}

type DeleteCategoryResponseFormat struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
