package sqlite

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	TGID      int64  `gorm:"uniqueIndex"` // ID пользователя в Telegram
	Username  string `gorm:"size:100"`    // Никнейм пользователя
	CreatedAt time.Time
}
