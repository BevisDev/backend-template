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
	//_ "github.com/lib/pq"
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
	db         *sqlx.DB
	Profile    string
	TimeoutSec int
}

func NewDB(cf *ConfigDB) (*Database, error) {
	database := &Database{
		Profile:    cf.Profile,
		TimeoutSec: cf.TimeoutSec,
	}
	db, err := database.newConnection(cf)
	database.db = db
	return database, err
}

func (d *Database) newConnection(cf *ConfigDB) (*sqlx.DB, error) {
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
	case consts.Postgres:
		connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cf.Host, cf.Port, cf.Username, cf.Password, cf.Schema)
		driver = "postgres"
	case consts.Oracle:
		connStr = fmt.Sprintf("user=%s password=%s connectString=%s:%d/%s",
			cf.Username, cf.Password, cf.Host, cf.Port, cf.Schema)
		driver = "godror"
		break
	default:
		return nil, errors.New("unsupported database kind " + cf.Kind)
	}

	// connect
	db, err = sqlx.Connect(driver, connStr)
	if err != nil {
		return nil, err
	}

	// set pool
	db.SetMaxOpenConns(cf.MaxOpenConns)
	db.SetMaxIdleConns(cf.MaxIdleConns)
	db.SetConnMaxIdleTime(time.Duration(cf.MaxLifeTime) * time.Second)

	// ping check connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Printf("connect db %s success \n", cf.Schema)
	return db, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) logQuery(query string) {
	if d.Profile == "dev" {
		log.Printf("Query: %s\n", query)
	}
}

func (d *Database) GetList(c context.Context, dest interface{}, query string, args ...interface{}) error {
	var err error
	d.logQuery(query)
	ctx, cancel := utils.CreateCtxTimeout(c, d.TimeoutSec)
	defer cancel()

	if utils.IsNilOrEmpty(args) {
		err = d.db.SelectContext(ctx, dest, query)
	} else {
		err = d.db.SelectContext(ctx, dest, query, args...)
	}
	return err
}

func (d *Database) GetUsingNamed(c context.Context, dest interface{}, query string, args interface{}) error {
	var err error
	d.logQuery(query)
	ctx, cancel := utils.CreateCtxTimeout(c, d.TimeoutSec)
	defer cancel()

	if utils.IsNilOrEmpty(args) {
		err = d.db.GetContext(ctx, dest, query)
	} else {
		err = d.db.GetContext(ctx, dest, query, args)
	}
	return err
}

func (d *Database) GetUsingArgs(c context.Context, dest interface{}, query string, args ...interface{}) error {
	var err error
	d.logQuery(query)
	ctx, cancel := utils.CreateCtxTimeout(c, d.TimeoutSec)
	defer cancel()

	if utils.IsNilOrEmpty(args) {
		err = d.db.GetContext(ctx, dest, query)
	} else {
		err = d.db.GetContext(ctx, dest, query, args...)
	}
	return err
}

func (d *Database) ExecQuery(c context.Context, isSelect bool, query string, args ...interface{}) error {
	var (
		err error
		tx  *sqlx.Tx
	)
	d.logQuery(query)
	ctx, cancel := utils.CreateCtxTimeout(c, d.TimeoutSec)
	defer cancel()

	if !isSelect {
		tx, err = d.db.BeginTxx(ctx, nil)
		if err != nil {
			return err
		}
	}

	_, err = d.db.ExecContext(ctx, query, args...)
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

func (d *Database) Insert(c context.Context, query string, args interface{}) error {
	var (
		err error
		tx  *sqlx.Tx
	)
	d.logQuery(query)
	ctx, cancel := utils.CreateCtxTimeout(c, d.TimeoutSec)
	defer cancel()

	tx, err = d.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = d.db.NamedExecContext(ctx, query, args)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (d *Database) InsertedId(c context.Context, query string, args ...interface{}) (int, error) {
	var (
		id  int
		tx  *sqlx.Tx
		err error
	)
	d.logQuery(query)
	ctx, cancel := utils.CreateCtxTimeout(c, d.TimeoutSec)
	defer cancel()

	tx, err = d.db.BeginTxx(ctx, nil)
	if err != nil {
		return id, err
	}

	err = d.db.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		tx.Rollback()
		return id, err
	}
	tx.Commit()
	return id, err
}

func (d *Database) Update(c context.Context, query string, args interface{}) error {
	var (
		err error
		tx  *sqlx.Tx
	)
	d.logQuery(query)
	ctx, cancel := utils.CreateCtxTimeout(c, d.TimeoutSec)
	defer cancel()

	tx, err = d.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = d.db.NamedExecContext(ctx, query, args)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (d *Database) Delete(c context.Context, query string, args interface{}) error {
	var (
		err error
		tx  *sqlx.Tx
	)
	d.logQuery(query)
	ctx, cancel := utils.CreateCtxTimeout(c, d.TimeoutSec)
	defer cancel()

	tx, err = d.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = d.db.NamedExecContext(ctx, query, args)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}
