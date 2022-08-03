package database

import (
	"github.com/furkanarsl/pf-final/pkg/queries"
	"github.com/jackc/pgx/v4"
)

type DbQueries struct {
	*queries.Queries
	db *pgx.Conn
}

func NewQueries(db *pgx.Conn) DbQueries {
	return DbQueries{
		Queries: queries.New(db),
		db:      db,
	}
}
