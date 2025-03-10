package Records

import (
	"database/sql"
	"fmt"
	"github.com/hell077/DiabetesHealthBot/internal/handlers"
	"strconv"
	"time"

	"gopkg.in/telebot.v3"
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
	handlers.Markup.Reply(
		handlers.Markup.Row(handlers.ToMainMenu),
	)
	insulinKeyboard := &telebot.ReplyMarkup{}
	btnLong := insulinKeyboard.Data("Длинный инсулин", "insulin_long")
	btnShort := insulinKeyboard.Data("Короткий инсулин", "insulin_short")

	insulinKeyboard.Inline(
		insulinKeyboard.Row(btnLong, btnShort),
	)

	_ = ctx.Send("Выберите тип инсулина:", insulinKeyboard, handlers.Markup)

	bot.Handle(&btnLong, func(ctx telebot.Context) error {
		insulinEntries[userID] = InsulinEntry{
			UserID:      userID,
			InsulinType: 2, // Длинный инсулин
			CreatedAt:   time.Now(),
		}
		return ctx.Send("✅ Вы выбрали длинный инсулин. Введите название инсулина:")
	})

	bot.Handle(&btnShort, func(ctx telebot.Context) error {
		insulinEntries[userID] = InsulinEntry{
			UserID:      userID,
			InsulinType: 1, // Короткий инсулин
			CreatedAt:   time.Now(),
		}
		return ctx.Send("✅ Вы выбрали короткий инсулин. Введите название инсулина:")
	})

	bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
		return recordInsulinName(ctx, bot, db)
	})

	return nil
}

func recordInsulinName(ctx telebot.Context, bot *telebot.Bot, db *sql.DB) error {
	userID := ctx.Sender().ID
	entry, exists := insulinEntries[userID]
	if !exists || entry.Name != "" {
		return nil
	}

	entry.Name = ctx.Text()
	insulinEntries[userID] = entry

	_ = ctx.Send("Введите количество единиц инсулина:")

	bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
		return recordInsulinUnits(ctx, bot, db)
	})

	return nil
}

func recordInsulinUnits(ctx telebot.Context, bot *telebot.Bot, db *sql.DB) error {
	userID := ctx.Sender().ID
	entry, exists := insulinEntries[userID]
	if !exists || entry.Unit != 0 {
		return nil
	}

	units, err := strconv.Atoi(ctx.Text())
	if err != nil || units <= 0 {
		_ = ctx.Send("Ошибка: введите корректное число единиц инсулина.")
		return nil
	}

	entry.Unit = uint8(units)
	insulinEntries[userID] = entry

	if err := saveInsulinEntry(db, entry); err != nil {
		_ = ctx.Send("Ошибка записи в базу данных: " + err.Error())
		return nil
	}

	message := fmt.Sprintf("✅ Запись сохранена!\n\n💉 Инсулин: %s\n🔢 Единицы: %d\n📅 Дата: %s",
		entry.Name, entry.Unit, entry.CreatedAt.Format("2006-01-02 15:04:05"))
	_ = ctx.Send(message)

	delete(insulinEntries, userID)
	return nil
}

func saveInsulinEntry(db *sql.DB, entry InsulinEntry) error {
	query := `INSERT INTO health_analytics.insulin_log (user_id, insulinType, name, unit, created_at) VALUES (?, ?, ?, ?, ?)`
	_, err := db.Exec(query, entry.UserID, entry.InsulinType, entry.Name, entry.Unit, entry.CreatedAt)
	return err
}
