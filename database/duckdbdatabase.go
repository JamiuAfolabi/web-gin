package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/marcboeker/go-duckdb"
)

func DuckDbDatabaseConnection() {
	db, err := sql.Open("duckdb", "sales_data.duckdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("This is duckdb connection")
	_, err = db.Exec(`CREATE TABLE if not exists people  (id INTEGER, name VARCHAR)`)
	if err != nil {
		log.Fatal(err)
	}
	sqlBytes, _ := ReadSqlFromFile("migrations/duckdbschema.sql")
	_, err = db.Exec(string(sqlBytes))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Required tables created")
	}

	_, err = db.Exec(`INSERT INTO people VALUES (42, 'John')`)
	if err != nil {
		log.Fatal(err)
	}

	var (
		id   int
		name string
	)
	row := db.QueryRow(`SELECT id, name FROM people`)
	err = row.Scan(&id, &name)
	if errors.Is(err, sql.ErrNoRows) {
		log.Println("no rows")
	} else if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("id: %d, name: %s\n", id, name)
}

func ReadSqlFromFile(filepath string) ([]byte, error) {
	sqlBytes, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Failed to read SQL file: %v", err)
		return nil, err
	}
	return sqlBytes, nil

}

func ExecuteByteSqlFile(db *sql.DB, filename string) {
	sqlBytes, _ := ReadSqlFromFile(filename)
	_, err := db.Exec(string(sqlBytes))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Required tables created")
	}

}
