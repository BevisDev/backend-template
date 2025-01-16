package repository

import (
	"context"
	"github.com/BevisDev/backend-template/src/main/helper"
)

type PingRepository struct {
}

func NewPingRepository() IPingRepository {
	return &PingRepository{}
}

func (p PingRepository) Get1MSSQL(ctx context.Context, schema string) bool {
	var result int
	helper.DBGetUsingNamed(ctx, &result, schema, "SELECT 1", nil)
	return result != 0
}

func (p PingRepository) Get1Orc(ctx context.Context, schema string) bool {
	var result int
	helper.DBGetUsingNamed(ctx, &result, schema, "SELECT 1 FROM DUAL", nil)
	return result != 0
}
