package repository

import "jamiuafolabi/web-service-gin/models"

// Interface abstraction for database operations
type DatabaseRepo interface {
	GetAllProductInformation() ([]models.Product, error)
	TotalSalesByCategory() ([]models.TotalSalesByCategory, error)
	RevenueTrendMonthly() ([]models.RevenueByMonth, error)
	Top5PerformingProduct() ([]models.Top5Revenue, error)
	SalesPerformanceMetric() ([]models.SalesPerformanceMetrics, error)
	SalesDistributionByCategory() ([]models.CategorySalesDistribution, error)
	ProductPerformanceByCategory() ([]models.ProductPerformanceByCategory, error)
	Top20CustomerPurchase() ([]models.CustomerPurchasePattern, error)
}
