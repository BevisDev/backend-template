package startup

import (
	"net/http"

	"github.com/BevisDev/backend-template/src/main/global"
	"github.com/BevisDev/backend-template/src/main/router"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	serverConfig := global.AppConfig.ServerConfig

	if serverConfig.Profile == "prod" {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	} else {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	}

	// use Middlewares

	// use Routers
	// ping to health check system
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	{
		// version 1
		v1 := api.Group("/v1")
		{
			// router auth
			router.InitAuthRouterGroup(v1)
		}
	}

	return r
}
