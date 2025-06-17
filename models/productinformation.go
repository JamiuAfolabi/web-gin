// models/productinformation.go
package models

import (
	"time"
)

type ProductInformation struct {
	ID        string
	Name      string
	Category  string
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
