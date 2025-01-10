package db

import (
	"github.com/jmoiron/sqlx"
)

type postgresDB struct {
	db *sqlx.DB
}
