package database

import (
	"context"
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
	onceDb        sync.Once
	connectionMap map[string]*sqlx.DB
	configDbMap   map[string]*config.Database
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
	if configDbMap == nil {
		configDbMap = make(map[string]*config.Database)
	}
	onceDb.Do(func() {
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
	configDbMap[schema] = cf
	logger.Info(state, "Connect db {} successful", schema)
}

func CloseAll() {
	for _, v := range connectionMap {
		v.Close()
	}
}

func GetDB(schema string) *sqlx.DB {
	return connectionMap[schema]
}

func GetDBAndConfig(schema string) (*sqlx.DB, *config.Database, bool) {
	if schema == "" {
		logger.Error("", "Error GetList: schema is empty", schema)
		return nil, nil, false
	}
	if utils.IsNilOrEmpty(connectionMap[schema]) ||
		utils.IsNilOrEmpty(configDbMap[schema]) {
		return nil, nil, false
	}
	return connectionMap[schema], configDbMap[schema], true
}

func GetList[T any](ctx context.Context, dest *T, schema, query string, args ...interface{}) {
	state := utils.GetState(ctx)
	db, cf, ok := GetDBAndConfig(schema)
	if !ok {
		logger.Error(state, "Error GetList: db or cf is nil with schema {}", schema)
		return
	}

	var timeout = time.Duration(cf.TimeoutSec) * time.Second
	var err error
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if utils.IsNilOrEmpty(args) {
		err = db.SelectContext(ctx, dest, query)
	} else {
		err = db.SelectContext(ctx, dest, query, args...)
	}

	if err != nil {
		logger.Error(state, "Error GetList: query failed {}", err.Error())
		return
	}
}

func GetUsingNamed[T any](ctx context.Context, dest *T, schema, query string, args interface{}) {
	state := utils.GetState(ctx)
	db, cf, ok := GetDBAndConfig(schema)
	if !ok {
		logger.Error(state, "Error GetUsingNamed: db or cf is nil with schema {}", schema)
		return
	}

	var timeout = time.Duration(cf.TimeoutSec) * time.Second
	var err error
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if utils.IsNilOrEmpty(args) {
		err = db.GetContext(ctx, dest, query)
	} else {
		err = db.GetContext(ctx, dest, query, args)
	}

	if err != nil {
		logger.Error(state, "Error GetUsingNamed query failed {}", err.Error())
		return
	}
}

func GetUsingArgs[T any](ctx context.Context, dest *T, schema, query string, args ...interface{}) {
	state := utils.GetState(ctx)
	db, cf, ok := GetDBAndConfig(schema)
	if !ok {
		logger.Error(state, "Error GetUsingArgs: db or cf is nil with schema {}", schema)
		return
	}

	var timeout = time.Duration(cf.TimeoutSec) * time.Second
	var err error
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if utils.IsNilOrEmpty(args) {
		err = db.GetContext(ctx, dest, query)
	} else {
		err = db.GetContext(ctx, dest, query, args...)
	}

	if err != nil {
		logger.Error(state, "Error GetUsingArgs query failed {}", err.Error())
		return
	}
}

func Insert(ctx context.Context, schema, query string, args interface{}) bool {
	state := utils.GetState(ctx)
	db, cf, ok := GetDBAndConfig(schema)
	if !ok {
		logger.Error(state, "Error Insert: db or cf is nil with schema {}", schema)
		return false
	}

	var timeout = time.Duration(cf.TimeoutSec) * time.Second
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		logger.Error(state, "Error BeginTxx in Insert method {}", err)
		return false
	}

	// rollback if has error
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	if _, err = db.NamedExecContext(ctx, query, args); err != nil {
		logger.Error(state, "Error Insert: query failed {}", err)
		return false
	}

	return true
}

func Update(ctx context.Context, schema, query string, args interface{}) bool {
	state := utils.GetState(ctx)
	db, cf, ok := GetDBAndConfig(schema)
	if !ok {
		logger.Error(state, "Error Update: db or cf is nil with schema {}", schema)
		return false
	}

	var timeout = time.Duration(cf.TimeoutSec) * time.Second
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		logger.Error(state, "Error BeginTxx in Update method {}", err)
		return false
	}

	// rollback if has error
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	if _, err = db.NamedExecContext(ctx, query, args); err != nil {
		logger.Error(state, "Error Update query failed {}", err)
		return false
	}

	return true
}

func Delete(ctx context.Context, schema, query string, args interface{}) bool {
	state := utils.GetState(ctx)
	db, cf, ok := GetDBAndConfig(schema)
	if !ok {
		logger.Error(state, "Error Delete: db or cf is nil with schema {}", schema)
		return false
	}

	var timeout = time.Duration(cf.TimeoutSec) * time.Second
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		logger.Error(state, "Error BeginTxx in Delete method {}", err)
		return false
	}

	// rollback if has error
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	if _, err = db.NamedExecContext(ctx, query, args); err != nil {
		logger.Error(state, "Error Delete: query failed {}", err)
		return false
	}

	return true
}
