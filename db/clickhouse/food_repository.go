package clickhouse

import (
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
