package models

import (
	"time"
)

// Category represents a product category
type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// Product represents a product in inventory
type Product struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	CategoryId int     `json:"category_id,omitempty"`
	Price      float64 `json:"price,omitempty"`
}

// Customer represents a customer in the system
type Customer struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Gender    string    `json:"gender,omitempty"`
	Location  string    `json:"location,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// SalesTransaction struct represents a sales record
type SalesTransaction struct {
	Id              int       `json:"id"`
	ProductId       int       `json:"product_id"`
	CustomerId      int       `json:"customer_id"`
	Quantity        int       `json:"quantity"`
	TransactionDate time.Time `json:"transaction_date"`
	UnitPrice       float64   `json:"unit_price"`
	TotalAmount     float64   `json:"total_amount"`
}

// TotalSalesByCategory struct represents sales data by product category
type TotalSalesByCategory struct {
	Category   string  `json:"category"`
	TotalSales float64 `json:"total_sales"`
}

// RevenueByMonth struct represents monthly revenue trends
type RevenueByMonth struct {
	Month   time.Time `json:"month"`
	Revenue float64   `json:"revenue"`
}

// Top5Revenue struct represents highest-performing products by revenue struct
type Top5Revenue struct {
	Product      string  `json:"product"`
	TotalRevenue float64 `json:"total_revenue"`
}

// SalesPerformanceMetrics struct represents key metrics in sales
type SalesPerformanceMetrics struct {
	TotalOrders    int64   `json:"total_orders"`
	TotalItemsSold int64   `json:"total_items_sold"`
	TotalRevenue   float64 `json:"total_revenue"`
	AvgOrderValue  float64 `json:"avg_order_value"`
}

// CategorySalesDistribution struct represents sales distribution across categories
type CategorySalesDistribution struct {
	Category      string  `json:"category"`
	CategorySales float64 `json:"category_sales"`
	Percentage    float64 `json:"percentage"`
}


// ProductPerformanceByCategory represents product performance within categories
type ProductPerformanceByCategory struct {
	Category string  `json:"category"`
	Product  string  `json:"product"`
	Revenue  float64 `json:"revenue"`
}

// CustomerPurchasePattern represents customer buying behavior
type CustomerPurchasePattern struct {
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Orders     int64   `json:"orders"`
	TotalSpent float64 `json:"total_spent"`
}
