package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

func ConnectDatabase() {
	var err error

	// Database connection details
	connStr := "host=localhost user=postgres password=postgres dbname=ginsample sslmode=disable"

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	fmt.Println("Database connected successfully!")
}
