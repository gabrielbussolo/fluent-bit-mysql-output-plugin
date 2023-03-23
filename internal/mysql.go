package internal

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	dsn   string
	table string
}

func New(dsn, table string) (*Mysql, error) {
	_, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &Mysql{
		dsn:   dsn,
		table: table,
	}, nil
}

func (m *Mysql) Write(ctx context.Context, time, tag, data string) error {
	// TODO
	return nil
}

func (m *Mysql) WriteAsync(ctx context.Context, time, tag, data string) error {
	// TODO
	return nil
}
