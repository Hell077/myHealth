package main

import (
	"github.com/hell077/DiabetesHealthBot/db"
	"github.com/joho/godotenv"
	"log"
)

var (
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = db.Migrate()
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
}
