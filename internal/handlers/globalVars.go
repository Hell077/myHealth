package handlers

import "gopkg.in/telebot.v3"

var Markup = &telebot.ReplyMarkup{
	ResizeKeyboard:  true,
	OneTimeKeyboard: false, // Клавиатура остается открытой после нажатия
}

var (
	// Начальные кнопки для входа/регистрации
	RegBtn   = Markup.Text("👤 Создать аккаунт")
	AuthBtn  = Markup.Text("🔓 Войти в аккаунт")
	GuestBtn = Markup.Text("🔍 Гостевой режим")
	HelpBtn  = Markup.Text("❓ Помощь")

	// Основные функции
	RecordInsulinEntry = Markup.Text("💉 Записать дозу инсулина")
	RecordBloodSugar   = Markup.Text("🩸 Записать уровень сахара")
	RecordFoodEntry    = Markup.Text("🍔 Записать прием пищи")

	// Статистика
	GetDailyStats = Markup.Text("📅 Статистика за день")
	GetMonthStats = Markup.Text("📊 Статистика за месяц")

	// Настройки
	Settings = Markup.Text("⚙ Настройки")

	ToMainMenu = Markup.Text("В меню🔙")
)
