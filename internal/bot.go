package internal

import (
	"errors"
	"github.com/hell077/DiabetesHealthBot/db/clickhouse"
	"github.com/hell077/DiabetesHealthBot/internal/handlers"
	"gopkg.in/telebot.v3"
	"os"
	"time"
)

func RunBot() error {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		return errors.New("BOT_TOKEN environment variable not set")
	}
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}
	bot, err := telebot.NewBot(pref)
	if err != nil {
		return err
	}

	RegisterHandlers(bot)

	bot.Start()
	return nil
}

func RegisterHandlers(bot *telebot.Bot) {
	bot.Handle("/help", handlers.HelpHandler)
	bot.Handle("/start", handlers.StartHandle)
	bot.Handle("👤 Создать аккаунт", handlers.RegisterAccount)
	bot.Handle("🔓 Войти в аккаунт", handlers.AuthHandler)

	bot.Handle("В меню🔙", handlers.ToMenuBtn)
	bot.Handle("⚙ Настройки", handlers.SettingHandler)

	bot.Handle("🩸 Записать уровень сахара", func(ctx telebot.Context) error {
		return handlers.RecordBloodSugarHandler(ctx, bot, clickhouse.CH)
	})
	bot.Handle("🍔 Записать прием пищи", func(ctx telebot.Context) error {
		return handlers.RecordFoodEntryHandler(ctx, bot, clickhouse.CH)
	})
	bot.Handle("💉 Записать дозу инсулина", func(ctx telebot.Context) error {
		return handlers.RecordInsulinEntryHandler(ctx, bot, clickhouse.CH)
	})
}
