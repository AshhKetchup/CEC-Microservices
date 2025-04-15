package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbInstance *sql.DB
)

// InitDB initializes and returns a database connection
func InitDB() (*sql.DB, error) {
	// err := godotenv.Load("product.env")
	// if err != nil {
	// 	return nil, fmt.Errorf("no env found")
	// }

	if dbInstance == nil {
		// Get configuration from environment variables
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")

		// Create connection string
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			dbUser, dbPass, dbHost, dbPort, dbName)

		// Open connection
		fmt.Printf("Connecting to database...dns: %s", dsn)
		var err error
		dbInstance, err = sql.Open("mysql", dsn)
		if err != nil {
			return nil, fmt.Errorf("failed to open database: %w", err)
		}

		// Configure connection pool
		// dbInstance.SetMaxOpenConns(25)
		// dbInstance.SetMaxIdleConns(25)
		// dbInstance.SetConnMaxLifetime(5 * time.Minute)

		// Verify connection
		print("Verifying database connection...")
		if err = dbInstance.Ping(); err != nil {
			return nil, fmt.Errorf("database ping failed: %w", err)
		}

		print("Database connection established successfully.")
		// Initialize schema
		if err = createTables(); err != nil {
			return nil, fmt.Errorf("schema initialization failed: %w", err)
		}

		print("Database schema initialized successfully.")
	}
	return dbInstance, nil
}

func createTables() error {
	_, err := dbInstance.Exec(`CREATE TABLE IF NOT EXISTS products (
		id VARCHAR(36) PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		price DECIMAL(10,2) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	return err
}
