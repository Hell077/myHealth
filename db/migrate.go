package db

import (
	"fmt"
	sqlite2 "github.com/hell077/DiabetesHealthBot/db/sqlite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Migrate() error {
	err := sqlLiteMigrate()
	if err != nil {
		return err
	}
	return nil
}

func sqlLiteMigrate() error {
	db, err := gorm.Open(sqlite.Open("health.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}
	err = db.AutoMigrate(&sqlite2.User{})
	if err != nil {
		return fmt.Errorf("failed to auto migrate database: %w", err)
	}
	return nil
}
