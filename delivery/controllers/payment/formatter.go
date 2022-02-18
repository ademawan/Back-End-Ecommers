package payment

import "time"

type RequestPayment struct {
	Name string `json:"name" form:"name"`
}

type ResponsePayment struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Name string `json:"name"`
}

type UpdateUserRequestFormat struct {
	Name string `json:"name" form:"name"`
}
