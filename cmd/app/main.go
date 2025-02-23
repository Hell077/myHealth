package main

import (
	"github.com/hell077/DiabetesHealthBot/db"
	"github.com/hell077/DiabetesHealthBot/db/sqlite"
	"github.com/hell077/DiabetesHealthBot/internal"
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
	err = sqlite.InitDatabase()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	if err = internal.RunBot(); err != nil {
		log.Fatalf("Error running bot: %v", err)
	}
}
