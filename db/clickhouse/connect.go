package clickhouse

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	_ "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"os"
)

type ch struct {
	dsn string
	err error
}

type ClickhouseDB struct {
	conn driver.Conn
}

func clickhouseDsn(dsn string) *ch {
	return &ch{dsn: dsn}
}

func (db *ClickhouseDB) Close() {
	db.conn.Close()
}

func NewClickhouse() (*ClickhouseDB, error) {
	clh := clickhouseDsn(os.Getenv("CLICKHOUSE_ADDR"))
	var (
		ctx       = context.Background()
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{clh.dsn},
			Auth: clickhouse.Auth{
				Username: os.Getenv("CLICKHOUSE_USERNAME"),
				Password: os.Getenv("CLICKHOUSE_PASSWORD"),
				Database: os.Getenv("CLICKHOUSE_DATABASE"),
			},
			ClientInfo: clickhouse.ClientInfo{
				Products: []struct {
					Name    string
					Version string
				}{
					{Name: "an-example-go-client", Version: "0.1"},
				},
			},
			Debugf: func(format string, v ...interface{}) {
				fmt.Printf(format, v)
			},
			TLS: &tls.Config{
				InsecureSkipVerify: true,
			},
		})
	)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, err
	}
	return &ClickhouseDB{conn: conn}, nil
}
