package router

import (
	"github.com/BevisDev/backend-template/src/main/api"
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/BevisDev/backend-template/src/main/consts"
	"github.com/BevisDev/backend-template/src/main/di"
	"github.com/BevisDev/backend-template/src/main/dto/response"
	"github.com/BevisDev/backend-template/src/main/middleware"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func RegisterRouter() *gin.Engine {
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
	pingC := di.NewPingDI()
	r.GET("/ping", pingC.Ping)
	r.GET("/db", pingC.PingDB)
	r.GET("/redis", pingC.PingRedis)

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// use middleware
	middleware.RegisterMiddleware(r)

	// define group api
	apiGr := r.Group("/api")
	{
		api.RegisterAPIs(apiGr)
	}

	// handler no route
	r.NoRoute(func(c *gin.Context) {
		response.SetError(c, 404, consts.NotFound)
	})
	return r
}
