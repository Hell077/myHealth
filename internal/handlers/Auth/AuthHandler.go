package Auth

import (
	"github.com/hell077/DiabetesHealthBot/db/sqlite"
	"github.com/hell077/DiabetesHealthBot/internal/handlers"
	"gopkg.in/telebot.v3"
)

func AuthHandler(ctx telebot.Context) error {
	userID := ctx.Sender().ID
	user := sqlite.User{}

	if status, err := user.ExistsByTelegramID(sqlite.DB, userID); err != nil {
		return ctx.Send("Произошла ошибка на сервере")
	} else if !status {
		handlers.Markup.Reply(
			handlers.Markup.Row(handlers.RegBtn, handlers.AuthBtn),
			handlers.Markup.Row(handlers.GuestBtn, handlers.HelpBtn),
		)
		return ctx.Send("Пользователь не найден, зарегистрируйтесь или войдите как гость", handlers.Markup)
	}

	handlers.Markup.Reply(
		handlers.Markup.Row(handlers.RecordFoodEntry, handlers.RecordInsulinEntry, handlers.RecordBloodSugar),
		handlers.Markup.Row(handlers.GetDailyStats, handlers.GetMonthStats),
		handlers.Markup.Row(handlers.Settings),
	)
	return ctx.Send("Добро пожаловать! Выберите действие:", handlers.Markup)
}
