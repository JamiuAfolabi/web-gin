package main

import (
	"fmt"
	"jamiuafolabi/web-service-gin/config"
	"jamiuafolabi/web-service-gin/database"
	"jamiuafolabi/web-service-gin/databasedriver"
	"jamiuafolabi/web-service-gin/handlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

var app config.AppConfig
var infoLog *log.Logger
var dbConfig = config.DatabaseConfig{}

func helloWorld(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func run() {
	// Connect to database
	// database.DatabaseConnection()
	app.DSN = "duckdb"
	app.DBSchemaFile = "migrations/duckdbschema.sql"
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

	databasedriver.ConnectSQL(duckdbConnString, app.DSN)

	databasedriver.ExecuteByteSqlFile(databasedriver.DbConn.SQL, app.DBSchemaFile)

	database.InsertRowProductInformation(databasedriver.DbConn.SQL)
	database.InsertSalesTransaction(databasedriver.DbConn.SQL)
	errRows := database.GetAllRows(databasedriver.DbConn.SQL)

	if errRows != nil {
		log.Println("There is an error connecting")
	}

	repo := handlers.NewRepo(&app, databasedriver.DbConn)
	handlers.NewHandlers(repo)
}

func main() {
	run()

	route := gin.Default()
	defer databasedriver.DbConn.SQL.Close()

	route.GET("/ping", handlers.Repo.HelloWorld)

	err := route.Run(":8080")
	if err != nil {
		panic(err)
	}

}
