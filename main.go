package main

import (
	"fmt"
	"jamiuafolabi/web-service-gin/config"
	"jamiuafolabi/web-service-gin/database"
	"jamiuafolabi/web-service-gin/databasedriver"
	"jamiuafolabi/web-service-gin/handlers"
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

// setup of the required configs
var app config.AppConfig
var infoLog *log.Logger
var dbConfig = config.DatabaseConfig{}

// connect and initializes database with required data
func run() {
	// Connect to database
	// database.DatabaseConnection()
	dbConfig.DSN = "duckdb"
	dbConfig.DBSchemaFile = "migrations/duckdbschema.sql"
	// database.DuckDbDatabaseConnection()

	dbConfig.Host = "localhost"
	dbConfig.Port = 5432
	dbConfig.Dbname = "postgres"
	dbConfig.Username = ""
	dbConfig.Password = ""

	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Dbname, dbConfig.Password)

	fmt.Println(psqlSetup)

	duckdbConnString := "sales_data.duckdb"

	databasedriver.ConnectSQL(duckdbConnString, dbConfig.DSN)

	databasedriver.ExecuteByteSqlFile(databasedriver.DbConn.SQL, dbConfig.DBSchemaFile)

	database.InsertCategories(databasedriver.DbConn.SQL)
	database.InsertProducts(databasedriver.DbConn.SQL)
	database.InsertCustomers(databasedriver.DbConn.SQL)
	database.InsertSalesTransactions(databasedriver.DbConn.SQL)

	repo := handlers.NewRepo(&app, databasedriver.DbConn)
	handlers.NewHandlers(repo)
}

// Application entry point
func main() {
	run()

	route := gin.Default()
	defer databasedriver.DbConn.SQL.Close()

	route.GET("/getproducts", handlers.Repo.GetAllProduct)
	route.GET("/finance/totalsalesbycat", handlers.Repo.GetTotalSalesByCategory)
	route.GET("/finance/revenuebymonth", handlers.Repo.GetTotalRevenueByMonth)
	route.GET("/finance/performingproducts", handlers.Repo.GetTop5Products)
	route.GET("/finance/salesperformance", handlers.Repo.GetSalesPerformance)
	route.GET("/sales/salesdistributionbycat", handlers.Repo.GetSalesDistributionByCategory)
	route.GET("/sales/productperformancebycat", handlers.Repo.GetProductPerformanceByCat)
	route.GET("/sales/topcustomerpurchase", handlers.Repo.GetTop20PurchasePatter)

	err := route.Run(":8080")
	if err != nil {
		panic(err)
	}

}
