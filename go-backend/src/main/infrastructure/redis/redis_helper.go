package redis

import (
	"context"
	"github.com/BevisDev/backend-template/src/main/common/utils"
	"github.com/BevisDev/backend-template/src/main/infrastructure/logger"
	"reflect"
	"time"
)

func Set(ctx context.Context, key string, value interface{}, expiredTimeSec int) bool {
	state := utils.GetState(ctx)
	err := redisClient.Set(ctx, key, marshalValue(value), time.Duration(expiredTimeSec)*time.Second).Err()
	if err != nil {
		logger.Error(state, "Error Redis set failed {}", err)
		return false
	}
	return true
}

func marshalValue(value interface{}) interface{} {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr ||
		v.Kind() == reflect.Struct ||
		v.Kind() == reflect.Map ||
		v.Kind() == reflect.Slice ||
		v.Kind() == reflect.Array {
		return utils.ToJSON(value)
	}
	return value
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
