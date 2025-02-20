package handlers

import "gopkg.in/telebot.v3"

func StartHandle(c telebot.Context) error {

	return c.Send("123")
}
