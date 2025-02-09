package db

import (
	"time"
)

// User - таблица пользователей
type User struct {
	ID        uint   `gorm:"primaryKey"`
	TGID      int64  `gorm:"uniqueIndex"` // ID пользователя в Telegram
	Username  string `gorm:"size:100"`    // Никнейм пользователя
	CreatedAt time.Time
}

// SugarLevel - таблица для хранения уровня сахара
type SugarLevel struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"index"`
	Level      float64   // Уровень сахара в крови
	RecordedAt time.Time `gorm:"autoCreateTime"` // Время записи
	User       User      `gorm:"foreignKey:UserID"`
}

// FoodEntry - таблица для хранения информации о еде
type FoodEntry struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"index"`
	Food       string    `gorm:"size:255"` // Что съел
	Carbs      float64   // Углеводы (г)
	Proteins   float64   // Белки (г)
	Fats       float64   // Жиры (г)
	RecordedAt time.Time `gorm:"autoCreateTime"`
	User       User      `gorm:"foreignKey:UserID"`
}

// InsulinEntry - таблица для хранения информации об инсулине
type InsulinEntry struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"index"`
	Type       string    `gorm:"size:50"` // Тип инсулина (быстрый, базальный)
	Dose       float64   // Доза (ед.)
	RecordedAt time.Time `gorm:"autoCreateTime"`
	User       User      `gorm:"foreignKey:UserID"`
}

// ActivityEntry - таблица для хранения физической активности
type ActivityEntry struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"index"`
	Type       string    `gorm:"size:100"` // Тип активности (ходьба, бег и т.д.)
	Duration   int       // Длительность в минутах
	RecordedAt time.Time `gorm:"autoCreateTime"`
	User       User      `gorm:"foreignKey:UserID"`
}

// HealthStatus - таблица для хранения самочувствия
type HealthStatus struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"index"`
	Symptoms   string    `gorm:"size:255"` // Описание симптомов
	RecordedAt time.Time `gorm:"autoCreateTime"`
	User       User      `gorm:"foreignKey:UserID"`
}
