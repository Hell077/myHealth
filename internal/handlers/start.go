package handlers

import (
	"github.com/hell077/DiabetesHealthBot/db/sqlite"
	"gopkg.in/telebot.v3"
)

func StartHandle(c telebot.Context) error {
	userID := c.Sender().ID
	newUser := sqlite.User{}

	exists, err := newUser.ExistsByTelegramID(sqlite.DB, userID)
	if err != nil {
		return c.Send("Произошла ошибка при проверке пользователя.")
	}
	if exists {
		Markup.Reply(
			Markup.Row(AuthBtn),
			Markup.Row(GuestBtn, HelpBtn),
		)
	} else {
		Markup.Reply(
			Markup.Row(RegBtn, AuthBtn),
			Markup.Row(GuestBtn, HelpBtn),
		)
	}

	return c.Send("Выберите опцию:", Markup)
}
