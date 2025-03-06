package clickhouse

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Eat struct {
	id                 uuid.UUID
	userID             uint64
	weight             float32
	mealTime           time.Time
	foodName           string
	carb, protein, fat float32
}

func getFoodLogByDay(db *sql.DB, userID uint64, date time.Time)
