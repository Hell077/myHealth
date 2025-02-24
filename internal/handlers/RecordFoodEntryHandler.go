package handlers

import (
	"database/sql"
	"fmt"
	"gopkg.in/telebot.v3"
	"strconv"
	"time"
)

var foodEntries = make(map[int64]map[string]string)

type FoodLog struct {
	UserID   int64     `db:"user_id"`
	MealTime time.Time `db:"meal_time"`
	FoodName string    `db:"food_name"`
	Carbs    float64   `db:"carbs"`
	Fats     float64   `db:"fats"`
	Protein  float64   `db:"protein"`
	Weight   float64   `db:"weight"`
}

func RecordFoodEntryHandler(ctx telebot.Context, b *telebot.Bot, db *sql.DB) error {
	userID := ctx.Sender().ID
	foodEntries[userID] = make(map[string]string)

	_ = ctx.Send("Введите название блюда:")
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
	userID := ctx.Sender().ID
	weight, err := strconv.ParseFloat(ctx.Text(), 64)
	if err != nil {
		_ = ctx.Send("Ошибка: введите число для веса блюда в граммах.")
		return nil
	}
	foodEntries[userID]["weight"] = fmt.Sprintf("%.2f", weight)

	_ = ctx.Send("Введите количество углеводов на 100г:")
	b.Handle(telebot.OnText, func(ctx telebot.Context) error {
		return recordCarbs(ctx, b, db)
	})
	return nil
}

func recordCarbs(ctx telebot.Context, b *telebot.Bot, db *sql.DB) error {
	userID := ctx.Sender().ID
	carbs, err := strconv.ParseFloat(ctx.Text(), 64)
	if err != nil {
		_ = ctx.Send("Ошибка: введите число для углеводов.")
		return nil
	}
	foodEntries[userID]["carbs"] = fmt.Sprintf("%.2f", carbs)

	_ = ctx.Send("Введите количество жиров на 100г:")
	b.Handle(telebot.OnText, func(ctx telebot.Context) error {
		return recordFats(ctx, b, db)
	})
	return nil
}

func recordFats(ctx telebot.Context, b *telebot.Bot, db *sql.DB) error {
	userID := ctx.Sender().ID
	fats, err := strconv.ParseFloat(ctx.Text(), 64)
	if err != nil {
		_ = ctx.Send("Ошибка: введите число для жиров.")
		return nil
	}
	foodEntries[userID]["fats"] = fmt.Sprintf("%.2f", fats)

	_ = ctx.Send("Введите количество белков на 100г:")
	b.Handle(telebot.OnText, func(ctx telebot.Context) error {
		return recordProteins(ctx, b, db)
	})
	return nil
}

func recordProteins(ctx telebot.Context, b *telebot.Bot, db *sql.DB) error {
	userID := ctx.Sender().ID
	protein, err := strconv.ParseFloat(ctx.Text(), 64)
	if err != nil {
		_ = ctx.Send("Ошибка: введите число для белков.")
		return nil
	}
	foodEntries[userID]["proteins"] = fmt.Sprintf("%.2f", protein)

	weight, _ := strconv.ParseFloat(foodEntries[userID]["weight"], 64)
	carbs, _ := strconv.ParseFloat(foodEntries[userID]["carbs"], 64)
	fats, _ := strconv.ParseFloat(foodEntries[userID]["fats"], 64)
	protein, _ = strconv.ParseFloat(foodEntries[userID]["proteins"], 64)

	finalCarbs := (carbs / 100) * weight
	finalFats := (fats / 100) * weight
	finalProtein := (protein / 100) * weight

	mealTime := time.Now()
	foodName := foodEntries[userID]["name"]

	query := `INSERT INTO food_log (user_id, meal_time, food_name, carbs, fat, protein, weight) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err = db.Exec(query, userID, mealTime, foodName, finalCarbs, finalFats, finalProtein, weight)

	if err != nil {
		_ = ctx.Send("Ошибка записи в БД: " + err.Error())
		return nil
	}

	_ = ctx.Send(fmt.Sprintf("Запись сохранена! ✅\n\n🍽 Блюдо: %s\n⚖️ Вес: %.2f г\n🥔 Углеводы: %.2f г\n🥑 Жиры: %.2f г\n🍗 Белки: %.2f г",
		foodName, weight, finalCarbs, finalFats, finalProtein))

	delete(foodEntries, userID)
	return nil
}
