package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Migrate() error {
	db, err := gorm.Open(sqlite.Open("health.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database")
	}
	err = db.AutoMigrate(&User{}, SugarLevel{}, FoodEntry{}, InsulinEntry{}, ActivityEntry{}, HealthStatus{})
	if err != nil {
		return fmt.Errorf("failed to connect database")
	}
	return nil
}
