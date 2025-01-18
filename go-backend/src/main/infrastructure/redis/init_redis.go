package redis

import (
	"context"
	"fmt"
	"github.com/BevisDev/backend-template/src/main/common/utils"
	"github.com/BevisDev/backend-template/src/main/infrastructure/config"
	"github.com/BevisDev/backend-template/src/main/infrastructure/logger"
	"github.com/redis/go-redis/v9"
	"sync"
)

var (
	redisOnce   sync.Once
	redisClient *redis.Client
)

func InitRedis(state string) *redis.Client {
	redisOnce.Do(func() {
		redisClient = newRedis(state)
	})
	return redisClient
}

func newRedis(state string) *redis.Client {
	appConfig := config.AppConfig
	if utils.IsNilOrEmpty(appConfig) ||
		utils.IsNilOrEmpty(appConfig.RedisConfig) {
		logger.Fatal(state, "Error Config Redis is not initialized")
		return nil
	}
	cf := appConfig.RedisConfig
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cf.Host, cf.Port),
		Password: cf.Password,
		DB:       cf.Index,
		PoolSize: cf.PoolSize,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logger.Fatal(state, "Error Redis ping failed {}", err)
		return nil
	}
	logger.Info(state, "Connect Redis successful")
	return rdb
}

func Close() {
	if redisClient != nil {
		redisClient.Close()
	}
}
