package ping

import (
	"context"
	"github.com/BevisDev/backend-template/src/main/helper"
	"github.com/BevisDev/backend-template/src/main/repository"
)

type PingServiceImpl struct {
	pingRepository repository.IPingRepository
}

func NewPingServiceImpl(
	pingRepository repository.IPingRepository,
) IPingService {
	return &PingServiceImpl{
		pingRepository: pingRepository,
	}
}

func (impl *PingServiceImpl) PingDB(ctx context.Context) map[string]bool {
	return map[string]bool{
		"Schema1": impl.pingRepository.Get1MSSQL(ctx, "Schema1"),
		"Schema2": impl.pingRepository.Get1Orc(ctx, "Schema2"),
	}
}

func (impl *PingServiceImpl) PingRedis(ctx context.Context) map[string]bool {
	var resp = make(map[string]bool)
	if !helper.RedisSet(ctx, "key1", 1, 10) {
		resp["Redis"] = false
		return resp
	}
	var rs int
	if !helper.RedisGet(ctx, "key1", &rs) {
		resp["Redis"] = false
		return resp
	}
	resp["Redis"] = true
	return resp
}
