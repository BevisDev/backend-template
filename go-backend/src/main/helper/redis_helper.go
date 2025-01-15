package helper

import (
	"context"
	"fmt"
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/redis/go-redis/v9"
	"reflect"
	"sync"
	"time"
)

var (
	onceRedis   sync.Once
	redisClient *redis.Client
)

func InitRedis(state string) *redis.Client {
	onceRedis.Do(func() {
		redisClient = newRedis(state)
	})
	return redisClient
}

func newRedis(state string) *redis.Client {
	appConfig := config.AppConfig
	if IsNilOrEmpty(appConfig) ||
		IsNilOrEmpty(appConfig.RedisConfig) {
		LogFatal(state, "Error Config Redis is not initialized")
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
		LogFatal(state, "Error Redis ping failed {}", err)
		return nil
	}
	LogInfo(state, "Connect Redis successful")
	return rdb
}

func Close() {
	if redisClient != nil {
		redisClient.Close()
	}
}

func RedisSet(ctx context.Context, key string, value interface{}, expiredTimeSec int) bool {
	state := GetState(ctx)
	var v interface{}
	if IsPtrOrStruct(value) {
		v = ToJSON(value)
	} else {
		v = value
	}
	err := redisClient.Set(ctx, key, v, time.Duration(expiredTimeSec)*time.Second).Err()
	if err != nil {
		LogError(state, "Error Redis set failed {}", err)
		return false
	}
	return true
}

func RedisGet(ctx context.Context, key string, result interface{}) bool {
	state := GetState(ctx)
	val, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		LogError(state, "Error Redis get failed {}", err)
		return false
	}
	if val == "" {
		LogError(state, "Error get value in Redis with key {} is empty", key)
		return false
	}
	if err = FromJSONStr(val, &result); err != nil {
		LogError(state, "Error deserialize JSON with type result {}, err {}", reflect.TypeOf(result), err)
		return false
	}
	return true
}

func RedisDelete(ctx context.Context, key string) bool {
	state := GetState(ctx)
	err := redisClient.Del(ctx, key).Err()
	if err != nil {
		LogError(state, "Error Redis delete {} failed {}", key, err)
		return false
	}
	return true
}
