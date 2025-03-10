package Records

import (
	"database/sql"
	"fmt"
	"github.com/hell077/DiabetesHealthBot/internal/handlers"
	"gopkg.in/telebot.v3"
	"strconv"
	"time"
)

var foodEntries = make(map[int64]map[string]string)

type FoodLog struct {
	UserID   uint8     `db:"user_id"`
	MealTime time.Time `db:"meal_time"`
	FoodName string    `db:"food_name"`
	Carbs    float64   `db:"carbs"`
	Fats     float64   `db:"fats"`
	Protein  float64   `db:"protein"`
	Weight   float64   `db:"weight"`
}

func RecordFoodEntryHandler(ctx telebot.Context, b *telebot.Bot, db *sql.DB) error {
	userID := ctx.Sender().ID
	handlers.Markup.Reply(
		handlers.Markup.Row(handlers.ToMainMenu),
	)
	foodEntries[userID] = make(map[string]string)

	_ = ctx.Send("Введите название блюда:", handlers.Markup)
	b.Handle(telebot.OnText, func(ctx telebot.Context) error {
		return recordFoodName(ctx, b, db)
	})
	return nil
}

func recordFoodName(ctx telebot.Context, b *telebot.Bot, db *sql.DB) error {
	userID := ctx.Sender().ID
	foodEntries[userID]["name"] = ctx.Text()
	_ = ctx.Send("Введите вес блюда в граммах:")
	b.Handle(telebot.OnText, func(ctx telebot.Context) error {
		return recordWeight(ctx, b, db)
	})
	return nil
}

func recordWeight(ctx telebot.Context, b *telebot.Bot, db *sql.DB) error {
	return recordNutrient(ctx, b, db, "weight", "Введите количество углеводов на 100г:", recordCarbs)
}

func recordCarbs(ctx telebot.Context, b *telebot.Bot, db *sql.DB) error {
	return recordNutrient(ctx, b, db, "carbs", "Введите количество жиров на 100г:", recordFats)
}

func recordFats(ctx telebot.Context, b *telebot.Bot, db *sql.DB) error {
	return recordNutrient(ctx, b, db, "fats", "Введите количество белков на 100г:", recordProteins)
}

func recordProteins(ctx telebot.Context, b *telebot.Bot, db *sql.DB) error {
	return recordNutrient(ctx, b, db, "proteins", "", saveFoodLog)
}

func recordNutrient(ctx telebot.Context, b *telebot.Bot, db *sql.DB, key, nextPrompt string, nextStep func(telebot.Context, *telebot.Bot, *sql.DB) error) error {
	userID := ctx.Sender().ID
	value, err := strconv.ParseFloat(ctx.Text(), 64)
	if err != nil {
		_ = ctx.Send(fmt.Sprintf("Ошибка: введите число для %s.", key))
		return nil
	}
	foodEntries[userID][key] = fmt.Sprintf("%.2f", value)

	if nextPrompt != "" {
		_ = ctx.Send(nextPrompt)
		b.Handle(telebot.OnText, func(ctx telebot.Context) error {
			return nextStep(ctx, b, db)
		})
	} else {
		return nextStep(ctx, b, db)
	}

	return nil
}

func saveFoodLog(ctx telebot.Context, b *telebot.Bot, db *sql.DB) error {
	userID := ctx.Sender().ID

	weight, _ := strconv.ParseFloat(foodEntries[userID]["weight"], 64)
	carbs, _ := strconv.ParseFloat(foodEntries[userID]["carbs"], 64)
	fats, _ := strconv.ParseFloat(foodEntries[userID]["fats"], 64)
	protein, _ := strconv.ParseFloat(foodEntries[userID]["proteins"], 64)

	finalCarbs := (carbs / 100) * weight
	finalFats := (fats / 100) * weight
	finalProtein := (protein / 100) * weight

	mealTime := time.Now()
	foodName := foodEntries[userID]["name"]

	query := `INSERT INTO health_analytics.food_log (user_id, meal_time, food_name, carbs, fat, protein, weight) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, userID, mealTime, foodName, finalCarbs, finalFats, finalProtein, weight)

	if err != nil {
		_ = ctx.Send("Ошибка записи в БД: " + err.Error())
		return nil
	}

	_ = ctx.Send(fmt.Sprintf(
		"✅ Запись сохранена!\n\n🍽 Блюдо: %s\n⚖️ Вес: %.2f г\n🥔 Углеводы: %.2f г\n🥑 Жиры: %.2f г\n🍗 Белки: %.2f г",
		foodName, weight, finalCarbs, finalFats, finalProtein,
	))

	delete(foodEntries, userID)
	return nil
}
