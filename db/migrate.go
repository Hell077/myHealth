package db

import (
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/clickhouse"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	sqlite2 "github.com/hell077/DiabetesHealthBot/db/sqlite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

func Migrate() error {
	if err := sqlLiteMigrate(); err != nil {
		return err
	}
	if err := clickhouseMigrate(); err != nil {
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
	fmt.Println("sqlite migrate applied successfully")
	return nil
}

func clickhouseMigrate() error {
	migrationsPath, _ := filepath.Abs("./db/clickhouse/migrations")
	migrationsPath = filepath.ToSlash(migrationsPath)

	clickhouseDSN := os.Getenv("CLICKHOUSE_DSN")

	m, err := migrate.New("file://"+migrationsPath, clickhouseDSN)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}
	defer func() {
		srcErr, dbErr := m.Close()
		if srcErr != nil {
			log.Printf("migration source close error: %v", srcErr)
		}
		if dbErr != nil {
			log.Printf("migration database close error: %v", dbErr)
		}
	}()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %w", err)
	}

	log.Println("clickhouse migrate applied successfully")
	return nil
}

func DownClickhouseMigrate() error {
	migrationsPath, _ := filepath.Abs("./db/clickhouse/migrations")
	migrationsPath = filepath.ToSlash(migrationsPath)

	clickhouseDSN := os.Getenv("CLICKHOUSE_DSN")
	m, err := migrate.New("file://"+migrationsPath, clickhouseDSN)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}
	defer func() {
		srcErr, dbErr := m.Close()
		if srcErr != nil {
			log.Printf("migration source close error: %v", srcErr)
		}
		if dbErr != nil {
			log.Printf("migration database close error: %v", dbErr)
		}
	}()
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %w", err)
	}
	log.Println("Migration down successfully")
	return nil
}

func DownSqliteMigrate() error {
	db, err := gorm.Open(sqlite.Open("health.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}
	err = db.Migrator().DropTable(&sqlite2.User{})
	if err != nil {
		return fmt.Errorf("failed to drop table: %w", err)
	}
	fmt.Println("Table dropped successfully")
	return nil
}
