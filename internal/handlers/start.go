package handlers

import "gopkg.in/telebot.v3"

func StartHandle(c telebot.Context) error {

	guest := Markup.Text("🔍 Гостевой режим")
	reg := Markup.Text("👤 Создать аккаунт")
	login := Markup.Text("🔓 Войти в аккаунт")
	help := Markup.Text("❓ Помощь")

	Markup.Reply(
		Markup.Row(reg, login),
		Markup.Row(guest, help),
	)

	return c.Send("Выберите опцию:", Markup)
}
