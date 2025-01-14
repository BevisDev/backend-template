package service

import "context"

type IPingService interface {
	PingDB(ctx context.Context) map[string]bool
	PingRedis(ctx context.Context) map[string]bool
}
