package address

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Street      string
	City        string
	Region      string
	Postal_code int
}
