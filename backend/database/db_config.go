package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func GetConnection() (*sql.DB, error) {
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("Error loading .env file")
		}
	}

	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		DBUser, DBPass, DBHost, DBPort, DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	log.Println("Connecting to database...")
	time.Sleep(5 * time.Second) // Simulate delay for connection

	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxIdleConns(70)
	db.SetMaxOpenConns(100)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed connecting to database: %w", err)
	}

	return db, nil
}
