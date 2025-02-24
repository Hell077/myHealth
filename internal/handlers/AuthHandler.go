package handlers

import (
	"github.com/hell077/DiabetesHealthBot/db/sqlite"
	"gopkg.in/telebot.v3"
)

func AuthHandler(ctx telebot.Context) error {
	userID := ctx.Sender().ID
	user := sqlite.User{}

	if status, err := user.ExistsByTelegramID(sqlite.DB, userID); err != nil {
		return ctx.Send("Произошла ошибка на сервере")
	} else if !status {
		Markup.Reply(
			Markup.Row(RegBtn, AuthBtn),
			Markup.Row(GuestBtn, HelpBtn),
		)
		return ctx.Send("Пользователь не найден, зарегистрируйтесь или войдите как гость", Markup)
	}

	Markup.Reply(
		Markup.Row(RecordFoodEntry, RecordInsulinEntry, RecordBloodSugar),
		Markup.Row(GetDailyStats, GetMonthStats),
		Markup.Row(Settings),
	)
	return ctx.Send("Добро пожаловать! Выберите действие:", Markup)
}
