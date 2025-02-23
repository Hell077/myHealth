package sqlite

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	TGID      int64  `gorm:"uniqueIndex"`
	Username  string `gorm:"size:100"`
	Name      string `gorm:"size:100"`
	CreatedAt time.Time
}
