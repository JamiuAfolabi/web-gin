package database

import (
	"database/sql"
	"fmt"
	"jamiuafolabi/web-service-gin/config"
	"jamiuafolabi/web-service-gin/models"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var dbConfig = config.DatabaseConfig{}
var Db *sql.DB

func DatabaseConnection() {
	dbConfig.Host = "localhost"
	dbConfig.Port = 5432
	dbConfig.Dbname = "postgres"
	dbConfig.Username = "jamiuafolabi"
	dbConfig.Password = "jamiuafolabi"

	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Dbname, dbConfig.Password)

	db, errSql := sql.Open("postgres", psqlSetup)
	if errSql != nil {
		log.Println("There is an error while connecting to the database ", errSql)
		panic(errSql)
	} else {
		Db = db
		log.Println("Successfully connected to database!")
	}

	// defer db.Close()

}

func GetAllRows(conn *sql.DB) error {
	rows, err := conn.Query("select id, name from productinformation")

	if err != nil {
		log.Println("error in connecting to the productinformation table")
		return err
	} else {
		log.Println("connecting to the productinformation table")
	}
	defer rows.Close()

	var id, name string

	for rows.Next() {
		err := rows.Scan(&id, &name)

		if err != nil {
			log.Println("error in retrieving data")
			return err
		}
		log.Println("Record is ", id, name)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning row", err)
	}
	return nil
}

func InsertRowProductInformation(conn *sql.DB) error {
	products := []models.ProductInformation{
		{ID: "1", Name: "Jamiu", Category: "ASP", Price: 123, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: "2", Name: "Alice", Category: "Electronics", Price: 299.99, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: "3", Name: "Bob", Category: "Books", Price: 15.50, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: "4", Name: "Charlie", Category: "Home", Price: 79.95, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: "5", Name: "Diana", Category: "Fitness", Price: 45.00, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	query := `
        INSERT INTO productinformation (id, name, category, price, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6)
    `

	for _, product := range products {
		_, err := conn.Exec(query, product.ID, product.Name, product.Category, product.Price, product.CreatedAt, product.UpdatedAt)

		if err != nil {
			log.Printf("error in inserting product %s", err)
			return err
		}
	}

	return nil

}

func InsertSalesTransaction(conn *sql.DB) error {
	orders := []models.Transactions{
		{ID: "1", Product_ID: "1", Quantity: 2, Date: time.Now(), CustomerInfo: "Alice Smith - alice@example.com", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: "2", Product_ID: "2", Quantity: 2, Date: time.Now(), CustomerInfo: "Bob Johnson - bob@example.com", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: "3", Product_ID: "3", Quantity: 2, Date: time.Now(), CustomerInfo: "John Doe - johndoe@example.com", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: "4", Product_ID: "4", Quantity: 2, Date: time.Now(), CustomerInfo: "Charlie Davis - charlie@example.com", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: "5", Product_ID: "5", Quantity: 2, Date: time.Now(), CustomerInfo: "Diana Prince - diana@example.com", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	query := `
        INSERT INTO salestransaction (id, product_id, quantity, date, customer_info, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `

	for _, order := range orders {
		_, err := conn.Exec(query, order.ID, order.Product_ID, order.Quantity, order.Date, order.CustomerInfo, order.CreatedAt, order.UpdatedAt)

		if err != nil {
			log.Printf("error in inserting sales transaction %s", err)
			return err
		}
	}

	return nil

}
