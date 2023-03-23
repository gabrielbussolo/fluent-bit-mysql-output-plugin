package internal

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Mysql struct {
	stmt *sql.Stmt
}

func New(dsn, table string) (*Mysql, error) {
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	stmt, err := conn.Prepare("INSERT INTO " + table + " (datetime, tag, data) VALUES (?, ?, ?)")
	if err != nil {
		return nil, err
	}
	return &Mysql{
		stmt: stmt,
	}, nil
}

func (m *Mysql) Write(ctx context.Context, time time.Time, tag string, data []byte) error {
	_, err := m.stmt.ExecContext(ctx, time, tag, data)
	if err != nil {
		return err
	}
	return nil
}

func (m *Mysql) WriteAsync(ctx context.Context, time, tag, data string) error {
	// TODO
	return nil
}
