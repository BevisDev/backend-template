package startup

import (
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/BevisDev/backend-template/src/main/helper"
	"github.com/BevisDev/backend-template/src/main/router"
	"github.com/gin-gonic/gin"
)

func startConfig() {
	config.LoadConfig()
}

func startDB(state string) {
	helper.InitDB(state)
}

func startLogger(state string) {
	helper.InitLogger()
	helper.LogInfo(state, "LOGGER is started successful {}", true)
}

func startRedis(state string) {
	helper.InitRedis(state)
}

func startRestClient(state string) {
	helper.InitRestClient(state)
}

func startRouter() *gin.Engine {
	return router.RegisterRouter()
}
