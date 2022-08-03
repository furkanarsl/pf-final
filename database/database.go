package database

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func Open(dataSourceName string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), dataSourceName)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
