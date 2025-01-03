package db

import (
	"database/sql"
	"fmt"
	"github.com/BevisDev/backend-template/src/main/consts"
	"github.com/BevisDev/backend-template/src/main/helper/logger"
)

type DB interface {
	Open() *sql.DB
	Close() error
	Begin() (*sql.Tx, error)
	Commit(tx *sql.Tx) error
	Rollback(tx *sql.Tx) error
}

func InitDB(port int, kind, host, username, password, schema string) (*sql.DB, error) {
	var connStr string

	// build connectionString
	switch kind {
	case consts.SQLServer:
		connStr = fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s",
			host, port, username, password, schema)
		break
	case consts.Postgres:
		connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, username, password, schema)
		break
	case consts.Oracle:
		connStr = fmt.Sprintf("%v/%v@%v:%v/%v", username, password, host, port, schema)
		break
	}

	// check err while opening connection
	db, err := sql.Open(kind, connStr)
	if err != nil {
		logger.Panic("", "Error open connection: {}", err)
		return nil, err
	}

	// ping check connection
	if err = db.Ping(); err != nil {
		logger.Panic("", "Error connect to database: {}", err)
		return nil, err
	}

	return db, nil
}
