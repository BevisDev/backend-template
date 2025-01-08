package db

import (
	"fmt"
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/BevisDev/backend-template/src/main/consts"
	"github.com/BevisDev/backend-template/src/main/helper/logger"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

var (
	Connections map[string]*sqlx.DB
	ConfigDb    map[string]config.Database
)

func InitConnections(state string) {
	Connections = make(map[string]*sqlx.DB)
	ConfigDb = make(map[string]config.Database)
	databases := config.AppConfig.Databases
	for _, db := range databases {
		for _, schema := range db.Schema {
			NewDb(db, schema, state)
		}
	}
}

func NewDb(cf config.Database, schema, state string) {
	var connStr string
	var db *sqlx.DB
	var err error

	// build connectionString
	switch cf.Driver {
	case consts.SQLServerDriver:
		connStr = fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s",
			cf.Host, cf.Port, cf.Username, cf.Password, schema)
		db, err = sqlx.Connect(consts.SQLServerDriver, connStr)
		if err != nil {
			logger.Panic(state, "Error open connection to SQLServer: {}", err)
			return
		}
		break
	case consts.PostgresDriver:
		connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cf.Host, cf.Port, cf.Username, cf.Password, schema)
		db, err = sqlx.Connect(cf.Driver, connStr)
		if err != nil {
			logger.Panic(state, "Error open connection to Postgres: {}", err)
			return
		}
		break
	case consts.OracleDriver:
		connStr = fmt.Sprintf("%v/%v@%v:%v/%v",
			cf.Username, cf.Password, cf.Host, cf.Port, schema)
		db, err = sqlx.Connect(cf.Driver, connStr)
		if err != nil {
			logger.Panic(state, "Error open connection to Oracle: {}", err)
			return
		}
		break
	}

	if db == nil {
		logger.Panic(state, "Error db is nil")
		return
	}

	// ping check connection
	if err = db.Ping(); err != nil {
		logger.Panic("", "Error connect to database: {}", err)
		return
	}

	ConfigDb[cf.Driver] = cf
	Connections[schema] = db
}
