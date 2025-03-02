package handlers

import (
	"database/sql"
	"gopkg.in/telebot.v3"
	"time"
)

type InsulinEntry struct {
	UserID      int64     `db:"user_id"`
	InsulinType uint8     `db:"insulinType"`
	Name        string    `db:"name"`
	Unit        uint8     `db:"unit"`
	CreatedAt   time.Time `db:"created_at"`
}

var insulinEntries = make(map[int64]InsulinEntry)

// Обработчик выбора инсулина
func RecordInsulinEntryHandler(ctx telebot.Context, bot *telebot.Bot, db *sql.DB) error {
	userID := ctx.Sender().ID

	insulinKeyboard := &telebot.ReplyMarkup{}
	btnLong := insulinKeyboard.Data("Длинный инсулин", "insulin_long")
	btnShort := insulinKeyboard.Data("Короткий инсулин", "insulin_short")

	insulinKeyboard.Inline(
		insulinKeyboard.Row(btnLong, btnShort),
	)

	_ = ctx.Send("Выберите тип инсулина:", insulinKeyboard)

	bot.Handle(&btnLong, func(ctx telebot.Context) error {
		insulinEntries[userID] = InsulinEntry{
			UserID:      userID,
			InsulinType: 2, // Длинный инсулин
			CreatedAt:   time.Now(),
		}
		return ctx.Send("✅ Вы выбрали длинный инсулин")
	})

	// Обработчик выбора короткого инсулина
	bot.Handle(&btnShort, func(ctx telebot.Context) error {
		insulinEntries[userID] = InsulinEntry{
			UserID:      userID,
			InsulinType: 1,
			CreatedAt:   time.Now(),
		}
		return ctx.Send("✅ Вы выбрали короткий инсулин")
	})

	return nil
}

func recordInsulinName(ctx telebot.Context, bot *telebot.Bot, db *sql.DB) error {}
