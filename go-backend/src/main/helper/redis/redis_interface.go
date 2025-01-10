package redis

type IRedis interface {
	Set(state, key string, value interface{}, expiredTimeSec int) bool
	Get(state, key string, result interface{}) bool
	Delete(state, key string) bool
}
