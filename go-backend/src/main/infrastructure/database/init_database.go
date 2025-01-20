package database

import (
	"fmt"
	"github.com/BevisDev/backend-template/src/main/common/consts"
	"github.com/BevisDev/backend-template/src/main/common/utils"
	"github.com/BevisDev/backend-template/src/main/infrastructure/config"
	"github.com/BevisDev/backend-template/src/main/infrastructure/logger"
	"sync"
	"time"

	//_ "github.com/denisenkom/go-mssqldb"
	//_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"
)

var (
	dbOnce        sync.Once
	connectionMap map[string]*sqlx.DB
	dbConfigMap   map[string]*config.Database
)

func InitDB(state string) {
	appConfig := config.AppConfig
	if utils.IsNilOrEmpty(appConfig) ||
		utils.IsNilOrEmpty(appConfig.Databases) {
		logger.Fatal(state, "Error Config DB is not initialized")
		return
	}
	if connectionMap == nil {
		connectionMap = make(map[string]*sqlx.DB)
	}
	if dbConfigMap == nil {
		dbConfigMap = make(map[string]*config.Database)
	}
	dbOnce.Do(func() {
		for _, db := range config.AppConfig.Databases {
			for _, schema := range db.Schema {
				newConnection(&db, schema, state)
			}
		}
	})
}

func newConnection(cf *config.Database, schema, state string) {
	var connStr string
	var db *sqlx.DB
	var err error

	// build connectionString
	switch cf.Kind {
	case consts.SQLServer:
		connStr = fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s",
			cf.Host, cf.Port, cf.Username, cf.Password, schema)
		db, err = sqlx.Connect("sqlserver", connStr)
		if err != nil {
			logger.Fatal(state, "Error open connection to SQLServer: {}", err)
			return
		}
		break
	case consts.Oracle:
		connStr = fmt.Sprintf("user=%s password=%s connectString=%s:%d/%s",
			cf.Username, cf.Password, cf.Host, cf.Port, schema)
		db, err = sqlx.Connect("godror", connStr)
		if err != nil {
			logger.Fatal(state, "Error open connection to Oracle: {}", err)
			return
		}
		break
	default:
		logger.Fatal(state, "Kind db {} not supported", cf.Kind)
		return
	}

	if db == nil {
		logger.Fatal(state, "Error connect db {} is nil", schema)
		return
	}

	// set pool
	db.SetMaxOpenConns(cf.MaxOpenConns)
	db.SetMaxIdleConns(cf.MaxIdleConns)
	db.SetConnMaxIdleTime(time.Duration(cf.ConnMaxLifeTime) * time.Second)

	// ping check connection
	if err = db.Ping(); err != nil {
		logger.Fatal("", "Error ping db {}: {}", schema, err)
		return
	}

	connectionMap[schema] = db
	dbConfigMap[schema] = cf
	logger.Info(state, "Connect db {} successful", schema)
}
