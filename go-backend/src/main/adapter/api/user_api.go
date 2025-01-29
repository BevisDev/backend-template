package api

import "github.com/gin-gonic/gin"

func UserAPI(g *gin.RouterGroup) {
	g.Group("/user")
	{
		g.POST("/")
		g.POST("/role")
	}
}
