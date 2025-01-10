package startup

import "github.com/BevisDev/backend-template/src/main/helper/redis"

func startRedis(state string) {
	redis.NewRedis(state)
}
