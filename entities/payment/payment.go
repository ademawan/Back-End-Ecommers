package payment

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Name string
}
