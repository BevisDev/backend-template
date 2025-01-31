package redis

import (
	"context"
	"fmt"
	"github.com/BevisDev/backend-template/utils"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
	"time"
)

var (
	redisOnce   sync.Once
	redisClient *RedisClient
)

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
	PoolSize int
}

type RedisClient struct {
	client *redis.Client
}

func InitRedis(cf *RedisConfig) *RedisClient {
	redisOnce.Do(func() {
		r := &RedisClient{}
		r.client = newClient(cf)
		redisClient = r
	})
	return redisClient
}

func newClient(cf *RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cf.Host, cf.Port),
		Password: cf.Password,
		DB:       cf.DB,
		PoolSize: cf.PoolSize,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Redis connect fail")
		return nil
	}
	log.Println("Redis connect success")
	return rdb
}

func (r *RedisClient) Close() {
	if redisClient != nil {
		redisClient.Close()
	}
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiredTimeSec int) error {
	err := r.client.Set(ctx, key, utils.CheckTypeAndConvert(value), time.Duration(expiredTimeSec)*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) Get(ctx context.Context, key string, result interface{}) error {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	if err = utils.FromJSONStr(val, &result); err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) Delete(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
