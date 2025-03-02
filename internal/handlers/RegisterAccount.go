package handlers

import (
	"github.com/hell077/DiabetesHealthBot/db/clickhouse"
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
			return StartHandle(ctx)
		}

		inputText := ctx.Text()
		userRecord := sqlite.User{}
		senderInfo := ctx.Sender()
		telegramID := senderInfo.ID
		telegramUsername := senderInfo.Username

		exists, err := userRecord.ExistsByTelegramID(sqlite.DB, telegramID)
		if err != nil {
			return err
		}
		if exists {
			return ctx.Send("Вы уже зарегистрированы!")
		}

		if err := userRecord.NewUser(sqlite.DB, telegramUsername, inputText, telegramID); err != nil {
			return err
		}
		clickhouse.CH.Exec("INSERT INTO health_analytics.users (tgID) values (?)", telegramID)
		return ctx.Send("Ваше имя сохранено: " + inputText)
	})

	c.Bot().Handle(&btnBack, func(ctx telebot.Context) error {
		return StartHandle(ctx)
	})

	return nil
}
