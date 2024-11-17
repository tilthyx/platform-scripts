package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connection() (*sql.DB, error) {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	// Get values from the environment
	host := os.Getenv("ORIGIN_HOST")
	port := os.Getenv("ORIGIN_POSTGRES_PORT")
	user := os.Getenv("ORIGIN_USER")
	password := os.Getenv("ORIGIN_PASSWORD")
	dbname := os.Getenv("ORIGIN_DB")

	// Create the connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a connection to the database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to the database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("unable to ping the database: %v", err)
	}

	return db, nil
}
