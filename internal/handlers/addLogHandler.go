package handlers

import "gopkg.in/telebot.v3"

func AddLogHandler(ctx telebot.Context) error {
	Markup.Reply(
		Markup.Row(RecordInsulinEntry, RecordFoodEntry, RecordBloodSugar),
	)
	return nil
}
