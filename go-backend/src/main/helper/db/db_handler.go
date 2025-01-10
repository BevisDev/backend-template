package db

import (
	"fmt"
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/BevisDev/backend-template/src/main/consts"
	"github.com/BevisDev/backend-template/src/main/helper/logger"
	"github.com/BevisDev/backend-template/src/main/helper/utils"

	//_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

var (
	Connections map[string]*sqlx.DB
	ConfigDb    map[string]config.Database
)

func NewDb(state string) {
	databases := config.AppConfig.Databases
	if utils.IsNilOrEmpty(databases) {
		logger.Fatal(state, "Error Config DB is not initialized")
		return
	}
	Connections = make(map[string]*sqlx.DB)
	ConfigDb = make(map[string]config.Database)
	for _, db := range databases {
		for _, schema := range db.Schema {
			newConnection(db, schema, state)
		}
	}
}

func newConnection(cf config.Database, schema, state string) {
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
			logger.Fatal(state, "Error open connection to SQLServer: {}", err)
			return
		}
		break
	case consts.PostgresDriver:
		connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cf.Host, cf.Port, cf.Username, cf.Password, schema)
		db, err = sqlx.Connect(cf.Driver, connStr)
		if err != nil {
			logger.Fatal(state, "Error open connection to Postgres: {}", err)
			return
		}
		break
	case consts.OracleDriver:
		connStr = fmt.Sprintf("%v/%v@%v:%v/%v",
			cf.Username, cf.Password, cf.Host, cf.Port, schema)
		db, err = sqlx.Connect(cf.Driver, connStr)
		if err != nil {
			logger.Fatal(state, "Error open connection to Oracle: {}", err)
			return
		}
		break
	}

	if db == nil {
		logger.Fatal(state, "Error db is nil")
		return
	}

	// ping check connection
	if err = db.Ping(); err != nil {
		logger.Fatal("", "Error connect to database: {}", err)
		return
	}

	ConfigDb[cf.Driver] = cf
	Connections[schema] = db
}
