package redis

import (
	"context"
	"fmt"
	"github.com/BevisDev/backend-template/utils"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
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

func NewRedis(cf *RedisConfig) (*RedisClient, error) {
	rdb, err := newClient(cf)
	return &RedisClient{
		client: rdb,
	}, err
}

func newClient(cf *RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cf.Host, cf.Port),
		Password: cf.Password,
		DB:       cf.DB,
		PoolSize: cf.PoolSize,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	log.Println("Redis connect success")
	return rdb, err
}

func (r *RedisClient) Close() {
	if r.client != nil {
		r.client.Close()
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
