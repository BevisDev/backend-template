package db

import (
	"database/sql"
	"github.com/BevisDev/backend-template/src/main/consts"
)

type sqlServerDB struct {
	db *sql.DB
}

func NewSQLServerDB(port int, host, username, password, database string) DB {
	db, _ := InitDB(port, consts.SQLServer, host, username, password, database)
	return &sqlServerDB{db: db}
}

func (s *sqlServerDB) Open() *sql.DB {
	return s.db
}

func (s *sqlServerDB) Close() error {
	return s.db.Close()
}

func (s *sqlServerDB) Begin() (*sql.Tx, error) {
	return s.db.Begin()
}

func (s *sqlServerDB) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

func (s *sqlServerDB) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}
