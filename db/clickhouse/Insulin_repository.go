package clickhouse

import (
	"context"
	"fmt"
)

const (
	shortInsulin = iota + 1
	longInsulin
)

type insulin struct {
	insulinType string
	name        string
	unit        uint8
}

func (db *ClickhouseDB) NewInsulinLog(ins insulin, userID string) error {
	query := `
		INSERT INTO insulin_log (user_id, insulinType, name, unit) 
		VALUES (?, ?, ?, ?)
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
