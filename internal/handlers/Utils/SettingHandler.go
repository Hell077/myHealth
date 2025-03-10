package Utils

import (
	"database/sql"
	"fmt"
	"github.com/hell077/DiabetesHealthBot/db/clickhouse"
	"github.com/hell077/DiabetesHealthBot/internal/handlers/Auth"

	"github.com/google/uuid"
	"gopkg.in/telebot.v3"
)

func SettingHandler(ctx telebot.Context) error {
	tgID := ctx.Sender().ID
	id, err := getUserUUID(tgID)
	if err != nil {
		return err
	}

	if id == uuid.Nil {
		return ctx.Send("❌ Ваш аккаунт не найден в системе.")
	}

	return ctx.Send(fmt.Sprintf("✅ Ваш ID для подключения аккаунта: `%s`", id))
}

func getUserUUID(tgID int64) (uuid.UUID, error) {
	var userUUID uuid.UUID

	err := clickhouse.CH.QueryRow("SELECT id FROM health_analytics.users WHERE tgID = ?", tgID).Scan(&userUUID)
	if err != nil {
		if err == sql.ErrNoRows {
			return uuid.Nil, nil
		}
		return uuid.Nil, err
	}

	return userUUID, nil
}

func ToMenuBtn(ctx telebot.Context) error {
	return Auth.AuthHandler(ctx)
}
