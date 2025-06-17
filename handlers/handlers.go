package handlers

import (
	"jamiuafolabi/web-service-gin/config"
	"jamiuafolabi/web-service-gin/databasedriver"
	"jamiuafolabi/web-service-gin/repository"
	"jamiuafolabi/web-service-gin/repository/dbrepo"

	"github.com/gin-gonic/gin"
)

// Repo the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

func NewRepo(app *config.AppConfig, db *databasedriver.DB) *Repository {
	return &Repository{
		App: app,
		DB:  dbrepo.NewDBRepo(db.SQL, app),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) GetAllProduct(context *gin.Context) {
	products, err := m.DB.GetAllProductInformation()
	// m.DB.GetAllRows()

	if err != nil {
		context.JSON(500, gin.H{"error": "Unable to fetch products"})
		return
	}

	context.IndentedJSON(200, products)
}

func (m *Repository) GetTotalSalesByCategory(context *gin.Context) {
	products, err := m.DB.TotalSalesByCategory()
	// m.DB.GetAllRows()

	if err != nil {
		context.JSON(500, gin.H{"error": "Unable to fetch total sales"})
		return
	}

	context.IndentedJSON(200, products)
}

func (m *Repository) GetTotalRevenueByMonth(context *gin.Context) {
	products, err := m.DB.TotalSalesByCategory()
	// m.DB.GetAllRows()

	if err != nil {
		context.JSON(500, gin.H{"error": "Unable to fetch total sales"})
		return
	}

	context.IndentedJSON(200, products)
}

func (m *Repository) GetTop5Products(context *gin.Context) {
	products, err := m.DB.Top5PerformingProduct()
	// m.DB.GetAllRows()

	if err != nil {
		context.JSON(500, gin.H{"error": "Unable to fetch total sales"})
		return
	}

	context.IndentedJSON(200, products)
}

func (m *Repository) GetSalesPerformance(context *gin.Context) {
	products, err := m.DB.SalesPerformanceMetric()
	// m.DB.GetAllRows()

	if err != nil {
		context.JSON(500, gin.H{"error": "Unable to fetch total sales"})
		return
	}

	context.IndentedJSON(200, products)
}

func (m *Repository) GetSalesDistributionByCategory(context *gin.Context) {
	products, err := m.DB.SalesDistributionByCategory()
	// m.DB.GetAllRows()

	if err != nil {
		context.JSON(500, gin.H{"error": "Unable to fetch total sales"})
		return
	}

	context.IndentedJSON(200, products)
}

func (m *Repository) GetProductPerformanceByCat(context *gin.Context) {
	products, err := m.DB.ProductPerformanceByCategory()
	// m.DB.GetAllRows()

	if err != nil {
		context.JSON(500, gin.H{"error": "Unable to fetch total sales"})
		return
	}

	context.IndentedJSON(200, products)
}

func (m *Repository) GetTop20PurchasePatter(context *gin.Context) {
	products, err := m.DB.Top20CustomerPurchase()
	// m.DB.GetAllRows()

	if err != nil {
		context.JSON(500, gin.H{"error": "Unable to fetch total sales"})
		return
	}

	context.IndentedJSON(200, products)
}
