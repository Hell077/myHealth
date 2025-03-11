package clickhouse

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/google/uuid"
	"log"
	"time"
)

type SugarLog struct {
	Id       uuid.UUID
	UserID   uint64
	SugarLvl float64
	MealTime string
}

func NewSugarLog(sLog SugarLog, db clickhouse.Conn) error {
	query := `
		INSERT INTO health_analytics.sugar_log (user_id, sugar_value, meal_time) 
		VALUES (?, ?, ?)
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	batch, err := db.PrepareBatch(ctx, query)
	if err != nil {
		log.Printf("Batch preparation failed: %v", err)
		return fmt.Errorf("failed to prepare batch: %w", err)
	}
	defer batch.Abort()

	if err := batch.Append(sLog.UserID, sLog.SugarLvl, sLog.MealTime); err != nil {
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

func GetSugarLogByDay(userID uint64, db *sql.DB, date time.Time) ([]SugarLog, error) {
	query := `
		SELECT id, user_id, sugar_value, meal_time 
		FROM health_analytics.sugar_log 
		WHERE user_id = ? 
		AND meal_time >= toDate(?) 
		AND meal_time < toDate(?) + 1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, query, userID, date, date)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var logs []SugarLog
	for rows.Next() {
		var log SugarLog
		if err := rows.Scan(&log.Id, &log.UserID, &log.SugarLvl, &log.MealTime); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		logs = append(logs, log)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return logs, nil
}

func GetSugarLogByMonth(userID uint64, db *sql.DB, date time.Time) ([]SugarLog, error) {
	query := `
		SELECT id, user_id, sugar_value, meal_time 
		FROM health_analytics.sugar_log 
		WHERE user_id = ? 
		AND meal_time >= toStartOfMonth(toDate(?)) 
		AND meal_time < toStartOfMonth(toDate(?)) + INTERVAL (1) MONTH
		ORDER BY meal_time
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, query, userID, date, date)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var logs []SugarLog
	for rows.Next() {
		var log SugarLog
		if err := rows.Scan(&log.Id, &log.UserID, &log.SugarLvl, &log.MealTime); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		logs = append(logs, log)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return logs, nil
}
