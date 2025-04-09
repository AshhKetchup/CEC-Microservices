package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "cec-project/delivery-app/mysql"
)

var (
	dbInstance *sql.DB
	dbOnce     sync.Once
)

// InitDB initializes the database connection (call once at startup)
func InitDB() (*sql.DB, error) {
	var initErr error

	dbOnce.Do(func() {
		// Get environment variables
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")

		// Create connection string
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			user, password, host, port, dbname)

		// Open connection
		dbInstance, initErr = sql.Open("mysql", dsn)
		if initErr != nil {
			return
		}

		// Configure connection pool
		dbInstance.SetMaxOpenConns(25)
		dbInstance.SetMaxIdleConns(25)
		dbInstance.SetConnMaxLifetime(5 * time.Minute)

		// Verify connection
		initErr = dbInstance.Ping()
		if initErr != nil {
			log.Printf("Failed to ping database: %v", initErr)
			return
		}

		log.Println("Successfully connected to database")
	})

	return dbInstance, initErr
}

// GetDB returns the singleton database instance
func GetDB() *sql.DB {
	return dbInstance
}
