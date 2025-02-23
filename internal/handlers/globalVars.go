package handlers

import "gopkg.in/telebot.v3"

var Markup = &telebot.ReplyMarkup{
	ResizeKeyboard: true,
}

var (
	RegBtn             = Markup.Text("👤 Создать аккаунт")
	AuthBtn            = Markup.Text("🔓 Войти в аккаунт")
	HelpBtn            = Markup.Text("❓ Помощь")
	GuestBtn           = Markup.Text("🔍 Гостевой режим")
	RecordInsulinEntry = Markup.Text("💉 Записать дозу инсулина")
	RecordBloodSugar   = Markup.Text("🩸 Записать уровень сахара")
	RecordFoodEntry    = Markup.Text("🍔 Записать прием пищи")
)
