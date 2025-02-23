package handlers

import (
	"github.com/hell077/DiabetesHealthBot/db/sqlite"
	"gopkg.in/telebot.v3"
)

func AuthHandler(ctx telebot.Context) error {
	userID := ctx.Sender().ID
	user := sqlite.User{}
	if status, err := user.ExistsByTelegramID(sqlite.DB, userID); err != nil {
		if e := ctx.Send("Произошла ошибка на сервере"); e != nil {
			return e
		}
	} else if status == false {
		_ = ctx.Send("Пользователь не найден, зарегистрируйтесь или войдите как гость")
		Markup.Reply(
			Markup.Row(AuthBtn),
			Markup.Row(GuestBtn),
		)
	} else if status {
		_ = ctx.Send("Успешный вход")
		Markup.Reply()
	}
	return nil
}
