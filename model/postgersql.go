package model

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func connect() (*pgx.Conn, error) {
	var c Config
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.Username, c.Password, c.Host, c.Port, c.DBName, c.SSLMode,
	))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
