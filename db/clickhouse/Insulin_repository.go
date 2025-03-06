package clickhouse

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"time"
)

const (
	shortInsulin = iota + 1
	longInsulin
)

type Insulin struct {
	insulinType string
	name        string
	unit        uint8
}

func (db *ClickhouseDB) NewInsulinLog(ins Insulin, userID string) error {
	query := `
		INSERT INTO insulin_log (user_id, insulinType, name, unit, created_at) 
		VALUES (?, ?, ?, ?, now())
	`
	ctx := context.Background()
	batch, err := db.conn.PrepareBatch(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to prepare batch: %w", err)
	}

	if err := batch.Append(userID, ins.insulinType, ins.name, ins.unit); err != nil {
		return fmt.Errorf("failed to append values: %w", err)
	}

	if err := batch.Send(); err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}

type InsulinLog struct {
	ID          uuid.UUID
	UserID      uint64
	CreatedAt   time.Time
	InsulinType uint8
	Unit        uint8
}

type DayInsulinLog struct {
	Day       time.Time
	TotalUnit int32
}

func GetInsulinLogByDay(db *sql.DB, userID uint64, date time.Time) ([]InsulinLog, error) {
	query := `
    	SELECT id, user_id, created_at, insulinType, unit 
    	FROM health_analytics.insulin_log 
    	WHERE user_id = ? 
    	AND created_at >= toDate(?) 
    	AND created_at < toDate(?) + 1
	`

	ctx := context.Background()
	rows, err := db.QueryContext(ctx, query, userID, date, date)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var logs []InsulinLog
	for rows.Next() {
		var log InsulinLog
		if err := rows.Scan(&log.ID, &log.UserID, &log.CreatedAt, &log.InsulinType, &log.Unit); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		logs = append(logs, log)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return logs, nil
}

// GetInsulinLogByMonth - получение среднего значения инсулина за месяц по каждому дню
func GetInsulinLogByMonth(db *sql.DB, userID uint64, date time.Time) ([]InsulinLog, error) {
	query := `
			SELECT 
    		toDate(created_at) AS day, 
    		SUM(unit) AS total_unit
			FROM health_analytics.insulin_log
			WHERE user_id = ?
  			AND created_at >= toStartOfMonth(toDate(?))
  			AND created_at < toStartOfMonth(toDate(?)) + INTERVAL (1) MONTH
			GROUP BY day
			ORDER BY day;
		`

	ctx := context.Background()
	rows, err := db.QueryContext(ctx, query, userID, date, date)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var logs []InsulinLog
	for rows.Next() {
		var log InsulinLog
		var avgUnit float64
		var day time.Time

		if err := rows.Scan(&day, &avgUnit); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		log.CreatedAt = day
		log.Unit = uint8(avgUnit) // Приведение среднего значения к uint8
		logs = append(logs, log)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return logs, nil
}
