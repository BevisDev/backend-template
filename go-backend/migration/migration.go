package migration

import (
	"github.com/BevisDev/backend-template/consts"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"os"
)

type Migration struct {
	Dir     string
	sqlx    *sqlx.DB
	TypeSQL string
}

func (m *Migration) Init() error {
	var dialect string
	switch m.TypeSQL {
	case consts.SQLServer:
		dialect = "mssql"
		break
	case consts.Postgres:
		dialect = "postgres"
		break
	}
	if err := goose.SetDialect(dialect); err != nil {
		return err
	}
	if _, err := os.Stat(m.Dir); os.IsNotExist(err) {
		return err
	}
	return nil
}

func (m *Migration) RunMigration() error {
	if err := m.Init(); err != nil {
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
