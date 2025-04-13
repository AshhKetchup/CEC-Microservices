package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var (
	dbInstance *sql.DB
)

// InitDB initializes and returns a database connection
func InitDB() (*sql.DB, error) {
	err := godotenv.Load("auth.env")
	if err != nil {
		return nil, fmt.Errorf("no env found")
	}

	if dbInstance == nil {
		// Get configuration from environment variables
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")
		sslMode := os.Getenv("DB_SSLMODE")

		// Create connection string
		dsn := fmt.Sprintf("host=%s port=%s dbname=%s sslmode=%s",
			dbHost, dbPort, dbName, sslMode)

		// Open connection
		dbInstance, err = sql.Open("postgres", dsn)
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
		if err = createTables(dbInstance); err != nil {
			return nil, fmt.Errorf("schema initialization failed: %w", err)
		}
	}
	print("Database connected!\n")
	return dbInstance, nil
}

// Add these database initialization steps in your application setup
func createTables(dbInstance *sql.DB) error {
	_, err := dbInstance.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto";`)
	if err != nil {
		return fmt.Errorf("failed to enable pgcrypto extension: %w", err)
	}
	_, err = dbInstance.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			email TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			role TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT now()
		);
	`)
	return err
}
