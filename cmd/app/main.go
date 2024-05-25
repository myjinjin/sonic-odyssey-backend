package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}
	db, err := database.New(
		database.WithHost(os.Getenv("DB_HOST")),
		database.WithPort(os.Getenv("DB_PORT")),
		database.WithUsername(os.Getenv("DB_USERNAME")),
		database.WithPassword(os.Getenv("DB_PASSWORD")),
		database.WithDBName(os.Getenv("DB_NAME")),
	)
	if err != nil {
		log.Fatalf("faied to connect to the database: %v", err)
	}
	defer db.Close()
}
