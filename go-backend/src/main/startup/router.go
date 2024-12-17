package startup

import (
	"net/http"

	"github.com/BevisDev/backend-template/src/main/global"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.AppConfig.ServerConfig.Profile == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// use Middlewares

	// use Routers
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			router.InitUserGroup.InitUserGroup(v1)
		}

		v2 := api.Group("/v2")
		{
			v2.GET("/users", getNewUsersHandler) // /api/v2/users
		}
	}

	return r
}
