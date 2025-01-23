package database

import (
	"context"
	"github.com/BevisDev/backend-template/src/main/common/utils"
	"github.com/BevisDev/backend-template/src/main/infrastructure/config"
	"github.com/BevisDev/backend-template/src/main/infrastructure/logger"
	"github.com/jmoiron/sqlx"
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

func logQuery(state, query string) {
	if isDev {
		logger.Info(state, "Query: {}", query)
	}
}

func GetList[T any](c context.Context, dest *T, schema, query string, args ...interface{}) {
	var (
		state  = utils.GetState(c)
		db, cf = getDBAndConfig(schema)
		err    error
	)
	if db == nil || cf == nil {
		logger.Error(state, "Error GetList: db or cf is nil with schema {}", schema)
		return
	}
	logQuery(state, query)
	ctx, cancel := utils.CreateCtxTimeout(c, cf.TimeoutSec)
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

func GetUsingNamed[T any](c context.Context, dest *T, schema, query string, args interface{}) {
	var (
		state  = utils.GetState(c)
		db, cf = getDBAndConfig(schema)
		err    error
	)
	if db == nil || cf == nil {
		logger.Error(state, "Error GetUsingNamed: db or cf is nil with schema {}", schema)
		return
	}

	logQuery(state, query)
	ctx, cancel := utils.CreateCtxTimeout(c, cf.TimeoutSec)
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

func GetUsingArgs[T any](c context.Context, dest *T, schema, query string, args ...interface{}) {
	var (
		state  = utils.GetState(c)
		db, cf = getDBAndConfig(schema)
		err    error
	)
	if db == nil || cf == nil {
		logger.Error(state, "Error GetUsingArgs: db or cf is nil with schema {}", schema)
		return
	}

	logQuery(state, query)
	ctx, cancel := utils.CreateCtxTimeout(c, cf.TimeoutSec)
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

func ExecQuery(c context.Context, isSelect bool, schema, query string, args ...interface{}) bool {
	var (
		state  = utils.GetState(c)
		db, cf = getDBAndConfig(schema)
		err    error
		tx     *sqlx.Tx
	)
	if db == nil || cf == nil {
		logger.Error(state, "Error ExecQuery: db or cf is nil with schema {}", schema)
		return false
	}
	logQuery(state, query)

	ctx, cancel := utils.CreateCtxTimeout(c, cf.TimeoutSec)
	defer cancel()

	if !isSelect {
		tx, err = db.BeginTxx(ctx, nil)
		if err != nil {
			logger.Error(state, "Error BeginTxx in ExecQuery method {}", err)
			return false

		}
	}

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		logger.Error(state, "Error ExecQuery {} ", err)
		if !isSelect {
			tx.Rollback()
		}
		return false
	}
	if !isSelect {
		err = tx.Commit()
	}
	return true
}

func Insert(c context.Context, schema, query string, args interface{}) bool {
	var (
		state  = utils.GetState(c)
		db, cf = getDBAndConfig(schema)
	)
	if db == nil || cf == nil {
		logger.Error(state, "Error Insert: db or cf is nil with schema {}", schema)
		return false
	}

	logQuery(state, query)
	ctx, cancel := utils.CreateCtxTimeout(c, cf.TimeoutSec)
	defer cancel()

	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		logger.Error(state, "Error BeginTxx in Insert method {}", err)
		return false
	}

	_, err = db.NamedExecContext(ctx, query, args)
	if err != nil {
		logger.Error(state, "Error Insert {} ", err)
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func InsertedId(c context.Context, col int, schema, query string, args ...interface{}) (int, bool) {
	var (
		state  = utils.GetState(c)
		id     int
		db, cf = getDBAndConfig(schema)
	)
	if db == nil || cf == nil {
		logger.Error(state, "Error Insert: db or cf is nil with schema {}", schema)
		return id, false
	}

	logQuery(state, query)
	ctx, cancel := utils.CreateCtxTimeout(c, cf.TimeoutSec)
	defer cancel()

	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		logger.Error(state, "Error BeginTxx in Insert method {}", err)
		return id, false
	}

	err = db.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		logger.Error(state, "Error InsertedId {} ", err)
		tx.Rollback()
		return id, false
	}
	tx.Commit()
	return id, true
}

func Update(c context.Context, schema, query string, args interface{}) bool {
	var (
		state  = utils.GetState(c)
		db, cf = getDBAndConfig(schema)
	)
	if db == nil || cf == nil {
		logger.Error(state, "Error Update: db or cf is nil with schema {}", schema)
		return false
	}

	logQuery(state, query)
	ctx, cancel := utils.CreateCtxTimeout(c, cf.TimeoutSec)
	defer cancel()

	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		logger.Error(state, "Error BeginTxx in Update method {}", err)
		return false
	}

	_, err = db.NamedExecContext(ctx, query, args)
	if err != nil {
		logger.Error(state, "Error Update {} ", err)
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func Delete(c context.Context, schema, query string, args interface{}) bool {
	var (
		state  = utils.GetState(c)
		db, cf = getDBAndConfig(schema)
	)
	if db == nil || cf == nil {
		logger.Error(state, "Error Delete: db or cf is nil with schema {}", schema)
		return false
	}

	logQuery(state, query)
	ctx, cancel := utils.CreateCtxTimeout(c, cf.TimeoutSec)
	defer cancel()

	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		logger.Error(state, "Error BeginTxx in Delete method {}", err)
		return false
	}

	_, err = db.NamedExecContext(ctx, query, args)
	if err != nil {
		logger.Error(state, "Error Delete {} ", err)
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}
