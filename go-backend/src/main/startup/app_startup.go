package startup

import (
	"github.com/BevisDev/backend-template/src/main/api"
	"github.com/BevisDev/backend-template/src/main/api/healthcheck"
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/BevisDev/backend-template/src/main/consts"
	"github.com/BevisDev/backend-template/src/main/dto/response"
	"github.com/BevisDev/backend-template/src/main/helper/db"
	"github.com/BevisDev/backend-template/src/main/helper/logger"
	"github.com/BevisDev/backend-template/src/main/helper/redis"
	"github.com/BevisDev/backend-template/src/main/helper/rest"
	"github.com/BevisDev/backend-template/src/main/middleware"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
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
	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// use middleware
	startMiddleware(r)

	// define group api
	apiGr := r.Group("/api")
	{
		api.APIs(apiGr)
	}

	// handler no route
	r.NoRoute(func(c *gin.Context) {
		response.SetError(c, 404, consts.NotFound)
	})
	return r
}
