package startup

import (
	"github.com/BevisDev/backend-template/src/main/api/v1"
	"github.com/BevisDev/backend-template/src/main/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
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
		version1 := api.Group("/v1")
		{
			// router auth
			v1.AuthApiGroup(version1)
		}
	}

	return r
}
