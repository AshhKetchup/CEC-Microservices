package db

import (
	"database/sql"
	"fmt"

	"os"
	"time"

	_ "github.com/lib/pq"
)

var (
	dbInstance *sql.DB
)

func InitDB() (*sql.DB, error) {
	//err := godotenv.Load("auth.env")
	//if err != nil {
	//	return nil, fmt.Errorf("no env found")
	//}

	if dbInstance == nil {
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")
		dbUser := os.Getenv("DB_USER")
		sslMode := os.Getenv("DB_SSLMODE")
		fmt.Printf("DB_HOST: %s, DB_PORT: %s\n, DB_USER", dbHost, dbPort, dbName)

		dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s sslmode=%s",
			dbHost, dbPort, dbName, dbUser, sslMode)

		// Open connection
		var err error
		dbInstance, err = sql.Open("postgres", dsn)
		if err != nil {
			return nil, fmt.Errorf("failed to open database: %w", err)
		}

		dbInstance.SetMaxOpenConns(25)
		dbInstance.SetMaxIdleConns(25)
		dbInstance.SetConnMaxLifetime(5 * time.Minute)

		if err = dbInstance.Ping(); err != nil {
			return nil, fmt.Errorf("database ping failed: %w", err)
		}

		if err = createTables(dbInstance); err != nil {
			return nil, fmt.Errorf("schema initialization failed: %w", err)
		}
	}
	print("Database connected!\n")
	return dbInstance, nil
}

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
