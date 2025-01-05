package db

import (
	"database/sql"
	"github.com/BevisDev/backend-template/src/main/consts"
)

type oracleDB struct {
	db *sql.DB
}

func NewOracleDB(port int, host, username, password, database string) DB {
	db, _ := NewDB(port, consts.Oracle, host, username, password, database)
	return &oracleDB{db: db}
}

func (h *oracleDB) Open() *sql.DB {
	return h.db
}

func (h *oracleDB) Close() error {
	return h.db.Close()
}

func (h *oracleDB) Begin() (*sql.Tx, error) {
	return h.db.Begin()
}

func (h *oracleDB) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

func (h *oracleDB) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}
