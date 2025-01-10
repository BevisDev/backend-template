package db

import "github.com/jmoiron/sqlx"

type oracleDB struct {
	db *sqlx.DB
}
