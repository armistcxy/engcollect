package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func SpawnDatabaseConnection(dsn string) *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", dsn)
	return db
}
