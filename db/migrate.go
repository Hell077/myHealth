package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	clickhouseMigrate "github.com/golang-migrate/migrate/v4/database/clickhouse"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

func Migrate() error {
	err := sqlLiteMigrate()
	if err != nil {
		return err
	}

	err = clickHouseMigrate()
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
	err = db.AutoMigrate(&User{}, SugarLevel{}, FoodEntry{}, InsulinEntry{}, ActivityEntry{}, HealthStatus{})
	if err != nil {
		return fmt.Errorf("failed to auto migrate database: %w", err)
	}

	return nil
}

func clickHouseMigrate() error {
	migrationsPath := "file://./clickhouse"

	dsn := os.Getenv("CLICKHOUSE_DSN")
	if dsn == "" {
		return fmt.Errorf("строка подключения ClickHouse не задана в переменной окружения CLICKHOUSE_DSN")
	}

	db, err := sql.Open("clickhouse", dsn)
	if err != nil {
		return fmt.Errorf("не удалось подключиться к ClickHouse: %w", err)
	}
	defer db.Close()

	driver, err := clickhouseMigrate.WithInstance(db, &clickhouseMigrate.Config{})
	if err != nil {
		return fmt.Errorf("не удалось создать драйвер миграции: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "clickhouse", driver)
	if err != nil {
		return fmt.Errorf("не удалось создать миграции: %w", err)
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("не удалось выполнить миграцию: %w", err)
	}

	log.Println("Миграция ClickHouse выполнена успешно")
	return nil
}
