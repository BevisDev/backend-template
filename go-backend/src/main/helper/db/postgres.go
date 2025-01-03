package db

import (
	"database/sql"
	"github.com/BevisDev/backend-template/src/main/consts"
)

type postgresDB struct {
	db *sql.DB
}

func NewPostgresDB(port int, host, username, password, database string) DB {
	db, _ := InitDB(port, consts.Postgres, host, username, password, database)
	return &postgresDB{db: db}
}

func (s *postgresDB) Open() *sql.DB {
	return s.db
}

func (s *postgresDB) Close() error {
	return s.db.Close()
}

func (s *postgresDB) Begin() (*sql.Tx, error) {
	return s.db.Begin()
}

func (s *postgresDB) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

func (s *postgresDB) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}
