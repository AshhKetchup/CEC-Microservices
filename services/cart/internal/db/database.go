package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbInstance *sql.DB
)

func InitDB() (*sql.DB, error) {
	// err := godotenv.Load("cart.env")
	// if err != nil {
	// 	return nil, fmt.Errorf("error loading .env file: %w", err)
	// }

	if dbInstance == nil {
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")

		// Add multiStatements=true to support multiple SQL commands
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true",
			dbUser, dbPass, dbHost, dbPort, dbName)

		var err error
		dbInstance, err = sql.Open("mysql", dsn)
		if err != nil {
			return nil, fmt.Errorf("failed to open database: %w", err)
		}

		// Configure connection pool
		dbInstance.SetMaxOpenConns(25)
		dbInstance.SetMaxIdleConns(25)
		dbInstance.SetConnMaxLifetime(5 * time.Minute)

		// Verify connection
		if err = dbInstance.Ping(); err != nil {
			return nil, fmt.Errorf("database ping failed: %w", err)
		}

		// Initialize schema
		if err = createTables(); err != nil {
			return nil, fmt.Errorf("schema initialization failed: %w", err)
		}
	}
	return dbInstance, nil
}

func createTables() error {
	// Split table creation into separate statements
	_, err := dbInstance.Exec(`
		CREATE TABLE IF NOT EXISTS cart_items (
			user_id VARCHAR(36) NOT NULL,
			product_id VARCHAR(36) NOT NULL,
			quantity INT NOT NULL DEFAULT 1,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			PRIMARY KEY (user_id, product_id)
		);
	`)
	if err != nil {
		return fmt.Errorf("failed to create cart_items table: %w", err)
	}

	_, err = dbInstance.Exec(`
		CREATE TABLE IF NOT EXISTS payments (
			transaction_id VARCHAR(36) PRIMARY KEY,
			user_id VARCHAR(36) NOT NULL,
			amount DECIMAL(10,2) UNSIGNED NOT NULL,
			currency CHAR(3) NOT NULL,
			status ENUM('PENDING', 'COMPLETED', 'FAILED') NOT NULL DEFAULT 'PENDING',
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			INDEX idx_user (user_id)
		);
	`)
	return err
}
