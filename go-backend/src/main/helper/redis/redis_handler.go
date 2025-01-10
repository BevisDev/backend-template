package redis

import (
	"context"
	"fmt"
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/BevisDev/backend-template/src/main/helper/json"
	"github.com/BevisDev/backend-template/src/main/helper/logger"
	"github.com/BevisDev/backend-template/src/main/helper/utils"
	"github.com/redis/go-redis/v9"
	"reflect"
	"sync"
	"time"
)

var (
	onceRedis     sync.Once
	instanceRedis IRedis
)

type Redis struct {
	client *redis.Client
}

func NewRedis(state string) IRedis {
	onceRedis.Do(func() {
		instanceRedis = &Redis{
			client: initRedis(state),
		}
	})
	return instanceRedis
}

func initRedis(state string) *redis.Client {
	cf := config.AppConfig.RedisConfig
	if utils.IsNilOrEmpty(cf) {
		logger.Fatal(state, "Error Config Redis is not initialized")
		return nil
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cf.Host, cf.Port),
		Password: cf.Password,
		DB:       cf.Index,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logger.Fatal(state, "Error Redis ping failed {}", err)
		return nil
	}

	return rdb
}

func (r *Redis) Set(state, key string, value interface{}, expiredTimeSec int) bool {
	var v interface{}
	if utils.IsPtrOrStruct(value) {
		v = json.ToJSON(value)
	}
	err := r.client.Set(context.Background(), key, v, time.Duration(expiredTimeSec)*time.Second).Err()
	if err != nil {
		logger.Error(state, "Error Redis set failed {}", err)
		return false
	}
	return true
}

func (r *Redis) Get(state, key string, result interface{}) bool {
	val, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		logger.Error(state, "Error Redis get failed {}", err)
		return false
	}
	if val == "" {
		logger.Error(state, "Error get value in Redis with key {} is empty", key)
		return false
	}
	if err = json.FromJSONStr(val, &result); err != nil {
		logger.Error(state, "Error deserialize JSON with type result {}, err {}", reflect.TypeOf(result), err)
		return false
	}
	return true
}

func (r *Redis) Delete(state, key string) bool {
	err := r.client.Del(context.Background(), key).Err()
	if err != nil {
		logger.Error(state, "Error Redis delete {} failed {}", key, err)
		return false
	}
	return true
}
