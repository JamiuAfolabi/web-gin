package database

import (
	"database/sql"
	"fmt"
	"jamiuafolabi/web-service-gin/models"
	"log"
	"reflect"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

// Converts camelCase strings to snake_case
func toSnakeCase(str string) string {
	var sb strings.Builder
	for i, r := range str {
		if i > 0 && r >= 'A' && r <= 'Z' {
			sb.WriteByte('_')
		}
		sb.WriteRune(r)
	}
	return strings.ToLower(sb.String())
}

// performs bulk insert operations for any slice of structs
func InsertRowsGeneric(conn *sql.DB, tableName string, rows any) error {
	slice := reflect.ValueOf(rows)

	if slice.Kind() != reflect.Slice || slice.Len() == 0 {
		return fmt.Errorf("expected non-empty slice")
	}

	elem := slice.Index(0)
	elemType := elem.Type()

	var columns []string
	for i := 0; i < elemType.NumField(); i++ {
		columns = append(columns, toSnakeCase(elemType.Field(i).Name))
	}

	placeholder := make([]string, len(columns))
	for i := range columns {
		placeholder[i] = fmt.Sprintf("$%d", i+1)
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		tableName,
		strings.Join(columns, ", "),
		strings.Join(placeholder, ", "),
	)

	for i := 0; i < slice.Len(); i++ {
		item := slice.Index(i)
		var values []interface{}
		for j := 0; j < item.NumField(); j++ {
			values = append(values, item.Field(j).Interface())
		}

		if _, err := conn.Exec(query, values...); err != nil {
			log.Printf("Insert error into %s: %v", tableName, err)
			return err
		}
	}

	return nil
}

// populates the category table with sample data
func InsertCategories(conn *sql.DB) error {
	categories := []models.Category{
		{Id: 1, Name: "Electronics"},
		{Id: 2, Name: "Books"},
	}
	return InsertRowsGeneric(conn, "category", categories)
}

// populates the products table with sample data
func InsertProducts(conn *sql.DB) error {
	products := []models.Product{
		{Id: 1, Name: "Laptop", CategoryId: 1, Price: 1200.00},
		{Id: 2, Name: "Headphones", CategoryId: 1, Price: 199.99},
		{Id: 3, Name: "Notebook", CategoryId: 2, Price: 5.50},
	}
	return InsertRowsGeneric(conn, "product", products)
}

// populates the customers table with sample data
func InsertCustomers(conn *sql.DB) error {
	now := time.Now()
	customers := []models.Customer{
		{Id: 1, Name: "Alice Smith", Email: "alice@example.com", Gender: "Female", Location: "NY", CreatedAt: now},
		{Id: 2, Name: "Bob Johnson", Email: "bob@example.com", Gender: "Male", Location: "CA", CreatedAt: now},
	}
	return InsertRowsGeneric(conn, "customer", customers)
}

// populates the transactions table with sample data
func InsertSalesTransactions(conn *sql.DB) error {
	now := time.Now()
	transactions := []models.SalesTransaction{
		{Id: 1, ProductId: 1, CustomerId: 1, Quantity: 2, TransactionDate: now, UnitPrice: 1200.00, TotalAmount: 2400},
		{Id: 2, ProductId: 2, CustomerId: 2, Quantity: 1, TransactionDate: now, UnitPrice: 199.99, TotalAmount: 199.99},
	}
	return InsertRowsGeneric(conn, "sales_transaction", transactions)
}
