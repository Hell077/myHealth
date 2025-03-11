package Auth

import (
	"github.com/hell077/DiabetesHealthBot/db/clickhouse"
	"github.com/hell077/DiabetesHealthBot/db/sqlite"
	"github.com/hell077/DiabetesHealthBot/internal/handlers"
	"gopkg.in/telebot.v3"
)

func RegisterAccount(c telebot.Context) error {
	btnBack := handlers.Markup.Text("🔙 Назад")
	handlers.Markup.Reply(handlers.Markup.Row(btnBack))

	err := c.Send("Введите свое имя:", handlers.Markup)
	if err != nil {
		return err
	}

	c.Bot().Handle(telebot.OnText, handleUserInput)
	c.Bot().Handle(&btnBack, func(ctx telebot.Context) error {
		return handlers.StartHandle(ctx)
	})

	return nil
}

func handleUserInput(ctx telebot.Context) error {
	if ctx.Text() == "🔙 Назад" {
		return handlers.StartHandle(ctx)
	}

	inputText := ctx.Text()
	senderInfo := ctx.Sender()
	telegramID := senderInfo.ID
	telegramUsername := senderInfo.Username

	exists, err := checkUserExists(telegramID)
	if err != nil {
		return err
	}
	if exists {
		return ctx.Send("Вы уже зарегистрированы!")
	}

	if err := createNewUser(telegramUsername, inputText, telegramID); err != nil {
		return err
	}

	if err := insertIntoAnalytics(telegramID); err != nil {
		return err
	}

	setupUserMenu()
	return ctx.Send("Ваше имя сохранено: "+inputText, handlers.Markup)
}

func checkUserExists(telegramID int64) (bool, error) {
	userRecord := sqlite.User{}
	return userRecord.ExistsByTelegramID(sqlite.DB, telegramID)
}

func createNewUser(username, name string, telegramID int64) error {
	userRecord := sqlite.User{}
	return userRecord.NewUser(sqlite.DB, username, name, telegramID)
}

func insertIntoAnalytics(telegramID int64) error {
	_, err := clickhouse.CH.Exec("INSERT INTO health_analytics.users (tgID) values (?)", telegramID)
	return err
}

func setupUserMenu() {
	handlers.Markup.Reply(
		handlers.Markup.Row(handlers.RecordFoodEntry, handlers.RecordInsulinEntry, handlers.RecordBloodSugar),
		handlers.Markup.Row(handlers.GetDailyStats, handlers.GetMonthStats),
		handlers.Markup.Row(handlers.Settings),
	)
}
