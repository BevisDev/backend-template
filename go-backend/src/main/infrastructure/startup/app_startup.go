package startup

import (
	"github.com/BevisDev/backend-template/src/main/infrastructure/config"
	"github.com/BevisDev/backend-template/src/main/infrastructure/database"
	"github.com/BevisDev/backend-template/src/main/infrastructure/logger"
	"github.com/BevisDev/backend-template/src/main/infrastructure/redis"
	"github.com/BevisDev/backend-template/src/main/infrastructure/rest"
	"github.com/BevisDev/backend-template/src/main/infrastructure/router"
	"github.com/gin-gonic/gin"
)

func startConfig() {
	config.LoadConfig()
}

func startDB(state string) {
	database.InitDB(state)
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
	return router.InitRouter()
}
