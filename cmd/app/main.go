package main

import (
	"github.com/hell077/DiabetesHealthBot/db"
	"github.com/hell077/DiabetesHealthBot/db/clickhouse"
	"github.com/hell077/DiabetesHealthBot/db/sqlite"
	"github.com/hell077/DiabetesHealthBot/internal"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	err error
)

func main() {
	if os.Getenv("GITHUB_ACTIONS") == "" {
		if err := godotenv.Load("config/.env"); err != nil {
			log.Println("Warning: .env file not found, using system environment variables")
		}
	}

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = db.Migrate()
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
	err = sqlite.InitDatabase()
	err = clickhouse.InitCH()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	if err = internal.RunBot(); err != nil {
		log.Fatalf("Error running bot: %v", err)
	}
}

