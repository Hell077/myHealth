package handlers

import "gopkg.in/telebot.v3"

func HelpHandler(c telebot.Context) error {
	if err := c.Send("Этот бот предназначен для отслеживания здоровья пользователей с диабетом. Он помогает вести учет уровня сахара в крови, приема пищи, физической активности и введения инсулина."); err != nil {
		return err
	}
	return nil
}
