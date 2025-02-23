package internal

import (
	"errors"
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

}
