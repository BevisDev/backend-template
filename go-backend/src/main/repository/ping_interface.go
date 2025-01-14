package repository

import "context"

type IPingRepository interface {
	Get1MSSQL(ctx context.Context, schema string) bool
	Get1Orc(ctx context.Context, schema string) bool
}
