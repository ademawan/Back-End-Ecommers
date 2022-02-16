package category

import "time"

type RequestCategory struct {
	Name string `json:"name"`
}

type ResponseCategory struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Name string `json:"name"`
}
