package migration

import (
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"os"
)

type Migration struct {
	Dir  string
	sqlx *sqlx.DB
}

func (m *Migration) RunMigration() error {
	if _, err := os.Stat(m.Dir); os.IsNotExist(err) {
		return err
	}

	if err := goose.Up(m.sqlx.DB, m.Dir); err != nil {
		return err
	}

	if err := goose.Status(m.sqlx.DB, m.Dir); err != nil {
		return err
	}
	return nil
}
