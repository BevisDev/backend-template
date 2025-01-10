package redis

func Set(state, key string, value interface{}, expiredTimeSec int) bool {
	return NewRedis(state).Set(state, key, value, expiredTimeSec)
}

func Get(state, key string, result interface{}) bool {
	return NewRedis(state).Get(state, key, result)
}

func Delete(state, key string) bool {
	return NewRedis(state).Delete(state, key)
}
