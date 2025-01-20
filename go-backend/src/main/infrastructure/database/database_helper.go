package database

import (
	"context"
	"github.com/BevisDev/backend-template/src/main/common/utils"
	"github.com/BevisDev/backend-template/src/main/infrastructure/config"
	"github.com/BevisDev/backend-template/src/main/infrastructure/logger"
	"github.com/jmoiron/sqlx"
	"time"
)

func CloseAll() {
	for _, v := range connectionMap {
		v.Close()
	}
}

func GetDB(schema string) *sqlx.DB {
	return connectionMap[schema]
}

func getDBAndConfig(schema string) (*sqlx.DB, *config.Database) {
	if schema == "" {
		logger.Error("", "Error getDBAndConfig: schema is empty", schema)
		return nil, nil
	}
	if utils.IsNilOrEmpty(connectionMap[schema]) ||
		utils.IsNilOrEmpty(dbConfigMap[schema]) {
		return nil, nil
	}
	return connectionMap[schema], dbConfigMap[schema]
}

func GetList[T any](ctx context.Context, dest *T, schema, query string, args ...interface{}) {
	state := utils.GetState(ctx)
	db, cf := getDBAndConfig(schema)
	if db == nil || cf == nil {
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
	db, cf := getDBAndConfig(schema)
	if db == nil || cf == nil {
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
	db, cf := getDBAndConfig(schema)
	if db == nil || cf == nil {
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
	db, cf := getDBAndConfig(schema)
	if db == nil || cf == nil {
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
	db, cf := getDBAndConfig(schema)
	if db == nil || cf == nil {
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
	db, cf := getDBAndConfig(schema)
	if db == nil || cf == nil {
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
