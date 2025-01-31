package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/BevisDev/backend-template/consts"
	"github.com/BevisDev/backend-template/utils"
	"log"
	"time"

	//_ "github.com/denisenkom/go-mssqldb"
	//_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"
)

type ConfigDB struct {
	Profile      string
	Kind         string
	Schema       string
	TimeoutSec   int
	Host         string
	Port         int
	Username     string
	Password     string
	MaxOpenConns int
	MaxIdleConns int
	MaxLifeTime  int
}

type Database struct {
	connectionMap map[string]*sqlx.DB
	dbConfigMap   map[string]*ConfigDB
	Profile       string
}

func NewDB(cf *ConfigDB) (*Database, error) {
	db := &Database{
		Profile:       cf.Profile,
		connectionMap: make(map[string]*sqlx.DB),
		dbConfigMap:   make(map[string]*ConfigDB),
	}
	err := db.newConnection(cf)
	return db, err
}

func (d *Database) newConnection(cf *ConfigDB) error {
	var (
		connStr string
		db      *sqlx.DB
		err     error
		driver  string
	)

	// build connectionString
	switch cf.Kind {
	case consts.SQLServer:
		connStr = fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s",
			cf.Host, cf.Port, cf.Username, cf.Password, cf.Schema)
		driver = "sqlserver"
		break
	case consts.Oracle:
		connStr = fmt.Sprintf("user=%s password=%s connectString=%s:%d/%s",
			cf.Username, cf.Password, cf.Host, cf.Port, cf.Schema)
		driver = "godror"
		break
	default:
		return errors.New("unsupported database kind " + cf.Kind)
	}

	// connect
	db, err = sqlx.Connect(driver, connStr)
	if err != nil {
		return err
	}

	// set pool
	db.SetMaxOpenConns(cf.MaxOpenConns)
	db.SetMaxIdleConns(cf.MaxIdleConns)
	db.SetConnMaxIdleTime(time.Duration(cf.MaxLifeTime) * time.Second)

	// ping check connection
	if err = db.Ping(); err != nil {
		return err
	}

	d.connectionMap[cf.Schema] = db
	d.dbConfigMap[cf.Schema] = cf
	log.Printf("connect db %s success \n", cf.Schema)
	return nil
}

func (d *Database) CloseAll() {
	for _, v := range d.connectionMap {
		v.Close()
	}
}

func (d *Database) GetDB(schema string) *sqlx.DB {
	return d.connectionMap[schema]
}

func (d *Database) getDBAndConfig(schema string) (*sqlx.DB, *ConfigDB) {
	if schema == "" {
		return nil, nil
	}
	return d.connectionMap[schema], d.dbConfigMap[schema]
}

func (d *Database) logQuery(query string) {
	if d.Profile == "dev" {
		log.Printf("Query: %s\n", query)
	}
}

func (d *Database) GetList(c context.Context, dest interface{}, schema, query string, args ...interface{}) error {
	var (
		db, cf = d.getDBAndConfig(schema)
		err    error
	)
	if db == nil || cf == nil {
		return errors.New("db or cf is nil")
	}

	d.logQuery(query)

	ctx, cancel := utils.CreateCtxTimeout(c, cf.TimeoutSec)
	defer cancel()

	if utils.IsNilOrEmpty(args) {
		err = db.SelectContext(ctx, dest, query)
	} else {
		err = db.SelectContext(ctx, dest, query, args...)
	}
	return err
}

func (d *Database) GetUsingNamed(c context.Context, dest interface{}, schema, query string, args interface{}) error {
	var (
		db, cf = d.getDBAndConfig(schema)
		err    error
	)
	if db == nil || cf == nil {
		return errors.New("db or cf is nil")
	}

	d.logQuery(query)

	ctx, cancel := utils.CreateCtxTimeout(c, cf.TimeoutSec)
	defer cancel()

	if utils.IsNilOrEmpty(args) {
		err = db.GetContext(ctx, dest, query)
	} else {
		err = db.GetContext(ctx, dest, query, args)
	}
	return err
}

func (d *Database) GetUsingArgs(c context.Context, dest interface{}, schema, query string, args ...interface{}) error {
	var (
		db, cf = d.getDBAndConfig(schema)
		err    error
	)
	if db == nil || cf == nil {
		return errors.New("db or cf is nil")
	}

	d.logQuery(query)

	ctx, cancel := utils.CreateCtxTimeout(c, cf.TimeoutSec)
	defer cancel()

	if utils.IsNilOrEmpty(args) {
		err = db.GetContext(ctx, dest, query)
	} else {
		err = db.GetContext(ctx, dest, query, args...)
	}
	return err
}

func (d *Database) ExecQuery(c context.Context, isSelect bool, schema, query string, args ...interface{}) error {
	var (
		db, cf = d.getDBAndConfig(schema)
		err    error
		tx     *sqlx.Tx
	)
	if db == nil || cf == nil {
		return errors.New("db or cf is nil")
	}

	d.logQuery(query)

	ctx, cancel := utils.CreateCtxTimeout(c, cf.TimeoutSec)
	defer cancel()

	if !isSelect {
		tx, err = db.BeginTxx(ctx, nil)
		if err != nil {
			return err
		}
	}

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		if !isSelect {
			tx.Rollback()
		}
		return err
	}
	if !isSelect {
		err = tx.Commit()
	}
	return err
}

func (d *Database) Insert(c context.Context, schema, query string, args interface{}) error {
	var (
		db, cf = d.getDBAndConfig(schema)
	)
	if db == nil || cf == nil {
		return errors.New("db or cf is nil")
	}

	d.logQuery(query)

	ctx, cancel := utils.CreateCtxTimeout(c, cf.TimeoutSec)
	defer cancel()

	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = db.NamedExecContext(ctx, query, args)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (d *Database) InsertedId(c context.Context, schema, query string, args ...interface{}) (int, error) {
	var (
		id     int
		db, cf = d.getDBAndConfig(schema)
	)
	if db == nil || cf == nil {
		return id, errors.New("db or cf is nil")
	}

	d.logQuery(query)

	ctx, cancel := utils.CreateCtxTimeout(c, cf.TimeoutSec)
	defer cancel()

	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return id, err
	}

	err = db.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		tx.Rollback()
		return id, err
	}
	tx.Commit()
	return id, err
}

func (d *Database) Update(c context.Context, schema, query string, args interface{}) error {
	var (
		db, cf = d.getDBAndConfig(schema)
	)
	if db == nil || cf == nil {
		return errors.New("db or cf is nil")
	}

	d.logQuery(query)

	ctx, cancel := utils.CreateCtxTimeout(c, cf.TimeoutSec)
	defer cancel()

	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = db.NamedExecContext(ctx, query, args)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (d *Database) Delete(c context.Context, schema, query string, args interface{}) error {
	var (
		db, cf = d.getDBAndConfig(schema)
	)
	if db == nil || cf == nil {
		return errors.New("db or cf is nil")
	}

	d.logQuery(query)

	ctx, cancel := utils.CreateCtxTimeout(c, cf.TimeoutSec)
	defer cancel()

	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = db.NamedExecContext(ctx, query, args)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}
