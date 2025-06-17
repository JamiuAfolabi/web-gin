package models

import "time"

type Transactions struct {
	ID           string
	Product_ID   string
	Quantity     float64
	Date         time.Time
	CustomerInfo string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
