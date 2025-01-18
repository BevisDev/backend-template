package redis

import (
	"context"
	"fmt"
	"github.com/BevisDev/backend-template/src/main/common/utils"
	"github.com/BevisDev/backend-template/src/main/infrastructure/config"
	"github.com/BevisDev/backend-template/src/main/infrastructure/logger"
	"github.com/redis/go-redis/v9"
	"reflect"
	"sync"
	"time"
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

func Set(ctx context.Context, key string, value interface{}, expiredTimeSec int) bool {
	state := utils.GetState(ctx)
	var v interface{}
	if utils.IsPtrOrStruct(value) {
		v = utils.ToJSON(value)
	} else {
		v = value
	}
	err := redisClient.Set(ctx, key, v, time.Duration(expiredTimeSec)*time.Second).Err()
	if err != nil {
		logger.Error(state, "Error Redis set failed {}", err)
		return false
	}
	return true
}

func Get(ctx context.Context, key string, result interface{}) bool {
	state := utils.GetState(ctx)
	val, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		logger.Error(state, "Error Redis get failed {}", err)
		return false
	}
	if val == "" {
		logger.Error(state, "Error get value in Redis with key {} is empty", key)
		return false
	}
	if err = utils.FromJSONStr(val, &result); err != nil {
		logger.Error(state, "Error deserialize JSON with type result {}, err {}", reflect.TypeOf(result), err)
		return false
	}
	return true
}

func Delete(ctx context.Context, key string) bool {
	state := utils.GetState(ctx)
	err := redisClient.Del(ctx, key).Err()
	if err != nil {
		logger.Error(state, "Error Redis delete {} failed {}", key, err)
		return false
	}
	return true
}
