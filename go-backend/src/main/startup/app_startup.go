package startup

import (
	"github.com/BevisDev/backend-template/src/main/api"
	"github.com/BevisDev/backend-template/src/main/api/healthcheck"
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/BevisDev/backend-template/src/main/helper/db"
	"github.com/BevisDev/backend-template/src/main/helper/logger"
	"github.com/BevisDev/backend-template/src/main/helper/redis"
	"github.com/BevisDev/backend-template/src/main/helper/rest"
	"github.com/BevisDev/backend-template/src/main/middleware"
	"github.com/gin-gonic/gin"
)

func startConfig() {
	config.LoadConfig()
}

func startDB(state string) {
	db.NewDb(state)
}

func startLogger(state string) {
	logger.Init()
	logger.Info(state, "LOGGER is started successfully {}", true)
}

func startRedis(state string) {
	redis.NewRedis(state)
}

func startRestClient() {
	rest.NewRestClient()
}

func startMiddleware(r *gin.Engine) {
	r.Use(middleware.LoggerHandler())
	r.Use(middleware.AuthHandler())
	r.Use(middleware.ErrorHandler())
}

func startRouter() *gin.Engine {
	var r *gin.Engine
	serverConfig := config.AppConfig.ServerConfig

	if serverConfig.Profile == "prod" {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	} else {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	}

	// use Routers
	// ping to health check system
	healthcheck.Ping(r)

	// use middleware
	startMiddleware(r)

	// define group api
	apiGr := r.Group("/api")
	{
		api.APIs(apiGr)
	}

	return r
}
