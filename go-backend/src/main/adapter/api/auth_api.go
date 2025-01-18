package api

import (
	"github.com/BevisDev/backend-template/src/main/adapter/controller"
	"github.com/gin-gonic/gin"
)

func AuthAPI(g *gin.RouterGroup) {
	g.GET("/signin", controller.SignIn)
	g.GET("/signup", controller.SignUp)
}
