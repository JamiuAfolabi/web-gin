package databasedriver

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/marcboeker/go-duckdb"
)

// Database interface to allow easy switching
type DB struct {
	SQL *sql.DB
}

// Database connection that can accessed globally
var DbConn = &DB{}

// Iniatiates database connection
func ConnectSQL(connstring string, db_parameter string) {
	db, err := NewDatabase(connstring, db_parameter)

	fmt.Println(db_parameter)
	if err != nil {
		log.Println("There is an error ", err)
		panic(err)
	}
	log.Println("Database successfully connected to")
	DbConn.SQL = db

}

// creates a new database connection and verifies connectivity
func NewDatabase(connstring string, db_parameter string) (*sql.DB, error) {
	log.Println(db_parameter)

	db, errSql := sql.Open(db_parameter, connstring)
	if errSql != nil {
		log.Println("There is an error while connecting to the database ", errSql)
		return nil, errSql
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}

// reads SQL from a file and returns it as bytes
func ReadSqlFromFile(filepath string) ([]byte, error) {
	sqlBytes, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Failed to read SQL file: %v", err)
		return nil, err
	}
	return sqlBytes, nil

}

// reads and executes SQL statements from a file
func ExecuteByteSqlFile(db *sql.DB, filename string) {
	sqlBytes, _ := ReadSqlFromFile(filename)
	_, err := db.Exec(string(sqlBytes))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Required tables created")
	}

}
