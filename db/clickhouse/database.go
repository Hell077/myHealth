package clickhouse

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/ClickHouse/clickhouse-go/v2"
)

var CH *sql.DB

func InitCH() error {
	dsn := fmt.Sprintf("clickhouse://localhost:9000?username=%s&password=%s&x-multi-statement=true", os.Getenv("CLICKHOUSE_USER"), os.Getenv("CLICKHOUSE_PASS"))

	CH, err := sql.Open("clickhouse", dsn)
	if err != nil {
		return err
	}

	if err := CH.Ping(); err != nil {
		log.Printf("Failed to ping ClickHouse: %v", err)
		return err
	}

	fmt.Println("Connected to ClickHouse successfully!")
	return nil
}

func CloseCH() {
	if CH != nil {
		CH.Close()
		fmt.Println("ClickHouse connection closed.")
	}
}
