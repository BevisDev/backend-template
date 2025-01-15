package helper

import (
	"context"
	"fmt"
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/BevisDev/backend-template/src/main/consts"
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
	if IsNilOrEmpty(appConfig) ||
		IsNilOrEmpty(appConfig.Databases) {
		LogFatal(state, "Error Config DB is not initialized")
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
			LogFatal(state, "Error open connection to SQLServer: {}", err)
			return
		}
		break
	case consts.Oracle:
		connStr = fmt.Sprintf("user=%s password=%s connectString=%s:%d/%s",
			cf.Username, cf.Password, cf.Host, cf.Port, schema)
		db, err = sqlx.Connect("godror", connStr)
		if err != nil {
			LogFatal(state, "Error open connection to Oracle: {}", err)
			return
		}
		break
	default:
		LogFatal(state, "Kind db {} not supported", cf.Kind)
		return
	}

	if db == nil {
		LogFatal(state, "Error connect db {} is nil", schema)
		return
	}

	// set pool
	db.SetMaxOpenConns(cf.MaxOpenConns)
	db.SetMaxIdleConns(cf.MaxIdleConns)
	db.SetConnMaxIdleTime(time.Duration(cf.ConnMaxLifeTime) * time.Second)

	// ping check connection
	if err = db.Ping(); err != nil {
		LogFatal("", "Error ping db {}: {}", schema, err)
		return
	}

	connectionMap[schema] = db
	configDbMap[schema] = cf
	LogInfo(state, "Connect db {} successful", schema)
}

func CloseAll() {
	for _, v := range connectionMap {
		v.Close()
	}
}

func GetDB(schema string) *sqlx.DB {
	return connectionMap[schema]
}

func GetDBInfo(schema string) (*sqlx.DB, *config.Database, bool) {
	if schema == "" {
		LogError("", "Error GetList: schema is empty", schema)
		return nil, nil, false
	}
	if IsNilOrEmpty(connectionMap[schema]) ||
		IsNilOrEmpty(configDbMap[schema]) {
		return nil, nil, false
	}
	return connectionMap[schema], configDbMap[schema], true
}

func GetList(ctx context.Context, dest interface{}, schema, query string, args ...interface{}) bool {
	state := GetState(ctx)
	db, cf, ok := GetDBInfo(schema)
	if !ok {
		LogError(state, "Error GetList: db or cf is nil with schema {}", schema)
		return false
	}

	var timeout = time.Duration(cf.TimeoutSec) * time.Second
	var err error
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if IsNilOrEmpty(args) {
		err = db.SelectContext(ctx, dest, query)
	} else {
		err = db.SelectContext(ctx, dest, query, args...)
	}

	if err != nil {
		LogError(state, "Error GetList: query failed {}", err.Error())
		return false
	}

	return true
}

func GetUsingNamed(ctx context.Context, dest interface{}, schema, query string, args interface{}) bool {
	state := GetState(ctx)
	db, cf, ok := GetDBInfo(schema)
	if !ok {
		LogError(state, "Error GetUsingNamed: db or cf is nil with schema {}", schema)
		return false
	}

	var timeout = time.Duration(cf.TimeoutSec) * time.Second
	var err error
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if IsNilOrEmpty(args) {
		err = db.GetContext(ctx, dest, query)
	} else {
		err = db.GetContext(ctx, dest, query, args)
	}

	if err != nil {
		LogError(state, "Error GetUsingNamed query failed {}", err.Error())
		return false
	}
	return true
}

func GetUsingArgs(ctx context.Context, dest interface{}, schema, query string, args ...interface{}) bool {
	state := GetState(ctx)
	db, cf, ok := GetDBInfo(schema)
	if !ok {
		LogError(state, "Error GetUsingArgs: db or cf is nil with schema {}", schema)
		return false
	}

	var timeout = time.Duration(cf.TimeoutSec) * time.Second
	var err error
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if IsNilOrEmpty(args) {
		err = db.GetContext(ctx, dest, query)
	} else {
		err = db.GetContext(ctx, dest, query, args...)
	}

	if err != nil {
		LogError(state, "Error GetUsingArgs query failed {}", err.Error())
		return false
	}
	return true
}

func DBInsert(ctx context.Context, schema, query string, args interface{}) bool {
	state := GetState(ctx)
	db, cf, ok := GetDBInfo(schema)
	if !ok {
		LogError(state, "Error Insert: db or cf is nil with schema {}", schema)
		return false
	}

	var timeout = time.Duration(cf.TimeoutSec) * time.Second
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		LogError(state, "Error BeginTxx in Insert method {}", err)
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
		LogError(state, "Error Insert: query failed {}", err)
		return false
	}

	return true
}

func DBUpdate(ctx context.Context, schema, query string, args interface{}) bool {
	state := GetState(ctx)
	db, cf, ok := GetDBInfo(schema)
	if !ok {
		LogError(state, "Error DBUpdate: db or cf is nil with schema {}", schema)
		return false
	}

	var timeout = time.Duration(cf.TimeoutSec) * time.Second
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		LogError(state, "Error BeginTxx in DBUpdate method {}", err)
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
		LogError(state, "Error DBUpdate query failed {}", err)
		return false
	}

	return true
}

func DBDelete(ctx context.Context, schema, query string, args interface{}) bool {
	state := GetState(ctx)
	db, cf, ok := GetDBInfo(schema)
	if !ok {
		LogError(state, "Error DBDelete: db or cf is nil with schema {}", schema)
		return false
	}

	var timeout = time.Duration(cf.TimeoutSec) * time.Second
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		LogError(state, "Error BeginTxx in DBDelete method {}", err)
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
		LogError(state, "Error DBDelete: query failed {}", err)
		return false
	}

	return true
}
