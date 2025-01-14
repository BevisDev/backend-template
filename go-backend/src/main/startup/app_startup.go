package startup

import (
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/BevisDev/backend-template/src/main/helper/db"
	"github.com/BevisDev/backend-template/src/main/helper/logger"
	"github.com/BevisDev/backend-template/src/main/helper/redis"
	"github.com/BevisDev/backend-template/src/main/helper/rest"
	"github.com/BevisDev/backend-template/src/main/router"
	"github.com/gin-gonic/gin"
)

func startConfig() {
	config.LoadConfig()
}

func startDB(state string) {
	db.InitDB(state)
}

func startLogger(state string) {
	logger.InitLogger()
	logger.Info(state, "LOGGER is started successful {}", true)
}

func startRedis(state string) {
	redis.InitRedis(state)
}

func startRestClient(state string) {
	rest.InitRestClient(state)
}

func startRouter() *gin.Engine {
	return router.RegisterRouter()
}
