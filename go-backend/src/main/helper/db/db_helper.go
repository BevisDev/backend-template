package db

import (
	"context"
	"fmt"
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/BevisDev/backend-template/src/main/consts"
	"github.com/BevisDev/backend-template/src/main/helper/logger"
	"github.com/jmoiron/sqlx"
)

var (
	Connections  map[string]*sqlx.DB
	CtxSqlServer context.Context
	CtxPostgres  context.Context
	CtxOracle    context.Context
)

type IDb interface {
	SelectList(ctx context.Context, state, query string, target interface{}, args ...interface{}) bool
	SelectOne(ctx context.Context, state, query string, args ...interface{}) bool
	Insert(ctx context.Context, query string, args ...interface{}) error
	Update(ctx context.Context, query string, args ...interface{}) error
	Delete(ctx context.Context, query string, args ...interface{}) error
}

type properties struct {
	state    string
	host     string
	port     int
	driver   string
	username string
	password string
	schema   string
	maxConns int
	maxIdle  int
	maxLife  int
}

func InitConnections(state string) {
	databases := config.AppConfig.DBConfig.Databases
	for _, db := range databases {
		props := &properties{
			state:    state,
			host:     db.Host,
			port:     db.Port,
			driver:   db.Driver,
			username: db.Username,
			password: db.Password,
		}
		for _, schema := range db.Schema {
			props.schema = schema
			NewDb(props)
		}
	}
}

func NewDb(props *properties) {
	var connStr string
	var db *sqlx.DB
	var err error

	// build connectionString
	switch props.driver {
	case consts.SQLServerDriver:
		connStr = fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s",
			props.host, props.port, props.username, props.password, props.schema)
		db, err = sqlx.Connect(props.driver, connStr)
		if err != nil {
			logger.Panic(props.state, "Error open connection to SQLServer: {}", err)
			return
		}
		break
	case consts.PostgresDriver:
		connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			props.host, props.port, props.username, props.password, props.schema)
		db, err = sqlx.Connect(props.driver, connStr)
		if err != nil {
			logger.Panic(props.state, "Error open connection to Postgres: {}", err)
			return
		}
		break
	case consts.OracleDriver:
		connStr = fmt.Sprintf("%v/%v@%v:%v/%v",
			props.username, props.password, props.host, props.port, props.schema)
		db, err = sqlx.Connect(props.driver, connStr)
		if err != nil {
			logger.Panic(props.state, "Error open connection to Oracle: {}", err)
			return
		}
		break
	}

	if db == nil {
		logger.Panic(props.state, "Error db is nil")
		return
	}

	// ping check connection
	if err = db.Ping(); err != nil {
		logger.Panic("", "Error connect to database: {}", err)
		return
	}

	Connections[props.schema] = db
}
