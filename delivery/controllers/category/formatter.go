package category

import "time"

type requestCategory struct {
	Name string `json:"name"`
}

type responseCategory struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Name string `json:"name"`
}
