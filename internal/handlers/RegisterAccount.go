package handlers

import (
	"github.com/hell077/DiabetesHealthBot/db/sqlite"
	"gopkg.in/telebot.v3"
)

func RegisterAccount(c telebot.Context) error {
	btnBack := Markup.Text("🔙 Назад")
	Markup.Reply(Markup.Row(btnBack))

	err := c.Send("Введите свое имя:", Markup)
	if err != nil {
		return err
	}

	c.Bot().Handle(telebot.OnText, func(ctx telebot.Context) error {
		if ctx.Text() == "🔙 Назад" {
			return ctx.Send("Вы вернулись в главное меню.")
		}

		inputText := ctx.Text()                 // Текст, введённый пользователем
		userRecord := sqlite.User{}             // Пустая запись пользователя для базы
		senderInfo := ctx.Sender()              // Данные о пользователе, отправившем сообщение
		telegramID := senderInfo.ID             // Уникальный ID пользователя в Telegram
		telegramUsername := senderInfo.Username // Username (@...) пользователя

		if err = userRecord.NewUser(sqlite.DB, telegramUsername, inputText, telegramID); err != nil {
			return err
		}

		return ctx.Send("Ваше имя сохранено: " + inputText)
	})

	c.Bot().Handle(&btnBack, func(ctx telebot.Context) error {
		return ctx.Send("Вы вернулись в главное меню.")
	})

	return nil
}
