package db

import (
	"context"
	"database/sql"
	"github.com/BevisDev/backend-template/src/main/helper/logger"
)

type sqlServerDB struct {
	db *sql.DB
}

func New(schema string) IDb {
	return &sqlServerDB{db: Connections[schema]}
}

func (s *sqlServerDB) SelectList(ctx context.Context, state, query string, target interface{}, args ...interface{}) bool {
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		logger.Error(state, "Error select query failed {}", err)
		return false
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&state); err != nil {
			return false
		}
	}

	if err = rows.Err(); err != nil {
		return false
	}
	return false
}

func (s *sqlServerDB) SelectOne(ctx context.Context, state, query string, args ...interface{}) bool {
	row := s.db.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&args); err != nil {
		logger.Error(state, "Error select query failed {}", err)
		return false
	}
	return true
}

func (s *sqlServerDB) Insert(ctx context.Context, query string, args ...interface{}) bool {
	if _, err := s.db.ExecContext(ctx, query, args...); err != nil {
		return false
	}
	return true
}

func (s *sqlServerDB) Update(ctx context.Context, query string, args ...interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (s *sqlServerDB) Delete(ctx context.Context, query string, args ...interface{}) error {
	//TODO implement me
	panic("implement me")
}
