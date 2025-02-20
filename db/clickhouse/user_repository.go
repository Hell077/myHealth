package clickhouse

import (
	"context"
	"log"
)

func (db *ClickhouseDB) CreateNewUser(login, id string) error {
	query := "INSERT INTO users (id, login) VALUES (?, ?)"
	if err := db.conn.Exec(context.Background(), query, id, login); err != nil {
		return err
	}
	log.Println("User inserted successfully")
	return nil
}
