package repository

import (
	"context"
	"github.com/BevisDev/backend-template/src/main/helper/db"
)

type PingRepository struct {
}

func NewPingRepository() IPingRepository {
	return &PingRepository{}
}

func (p PingRepository) Get1MSSQL(ctx context.Context, schema string) bool {
	var result int
	return db.GetUsingNamed(ctx, &result, schema, "SELECT 1", nil)
}

func (p PingRepository) Get1Orc(ctx context.Context, schema string) bool {
	var result int
	return db.GetUsingNamed(ctx, &result, schema, "SELECT 1 FROM DUAL", nil)
}
