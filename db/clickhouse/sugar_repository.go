package clickhouse

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
)

type Status int

const (
	beforeMeal Status = iota + 1
	afterMeal
	random
)

var Db *ClickhouseDB

func (db *ClickhouseDB) NewSugarLog(userID uuid.UUID, sugarLvl float32, mealTime Status) error {
	query := `
		INSERT INTO sugar_log (user_id, sugar_value, meal_time) 
		VALUES (?, ?, ?)
	`
	ctx := context.Background()

	batch, err := db.conn.PrepareBatch(ctx, query)
	if err != nil {
		log.Printf("Batch preparation failed: %v", err)
		return fmt.Errorf("failed to prepare batch: %w", err)
	}

	if err := batch.Append(userID, sugarLvl, mealTime); err != nil {
		log.Printf("Batch append failed: %v", err)
		return fmt.Errorf("failed to append values: %w", err)
	}

	if err := batch.Send(); err != nil {
		log.Printf("Batch execution failed: %v", err)
		return fmt.Errorf("failed to execute batch: %w", err)
	}

	log.Println("Log inserted successfully")
	return nil
}

func (db *ClickhouseDB) GetSugarLogByDay(userID uuid.UUID) {

}

func (db *ClickhouseDB) GetSugarLogBy(userID uuid.UUID) {}
