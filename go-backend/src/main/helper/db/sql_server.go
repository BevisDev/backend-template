package db

import (
	"context"
	"github.com/BevisDev/backend-template/src/main/consts"
	"github.com/BevisDev/backend-template/src/main/helper/logger"
	"github.com/BevisDev/backend-template/src/main/helper/utils"
	"github.com/jmoiron/sqlx"
	"sync"
	"time"
)

var onceSQlServer sync.Once
var instanceSQLServer IDb

type sqlServerDB struct {
	db         *sqlx.DB
	timeoutSec time.Duration
}

func NewSqlServer(schema string) IDb {
	onceSQlServer.Do(func() {
		instanceSQLServer = &sqlServerDB{
			db:         Connections[schema],
			timeoutSec: time.Duration(ConfigDb[consts.SQLServerDriver].TimeoutSec) * time.Second,
		}
	})
	return instanceSQLServer
}

func (s *sqlServerDB) GetList(dest interface{}, state, query string, args map[string]interface{}) bool {
	if !utils.IsPtrOrStruct(dest) {
		logger.Error(state, "dest must be a pointer")
		return false
	}
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), s.timeoutSec)
	defer cancel()

	if utils.IsNilOrEmpty(args) {
		err = s.db.SelectContext(ctx, dest, query)
	} else {
		err = s.db.SelectContext(ctx, dest, query, args)
	}

	if err != nil {
		logger.Error(state, "Error GetList query failed {}", err.Error())
		return false
	}

	return true
}

func (s *sqlServerDB) Get(dest interface{}, state, query string, args map[string]interface{}) bool {
	if !utils.IsPtrOrStruct(dest) {
		logger.Error(state, "dest must be a pointer")
		return false
	}
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), s.timeoutSec)
	defer cancel()

	if utils.IsNilOrEmpty(args) {
		err = s.db.GetContext(ctx, dest, query)
	} else {
		err = s.db.GetContext(ctx, dest, query, args)
	}

	if err != nil {
		logger.Error(state, "Error Get query failed {}", err.Error())
		return false
	}
	return true
}

func (s *sqlServerDB) Insert(state, query string, args interface{}) bool {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeoutSec)
	defer cancel()

	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		logger.Error(state, "Error begin trans in insert method {}", err)
		return false
	}

	// rollback if has error
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if _, err = s.db.NamedExecContext(ctx, query, args); err != nil {
		logger.Error(state, "Error insert query failed {}", err)
		return false
	}

	err = tx.Commit()
	if err != nil {
		logger.Error(state, "Error commit transaction in insert method {}", err)
		return false
	}

	return true
}

func (s *sqlServerDB) Update(state, query string, args interface{}) bool {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeoutSec)
	defer cancel()

	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		logger.Error(state, "Error begin trans in update method {}", err)
		return false
	}

	// rollback if has error
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if _, err = s.db.NamedExecContext(ctx, query, args); err != nil {
		logger.Error(state, "Error update query failed {}", err)
		return false
	}

	err = tx.Commit()
	if err != nil {
		logger.Error(state, "Error commit transaction in update method {}", err)
		return false
	}

	return true
}

func (s *sqlServerDB) Delete(state, query string, args interface{}) bool {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeoutSec)
	defer cancel()

	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		logger.Error(state, "Error begin trans in delete method {}", err)
		return false
	}

	// rollback if has error
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if _, err = s.db.NamedExecContext(ctx, query, args); err != nil {
		logger.Error(state, "Error delete query failed {}", err)
		return false
	}

	err = tx.Commit()
	if err != nil {
		logger.Error(state, "Error commit transaction in delete method {}", err)
		return false
	}

	return true
}
