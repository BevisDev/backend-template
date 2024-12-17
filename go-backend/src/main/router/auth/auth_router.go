package auth

import "github.com/gin-gonic/gin"

func InitAuthRouterGroup(group *gin.RouterGroup) {
	group.GET("/signin")
	group.POST("/signup")
}
