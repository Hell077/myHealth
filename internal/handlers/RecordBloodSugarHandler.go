package handlers

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"gopkg.in/telebot.v3"
)

var sugarEntries = make(map[int64]map[string]string)

type SugarLog struct {
	UserID          int64     `db:"user_id"`
	MeasurementTime time.Time `db:"measurement_time"`
	SugarValue      float64   `db:"sugar_value"`
	MealTime        string    `db:"meal_time"`
	CreatedAt       time.Time `db:"created_at"`
}

func RecordBloodSugarHandler(ctx telebot.Context, b *telebot.Bot, db *sql.DB) error {
	userID := ctx.Sender().ID
	Markup.Reply(
		Markup.Row(ToMainMenu),
	)
	sugarEntries[userID] = make(map[string]string)

	_ = ctx.Send("Введите уровень сахара в крови (ммоль/л):", Markup)
	b.Handle(telebot.OnText, func(ctx telebot.Context) error {
		return recordSugarValue(ctx, b, db)
	})
	return nil
}

func recordSugarValue(ctx telebot.Context, b *telebot.Bot, db *sql.DB) error {
	userID := ctx.Sender().ID
	value, err := strconv.ParseFloat(ctx.Text(), 64)
	if err != nil {
		_ = ctx.Send("Ошибка: введите корректное число для уровня сахара.")
		return nil
	}
	sugarEntries[userID]["sugar_value"] = fmt.Sprintf("%.2f", value)

	mealButtons := &telebot.ReplyMarkup{ResizeKeyboard: true}
	btnBeforeMeal := mealButtons.Data("До еды", "before_meal")
	btnAfterMeal := mealButtons.Data("После еды", "after_meal")
	btnRandom := mealButtons.Data("Случайное", "random")

	mealButtons.Inline(
		mealButtons.Row(btnBeforeMeal, btnAfterMeal, btnRandom),
	)

	_ = ctx.Send("Выберите время измерения:", mealButtons)

	b.Handle(&btnBeforeMeal, func(ctx telebot.Context) error {
		sugarEntries[userID]["meal_time"] = "before_meal"
		return saveBloodSugarLog(ctx, b, db)
	})
	b.Handle(&btnAfterMeal, func(ctx telebot.Context) error {
		sugarEntries[userID]["meal_time"] = "after_meal"
		return saveBloodSugarLog(ctx, b, db)
	})
	b.Handle(&btnRandom, func(ctx telebot.Context) error {
		sugarEntries[userID]["meal_time"] = "random"
		return saveBloodSugarLog(ctx, b, db)
	})

	return nil
}

func saveBloodSugarLog(ctx telebot.Context, b *telebot.Bot, db *sql.DB) error {
	userID := ctx.Sender().ID
	measurementTime := time.Now()
	sugarValue, _ := strconv.ParseFloat(sugarEntries[userID]["sugar_value"], 64)
	mealTime := sugarEntries[userID]["meal_time"]
	createdAt := time.Now()

	query := `INSERT INTO health_analytics.sugar_log (user_id, measurement_time, sugar_value, meal_time, created_at) VALUES (?, ?, ?, ?, ?)`
	_, err := db.Exec(query, userID, measurementTime, sugarValue, mealTime, createdAt)
	if err != nil {
		_ = ctx.Send("Ошибка записи в БД: " + err.Error())
		return nil
	}

	var mealTimeText string
	switch mealTime {
	case "before_meal":
		mealTimeText = "До еды"
	case "after_meal":
		mealTimeText = "После еды"
	case "random":
		mealTimeText = "Случайное"
	default:
		mealTimeText = "Неизвестно"
	}

	message := fmt.Sprintf(
		"✅ Запись сохранена!\n\n🩸 Уровень сахара: %.2f ммоль/л\n🍽 Время измерения: %s\n📅 Дата: %s",
		sugarValue, mealTimeText, measurementTime.Format("2006-01-02 15:04:05"),
	)
	_ = ctx.Send(message)

	delete(sugarEntries, userID)
	return nil
}
